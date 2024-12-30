// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetesingressnginx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ControllerAdmissionWebhooksServiceArgs extends com.pulumi.resources.ResourceArgs {

    public static final ControllerAdmissionWebhooksServiceArgs Empty = new ControllerAdmissionWebhooksServiceArgs();

    @Import(name="annotations")
    private @Nullable Output<Map<String,Map<String,String>>> annotations;

    public Optional<Output<Map<String,Map<String,String>>>> annotations() {
        return Optional.ofNullable(this.annotations);
    }

    @Import(name="clusterIP")
    private @Nullable Output<String> clusterIP;

    public Optional<Output<String>> clusterIP() {
        return Optional.ofNullable(this.clusterIP);
    }

    @Import(name="externalIPs")
    private @Nullable Output<List<String>> externalIPs;

    public Optional<Output<List<String>>> externalIPs() {
        return Optional.ofNullable(this.externalIPs);
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

    @Import(name="servicePort")
    private @Nullable Output<Integer> servicePort;

    public Optional<Output<Integer>> servicePort() {
        return Optional.ofNullable(this.servicePort);
    }

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private ControllerAdmissionWebhooksServiceArgs() {}

    private ControllerAdmissionWebhooksServiceArgs(ControllerAdmissionWebhooksServiceArgs $) {
        this.annotations = $.annotations;
        this.clusterIP = $.clusterIP;
        this.externalIPs = $.externalIPs;
        this.loadBalancerIPs = $.loadBalancerIPs;
        this.loadBalancerSourceRanges = $.loadBalancerSourceRanges;
        this.servicePort = $.servicePort;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ControllerAdmissionWebhooksServiceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ControllerAdmissionWebhooksServiceArgs $;

        public Builder() {
            $ = new ControllerAdmissionWebhooksServiceArgs();
        }

        public Builder(ControllerAdmissionWebhooksServiceArgs defaults) {
            $ = new ControllerAdmissionWebhooksServiceArgs(Objects.requireNonNull(defaults));
        }

        public Builder annotations(@Nullable Output<Map<String,Map<String,String>>> annotations) {
            $.annotations = annotations;
            return this;
        }

        public Builder annotations(Map<String,Map<String,String>> annotations) {
            return annotations(Output.of(annotations));
        }

        public Builder clusterIP(@Nullable Output<String> clusterIP) {
            $.clusterIP = clusterIP;
            return this;
        }

        public Builder clusterIP(String clusterIP) {
            return clusterIP(Output.of(clusterIP));
        }

        public Builder externalIPs(@Nullable Output<List<String>> externalIPs) {
            $.externalIPs = externalIPs;
            return this;
        }

        public Builder externalIPs(List<String> externalIPs) {
            return externalIPs(Output.of(externalIPs));
        }

        public Builder externalIPs(String... externalIPs) {
            return externalIPs(List.of(externalIPs));
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

        public Builder servicePort(@Nullable Output<Integer> servicePort) {
            $.servicePort = servicePort;
            return this;
        }

        public Builder servicePort(Integer servicePort) {
            return servicePort(Output.of(servicePort));
        }

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ControllerAdmissionWebhooksServiceArgs build() {
            return $;
        }
    }

}
