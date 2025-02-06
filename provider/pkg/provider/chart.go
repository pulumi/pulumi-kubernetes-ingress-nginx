// Copyright 2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	helmbase "github.com/pulumi/pulumi-go-helmbase"
	king "github.com/pulumi/pulumi-kubernetes-ingress-nginx/sdk/go/kubernetes-ingress-nginx"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IngressController provisions a NGINX reverse proxy and load balancer for Kubernetes.
type IngressController struct {
	pulumi.ResourceState
	Status helmv3.ReleaseStatusOutput `pulumi:"status" pschema:"out"`
}

func (c *IngressController) SetOutputs(out helmv3.ReleaseStatusOutput) { c.Status = out }
func (c *IngressController) Type() string                              { return ComponentName }
func (c *IngressController) DefaultChartName() string                  { return "ingress-nginx" }
func (c *IngressController) DefaultRepoURL() string {
	return "https://kubernetes.github.io/ingress-nginx"
}

// IngressControllerArgs contains the set of arguments for creating a IngressController component resource.
type IngressControllerArgs struct {
	// Overrides for generated resource names.
	NameOverride *string `pulumi:"nameOverride"`
	// Overrides for generated resource names.
	FullnameOverride *string     `pulumi:"fullnameOverride"`
	Controller       *Controller `pulumi:"controller"`
	// Rollback limit.
	RevisionHistoryLimit *int `pulumi:"revisionHistoryLimit"`
	// Default 404 backend.
	DefaultBackend *ControllerDefaultBackend `pulumi:"defaultBackend"`
	// Enable RBAC as per
	// https://github.com/kubernetes/ingress-nginx/blob/main/docs/deploy/rbac.md and
	// https://github.com/kubernetes/ingress-nginx/issues/266
	RBAC *ControllerRBAC `pulumi:"rbac"`
	// If true, create & use Pod Security Policy resources
	// https://kubernetes.io/docs/concepts/policy/pod-security-policy/
	PodSecurityPolicy *ControllerPodSecurityPolicy `pulumi:"podSecurityPolicy"`
	ServiceAccount    *ControllerServiceAccount    `pulumi:"serviceAccount"`
	// Optional array of imagePullSecrets containing private registry credentials
	// Ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/.
	ImagePullSecrets *[]corev1.LocalObjectReference `pulumi:"imagePullSecrets" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:LocalObjectReference"`
	// TCP service key:value pairs
	// Ref: https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/exposing-tcp-udp-services.md.
	TCP *map[string]interface{} `pulumi:"tcp"`
	// UDP service key:value pairs
	// Ref: https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/exposing-tcp-udp-services.md.
	UDP *map[string]interface{} `pulumi:"udp"`
	// A base64ed Diffie-Hellman parameter.
	// This can be generated with: openssl dhparam 4096 2> /dev/null | base64
	// Ref: https://github.com/kubernetes/ingress-nginx/tree/main/docs/examples/customization/ssl-dh-param.
	DHParam *string `pulumi:"dhParam"`

	// HelmOptions is an escape hatch that lets the end user control any aspect of the
	// Helm deployment. This exposes the entirety of the underlying Helm Release component args.
	HelmOptions *helmbase.ReleaseType `pulumi:"helmOptions" pschema:"ref=#/types/chart-ingress-nginx:index:Release" json:"-"`
}

func (args *IngressControllerArgs) R() **helmbase.ReleaseType { return &args.HelmOptions }

