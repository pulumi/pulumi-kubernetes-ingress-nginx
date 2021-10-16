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
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IngressController provisions a NGINX reverse proxy and load balancer for Kubernetes.
type IngressController struct {
	pulumi.ResourceState
	HelmRelease *helmv3.Release `pulumi:"helmRelease,out,ref=pulumi-kubernetes:helm/v3:Release"`
}

func (c *IngressController) SetOutputs(rel *helmv3.Release) { c.HelmRelease = rel }
func (c *IngressController) Type() string                   { return ComponentName }
func (c *IngressController) DefaultChartName() string       { return "ingress-nginx" }
func (c *IngressController) DefaultRepoURL() string {
	return "https://kubernetes.github.io/ingress-nginx"
}

// IngressControllerArgs contains the set of arguments for creating a IngressController component resource.
type IngressControllerArgs struct {
	// Overrides for generated resource names.
	NameOverride *string `pulumi:"nameOverride,omitempty"`
	// Overrides for generated resource names.
	FullnameOverride *string     `pulumi:"fullnameOverride,omitempty"`
	Controller       *Controller `pulumi:"controller" json:"controller,omitempty"`
	// Rollback limit.
	RevisionHistoryLimit *int `pulumi:"revisionHistoryLimit,omitempty"`
	// Default 404 backend.
	DefaultBackend *ControllerDefaultBackend `pulumi:"defaultBackend,omitempty"`
	// Enable RBAC as per
	// https://github.com/kubernetes/ingress-nginx/blob/main/docs/deploy/rbac.md and
	// https://github.com/kubernetes/ingress-nginx/issues/266
	RBAC *ControllerRBAC `pulumi:"rbac,omitempty"`
	// If true, create & use Pod Security Policy resources
	// https://kubernetes.io/docs/concepts/policy/pod-security-policy/
	PodSecurityPolicy *ControllerPodSecurityPolicy `pulumi:"podSecurityPolicy,omitempty"`
	ServiceAccount    *ControllerServiceAccount    `pulumi:"serviceAccount,omitempty"`
	// Optional array of imagePullSecrets containing private registry credentials
	// Ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/.
	ImagePullSecrets *[]corev1.LocalObjectReference `pulumi:"imagePullSecrets,ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:LocalObjectReference,omitempty"`
	// TCP service key:value pairs
	// Ref: https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/exposing-tcp-udp-services.md.
	TCP map[string]interface{} `pulumi:"tcp,omitempty"`
	// UDP service key:value pairs
	// Ref: https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/exposing-tcp-udp-services.md.
	UDP map[string]interface{} `pulumi:"udp,omitempty"`
	// A base64ed Diffie-Hellman parameter.
	// This can be generated with: openssl dhparam 4096 2> /dev/null | base64
	// Ref: https://github.com/kubernetes/ingress-nginx/tree/main/docs/examples/customization/ssl-dh-param.
	DHParam *string `pulumi:"dhParam,omitempty"`

	// HelmOptions is an escape hatch that lets the end user control any aspect of the
	// Helm deployment. This exposes the entirety of the underlying Helm Release component args.
	HelmOptions *HelmOptions `pulumi:"helmOptions" json:"-"`
}

// HelmOptions defines the underlying Helm deployment options that can be specified to override defaults.
type HelmOptions = helmv3.ReleaseType

func (args *IngressControllerArgs) R() **helmv3.ReleaseType { return &args.HelmOptions }

