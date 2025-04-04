// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetesingressnginx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetesingressnginx.inputs.ControllerPortArgs;
import com.pulumi.kubernetesingressnginx.inputs.ControllerServiceInternalArgs;
import com.pulumi.kubernetesingressnginx.inputs.ControllerServiceNodePortsArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ControllerServiceArgs extends com.pulumi.resources.ResourceArgs {

    public static final ControllerServiceArgs Empty = new ControllerServiceArgs();

    @Import(name="annotations")
    private @Nullable Output<Map<String,String>> annotations;

    public Optional<Output<Map<String,String>>> annotations() {
        return Optional.ofNullable(this.annotations);
    }

    @Import(name="clusterIP")
    private @Nullable Output<String> clusterIP;

    public Optional<Output<String>> clusterIP() {
        return Optional.ofNullable(this.clusterIP);
    }

    @Import(name="enableHttp")
    private @Nullable Output<Boolean> enableHttp;

    public Optional<Output<Boolean>> enableHttp() {
        return Optional.ofNullable(this.enableHttp);
    }

    @Import(name="enableHttps")
    private @Nullable Output<Boolean> enableHttps;

    public Optional<Output<Boolean>> enableHttps() {
        return Optional.ofNullable(this.enableHttps);
    }

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * List of IP addresses at which the controller services are available Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
     * 
     */
    @Import(name="externalIPs")
    private @Nullable Output<List<String>> externalIPs;

    /**
     * @return List of IP addresses at which the controller services are available Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
     * 
     */
    public Optional<Output<List<String>>> externalIPs() {
        return Optional.ofNullable(this.externalIPs);
    }

    /**
     * Set external traffic policy to: &#34;Local&#34; to preserve source IP on providers supporting it. Ref: https://kubernetes.io/docs/tutorials/services/source-ip/#source-ip-for-services-with-typeloadbalancer
     * 
     */
    @Import(name="externalTrafficPolicy")
    private @Nullable Output<String> externalTrafficPolicy;

    /**
     * @return Set external traffic policy to: &#34;Local&#34; to preserve source IP on providers supporting it. Ref: https://kubernetes.io/docs/tutorials/services/source-ip/#source-ip-for-services-with-typeloadbalancer
     * 
     */
    public Optional<Output<String>> externalTrafficPolicy() {
        return Optional.ofNullable(this.externalTrafficPolicy);
    }

    /**
     * specifies the health check node port (numeric port number) for the service. If healthCheckNodePort isn’t specified, the service controller allocates a port from your cluster’s NodePort range. Ref: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#preserving-the-client-source-ip
     * 
     */
    @Import(name="healthCheckNodePort")
    private @Nullable Output<Integer> healthCheckNodePort;

    /**
     * @return specifies the health check node port (numeric port number) for the service. If healthCheckNodePort isn’t specified, the service controller allocates a port from your cluster’s NodePort range. Ref: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#preserving-the-client-source-ip
     * 
     */
    public Optional<Output<Integer>> healthCheckNodePort() {
        return Optional.ofNullable(this.healthCheckNodePort);
    }

    /**
     * Enables an additional internal load balancer (besides the external one). Annotations are mandatory for the load balancer to come up. Varies with the cloud service.
     * 
     */
    @Import(name="internal")
    private @Nullable Output<ControllerServiceInternalArgs> internal;

    /**
     * @return Enables an additional internal load balancer (besides the external one). Annotations are mandatory for the load balancer to come up. Varies with the cloud service.
     * 
     */
    public Optional<Output<ControllerServiceInternalArgs>> internal() {
        return Optional.ofNullable(this.internal);
    }

    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    @Import(name="loadBalancerIP")
    private @Nullable Output<String> loadBalancerIP;

    public Optional<Output<String>> loadBalancerIP() {
        return Optional.ofNullable(this.loadBalancerIP);
    }

    @Import(name="loadBalancerIPs")
    private @Nullable Output<String> loadBalancerIPs;

    public Optional<Output<String>> loadBalancerIPs() {
        return Optional.ofNullable(this.loadBalancerIPs);
    }

    @Import(name="loadBalancerSourceRanges")
    private @Nullable Output<List<String>> loadBalancerSourceRanges;

    public Optional<Output<List<String>>> loadBalancerSourceRanges() {
        return Optional.ofNullable(this.loadBalancerSourceRanges);
    }

    @Import(name="nodePorts")
    private @Nullable Output<ControllerServiceNodePortsArgs> nodePorts;

    public Optional<Output<ControllerServiceNodePortsArgs>> nodePorts() {
        return Optional.ofNullable(this.nodePorts);
    }

    @Import(name="ports")
    private @Nullable Output<ControllerPortArgs> ports;

    public Optional<Output<ControllerPortArgs>> ports() {
        return Optional.ofNullable(this.ports);
    }

    /**
     * Must be either &#34;None&#34; or &#34;ClientIP&#34; if set. Kubernetes will default to &#34;None&#34;. Ref: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
     * 
     */
    @Import(name="sessionAffinity")
    private @Nullable Output<String> sessionAffinity;

    /**
     * @return Must be either &#34;None&#34; or &#34;ClientIP&#34; if set. Kubernetes will default to &#34;None&#34;. Ref: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
     * 
     */
    public Optional<Output<String>> sessionAffinity() {
        return Optional.ofNullable(this.sessionAffinity);
    }

    @Import(name="targetPorts")
    private @Nullable Output<ControllerPortArgs> targetPorts;

    public Optional<Output<ControllerPortArgs>> targetPorts() {
        return Optional.ofNullable(this.targetPorts);
    }

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private ControllerServiceArgs() {}

    private ControllerServiceArgs(ControllerServiceArgs $) {
        this.annotations = $.annotations;
        this.clusterIP = $.clusterIP;
        this.enableHttp = $.enableHttp;
        this.enableHttps = $.enableHttps;
        this.enabled = $.enabled;
        this.externalIPs = $.externalIPs;
        this.externalTrafficPolicy = $.externalTrafficPolicy;
        this.healthCheckNodePort = $.healthCheckNodePort;
        this.internal = $.internal;
        this.labels = $.labels;
        this.loadBalancerIP = $.loadBalancerIP;
        this.loadBalancerIPs = $.loadBalancerIPs;
        this.loadBalancerSourceRanges = $.loadBalancerSourceRanges;
        this.nodePorts = $.nodePorts;
        this.ports = $.ports;
        this.sessionAffinity = $.sessionAffinity;
        this.targetPorts = $.targetPorts;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ControllerServiceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ControllerServiceArgs $;

        public Builder() {
            $ = new ControllerServiceArgs();
        }

        public Builder(ControllerServiceArgs defaults) {
            $ = new ControllerServiceArgs(Objects.requireNonNull(defaults));
        }

        public Builder annotations(@Nullable Output<Map<String,String>> annotations) {
            $.annotations = annotations;
            return this;
        }

        public Builder annotations(Map<String,String> annotations) {
            return annotations(Output.of(annotations));
        }

        public Builder clusterIP(@Nullable Output<String> clusterIP) {
            $.clusterIP = clusterIP;
            return this;
        }

        public Builder clusterIP(String clusterIP) {
            return clusterIP(Output.of(clusterIP));
        }

        public Builder enableHttp(@Nullable Output<Boolean> enableHttp) {
            $.enableHttp = enableHttp;
            return this;
        }

        public Builder enableHttp(Boolean enableHttp) {
            return enableHttp(Output.of(enableHttp));
        }

        public Builder enableHttps(@Nullable Output<Boolean> enableHttps) {
            $.enableHttps = enableHttps;
            return this;
        }

        public Builder enableHttps(Boolean enableHttps) {
            return enableHttps(Output.of(enableHttps));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param externalIPs List of IP addresses at which the controller services are available Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
         * 
         * @return builder
         * 
         */
        public Builder externalIPs(@Nullable Output<List<String>> externalIPs) {
            $.externalIPs = externalIPs;
            return this;
        }

        /**
         * @param externalIPs List of IP addresses at which the controller services are available Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
         * 
         * @return builder
         * 
         */
        public Builder externalIPs(List<String> externalIPs) {
            return externalIPs(Output.of(externalIPs));
        }

        /**
         * @param externalIPs List of IP addresses at which the controller services are available Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
         * 
         * @return builder
         * 
         */
        public Builder externalIPs(String... externalIPs) {
            return externalIPs(List.of(externalIPs));
        }

        /**
         * @param externalTrafficPolicy Set external traffic policy to: &#34;Local&#34; to preserve source IP on providers supporting it. Ref: https://kubernetes.io/docs/tutorials/services/source-ip/#source-ip-for-services-with-typeloadbalancer
         * 
         * @return builder
         * 
         */
        public Builder externalTrafficPolicy(@Nullable Output<String> externalTrafficPolicy) {
            $.externalTrafficPolicy = externalTrafficPolicy;
            return this;
        }

        /**
         * @param externalTrafficPolicy Set external traffic policy to: &#34;Local&#34; to preserve source IP on providers supporting it. Ref: https://kubernetes.io/docs/tutorials/services/source-ip/#source-ip-for-services-with-typeloadbalancer
         * 
         * @return builder
         * 
         */
        public Builder externalTrafficPolicy(String externalTrafficPolicy) {
            return externalTrafficPolicy(Output.of(externalTrafficPolicy));
        }

        /**
         * @param healthCheckNodePort specifies the health check node port (numeric port number) for the service. If healthCheckNodePort isn’t specified, the service controller allocates a port from your cluster’s NodePort range. Ref: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#preserving-the-client-source-ip
         * 
         * @return builder
         * 
         */
        public Builder healthCheckNodePort(@Nullable Output<Integer> healthCheckNodePort) {
            $.healthCheckNodePort = healthCheckNodePort;
            return this;
        }

        /**
         * @param healthCheckNodePort specifies the health check node port (numeric port number) for the service. If healthCheckNodePort isn’t specified, the service controller allocates a port from your cluster’s NodePort range. Ref: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#preserving-the-client-source-ip
         * 
         * @return builder
         * 
         */
        public Builder healthCheckNodePort(Integer healthCheckNodePort) {
            return healthCheckNodePort(Output.of(healthCheckNodePort));
        }

        /**
         * @param internal Enables an additional internal load balancer (besides the external one). Annotations are mandatory for the load balancer to come up. Varies with the cloud service.
         * 
         * @return builder
         * 
         */
        public Builder internal(@Nullable Output<ControllerServiceInternalArgs> internal) {
            $.internal = internal;
            return this;
        }

        /**
         * @param internal Enables an additional internal load balancer (besides the external one). Annotations are mandatory for the load balancer to come up. Varies with the cloud service.
         * 
         * @return builder
         * 
         */
        public Builder internal(ControllerServiceInternalArgs internal) {
            return internal(Output.of(internal));
        }

        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        public Builder loadBalancerIP(@Nullable Output<String> loadBalancerIP) {
            $.loadBalancerIP = loadBalancerIP;
            return this;
        }

        public Builder loadBalancerIP(String loadBalancerIP) {
            return loadBalancerIP(Output.of(loadBalancerIP));
        }

        public Builder loadBalancerIPs(@Nullable Output<String> loadBalancerIPs) {
            $.loadBalancerIPs = loadBalancerIPs;
            return this;
        }

        public Builder loadBalancerIPs(String loadBalancerIPs) {
            return loadBalancerIPs(Output.of(loadBalancerIPs));
        }

        public Builder loadBalancerSourceRanges(@Nullable Output<List<String>> loadBalancerSourceRanges) {
            $.loadBalancerSourceRanges = loadBalancerSourceRanges;
            return this;
        }

        public Builder loadBalancerSourceRanges(List<String> loadBalancerSourceRanges) {
            return loadBalancerSourceRanges(Output.of(loadBalancerSourceRanges));
        }

        public Builder loadBalancerSourceRanges(String... loadBalancerSourceRanges) {
            return loadBalancerSourceRanges(List.of(loadBalancerSourceRanges));
        }

        public Builder nodePorts(@Nullable Output<ControllerServiceNodePortsArgs> nodePorts) {
            $.nodePorts = nodePorts;
            return this;
        }

        public Builder nodePorts(ControllerServiceNodePortsArgs nodePorts) {
            return nodePorts(Output.of(nodePorts));
        }

        public Builder ports(@Nullable Output<ControllerPortArgs> ports) {
            $.ports = ports;
            return this;
        }

        public Builder ports(ControllerPortArgs ports) {
            return ports(Output.of(ports));
        }

        /**
         * @param sessionAffinity Must be either &#34;None&#34; or &#34;ClientIP&#34; if set. Kubernetes will default to &#34;None&#34;. Ref: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
         * 
         * @return builder
         * 
         */
        public Builder sessionAffinity(@Nullable Output<String> sessionAffinity) {
            $.sessionAffinity = sessionAffinity;
            return this;
        }

        /**
         * @param sessionAffinity Must be either &#34;None&#34; or &#34;ClientIP&#34; if set. Kubernetes will default to &#34;None&#34;. Ref: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
         * 
         * @return builder
         * 
         */
        public Builder sessionAffinity(String sessionAffinity) {
            return sessionAffinity(Output.of(sessionAffinity));
        }

        public Builder targetPorts(@Nullable Output<ControllerPortArgs> targetPorts) {
            $.targetPorts = targetPorts;
            return this;
        }

        public Builder targetPorts(ControllerPortArgs targetPorts) {
            return targetPorts(Output.of(targetPorts));
        }

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ControllerServiceArgs build() {
            return $;
        }
    }

}
