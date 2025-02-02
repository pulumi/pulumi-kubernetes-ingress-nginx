// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetesingressnginx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetesingressnginx.inputs.ControllerMetricsPrometheusRulesArgs;
import com.pulumi.kubernetesingressnginx.inputs.ControllerMetricsServiceArgs;
import com.pulumi.kubernetesingressnginx.inputs.ControllerMetricsServiceMonitorArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ControllerMetricsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ControllerMetricsArgs Empty = new ControllerMetricsArgs();

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * if this port is changed, change healthz-port: in extraArgs: accordingly.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return if this port is changed, change healthz-port: in extraArgs: accordingly.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    @Import(name="prometheusRule")
    private @Nullable Output<ControllerMetricsPrometheusRulesArgs> prometheusRule;

    public Optional<Output<ControllerMetricsPrometheusRulesArgs>> prometheusRule() {
        return Optional.ofNullable(this.prometheusRule);
    }

    @Import(name="service")
    private @Nullable Output<ControllerMetricsServiceArgs> service;

    public Optional<Output<ControllerMetricsServiceArgs>> service() {
        return Optional.ofNullable(this.service);
    }

    @Import(name="serviceMonitor")
    private @Nullable Output<ControllerMetricsServiceMonitorArgs> serviceMonitor;

    public Optional<Output<ControllerMetricsServiceMonitorArgs>> serviceMonitor() {
        return Optional.ofNullable(this.serviceMonitor);
    }

    private ControllerMetricsArgs() {}

    private ControllerMetricsArgs(ControllerMetricsArgs $) {
        this.enabled = $.enabled;
        this.port = $.port;
        this.prometheusRule = $.prometheusRule;
        this.service = $.service;
        this.serviceMonitor = $.serviceMonitor;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ControllerMetricsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ControllerMetricsArgs $;

        public Builder() {
            $ = new ControllerMetricsArgs();
        }

        public Builder(ControllerMetricsArgs defaults) {
            $ = new ControllerMetricsArgs(Objects.requireNonNull(defaults));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param port if this port is changed, change healthz-port: in extraArgs: accordingly.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port if this port is changed, change healthz-port: in extraArgs: accordingly.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        public Builder prometheusRule(@Nullable Output<ControllerMetricsPrometheusRulesArgs> prometheusRule) {
            $.prometheusRule = prometheusRule;
            return this;
        }

        public Builder prometheusRule(ControllerMetricsPrometheusRulesArgs prometheusRule) {
            return prometheusRule(Output.of(prometheusRule));
        }

        public Builder service(@Nullable Output<ControllerMetricsServiceArgs> service) {
            $.service = service;
            return this;
        }

        public Builder service(ControllerMetricsServiceArgs service) {
            return service(Output.of(service));
        }

        public Builder serviceMonitor(@Nullable Output<ControllerMetricsServiceMonitorArgs> serviceMonitor) {
            $.serviceMonitor = serviceMonitor;
            return this;
        }

        public Builder serviceMonitor(ControllerMetricsServiceMonitorArgs serviceMonitor) {
            return serviceMonitor(Output.of(serviceMonitor));
        }

        public ControllerMetricsArgs build() {
            return $;
        }
    }

}