type Controller struct {
	Name  *string          `pulumi:"name,omitempty"`
	Image *ControllerImage `pulumi:"image,omitempty"`
	// Use an existing PSP instead of creating one.
	ExistingPsp *string `pulumi:"existingPsp,omitempty"`
	// Configures the controller container name.
	ContainerName *string `pulumi:"containerName,omitempty"`
	// Configures the ports the nginx-controller listens on.
	ContainerPort *ControllerPort `pulumi:"containerPort,omitempty"`
	// Will add custom configuration options to Nginx
	// https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/.
	Config *map[string]interface{} `pulumi:"config,omitempty"`
	// Annotations to be added to the controller config configuration configmap.
	ConfigAnnotations *map[string]interface{} `pulumi:"configAnnotations,omitempty"`
	// Will add custom headers before sending traffic to backends according to
	// https://github.com/kubernetes/ingress-nginx/tree/main/docs/examples/customization/custom-headers.
	ProxySetHeaders *map[string]interface{} `pulumi:"proxySetHeaders,omitempty"`
	// Will add custom headers before sending response traffic to the client according to:
	// https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#add-headers.
	AddHeaders *map[string]interface{} `pulumi:"addHeaders,omitempty"`
	// Optionally customize the pod dnsConfig.
	DnsConfig *map[string]interface{} `pulumi:"dnsConfig,omitempty"`
	// Optionally customize the pod hostname.
	Hostname *map[string]interface{} `pulumi:"hostname,omitempty"`
	// Optionally change this to ClusterFirstWithHostNet in case you have 'hostNetwork: true'.
	// By default, while using host network, name resolution uses the host's DNS. If you wish nginx-controller
	// to keep resolving names inside the k8s network, use ClusterFirstWithHostNet.
	DnsPolicy *string `pulumi:"dnsPolicy,omitempty"`
	// Bare-metal considerations via the host network https://kubernetes.github.io/ingress-nginx/deploy/baremetal/#via-the-host-network
	// Ingress status was blank because there is no Service exposing the NGINX Ingress controller in a
	// configuration using the host network, the default --publish-service flag used in standard cloud setups does not apply.
	ReportNodeInternalIp *bool `pulumi:"reportNodeInternalIp,omitempty"`
	// Process Ingress objects without ingressClass annotation/ingressClassName field.
	// Overrides value for --watch-ingress-without-class flag of the controller binary.
	// Defaults to false.
	WatchIngressWithoutClass *bool `pulumi:"watchIngressWithoutClass,omitempty"`
	// Process IngressClass per name (additionally as per spec.controller).
	IngressClassByName *bool `pulumi:"ingressClassByName,omitempty"`
	// This configuration defines if Ingress Controller should allow users to set
	// their own *-snippet annotations, otherwise this is forbidden / dropped when users add those annotations.
	// Global snippets in ConfigMap are still respected.
	AllowSnippetAnnotations *bool `pulumi:"allowSnippetAnnotations,omitempty"`
	// Required for use with CNI based kubernetes installations (such as ones set up by kubeadm),
	// since CNI and hostport don't mix yet. Can be deprecated once https://github.com/kubernetes/kubernetes/issues/23920
	// is merged.
	HostNetwork *bool `pulumi:"hostNetwork,omitempty"`
	// Use host ports 80 and 443. Disabled by default.
	HostPort *ControllerHostPort `pulumi:"hostPort,omitempty"`
	// Election ID to use for status update.
	ElectionID *string `pulumi:"electionID,omitempty"`
	// This section refers to the creation of the IngressClass resource.
	// IngressClass resources are supported since k8s >= 1.18 and required since k8s >= 1.19
	IngressClassResource *ControllerIngressClassResource `pulumi:"ingressClassResource,omitempty"`
	// labels to add to the pod container metadata.
	PodLabels *map[string]interface{} `pulumi:"podLabels,omitempty"`
	// Security Context policies for controller pods.
	PodSecurityContext *corev1.PodSecurityContext `pulumi:"podSecurityContext,ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:PodSecurityContext,omitempty"`
	// See https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/ for
	// notes on enabling and using sysctls.
	Sysctls *map[string]interface{} `pulumi:"sysctls,omitempty"`
	// Allows customization of the source of the IP address or FQDN to report
	// in the ingress status field. By default, it reads the information provided
	// by the service. If disable, the status field reports the IP address of the
	// node or nodes where an ingress controller pod is running.
	PublishService *ControllerPublishService `pulumi:"publishService,omitempty"`
	// Limit the scope of the controller.
	Scope *ControllerScope `pulumi:"scope,omitempty"`
	// Allows customization of the configmap / nginx-configmap namespace.
	ConfigMapNamespace *string `pulumi:"configMapNamespace,omitempty"`
	// Allows customization of the tcp-services-configmap.
	Tcp *ControllerTcp `pulumi:"tcp,omitempty"`
	Udp *ControllerUdp `pulumi:"udp,omitempty"`
	// Maxmind license key to download GeoLite2 Databases
	// https://blog.maxmind.com/2019/12/18/significant-changes-to-accessing-and-using-geolite2-databases.
	MaxmindLicenseKey *string `pulumi:"maxmindLicenseKey,omitempty"`
	// Additional command line arguments to pass to nginx-ingress-controller
	// E.g. to specify the default SSL certificate you can use `default-ssl-certificate: "<namespace>/<secret_name>"`.
	ExtraArgs *map[string]interface{} `pulumi:"extraArgs,omitempty"`
	// Additional environment variables to set.
	ExtraEnvs []*corev1.EnvVar `pulumi:"extraEnvs,omitempty"`
	// DaemonSet or Deployment.
	Kind *string `pulumi:"kind,omitempty"`
	// Annotations to be added to the controller Deployment or DaemonSet.
	Annotations *map[string]string `pulumi:"annotations,omitempty"`
	// Labels to be added to the controller Deployment or DaemonSet.
	Labels *map[string]string `pulumi:"labels,omitempty"`
	// The update strategy to apply to the Deployment or DaemonSet.
	UpdateStrategy *ControllerUpdateStrategy `pulumi:"updateStrategy,omitempty"`
	// minReadySeconds to avoid killing pods before we are ready.
	MinReadySeconds *int `pulumi:"minReadySeconds,omitempty"`
	// Node tolerations for server scheduling to nodes with taints
	// Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.
	Tolerations *map[string]string `pulumi:"labels,omitempty"`
	// Affinity and anti-affinity
	// Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity.
	Affinity *corev1.Affinity `pulumi:"affinity,ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Affinity,omitempty"`
	// Topology spread constraints rely on node labels to identify the topology domain(s) that each Node is in.
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/.
	TopologySpreadConstraints *[]corev1.TopologySpreadConstraint `pulumi:"topologySpreadConstraints,ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:TopologySpreadConstraint,omitempty"`
	// How long to wait for the drain of connections.
	TerminateGracePeriodSeconds *int `pulumi:"terminateGracePeriodSeconds,omitempty"`
	// Node labels for controller pod assignment
	// Ref: https://kubernetes.io/docs/user-guide/node-selection/.
	NodeSelector *map[string]string `pulumi:"nodeSelector,omitempty"`
	// Startup probe values
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
	StartupProbe *corev1.Probe `pulumi:"startupProbe,ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Probe,omitempty"`
	// Liveness probe values
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
	LivenessProbe *corev1.Probe `pulumi:"livenessProbe,ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Probe,omitempty"`
	// Readiness probe values
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
	ReadinessProbe *corev1.Probe `pulumi:"readinessProbe,ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Probe,omitempty"`
	// Path of the health check endpoint. All requests received on the port defined by
	// the healthz-port parameter are forwarded internally to this path.
	HealthCheckPath *string `pulumi:"healthCheckPath,omitempty"`
	// Address to bind the health check endpoint.
	// It is better to set this option to the internal node address
	// if the ingress nginx controller is running in the hostNetwork: true mode.
	HealthCheckHost *string `pulumi:"heathCheckHost,omitempty"`
	// Annotations to be added to controller pods.
	PodAnnotations *map[string]string `pulumi:"podAnnotations,omitempty"`
	ReplicaCount   *int               `pulumi:"replicaCount,omitempty"`
	MinAvailable   *int               `pulumi:"minAvailable,omitempty"`
	// Define requests resources to avoid probe issues due to CPU utilization in busy nodes
	// ref: https://github.com/kubernetes/ingress-nginx/issues/4735#issuecomment-551204903
	// Ideally, there should be no limits.
	// https://engineering.indeedblog.com/blog/2019/12/cpu-throttling-regression-fix/
	Resources *corev1.ResourceRequirements `pulumi:"resources,ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:ResourceRequirements,omitempty"`
	// Mutually exclusive with keda autoscaling.
	Autoscaling *Autoscaling `pulumi:"autoscaling,omitempty"`
	// Custom or additional autoscaling metrics
	// ref: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-custom-metrics
	AutoscalingTemplate *[]AutoscalingTemplate `pulumi:"autoscalingTemplate,omitempty"`
	// Mutually exclusive with hpa autoscaling.
	Keds *Keda `pulumi:"keda,omitempty"`
	// Enable mimalloc as a drop-in replacement for malloc.
	// ref: https://github.com/microsoft/mimalloc.
	EnableMimalloc *bool `pulumi:"enableMimalloc,omitempty"`
	// Override NGINX template.
	CustomTemplate *ControllerCustomTemplate `pulumi:"customTemplate,omitempty"`
	Service        *ControllerService        `pulumi:"service,omitempty"`
	// Additional containers to be added to the controller pod.
	// See https://github.com/lemonldap-ng-controller/lemonldap-ng-controller as example.
	ExtraContainers []*corev1.Container `pulumi:"extraContainers,omitempty"`
	// Additional volumeMounts to the controller main container.
	//  - name: copy-portal-skins
	//    mountPath: /var/lib/lemonldap-ng/portal/skins
	ExtraVolumeMounts []*corev1.VolumeMount `pulumi:"extraVolumeMounts,omitempty"`
	// Additional volumes to the controller pod.
	//  - name: copy-portal-skins
	//    emptyDir: {}
	ExtraVolumes []*corev1.Volume `pulumi:"extraVolume,omitempty"`
	// Containers, which are run before the app containers are started.
	// - name: init-myservice
	//   image: busybox
	//   command: ['sh', '-c', 'until nslookup myservice; do echo waiting for myservice; sleep 2; done;']
	ExtraInitContainers []*corev1.Container         `pulumi:"extraInitContainers,omitempty"`
	AdmissionWebhooks   *ContollerAdmissionWebhooks `pulumi:"admissionWebhooks,omitempty"`
	Metrics             *ControllerMetrics          `pulumi:"metrics,omitempty"`
	// Improve connection draining when ingress controller pod is deleted using a lifecycle hook:
	// With this new hook, we increased the default terminationGracePeriodSeconds from 30 seconds
	// to 300, allowing the draining of connections up to five minutes.
	// If the active connections end before that, the pod will terminate gracefully at that time.
	// To effectively take advantage of this feature, the Configmap feature
	// worker-shutdown-timeout new value is 240s instead of 10s.
	Lifecycle         *corev1.Lifecycle `pulumi:"lifecycle,omitempty"`
	PriorityClassName *string           `pulumi:"priorityClassName,omitempty"`
}