type Controller struct {
	Name  *string                      `pulumi:"name"`
	Image king.ControllerImagePtrInput `pulumi:"image"`
	// Use an existing PSP instead of creating one.
	ExistingPsp *string `pulumi:"existingPsp"`
	// Configures the controller container name.
	ContainerName *string `pulumi:"containerName"`
	// Configures the ports the nginx-controller listens on.
	ContainerPort king.ControllerPortPtrInput `pulumi:"containerPort"`
	// Will add custom configuration options to Nginx
	// https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/.
	Config *map[string]interface{} `pulumi:"config"`
	// Annotations to be added to the controller config configuration configmap.
	ConfigAnnotations pulumi.StringMapInput `pulumi:"configAnnotations"`
	// Will add custom headers before sending traffic to backends according to
	// https://github.com/kubernetes/ingress-nginx/tree/main/docs/examples/customization/custom-headers.
	ProxySetHeaders *map[string]interface{} `pulumi:"proxySetHeaders"`
	// Will add custom headers before sending response traffic to the client according to:
	// https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#add-headers.
	AddHeaders *map[string]interface{} `pulumi:"addHeaders"`
	// Optionally customize the pod dnsConfig.
	DnsConfig *map[string]interface{} `pulumi:"dnsConfig"`
	// Optionally customize the pod hostname.
	Hostname *map[string]interface{} `pulumi:"hostname"`
	// Optionally change this to ClusterFirstWithHostNet in case you have 'hostNetwork: true'.
	// By default, while using host network, name resolution uses the host's DNS. If you wish nginx-controller
	// to keep resolving names inside the k8s network, use ClusterFirstWithHostNet.
	DnsPolicy *string `pulumi:"dnsPolicy"`
	// Bare-metal considerations via the host network https://kubernetes.github.io/ingress-nginx/deploy/baremetal/#via-the-host-network
	// Ingress status was blank because there is no Service exposing the NGINX Ingress controller in a
	// configuration using the host network, the default --publish-service flag used in standard cloud setups does not apply.
	ReportNodeInternalIp *bool `pulumi:"reportNodeInternalIp"`
	// Process Ingress objects without ingressClass annotation/ingressClassName field.
	// Overrides value for --watch-ingress-without-class flag of the controller binary.
	// Defaults to false.
	WatchIngressWithoutClass *bool `pulumi:"watchIngressWithoutClass"`
	// Process IngressClass per name (additionally as per spec.controller).
	IngressClassByName *bool `pulumi:"ingressClassByName"`
	// This configuration defines if Ingress Controller should allow users to set
	// their own *-snippet annotations, otherwise this is forbidden / dropped when users add those annotations.
	// Global snippets in ConfigMap are still respected.
	AllowSnippetAnnotations *bool `pulumi:"allowSnippetAnnotations"`
	// Required for use with CNI based kubernetes installations (such as ones set up by kubeadm),
	// since CNI and hostport don't mix yet. Can be deprecated once https://github.com/kubernetes/kubernetes/issues/23920
	// is merged.
	HostNetwork *bool `pulumi:"hostNetwork"`
	// Use host ports 80 and 443. Disabled by default.
	HostPort *ControllerHostPort `pulumi:"hostPort"`
	// Election ID to use for status update.
	ElectionID *string `pulumi:"electionID"`
	// This section refers to the creation of the IngressClass resource.
	// IngressClass resources are supported since k8s >= 1.18 and required since k8s >= 1.19
	IngressClassResource *ControllerIngressClassResource `pulumi:"ingressClassResource"`
	// labels to add to the pod container metadata.
	PodLabels pulumi.StringMapInput `pulumi:"podLabels"`
	// Security Context policies for controller pods.
	PodSecurityContext *corev1.PodSecurityContext `pulumi:"podSecurityContext" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:PodSecurityContext"`
	// See https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/ for
	// notes on enabling and using sysctls.
	Sysctls *map[string]interface{} `pulumi:"sysctls"`
	// Allows customization of the source of the IP address or FQDN to report
	// in the ingress status field. By default, it reads the information provided
	// by the service. If disable, the status field reports the IP address of the
	// node or nodes where an ingress controller pod is running.
	PublishService *ControllerPublishService `pulumi:"publishService"`
	// Limit the scope of the controller.
	Scope *ControllerScope `pulumi:"scope"`
	// Allows customization of the configmap / nginx-configmap namespace.
	ConfigMapNamespace pulumi.StringPtrInput `pulumi:"configMapNamespace"`
	// Allows customization of the tcp-services-configmap.
	Tcp *ControllerTcp `pulumi:"tcp"`
	Udp *ControllerUdp `pulumi:"udp"`
	// Maxmind license key to download GeoLite2 Databases
	// https://blog.maxmind.com/2019/12/18/significant-changes-to-accessing-and-using-geolite2-databases.
	MaxmindLicenseKey *string `pulumi:"maxmindLicenseKey"`
	// Additional command line arguments to pass to nginx-ingress-controller
	// E.g. to specify the default SSL certificate you can use `default-ssl-certificate: "<namespace>/<secret_name>"`.
	ExtraArgs *map[string]interface{} `pulumi:"extraArgs"`
	// Additional environment variables to set.
	ExtraEnvs *[]corev1.EnvVar `pulumi:"extraEnvs" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:EnvVar"`
	// DaemonSet or Deployment.
	Kind *string `pulumi:"kind"`
	// Annotations to be added to the controller Deployment or DaemonSet.
	Annotations pulumi.StringMapInput `pulumi:"annotations"`
	// Labels to be added to the controller Deployment or DaemonSet.
	Labels pulumi.StringMapInput `pulumi:"labels"`
	// The update strategy to apply to the Deployment or DaemonSet.
	UpdateStrategy *ControllerUpdateStrategy `pulumi:"updateStrategy"`
	// minReadySeconds to avoid killing pods before we are ready.
	MinReadySeconds *int `pulumi:"minReadySeconds"`
	// Node tolerations for server scheduling to nodes with taints
	// Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.
	Tolerations *[]corev1.Toleration `pulumi:"tolerations" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Toleration"`
	// Affinity and anti-affinity
	// Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity.
	Affinity *corev1.Affinity `pulumi:"affinity" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Affinity"`
	// Topology spread constraints rely on node labels to identify the topology domain(s) that each Node is in.
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/.
	TopologySpreadConstraints *[]corev1.TopologySpreadConstraint `pulumi:"topologySpreadConstraints" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:TopologySpreadConstraint"`
	// How long to wait for the drain of connections.
	TerminateGracePeriodSeconds *int `pulumi:"terminateGracePeriodSeconds"`
	// Node labels for controller pod assignment
	// Ref: https://kubernetes.io/docs/user-guide/node-selection/.
	NodeSelector *map[string]string `pulumi:"nodeSelector"`
	// Startup probe values
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
	StartupProbe *corev1.Probe `pulumi:"startupProbe" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Probe"`
	// Liveness probe values
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
	LivenessProbe *corev1.Probe `pulumi:"livenessProbe" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Probe"`
	// Readiness probe values
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
	ReadinessProbe *corev1.Probe `pulumi:"readinessProbe" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Probe"`
	// Path of the health check endpoint. All requests received on the port defined by
	// the healthz-port parameter are forwarded internally to this path.
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// Address to bind the health check endpoint.
	// It is better to set this option to the internal node address
	// if the ingress nginx controller is running in the hostNetwork: true mode.
	HealthCheckHost *string `pulumi:"heathCheckHost"`
	// Annotations to be added to controller pods.
	PodAnnotations pulumi.StringMapInput `pulumi:"podAnnotations"`
	ReplicaCount   *int                  `pulumi:"replicaCount"`
	MinAvailable   *int                  `pulumi:"minAvailable"`
	// Define requests resources to avoid probe issues due to CPU utilization in busy nodes
	// ref: https://github.com/kubernetes/ingress-nginx/issues/4735#issuecomment-551204903
	// Ideally, there should be no limits.
	// https://engineering.indeedblog.com/blog/2019/12/cpu-throttling-regression-fix/
	Resources *corev1.ResourceRequirements `pulumi:"resources" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:ResourceRequirements"`
	// Mutually exclusive with keda autoscaling.
	Autoscaling king.AutoscalingPtrInput `pulumi:"autoscaling"`
	// Custom or additional autoscaling metrics
	// ref: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-custom-metrics
	AutoscalingTemplate *king.AutoscalingTemplateArrayInput `pulumi:"autoscalingTemplate"`
	// Mutually exclusive with hpa autoscaling.
	Keds *Keda `pulumi:"keda"`
	// Enable mimalloc as a drop-in replacement for malloc.
	// ref: https://github.com/microsoft/mimalloc.
	EnableMimalloc *bool `pulumi:"enableMimalloc"`
	// Override NGINX template.
	CustomTemplate king.ControllerCustomTemplatePtrInput `pulumi:"customTemplate"`
	Service        king.ControllerServicePtrInput        `pulumi:"service"`
	// Additional containers to be added to the controller pod.
	// See https://github.com/lemonldap-ng-controller/lemonldap-ng-controller as example.
	ExtraContainers *[]corev1.Container `pulumi:"extraContainers" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Container"`
	// Additional volumeMounts to the controller main container.
	//  - name: copy-portal-skins
	//    mountPath: /var/lib/lemonldap-ng/portal/skins
	ExtraVolumeMounts *[]corev1.VolumeMount `pulumi:"extraVolumeMounts" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:VolumeMount"`
	// Additional volumes to the controller pod.
	//  - name: copy-portal-skins
	//    emptyDir: {}
	ExtraVolumes *[]corev1.Volume `pulumi:"extraVolumes" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Volume"`
	// Containers, which are run before the app containers are started.
	// - name: init-myservice
	//   image: busybox
	//   command: ['sh', '-c', 'until nslookup myservice; do echo waiting for myservice; sleep 2; done;']
	ExtraInitContainers *[]corev1.Container                     `pulumi:"extraInitContainers" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Container"`
	AdmissionWebhooks   king.ContollerAdmissionWebhooksPtrInput `pulumi:"admissionWebhooks"`
	Metrics             king.ControllerMetricsPtrInput          `pulumi:"metrics"`
	// Improve connection draining when ingress controller pod is deleted using a lifecycle hook:
	// With this new hook, we increased the default terminationGracePeriodSeconds from 30 seconds
	// to 300, allowing the draining of connections up to five minutes.
	// If the active connections end before that, the pod will terminate gracefully at that time.
	// To effectively take advantage of this feature, the Configmap feature
	// worker-shutdown-timeout new value is 240s instead of 10s.
	Lifecycle         *corev1.Lifecycle `pulumi:"lifecycle" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Lifecycle"`
	PriorityClassName *string           `pulumi:"priorityClassName"`
}

