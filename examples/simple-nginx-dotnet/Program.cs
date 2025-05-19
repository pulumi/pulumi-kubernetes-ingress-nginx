using Pulumi;
using Pulumi.Kubernetes.Core.V1;
using Pulumi.Kubernetes.Apps.V1;
using Pulumi.Kubernetes.Networking.V1;
using System.Collections.Generic;
using Pulumi.Kubernetes.Types.Inputs.Core.V1;
using Pulumi.Kubernetes.Types.Inputs.Apps.V1;
using Pulumi.Kubernetes.Types.Inputs.Meta.V1;
using Pulumi.Kubernetes.Types.Inputs.Networking.V1;
using Pulumi.KubernetesIngressNginx;
using Pulumi.KubernetesIngressNginx.Inputs;

return await Pulumi.Deployment.RunAsync(() =>
{
    // Create a sandbox namespace
    var ns = new Namespace("sandbox-ns");

    // Install the NGINX ingress controller
    var ctrl = new IngressController("myctrl", new IngressControllerArgs
    {
        Controller = new ControllerArgs
        {
            Config = new Dictionary<string, string>
            {
                { "hsts-max-age", "31536000" }
            },
            HostPort = new ControllerHostPortArgs
            {
                Enabled = true
            },
            PublishService = new ControllerPublishServiceArgs
            {
                Enabled = true
            },
            Service = new ControllerServiceArgs
            {
                Annotations = new Dictionary<string, string>
                {
                    { "pulumi.com/test-annotation1", "test-value1" },
                    { "pulumi.com/test-annotation2", "test-value2" }
                },
                Type = "NodePort",
                ExternalTrafficPolicy = "Local"
            }
        },
        HelmOptions = new ReleaseArgs
        {
            Namespace = ns.Metadata.Apply(m => m.Name)
        }
    });

    // Deploy two applications
    var apps = new List<Output<Service>>();
    var appBase = "hello-k8s";
    var appNames = new[] { $"{appBase}-first", $"{appBase}-second" };

    foreach (var appName in appNames)
    {
        var appSvc = new Service($"{appName}-svc", new ServiceArgs
        {
            Metadata = new ObjectMetaArgs
            {
                Name = appName,
                Namespace = ns.Metadata.Apply(m => m.Name)
            },
            Spec = new ServiceSpecArgs
            {
                Type = "ClusterIP",
                Ports = new ServicePortArgs
                {
                    Port = 80,
                    TargetPort = 8080
                },
                Selector = new Dictionary<string, string>
                {
                    { "app", appName }
                }
            }
        });

        var appDep = new Pulumi.Kubernetes.Apps.V1.Deployment($"{appName}-dep", new DeploymentArgs
        {
            Metadata = new ObjectMetaArgs
            {
                Name = appName,
                Namespace = ns.Metadata.Apply(m => m.Name)
            },
            Spec = new DeploymentSpecArgs
            {
                Replicas = 3,
                Selector = new LabelSelectorArgs
                {
                    MatchLabels = new Dictionary<string, string>
                    {
                        { "app", appName }
                    }
                },
                Template = new PodTemplateSpecArgs
                {
                    Metadata = new ObjectMetaArgs
                    {
                        Labels = new Dictionary<string, string>
                        {
                            { "app", appName }
                        }
                    },
                    Spec = new PodSpecArgs
                    {
                        Containers =
                        {
                            new ContainerArgs
                            {
                                Name = appName,
                                Image = "paulbouwer/hello-kubernetes:1.8",
                                Ports =
                                {
                                    new ContainerPortArgs
                                    {
                                        ContainerPortValue = 8080
                                    }
                                },
                                Env =
                                {
                                    new EnvVarArgs
                                    {
                                        Name = "MESSAGE",
                                        Value = "Hello K8s!"
                                    }
                                }
                            }
                        }
                    }
                }
            }
        });

        apps.Add(Output.Create(appSvc));
    }

    // Create the Ingress
    var appIngress = new Ingress($"{appBase}-ingress", new IngressArgs
    {
        Metadata = new ObjectMetaArgs
        {
            Name = "hello-k8s-ingress",
            Annotations = new Dictionary<string, string>
            {
                { "kubernetes.io/ingress.class", "nginx" }
            },
            Namespace = ns.Metadata.Apply(m => m.Name)
        },
        Spec = new IngressSpecArgs
        {
            Rules =
            {
                new IngressRuleArgs
                {
                    Host = "myservicea.foo.org",
                    Http = new HTTPIngressRuleValueArgs
                    {
                        Paths =
                        {
                            new HTTPIngressPathArgs
                            {
                                PathType = "Prefix",
                                Path = "/",
                                Backend = new IngressBackendArgs
                                {
                                    Service = new IngressServiceBackendArgs
                                    {
                                        Name = appNames[0],
                                        Port = new ServiceBackendPortArgs
                                        {
                                            Number = 80
                                        }
                                    }
                                }
                            }
                        }
                    }
                },
                new IngressRuleArgs
                {
                    Host = "myserviceb.foo.org",
                    Http = new HTTPIngressRuleValueArgs
                    {
                        Paths =
                        {
                            new HTTPIngressPathArgs
                            {
                                PathType = "Prefix",
                                Path = "/",
                                Backend = new IngressBackendArgs
                                {
                                    Service = new IngressServiceBackendArgs
                                    {
                                        Name = appNames[1],
                                        Port = new ServiceBackendPortArgs
                                        {
                                            Number = 80
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    });

    return new Dictionary<string, object?>
    {
        ["appStatuses"] = apps,
        ["namespace"] = ns.Metadata.Apply(m => m.Name)
    };
});