type ControllerImage struct {
	Registry *string `pulumi:"registry,omitempty"`
	Image    *string `pulumi:"image,omitempty"`
	// for backwards compatibility consider setting the full image url via the repository value below
	// use *either* current default registry/image or repository format or installing will fail.
	Repository               *string `pulumi:"repository,omitempty"`
	Tag                      *string `pulumi:"tag,omitempty"`
	Digest                   *string `pulumi:"digest,omitempty"`
	PullPolicy               *string `pulumi:"pullPolicy,omitempty"`
	RunAsUser                *string `pulumi:"runAsUser,omitempty"`
	RunAsNonRoot             *bool   `pulumi:"runAsNonRoot,omitempty"`
	ReadOnlyRootFilesystem   *bool   `pulumi:"readOnlyRootFilesystem,omitempty"`
	AllowPrivilegeEscalation *bool   `pulumi:"allowPrivilegeEscalation,omitempty"`
}

type ControllerPort struct {
	Http  *int `pulumi:"http,omitempty"`
	Https *int `pulumi:"https,omitempty"`
}

type ControllerHostPort struct {
	Enabled *bool                    `pulumi:"enabled,omitempty"`
	Ports   *ControllerHostPortPorts `pulumi:"ports,omitempty"`
}

type ControllerHostPortPorts struct {
	Http  *int `pulumi:"http,omitempty"`
	Https *int `pulumi:"https,omitempty"`
}