type ControllerImage struct {
	Registry *string `pulumi:"registry"`
	Image    *string `pulumi:"image"`
	// for backwards compatibility consider setting the full image url via the repository value below
	// use *either* current default registry/image or repository format or installing will fail.
	Repository               *string `pulumi:"repository"`
	Tag                      *string `pulumi:"tag"`
	Digest                   *string `pulumi:"digest"`
	PullPolicy               *string `pulumi:"pullPolicy"`
	RunAsUser                *string `pulumi:"runAsUser"`
	RunAsNonRoot             *bool   `pulumi:"runAsNonRoot"`
	ReadOnlyRootFilesystem   *bool   `pulumi:"readOnlyRootFilesystem"`
	AllowPrivilegeEscalation *bool   `pulumi:"allowPrivilegeEscalation"`
}

type ControllerPort struct {
	Http  *int `pulumi:"http"`
	Https *int `pulumi:"https"`
}

type ControllerHostPort struct {
	Enabled *bool                    `pulumi:"enabled"`
	Ports   *ControllerHostPortPorts `pulumi:"ports"`
}

type ControllerHostPortPorts struct {
	Http  *int `pulumi:"http"`
	Https *int `pulumi:"https"`
}

type ControllerServiceAccount struct {
	Create                       *bool   `pulumi:"create"`
	Name                         *string `pulumi:"name"`
	AutomountServiceAccountToken *bool   `pulumi:"automountServiceAccountToken"`
}

