// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesIngressNginx.Inputs
{

    public sealed class ControllerMetricsArgs : global::Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// if this port is changed, change healthz-port: in extraArgs: accordingly.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("prometheusRule")]
        public Input<Inputs.ControllerMetricsPrometheusRulesArgs>? PrometheusRule { get; set; }

        [Input("service")]
        public Input<Inputs.ControllerMetricsServiceArgs>? Service { get; set; }

        [Input("serviceMonitor")]
        public Input<Inputs.ControllerMetricsServiceMonitorArgs>? ServiceMonitor { get; set; }

        public ControllerMetricsArgs()
        {
        }
        public static new ControllerMetricsArgs Empty => new ControllerMetricsArgs();
    }
}