type ControllerServiceAccount struct {
	Create                       *bool   `pulumi:"create,omitempty"`
	Name                         *string `pulumi:"name,omitempty"`
	AutomountServiceAccountToken *bool   `pulumi:"automountServiceAccountToken,omitempty"`
}

type ControllerIngressClassResource struct {
	Name            *string `pulumi:"name,omitempty"`
	Enabled         *bool   `pulumi:"enabled,omitempty"`
	Default         *bool   `pulumi:"default,omitempty"`
	ControllerValue *string `pulumi:"controllerValue,omitempty"`
	// Parameters is a link to a custom resource containing additional
	// configuration for the controller. This is optional if the controller
	// does not require extra parameters.
	Parameters *map[string]interface{} `pulumi:"parameters,omitempty"`
}

type ControllerPublishService struct {
	Enabled *bool `pulumi:"enabled,omitempty"`
	// Allows overriding of the publish service to bind to. Must be <namespace>/<service_name>.
	PathOverride *string `pulumi:"pathOverride,omitempty"`
}

type ControllerScope struct {
	Enabled   *bool   `pulumi:"enabled,omitempty"`
	Namespace *string `pulumi:"namespace,omitempty"`
}

type ControllerTcp struct {
	ConfigMapNamespace *string `pulumi:"configMapNamespace,omitempty"`
	// Annotations to be added to the tcp config configmap.
	Annotations *map[string]string `pulumi:"annotations,omitempty"`
}