type ControllerIngressClassResource struct {
	Name            *string `pulumi:"name"`
	Enabled         *bool   `pulumi:"enabled"`
	Default         *bool   `pulumi:"default"`
	ControllerValue *string `pulumi:"controllerValue"`
	// Parameters is a link to a custom resource containing additional
	// configuration for the controller. This is optional if the controller
	// does not require extra parameters.
	Parameters *map[string]interface{} `pulumi:"parameters"`
}

type ControllerPublishService struct {
	Enabled *bool `pulumi:"enabled"`
	// Allows overriding of the publish service to bind to. Must be <namespace>/<service_name>.
	PathOverride *string `pulumi:"pathOverride"`
}

type ControllerScope struct {
	Enabled   *bool   `pulumi:"enabled"`
	Namespace *string `pulumi:"namespace"`
}

type ControllerTcp struct {
	ConfigMapNamespace *string `pulumi:"configMapNamespace"`
	// Annotations to be added to the tcp config configmap.
	Annotations pulumi.StringMapInput `pulumi:"annotations"`
}

type ControllerUdp struct {
	ConfigMapNamespace *string `pulumi:"configMapNamespace"`
	// Annotations to be added to the udp config configmap.
	Annotations pulumi.StringMapInput `pulumi:"annotations"`
}

