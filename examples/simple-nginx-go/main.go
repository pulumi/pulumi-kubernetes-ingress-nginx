package main

import (
	ingressnginx "github.com/pulumi/pulumi-kubernetes-ingress-nginx/sdk/go/kubernetes-ingress-nginx"

	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	networkingv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/networking/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Create a sandbox namespace
		ns, err := corev1.NewNamespace(ctx, "sandbox-ns", &corev1.NamespaceArgs{})
		if err != nil {
			return err
		}

		// Install the NGINX ingress controller
		ctrl, err := ingressnginx.NewIngressController(ctx, "myctrl", &ingressnginx.IngressControllerArgs{
			Controller: &ingressnginx.ControllerArgs{
				HostPort: &ingressnginx.ControllerHostPortArgs{
					Enabled: pulumi.BoolPtr(true),
				},
				PublishService: &ingressnginx.ControllerPublishServiceArgs{
					Enabled: pulumi.BoolPtr(true),
				},
				Service: &ingressnginx.ControllerServiceArgs{
					Annotations:           pulumi.StringMap{"pulumi.com/test-annotation1": "test-value1", "pulumi.com/test-annotation2": "test-value2"},
					Type:                  pulumi.String("NodePort"),
					ExternalTrafficPolicy: pulumi.String("Local"),
				},
			},
			HelmOptions: &ingressnginx.ReleaseArgs{
				Namespace: ns.Metadata.Name().Elem(),
			},
		})
		if err != nil {
			return err
		}

		// Names for the two applications
		appBase := "hello-k8s"
		appNames := []string{appBase + "-first", appBase + "-second"}

		var appStatuses pulumi.Array
		for _, appName := range appNames {
			// Create the service for the app
			appSvc, err := corev1.NewService(ctx, appName+"-svc", &corev1.ServiceArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Name:      pulumi.String(appName),
					Namespace: ns.Metadata.Name().Elem(),
				},
				Spec: &corev1.ServiceSpecArgs{
					Type: pulumi.String("ClusterIP"),
					Ports: corev1.ServicePortArray{
						&corev1.ServicePortArgs{
							Port:       pulumi.Int(80),
							TargetPort: pulumi.Int(8080),
						},
					},
					Selector: pulumi.StringMap{
						"app": pulumi.String(appName),
					},
				},
			})
			if err != nil {
				return err
			}

			// Create the deployment for the app
			appsv1.NewDeployment(ctx, appName+"-dep", &appsv1.DeploymentArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Name:      pulumi.String(appName),
					Namespace: ns.Metadata.Name().Elem(),
				},
				Spec: &appsv1.DeploymentSpecArgs{
					Replicas: pulumi.Int(3),
					Selector: &metav1.LabelSelectorArgs{
						MatchLabels: pulumi.StringMap{
							"app": pulumi.String(appName),
						},
					},
					Template: &corev1.PodTemplateSpecArgs{
						Metadata: &metav1.ObjectMetaArgs{
							Labels: pulumi.StringMap{
								"app": pulumi.String(appName),
							},
						},
						Spec: &corev1.PodSpecArgs{
							Containers: corev1.ContainerArray{
								&corev1.ContainerArgs{
									Name:  pulumi.String(appName),
									Image: pulumi.String("paulbouwer/hello-kubernetes:1.8"),
									Ports: corev1.ContainerPortArray{
										&corev1.ContainerPortArgs{
											ContainerPort: pulumi.Int(8080),
										},
									},
									Env: corev1.EnvVarArray{
										&corev1.EnvVarArgs{
											Name:  pulumi.String("MESSAGE"),
											Value: pulumi.String("Hello K8s!"),
										},
									},
								},
							},
						},
					},
				},
			})
			if err != nil {
				return err
			}

			// Collect application statuses
			appStatuses = append(appStatuses, appSvc.Status)
		}

		// Create the ingress resource
		_, err = networkingv1.NewIngress(ctx, appBase+"-ingress", &networkingv1.IngressArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("hello-k8s-ingress"),
				Annotations: pulumi.StringMap{
					"kubernetes.io/ingress.class": pulumi.String("nginx"),
				},
				Namespace: ns.Metadata.Name().Elem(),
			},
			Spec: &networkingv1.IngressSpecArgs{
				Rules: networkingv1.IngressRuleArray{
					&networkingv1.IngressRuleArgs{
						Host: pulumi.String("myservicea.foo.org"),
						Http: &networkingv1.HTTPIngressRuleValueArgs{
							Paths: networkingv1.HTTPIngressPathArray{
								&networkingv1.HTTPIngressPathArgs{
									PathType: pulumi.String("Prefix"),
									Path:     pulumi.String("/"),
									Backend: &networkingv1.IngressBackendArgs{
										Service: &networkingv1.IngressServiceBackendArgs{
											Name: pulumi.String(appNames[0]),
											Port: &networkingv1.ServiceBackendPortArgs{
												Number: pulumi.Int(80),
											},
										},
									},
								},
							},
						},
					},
					&networkingv1.IngressRuleArgs{
						Host: pulumi.String("myserviceb.foo.org"),
						Http: &networkingv1.HTTPIngressRuleValueArgs{
							Paths: networkingv1.HTTPIngressPathArray{
								&networkingv1.HTTPIngressPathArgs{
									PathType: pulumi.String("Prefix"),
									Path:     pulumi.String("/"),
									Backend: &networkingv1.IngressBackendArgs{
										Service: &networkingv1.IngressServiceBackendArgs{
											Name: pulumi.String(appNames[1]),
											Port: &networkingv1.ServiceBackendPortArgs{
												Number: pulumi.Int(80),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}

		// Export the statuses of the apps and the ingress controller
		ctx.Export("appStatuses", appStatuses)
		ctx.Export("controllerStatus", ctrl.Status)

		return nil
	})
}