type ControllerUdp struct {
	ConfigMapNamespace *string `pulumi:"configMapNamespace,omitempty"`
	// Annotations to be added to the udp config configmap.
	Annotations *map[string]string `pulumi:"annotations,omitempty"`
}

type ControllerUpdateStrategy struct {
	RollingUpdate *ControllerRollingUpdate `pulumi:"rollingUpdate,omitempty"`
	Type          *string                  `pulumi:"type,omitempty"`
}

type ControllerRollingUpdate struct {
	MaxUnavailable *int `pulumi:"maxUnavailable,omitempty"`
}

type Autoscaling struct {
	Annotations                        *map[string]string   `pulumi:"annotations,omitempty"`
	Enabled                            *bool                `pulumi:"enabled,omitempty"`
	MinReplicas                        *int                 `pulumi:"minReplicas,omitempty"`
	MaxReplicas                        *int                 `pulumi:"maxReplicas,omitempty"`
	TargetCPUUtilizationPercentation   *int                 `pulumi:"targetCPUUtilizationPercentage,omitempty"`
	TargetMemoryUtilizationPercentatge *int                 `pulumi:"targetMemoryUtilizationPercentage,omitempty"`
	Behavior                           *AutoscalingBehavior `pulumi:"controllerAutoscalingBehavior,omitempty"`
}

type AutoscalingBehavior struct {
	ScaleDown *AutoscalingBehaviorScaling `pulumi:"scaleDown,omitempty"`
	ScaleUp   *AutoscalingBehaviorScaling `pulumi:"scaleUp,omitempty"`
}

type AutoscalingBehaviorScaling struct {
	StabilizationWindowSeconds *int                                `pulumi:"stabilizationWindowSeconds,omitempty"`
	Policies                   *[]AutoscalingBehaviorScalingPolicy `pulumi:"policies,omitempty"`
}

type AutoscalingBehaviorScalingPolicy struct {
	Type          *string `pulumi:"type,omitempty"`
	Value         *int    `pulumi:"value,omitempty"`
	PeriodSeconds *int    `pulumi:"periodSeconds,omitempty"`
}

type AutoscalingTemplate struct {
	Type *string                  `pulumi:"type,omitempty"`
	Pods *AutoscalingTemplatePods `pulumi:"pods,omitempty"`
}

type AutoscalingTemplatePods struct {
	Metric *AutoscalingTemplatePodsMetric `pulumi:"metric,omitempty"`
	Target *AutoscalingTemplatePodsTarget `pulumi:"target,omitempty"`
}

type AutoscalingTemplatePodsMetric struct {
	Name *string `pulumi:"name,omitempty"`
}

type AutoscalingTemplatePodsTarget struct {
	Type         *string `pulumi:"type,omitempty"`
	AverageValue *string `pulumi:"averageValue,omitempty"`
}

type Keda struct {
	// apiVersion changes with keda 1.x vs 2.x:
	// 2.x = keda.sh/v1alpha1,
	// 1.x = keda.k8s.io/v1alpha1.
	APIVersion                    *string              `pulumi:"apiVersion,omitempty"`
	Enabled                       *bool                `pulumi:"enabled,omitempty"`
	MinReplicas                   *int                 `pulumi:"minReplicas,omitempty"`
	MaxReplicas                   *int                 `pulumi:"maxReplicas,omitempty"`
	PollingInterval               *int                 `pulumi:"pollingInterval,omitempty"`
	CooldownPeriod                *int                 `pulumi:"cooldownPeriod,omitempty"`
	RestoreToOriginalReplicaCount *bool                `pulumi:"restoreToOriginalReplicaCount,omitempty"`
	ScaledObject                  *KedaScaledObject    `pulumi:"scaledObject,omitempty"`
	Triggers                      *[]KedaTrigger       `pulumi:"triggers,omitempty"`
	Behavior                      *AutoscalingBehavior `pulumi:"behavior,omitempty"`
}

type KedaScaledObject struct {
	// Custom annotations for ScaledObject resource.
	Annotations *map[string]string `pulumi:"annotations,omitempty"`
}

type KedaTrigger struct {
	Type     *string                 `pulumi:"type,omitempty"`
	Metadata *map[string]interface{} `pulumi:"metadata,omitempty"`
}