type ControllerUpdateStrategy struct {
	RollingUpdate *ControllerRollingUpdate `pulumi:"rollingUpdate"`
	Type          *string                  `pulumi:"type"`
}

type ControllerRollingUpdate struct {
	MaxUnavailable *int `pulumi:"maxUnavailable"`
}

type Autoscaling struct {
	Annotations                        pulumi.StringMapInput `pulumi:"annotations"`
	Enabled                            *bool                 `pulumi:"enabled"`
	MinReplicas                        *int                  `pulumi:"minReplicas"`
	MaxReplicas                        *int                  `pulumi:"maxReplicas"`
	TargetCPUUtilizationPercentation   *int                  `pulumi:"targetCPUUtilizationPercentage"`
	TargetMemoryUtilizationPercentatge *int                  `pulumi:"targetMemoryUtilizationPercentage"`
	Behavior                           *AutoscalingBehavior  `pulumi:"controllerAutoscalingBehavior"`
}

type AutoscalingBehavior struct {
	ScaleDown *AutoscalingBehaviorScaling `pulumi:"scaleDown"`
	ScaleUp   *AutoscalingBehaviorScaling `pulumi:"scaleUp"`
}

type AutoscalingBehaviorScaling struct {
	StabilizationWindowSeconds *int                                `pulumi:"stabilizationWindowSeconds"`
	Policies                   *[]AutoscalingBehaviorScalingPolicy `pulumi:"policies"`
}

type AutoscalingBehaviorScalingPolicy struct {
	Type          *string `pulumi:"type"`
	Value         *int    `pulumi:"value"`
	PeriodSeconds *int    `pulumi:"periodSeconds"`
}

type AutoscalingTemplate struct {
	Type *string                  `pulumi:"type"`
	Pods *AutoscalingTemplatePods `pulumi:"pods"`
}

type AutoscalingTemplatePods struct {
	Metric *AutoscalingTemplatePodsMetric `pulumi:"metric"`
	Target *AutoscalingTemplatePodsTarget `pulumi:"target"`
}

type AutoscalingTemplatePodsMetric struct {
	Name *string `pulumi:"name"`
}

type AutoscalingTemplatePodsTarget struct {
	Type         *string `pulumi:"type"`
	AverageValue *string `pulumi:"averageValue"`
}

type Keda struct {
	// apiVersion changes with keda 1.x vs 2.x:
	// 2.x = keda.sh/v1alpha1,
	// 1.x = keda.k8s.io/v1alpha1.
	APIVersion                    *string              `pulumi:"apiVersion"`
	Enabled                       *bool                `pulumi:"enabled"`
	MinReplicas                   *int                 `pulumi:"minReplicas"`
	MaxReplicas                   *int                 `pulumi:"maxReplicas"`
	PollingInterval               *int                 `pulumi:"pollingInterval"`
	CooldownPeriod                *int                 `pulumi:"cooldownPeriod"`
	RestoreToOriginalReplicaCount *bool                `pulumi:"restoreToOriginalReplicaCount"`
	ScaledObject                  *KedaScaledObject    `pulumi:"scaledObject"`
	Triggers                      *[]KedaTrigger       `pulumi:"triggers"`
	Behavior                      *AutoscalingBehavior `pulumi:"behavior"`
}

type KedaScaledObject struct {
	// Custom annotations for ScaledObject resource.
	Annotations pulumi.StringMapInput `pulumi:"annotations"`
}

