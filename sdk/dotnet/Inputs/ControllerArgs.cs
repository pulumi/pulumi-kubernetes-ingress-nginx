// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesIngressNginx.Inputs
{

    public sealed class ControllerArgs : global::Pulumi.ResourceArgs
    {
        [Input("addHeaders")]
        private InputMap<ImmutableDictionary<string, string>>? _addHeaders;

        /// <summary>
        /// Will add custom headers before sending response traffic to the client according to: https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#add-headers.
        /// </summary>
        public InputMap<ImmutableDictionary<string, string>> AddHeaders
        {
            get => _addHeaders ?? (_addHeaders = new InputMap<ImmutableDictionary<string, string>>());
            set => _addHeaders = value;
        }

        [Input("admissionWebhooks")]
        public Input<Inputs.ContollerAdmissionWebhooksArgs>? AdmissionWebhooks { get; set; }

        /// <summary>
        /// Affinity and anti-affinity Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity.
        /// </summary>
        [Input("affinity")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.AffinityArgs>? Affinity { get; set; }

        /// <summary>
        /// This configuration defines if Ingress Controller should allow users to set their own *-snippet annotations, otherwise this is forbidden / dropped when users add those annotations. Global snippets in ConfigMap are still respected.
        /// </summary>
        [Input("allowSnippetAnnotations")]
        public Input<bool>? AllowSnippetAnnotations { get; set; }

        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Annotations to be added to the controller Deployment or DaemonSet.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// Mutually exclusive with keda autoscaling.
        /// </summary>
        [Input("autoscaling")]
        public Input<Inputs.AutoscalingArgs>? Autoscaling { get; set; }

        [Input("autoscalingTemplate")]
        private InputList<Inputs.AutoscalingTemplateArgs>? _autoscalingTemplate;

        /// <summary>
        /// Custom or additional autoscaling metrics ref: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-custom-metrics
        /// </summary>
        public InputList<Inputs.AutoscalingTemplateArgs> AutoscalingTemplate
        {
            get => _autoscalingTemplate ?? (_autoscalingTemplate = new InputList<Inputs.AutoscalingTemplateArgs>());
            set => _autoscalingTemplate = value;
        }

        [Input("config")]
        private InputMap<string>? _config;

        /// <summary>
        /// Will add custom configuration options to Nginx https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/.
        /// </summary>
        public InputMap<string> Config
        {
            get => _config ?? (_config = new InputMap<string>());
            set => _config = value;
        }

        [Input("configAnnotations")]
        private InputMap<string>? _configAnnotations;

        /// <summary>
        /// Annotations to be added to the controller config configuration configmap.
        /// </summary>
        public InputMap<string> ConfigAnnotations
        {
            get => _configAnnotations ?? (_configAnnotations = new InputMap<string>());
            set => _configAnnotations = value;
        }

        /// <summary>
        /// Allows customization of the configmap / nginx-configmap namespace.
        /// </summary>
        [Input("configMapNamespace")]
        public Input<string>? ConfigMapNamespace { get; set; }

        /// <summary>
        /// Configures the controller container name.
        /// </summary>
        [Input("containerName")]
        public Input<string>? ContainerName { get; set; }

        /// <summary>
        /// Configures the ports the nginx-controller listens on.
        /// </summary>
        [Input("containerPort")]
        public Input<Inputs.ControllerPortArgs>? ContainerPort { get; set; }

        /// <summary>
        /// Override NGINX template.
        /// </summary>
        [Input("customTemplate")]
        public Input<Inputs.ControllerCustomTemplateArgs>? CustomTemplate { get; set; }

        [Input("dnsConfig")]
        private InputMap<string>? _dnsConfig;

        /// <summary>
        /// Optionally customize the pod dnsConfig.
        /// </summary>
        public InputMap<string> DnsConfig
        {
            get => _dnsConfig ?? (_dnsConfig = new InputMap<string>());
            set => _dnsConfig = value;
        }

        /// <summary>
        /// Optionally change this to ClusterFirstWithHostNet in case you have 'hostNetwork: true'. By default, while using host network, name resolution uses the host's DNS. If you wish nginx-controller to keep resolving names inside the k8s network, use ClusterFirstWithHostNet.
        /// </summary>
        [Input("dnsPolicy")]
        public Input<string>? DnsPolicy { get; set; }

        /// <summary>
        /// Election ID to use for status update.
        /// </summary>
        [Input("electionID")]
        public Input<string>? ElectionID { get; set; }

        /// <summary>
        /// Enable mimalloc as a drop-in replacement for malloc. ref: https://github.com/microsoft/mimalloc.
        /// </summary>
        [Input("enableMimalloc")]
        public Input<bool>? EnableMimalloc { get; set; }

        /// <summary>
        /// Use an existing PSP instead of creating one.
        /// </summary>
        [Input("existingPsp")]
        public Input<string>? ExistingPsp { get; set; }

        [Input("extraArgs")]
        private InputMap<string>? _extraArgs;

        /// <summary>
        /// Additional command line arguments to pass to nginx-ingress-controller E.g. to specify the default SSL certificate you can use `default-ssl-certificate: "&lt;namespace&gt;/&lt;secret_name&gt;"`.
        /// </summary>
        public InputMap<string> ExtraArgs
        {
            get => _extraArgs ?? (_extraArgs = new InputMap<string>());
            set => _extraArgs = value;
        }

        [Input("extraContainers")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerArgs>? _extraContainers;

        /// <summary>
        /// Additional containers to be added to the controller pod. See https://github.com/lemonldap-ng-controller/lemonldap-ng-controller as example.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerArgs> ExtraContainers
        {
            get => _extraContainers ?? (_extraContainers = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerArgs>());
            set => _extraContainers = value;
        }

        [Input("extraEnvs")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.EnvVarArgs>? _extraEnvs;

        /// <summary>
        /// Additional environment variables to set.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.EnvVarArgs> ExtraEnvs
        {
            get => _extraEnvs ?? (_extraEnvs = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.EnvVarArgs>());
            set => _extraEnvs = value;
        }

        [Input("extraInitContainers")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerArgs>? _extraInitContainers;

        /// <summary>
        /// Containers, which are run before the app containers are started. - name: init-myservice   image: busybox   command: ['sh', '-c', 'until nslookup myservice; do echo waiting for myservice; sleep 2; done;']
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerArgs> ExtraInitContainers
        {
            get => _extraInitContainers ?? (_extraInitContainers = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerArgs>());
            set => _extraInitContainers = value;
        }

        [Input("extraVolumeMounts")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs>? _extraVolumeMounts;

        /// <summary>
        /// Additional volumeMounts to the controller main container.  - name: copy-portal-skins    mountPath: /var/lib/lemonldap-ng/portal/skins
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs> ExtraVolumeMounts
        {
            get => _extraVolumeMounts ?? (_extraVolumeMounts = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs>());
            set => _extraVolumeMounts = value;
        }

        [Input("extraVolumes")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeArgs>? _extraVolumes;

        /// <summary>
        /// Additional volumes to the controller pod.  - name: copy-portal-skins    emptyDir: {}
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeArgs> ExtraVolumes
        {
            get => _extraVolumes ?? (_extraVolumes = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeArgs>());
            set => _extraVolumes = value;
        }

        /// <summary>
        /// Path of the health check endpoint. All requests received on the port defined by the healthz-port parameter are forwarded internally to this path.
        /// </summary>
        [Input("healthCheckPath")]
        public Input<string>? HealthCheckPath { get; set; }

        /// <summary>
        /// Address to bind the health check endpoint. It is better to set this option to the internal node address if the ingress nginx controller is running in the hostNetwork: true mode.
        /// </summary>
        [Input("heathCheckHost")]
        public Input<string>? HeathCheckHost { get; set; }

        /// <summary>
        /// Required for use with CNI based kubernetes installations (such as ones set up by kubeadm), since CNI and hostport don't mix yet. Can be deprecated once https://github.com/kubernetes/kubernetes/issues/23920 is merged.
        /// </summary>
        [Input("hostNetwork")]
        public Input<bool>? HostNetwork { get; set; }

        /// <summary>
        /// Use host ports 80 and 443. Disabled by default.
        /// </summary>
        [Input("hostPort")]
        public Input<Inputs.ControllerHostPortArgs>? HostPort { get; set; }

        [Input("hostname")]
        private InputMap<ImmutableDictionary<string, string>>? _hostname;

        /// <summary>
        /// Optionally customize the pod hostname.
        /// </summary>
        public InputMap<ImmutableDictionary<string, string>> Hostname
        {
            get => _hostname ?? (_hostname = new InputMap<ImmutableDictionary<string, string>>());
            set => _hostname = value;
        }

        [Input("image")]
        public Input<Inputs.ControllerImageArgs>? Image { get; set; }

        /// <summary>
        /// Process IngressClass per name (additionally as per spec.controller).
        /// </summary>
        [Input("ingressClassByName")]
        public Input<bool>? IngressClassByName { get; set; }

        /// <summary>
        /// This section refers to the creation of the IngressClass resource. IngressClass resources are supported since k8s &gt;= 1.18 and required since k8s &gt;= 1.19
        /// </summary>
        [Input("ingressClassResource")]
        public Input<Inputs.ControllerIngressClassResourceArgs>? IngressClassResource { get; set; }

        /// <summary>
        /// Mutually exclusive with hpa autoscaling.
        /// </summary>
        [Input("keda")]
        public Input<Inputs.KedaArgs>? Keda { get; set; }

        /// <summary>
        /// DaemonSet or Deployment.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Improve connection draining when ingress controller pod is deleted using a lifecycle hook: With this new hook, we increased the default terminationGracePeriodSeconds from 30 seconds to 300, allowing the draining of connections up to five minutes. If the active connections end before that, the pod will terminate gracefully at that time. To effectively take advantage of this feature, the Configmap feature worker-shutdown-timeout new value is 240s instead of 10s.
        /// </summary>
        [Input("lifecycle")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.LifecycleArgs>? Lifecycle { get; set; }

        /// <summary>
        /// Liveness probe values Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
        /// </summary>
        [Input("livenessProbe")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ProbeArgs>? LivenessProbe { get; set; }

        /// <summary>
        /// Maxmind license key to download GeoLite2 Databases https://blog.maxmind.com/2019/12/18/significant-changes-to-accessing-and-using-geolite2-databases.
        /// </summary>
        [Input("maxmindLicenseKey")]
        public Input<string>? MaxmindLicenseKey { get; set; }

        [Input("metrics")]
        public Input<Inputs.ControllerMetricsArgs>? Metrics { get; set; }

        [Input("minAvailable")]
        public Input<int>? MinAvailable { get; set; }

        /// <summary>
        /// minReadySeconds to avoid killing pods before we are ready.
        /// </summary>
        [Input("minReadySeconds")]
        public Input<int>? MinReadySeconds { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nodeSelector")]
        private InputMap<string>? _nodeSelector;

        /// <summary>
        /// Node labels for controller pod assignment Ref: https://kubernetes.io/docs/user-guide/node-selection/.
        /// </summary>
        public InputMap<string> NodeSelector
        {
            get => _nodeSelector ?? (_nodeSelector = new InputMap<string>());
            set => _nodeSelector = value;
        }

        [Input("podAnnotations")]
        private InputMap<string>? _podAnnotations;

        /// <summary>
        /// Annotations to be added to controller pods.
        /// </summary>
        public InputMap<string> PodAnnotations
        {
            get => _podAnnotations ?? (_podAnnotations = new InputMap<string>());
            set => _podAnnotations = value;
        }

        [Input("podLabels")]
        private InputMap<string>? _podLabels;

        /// <summary>
        /// labels to add to the pod container metadata.
        /// </summary>
        public InputMap<string> PodLabels
        {
            get => _podLabels ?? (_podLabels = new InputMap<string>());
            set => _podLabels = value;
        }

        /// <summary>
        /// Security Context policies for controller pods.
        /// </summary>
        [Input("podSecurityContext")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.PodSecurityContextArgs>? PodSecurityContext { get; set; }

        [Input("priorityClassName")]
        public Input<string>? PriorityClassName { get; set; }

        [Input("proxySetHeaders")]
        private InputMap<ImmutableDictionary<string, string>>? _proxySetHeaders;

        /// <summary>
        /// Will add custom headers before sending traffic to backends according to https://github.com/kubernetes/ingress-nginx/tree/main/docs/examples/customization/custom-headers.
        /// </summary>
        public InputMap<ImmutableDictionary<string, string>> ProxySetHeaders
        {
            get => _proxySetHeaders ?? (_proxySetHeaders = new InputMap<ImmutableDictionary<string, string>>());
            set => _proxySetHeaders = value;
        }

        /// <summary>
        /// Allows customization of the source of the IP address or FQDN to report in the ingress status field. By default, it reads the information provided by the service. If disable, the status field reports the IP address of the node or nodes where an ingress controller pod is running.
        /// </summary>
        [Input("publishService")]
        public Input<Inputs.ControllerPublishServiceArgs>? PublishService { get; set; }

        /// <summary>
        /// Readiness probe values Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
        /// </summary>
        [Input("readinessProbe")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ProbeArgs>? ReadinessProbe { get; set; }

        [Input("replicaCount")]
        public Input<int>? ReplicaCount { get; set; }

        /// <summary>
        /// Bare-metal considerations via the host network https://kubernetes.github.io/ingress-nginx/deploy/baremetal/#via-the-host-network Ingress status was blank because there is no Service exposing the NGINX Ingress controller in a configuration using the host network, the default --publish-service flag used in standard cloud setups does not apply.
        /// </summary>
        [Input("reportNodeInternalIp")]
        public Input<bool>? ReportNodeInternalIp { get; set; }

        /// <summary>
        /// Define requests resources to avoid probe issues due to CPU utilization in busy nodes ref: https://github.com/kubernetes/ingress-nginx/issues/4735#issuecomment-551204903 Ideally, there should be no limits. https://engineering.indeedblog.com/blog/2019/12/cpu-throttling-regression-fix/
        /// </summary>
        [Input("resources")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs>? Resources { get; set; }

        /// <summary>
        /// Limit the scope of the controller.
        /// </summary>
        [Input("scope")]
        public Input<Inputs.ControllerScopeArgs>? Scope { get; set; }

        [Input("service")]
        public Input<Inputs.ControllerServiceArgs>? Service { get; set; }

        /// <summary>
        /// Startup probe values Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
        /// </summary>
        [Input("startupProbe")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ProbeArgs>? StartupProbe { get; set; }

        [Input("sysctls")]
        private InputMap<ImmutableDictionary<string, string>>? _sysctls;

        /// <summary>
        /// See https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/ for notes on enabling and using sysctls.
        /// </summary>
        public InputMap<ImmutableDictionary<string, string>> Sysctls
        {
            get => _sysctls ?? (_sysctls = new InputMap<ImmutableDictionary<string, string>>());
            set => _sysctls = value;
        }

        /// <summary>
        /// Allows customization of the tcp-services-configmap.
        /// </summary>
        [Input("tcp")]
        public Input<Inputs.ControllerTcpArgs>? Tcp { get; set; }

        /// <summary>
        /// How long to wait for the drain of connections.
        /// </summary>
        [Input("terminateGracePeriodSeconds")]
        public Input<int>? TerminateGracePeriodSeconds { get; set; }

        /// <summary>
        /// Node tolerations for server scheduling to nodes with taints Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.
        /// </summary>
        [Input("tolerations")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.TolerationArgs>? Tolerations { get; set; }

        [Input("topologySpreadConstraints")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TopologySpreadConstraintArgs>? _topologySpreadConstraints;

        /// <summary>
        /// Topology spread constraints rely on node labels to identify the topology domain(s) that each Node is in. Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TopologySpreadConstraintArgs> TopologySpreadConstraints
        {
            get => _topologySpreadConstraints ?? (_topologySpreadConstraints = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TopologySpreadConstraintArgs>());
            set => _topologySpreadConstraints = value;
        }

        [Input("udp")]
        public Input<Inputs.ControllerUdpArgs>? Udp { get; set; }

        /// <summary>
        /// The update strategy to apply to the Deployment or DaemonSet.
        /// </summary>
        [Input("updateStrategy")]
        public Input<Inputs.ControllerUpdateStrategyArgs>? UpdateStrategy { get; set; }

        /// <summary>
        /// Process Ingress objects without ingressClass annotation/ingressClassName field. Overrides value for --watch-ingress-without-class flag of the controller binary. Defaults to false.
        /// </summary>
        [Input("watchIngressWithoutClass")]
        public Input<bool>? WatchIngressWithoutClass { get; set; }

        public ControllerArgs()
        {
        }
        public static new ControllerArgs Empty => new ControllerArgs();
    }
}