type ControllerCustomTemplate struct {
	ConfigMapName *string `pulumi:"configMapName,omitempty"`
	ConfigMapKey  *string `pulumi:"configMapKey,omitempty"`
}

type ControllerService struct {
	Enabled     *bool                   `pulumi:"enabled,omitempty"`
	Annotations *map[string]interface{} `pulumi:"annotations,omitempty"`
	Labels      *map[string]interface{} `pulumi:"labels,omitempty"`
	ClusterIP   *string                 `pulumi:"clusterIP,omitempty"`
	// List of IP addresses at which the controller services are available
	// Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
	ExternalIPs              *[]string `pulumi:"externalIPs,omitempty"`
	LoadBalancerIP           *string   `pulumi:"loadBalancerIPs,omitempty"`
	LoadBalancerSourceRanges *[]string `pulumi:"loadBalancerSourceRanges,omitempty"`
	EnableHttp               *bool     `pulumi:"enableHttp,omitempty"`
	EnableHttps              *bool     `pulumi:"enableHttps,omitempty"`
	// Set external traffic policy to: "Local" to preserve source IP on
	// providers supporting it.
	// Ref: https://kubernetes.io/docs/tutorials/services/source-ip/#source-ip-for-services-with-typeloadbalancer
	ExternalTrafficPolicy *string `pulumi:"externalTrafficPolicy,omitempty"`
	// Must be either "None" or "ClientIP" if set. Kubernetes will default to "None".
	// Ref: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	SessionAffinity *string `pulumi:"sessionAffinity,omitempty"`
	// specifies the health check node port (numeric port number) for the service. If healthCheckNodePort isn’t specified,
	// the service controller allocates a port from your cluster’s NodePort range.
	// Ref: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#preserving-the-client-source-ip
	HealthCheckNodePort *int                        `pulumi:"healthCheckNodePort,omitempty"`
	Ports               *ControllerPort             `pulumi:"ports,omitempty"`
	TargetPorts         *ControllerPort             `pulumi:"targetPorts,omitempty"`
	Type                *string                     `pulumi:"type,omitempty"`
	NodePorts           *ControllerServiceNodePorts `pulumi:"nodePorts,omitempty"`
	// Enables an additional internal load balancer (besides the external one).
	// Annotations are mandatory for the load balancer to come up. Varies with the cloud service.
	Internal *ControllerServiceInternal `pulumi:"internal,omitempty"`
}

type ControllerServiceNodePorts struct {
	Http  *string                 `pulumi:"http,omitempty"`
	Https *string                 `pulumi:"https,omitempty"`
	Tcp   *map[string]interface{} `pulumi:"tcp,omitempty"`
	Udp   *map[string]interface{} `pulumi:"udp,omitempty"`
}

type ControllerServiceInternal struct {
	Enabled        *bool                   `pulumi:"enabled,omitempty"`
	Annotations    *map[string]interface{} `pulumi:"annotations,omitempty"`
	Labels         *map[string]interface{} `pulumi:"labels,omitempty"`
	LoadBalancerIP *string                 `pulumi:"loadBalancerIPs,omitempty"`
	// Restrict access For LoadBalancer service. Defaults to 0.0.0.0/0.
	LoadBalancerSourceRanges *[]string `pulumi:"loadBalancerSourceRanges,omitempty"`
	// Set external traffic policy to: "Local" to preserve source IP on
	// providers supporting it.
	// Ref: https://kubernetes.io/docs/tutorials/services/source-ip/#source-ip-for-services-with-typeloadbalancer
	ExternalTrafficPolicy *string `pulumi:"externalTrafficPolicy,omitempty"`
}

type ContollerAdmissionWebhooks struct {
	Enabled           *bool                   `pulumi:"enabled,omitempty"`
	Annotations       *map[string]interface{} `pulumi:"annotations,omitempty"`
	FailurePolicy     *string                 `pulumi:"failurePolicy,omitempty"`
	TimeoutSeconds    *int                    `pulumi:"timeoutSeconds,omitempty"`
	Port              *int                    `pulumi:"port,omitempty"`
	Certificate       *string                 `pulumi:"certificate,omitempty"`
	Key               *string                 `pulumi:"key,omitempty"`
	NamespaceSelector *map[string]interface{} `pulumi:"namespaceSelector,omitempty"`
	ObjectSelector    *map[string]interface{} `pulumi:"objectSelector,omitempty"`
	// Use an existing PSP instead of creating one.
	ExistingPsp     *string                                      `pulumi:"existingPsp,omitempty"`
	Service         *ControllerAdmissionWebhooksService          `pulumi:"service,omitempty"`
	CreateSecretJob *ControllerAdmissionWebhooksCreateSecretJob  `pulumi:"createSecretJob,omitempty"`
	PatchWebhookJob *ControllerAdmissionWebhooksPatchWebhbookJob `pulumi:"patchWebhookJob,omitempty"`
	Patch           *ControllerAdmissionWebhooksPatch            `pulumi:"patch,omitempty"`
}