type KedaTrigger struct {
	Type     *string                 `pulumi:"type"`
	Metadata *map[string]interface{} `pulumi:"metadata"`
}

type ControllerCustomTemplate struct {
	ConfigMapName *string `pulumi:"configMapName"`
	ConfigMapKey  *string `pulumi:"configMapKey"`
}

type ControllerService struct {
	Enabled     *bool                 `pulumi:"enabled"`
	Annotations pulumi.StringMapInput `pulumi:"annotations"`
	Labels      pulumi.StringMapInput `pulumi:"labels"`
	ClusterIP   *string               `pulumi:"clusterIP"`
	// List of IP addresses at which the controller services are available
	// Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
	ExternalIPs              *[]string             `pulumi:"externalIPs"`
	LoadBalancerIP           pulumi.StringPtrInput `pulumi:"loadBalancerIPs"`
	LoadBalancerSourceRanges *[]string             `pulumi:"loadBalancerSourceRanges"`
	EnableHttp               *bool                 `pulumi:"enableHttp"`
	EnableHttps              *bool                 `pulumi:"enableHttps"`
	// Set external traffic policy to: "Local" to preserve source IP on
	// providers supporting it.
	// Ref: https://kubernetes.io/docs/tutorials/services/source-ip/#source-ip-for-services-with-typeloadbalancer
	ExternalTrafficPolicy *string `pulumi:"externalTrafficPolicy"`
	// Must be either "None" or "ClientIP" if set. Kubernetes will default to "None".
	// Ref: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	SessionAffinity *string `pulumi:"sessionAffinity"`
	// specifies the health check node port (numeric port number) for the service. If healthCheckNodePort isn’t specified,
	// the service controller allocates a port from your cluster’s NodePort range.
	// Ref: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#preserving-the-client-source-ip
	HealthCheckNodePort *int                        `pulumi:"healthCheckNodePort"`
	Ports               *ControllerPort             `pulumi:"ports"`
	TargetPorts         *ControllerPort             `pulumi:"targetPorts"`
	Type                *string                     `pulumi:"type"`
	NodePorts           *ControllerServiceNodePorts `pulumi:"nodePorts"`
	// Enables an additional internal load balancer (besides the external one).
	// Annotations are mandatory for the load balancer to come up. Varies with the cloud service.
	Internal *ControllerServiceInternal `pulumi:"internal"`
}

type ControllerServiceNodePorts struct {
	Http  *string                 `pulumi:"http"`
	Https *string                 `pulumi:"https"`
	Tcp   *map[string]interface{} `pulumi:"tcp"`
	Udp   *map[string]interface{} `pulumi:"udp"`
}

type ControllerServiceInternal struct {
	Enabled        *bool                 `pulumi:"enabled"`
	Annotations    pulumi.StringMapInput `pulumi:"annotations"`
	Labels         pulumi.StringMapInput `pulumi:"labels"`
	LoadBalancerIP *string               `pulumi:"loadBalancerIPs"`
	// Restrict access For LoadBalancer service. Defaults to 0.0.0.0/0.
	LoadBalancerSourceRanges *[]string `pulumi:"loadBalancerSourceRanges"`
	// Set external traffic policy to: "Local" to preserve source IP on
	// providers supporting it.
	// Ref: https://kubernetes.io/docs/tutorials/services/source-ip/#source-ip-for-services-with-typeloadbalancer
	ExternalTrafficPolicy *string `pulumi:"externalTrafficPolicy"`
}