type ControllerAdmissionWebhooksService struct {
	Annotations              *map[string]interface{} `pulumi:"annotations,omitempty"`
	ClusterIP                *string                 `pulumi:"clusterIP,omitempty"`
	ExternalIPs              *[]string               `pulumi:"externalIPs,omitempty"`
	LoadBalancerIP           *string                 `pulumi:"loadBalancerIPs,omitempty"`
	LoadBalancerSourceRanges *[]string               `pulumi:"loadBalancerSourceRanges,omitempty"`
	ServicePort              *int                    `pulumi:"servicePort,omitempty"`
	Type                     *string                 `pulumi:"type,omitempty"`
}

type ControllerAdmissionWebhooksCreateSecretJob struct {
	Resources *corev1.ResourceRequirements `pulumi:"resources,ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:ResourceRequirements,omitempty"`
}

type ControllerAdmissionWebhooksPatchWebhbookJob struct {
	Resources *corev1.ResourceRequirements `pulumi:"resources,ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:ResourceRequirements,omitempty"`
}

type ControllerAdmissionWebhooksPatch struct {
	Enabled *bool            `pulumi:"enabled,omitempty"`
	Image   *ControllerImage `pulumi:"image,omitempty"`
	// Provide a priority class name to the webhook patching job.
	PriorityClassName *string                 `pulumi:"priorityClassName,omitempty"`
	PodAnnotations    *map[string]interface{} `pulumi:"podAnnotations,omitempty"`
	NodeSelector      *map[string]string      `pulumi:"nodeSelector,omitempty"`
	Tolerations       *[]corev1.Toleration    `pulumi:"tolerations,omitempty"`
	RunAsUser         *int                    `pulumi:"runAsUser,omitempty"`
}

type ControllerMetrics struct {
	// if this port is changed, change healthz-port: in extraArgs: accordingly.
	Port           *int                              `pulumi:"port,omitempty"`
	Enabled        *bool                             `pulumi:"enabled,omitempty"`
	Service        *ControllerMetricsService         `pulumi:"service,omitempty"`
	ServiceMonitor *ControllerMetricsServiceMonitor  `pulumi:"serviceMonitor,omitempty"`
	PrometheusRule *ControllerMetricsPrometheusRules `pulumi:"prometheusRule,omitempty"`
}

type ControllerMetricsService struct {
	Annotations              *map[string]string `pulumi:"annotations,omitempty"`
	ClusterIP                *string            `pulumi:"clusterIP,omitempty"`
	ExternalIPs              *[]string          `pulumi:"externalIPs,omitempty"`
	LoadBalancerIP           *string            `pulumi:"loadBalancerIPs,omitempty"`
	LoadBalancerSourceRanges *[]string          `pulumi:"loadBalancerSourceRanges,omitempty"`
	ServicePort              *int               `pulumi:"servicePort,omitempty"`
	Type                     *string            `pulumi:"type,omitempty"`
	ExternalTrafficPolicy    *string            `pulumi:"externalTrafficPolicy,omitempty"`
	NodePort                 *string            `pulumi:"nodePort,omitempty"`
}

type ControllerMetricsServiceMonitor struct {
	Enabled          *bool                   `pulumi:"enabled,omitempty"`
	AdditionalLabels *map[string]interface{} `pulumi:"additionalLabels,omitempty"`
	// The label to use to retrieve the job name from.
	JobLabel          *string                 `pulumi:"jobLabel,omitempty"`
	Namespace         *string                 `pulumi:"namespace,omitempty"`
	NamespaceSelector *map[string]interface{} `pulumi:"namespaceSelector,omitempty"`
	ScrapeInterval    *string                 `pulumi:"scrapeInterval,omitempty"`
	HonorLabels       *bool                   `pulumi:"honorLabels,omitempty"`
	TargetLabels      *[]string               `pulumi:"targetLabels,omitempty"`
	MetricRelabelings *[]string               `pulumi:"metricRelabelings,omitempty"`
}