type ContollerAdmissionWebhooks struct {
	Enabled           *bool                   `pulumi:"enabled"`
	Annotations       pulumi.StringMapInput   `pulumi:"annotations"`
	FailurePolicy     *string                 `pulumi:"failurePolicy"`
	TimeoutSeconds    *int                    `pulumi:"timeoutSeconds"`
	Port              *int                    `pulumi:"port"`
	Certificate       *string                 `pulumi:"certificate"`
	Key               *string                 `pulumi:"key"`
	NamespaceSelector *map[string]interface{} `pulumi:"namespaceSelector"`
	ObjectSelector    *map[string]interface{} `pulumi:"objectSelector"`
	// Use an existing PSP instead of creating one.
	ExistingPsp     *string                                      `pulumi:"existingPsp"`
	Service         *ControllerAdmissionWebhooksService          `pulumi:"service"`
	CreateSecretJob *ControllerAdmissionWebhooksCreateSecretJob  `pulumi:"createSecretJob"`
	PatchWebhookJob *ControllerAdmissionWebhooksPatchWebhbookJob `pulumi:"patchWebhookJob"`
	Patch           *ControllerAdmissionWebhooksPatch            `pulumi:"patch"`
}

type ControllerAdmissionWebhooksService struct {
	Annotations              pulumi.StringMapInput `pulumi:"annotations"`
	ClusterIP                *string               `pulumi:"clusterIP"`
	ExternalIPs              *[]string             `pulumi:"externalIPs"`
	LoadBalancerIP           *string               `pulumi:"loadBalancerIPs"`
	LoadBalancerSourceRanges *[]string             `pulumi:"loadBalancerSourceRanges"`
	ServicePort              *int                  `pulumi:"servicePort"`
	Type                     *string               `pulumi:"type"`
}

type ControllerAdmissionWebhooksCreateSecretJob struct {
	Resources *corev1.ResourceRequirements `pulumi:"resources" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:ResourceRequirements"`
}

type ControllerAdmissionWebhooksPatchWebhbookJob struct {
	Resources *corev1.ResourceRequirements `pulumi:"resources" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:ResourceRequirements"`
}

type ControllerAdmissionWebhooksPatch struct {
	Enabled *bool            `pulumi:"enabled"`
	Image   *ControllerImage `pulumi:"image"`
	// Provide a priority class name to the webhook patching job.
	PriorityClassName *string               `pulumi:"priorityClassName"`
	PodAnnotations    pulumi.StringMapInput `pulumi:"podAnnotations"`
	NodeSelector      *map[string]string    `pulumi:"nodeSelector"`
	Tolerations       *[]corev1.Toleration  `pulumi:"tolerations" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Toleration"`
	RunAsUser         *int                  `pulumi:"runAsUser"`
}

type ControllerMetrics struct {
	// if this port is changed, change healthz-port: in extraArgs: accordingly.
	Port           *int                              `pulumi:"port"`
	Enabled        *bool                             `pulumi:"enabled"`
	Service        *ControllerMetricsService         `pulumi:"service"`
	ServiceMonitor *ControllerMetricsServiceMonitor  `pulumi:"serviceMonitor"`
	PrometheusRule *ControllerMetricsPrometheusRules `pulumi:"prometheusRule"`
}

type ControllerMetricsService struct {
	Annotations              pulumi.StringMapInput `pulumi:"annotations"`
	ClusterIP                *string               `pulumi:"clusterIP"`
	ExternalIPs              *[]string             `pulumi:"externalIPs"`
	LoadBalancerIP           *string               `pulumi:"loadBalancerIPs"`
	LoadBalancerSourceRanges *[]string             `pulumi:"loadBalancerSourceRanges"`
	ServicePort              *int                  `pulumi:"servicePort"`
	Type                     *string               `pulumi:"type"`
	ExternalTrafficPolicy    *string               `pulumi:"externalTrafficPolicy"`
	NodePort                 *string               `pulumi:"nodePort"`
}

type ControllerMetricsServiceMonitor struct {
	Enabled          *bool                 `pulumi:"enabled"`
	AdditionalLabels pulumi.StringMapInput `pulumi:"additionalLabels"`
	// The label to use to retrieve the job name from.
	JobLabel          *string               `pulumi:"jobLabel"`
	Namespace         *string               `pulumi:"namespace"`
	NamespaceSelector pulumi.StringMapInput `pulumi:"namespaceSelector"`
	ScrapeInterval    *string               `pulumi:"scrapeInterval"`
	HonorLabels       *bool                 `pulumi:"honorLabels"`
	TargetLabels      *[]string             `pulumi:"targetLabels"`
	MetricRelabelings *[]string             `pulumi:"metricRelabelings"`
}