type ControllerMetricsPrometheusRules struct {
	Enabled          *bool                   `pulumi:"enabled,omitempty"`
	AdditionalLabels *map[string]interface{} `pulumi:"additionalLabels,omitempty"`
	Namespace        *string                 `pulumi:"namespace,omitempty"`
	Rules            *[]interface{}          `pulumi:"rules,omitempty"`
}

type ControllerDefaultBackend struct {
	Enabled *bool            `pulumi:"enabled,omitempty"`
	Name    *string          `pulumi:"name,omitempty"`
	Image   *ControllerImage `pulumi:"image,omitempty"`
	// Use an existing PSP instead of creating one.
	ExistingPsp    *string                   `pulumi:"existingPsp,omitempty"`
	ExtraArgs      *map[string]interface{}   `pulumi:"extraArgs,omitempty"`
	ServiceAccount *ControllerServiceAccount `pulumi:"serviceAccount,omitempty"`
	ExtraEnvs      *[]corev1.EnvVar          `pulumi:"extraEnvs,omitempty"`
	Port           *int                      `pulumi:"port,omitempty"`
	// Liveness probe values for default backend.
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
	LivenessProbe *corev1.Probe `pulumi:"livenessProbe,omitempty"`
	// Readiness probe values for default backend.
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
	ReadinessProbe *corev1.Probe `pulumi:"readinessProbe,omitempty"`
	// Node tolerations for server scheduling to nodes with taints.
	// Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	Tolerations *[]corev1.Toleration `pulumi:"tolerations,omitempty"`
	Affinity    *corev1.Affinity     `pulumi:"affinity,omitempty"`
	// Security Context policies for controller pods.
	// See https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/ for
	// notes on enabling and using sysctls.
	PodSecurityContext *corev1.PodSecurityContext `pulumi:"podSecurityContext,omitempty"`
	// labels to add to the pod container metadata
	PodLabels *map[string]string `pulumi:"podLabels,omitempty"`
	// Node labels for default backend pod assignment
	// Ref: https://kubernetes.io/docs/user-guide/node-selection/.
	NodeSelector *map[string]string `pulumi:"nodeSelector,omitempty"`
	// Annotations to be added to default backend pods.
	PodAnnotations *map[string]string           `pulumi:"podAnnotations,omitempty"`
	ReplicaCount   *int                         `pulumi:"replicaCount,omitempty"`
	MinAvailable   *int                         `pulumi:"minAvailable,omitempty"`
	Resources      *corev1.ResourceRequirements `pulumi:"resources,ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:ResourceRequirements,omitempty"`
	// Additional volumeMounts to the default backend container.
	//  - name: copy-portal-skins
	//    mountPath: /var/lib/lemonldap-ng/portal/skins
	ExtraVolumeMounts []*corev1.VolumeMount `pulumi:"extraVolumeMounts,omitempty"`
	// Additional volumes to the default backend pod.
	//  - name: copy-portal-skins
	//    emptyDir: {}
	ExtraVolumes      []*corev1.Volume                 `pulumi:"extraVolume,omitempty"`
	Autoscaling       *Autoscaling                     `pulumi:"autoscaling,omitempty"`
	Service           *ControllerDefaultBackendService `pulumi:"service,omitempty"`
	PriorityClassName *string                          `pulumi:"priorityClassName,omitempty"`
}

type ControllerDefaultBackendService struct {
	Annoations *map[string]string `pulumi:"annotations,omitempty"`
	ClusterIP  *string            `pulumi:"clusterIP,omitempty"`
	// List of IP addresses at which the default backend service is available.
	// Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
	ExternalIPs              *[]string `pulumi:"externalIPs,omitempty"`
	LoadBalancerIP           *string   `pulumi:"loadBalancerIP,omitempty"`
	LoadBalancerSourceRanges *[]string `pulumi:"loadBalancerSourceRanges,omitempty"`
	ServicePort              *int      `pulumi:"servicePort,omitempty"`
	Type                     *string   `pulumi:"type,omitempty"`
}

type ControllerPodSecurityPolicy struct {
	Enabled *bool `pulumi:"enabled,omitempty"`
}

type ControllerRBAC struct {
	Create *bool `pulumi:"create,omitempty"`
	Scope  *bool `pulumi:"scope,omitempty"`
}