type ControllerMetricsPrometheusRules struct {
	Enabled          *bool                 `pulumi:"enabled"`
	AdditionalLabels pulumi.StringMapInput `pulumi:"additionalLabels"`
	Namespace        *string               `pulumi:"namespace"`
	Rules            *[]interface{}        `pulumi:"rules"`
}

type ControllerDefaultBackend struct {
	Enabled *bool            `pulumi:"enabled"`
	Name    *string          `pulumi:"name"`
	Image   *ControllerImage `pulumi:"image"`
	// Use an existing PSP instead of creating one.
	ExistingPsp    *string                   `pulumi:"existingPsp"`
	ExtraArgs      *map[string]interface{}   `pulumi:"extraArgs"`
	ServiceAccount *ControllerServiceAccount `pulumi:"serviceAccount"`
	ExtraEnvs      *[]corev1.EnvVar          `pulumi:"extraEnvs" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:EnvVar"`
	Port           *int                      `pulumi:"port"`
	// Liveness probe values for default backend.
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
	LivenessProbe *corev1.Probe `pulumi:"livenessProbe" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Probe"`
	// Readiness probe values for default backend.
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
	ReadinessProbe *corev1.Probe `pulumi:"readinessProbe" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Probe"`
	// Node tolerations for server scheduling to nodes with taints.
	// Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	Tolerations *[]corev1.Toleration `pulumi:"tolerations" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Toleration"`
	Affinity    *corev1.Affinity     `pulumi:"affinity" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Affinity"`
	// Security Context policies for controller pods.
	// See https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/ for
	// notes on enabling and using sysctls.
	PodSecurityContext *corev1.PodSecurityContext `pulumi:"podSecurityContext" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:PodSecurityContext"`
	// labels to add to the pod container metadata
	PodLabels pulumi.StringMapInput `pulumi:"podLabels"`
	// Node labels for default backend pod assignment
	// Ref: https://kubernetes.io/docs/user-guide/node-selection/.
	NodeSelector *map[string]string `pulumi:"nodeSelector"`
	// Annotations to be added to default backend pods.
	PodAnnotations pulumi.StringMapInput        `pulumi:"podAnnotations"`
	ReplicaCount   *int                         `pulumi:"replicaCount"`
	MinAvailable   *int                         `pulumi:"minAvailable"`
	Resources      *corev1.ResourceRequirements `pulumi:"resources" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:ResourceRequirements"`
	// Additional volumeMounts to the default backend container.
	//  - name: copy-portal-skins
	//    mountPath: /var/lib/lemonldap-ng/portal/skins
	ExtraVolumeMounts *[]corev1.VolumeMount `pulumi:"extraVolumeMounts" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:VolumeMount"`
	// Additional volumes to the default backend pod.
	//  - name: copy-portal-skins
	//    emptyDir: {}
	ExtraVolumes      *[]corev1.Volume                 `pulumi:"extraVolumes" pschema:"ref=/kubernetes/v4.21.0/schema.json#/types/kubernetes:core/v1:Volume"`
	Autoscaling       *Autoscaling                     `pulumi:"autoscaling"`
	Service           *ControllerDefaultBackendService `pulumi:"service"`
	PriorityClassName *string                          `pulumi:"priorityClassName"`
}

type ControllerDefaultBackendService struct {
	Annoations pulumi.StringMapInput `pulumi:"annotations"`
	ClusterIP  *string               `pulumi:"clusterIP"`
	// List of IP addresses at which the default backend service is available.
	// Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
	ExternalIPs              *[]string `pulumi:"externalIPs"`
	LoadBalancerIP           *string   `pulumi:"loadBalancerIP"`
	LoadBalancerSourceRanges *[]string `pulumi:"loadBalancerSourceRanges"`
	ServicePort              *int      `pulumi:"servicePort"`
	Type                     *string   `pulumi:"type"`
}

type ControllerPodSecurityPolicy struct {
	Enabled *bool `pulumi:"enabled"`
}

type ControllerRBAC struct {
	Create *bool `pulumi:"create"`
	Scope  *bool `pulumi:"scope"`
}
