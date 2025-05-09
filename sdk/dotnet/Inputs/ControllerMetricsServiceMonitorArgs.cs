// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesIngressNginx.Inputs
{

    public sealed class ControllerMetricsServiceMonitorArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalLabels")]
        private InputMap<string>? _additionalLabels;
        public InputMap<string> AdditionalLabels
        {
            get => _additionalLabels ?? (_additionalLabels = new InputMap<string>());
            set => _additionalLabels = value;
        }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("honorLabels")]
        public Input<bool>? HonorLabels { get; set; }

        /// <summary>
        /// The label to use to retrieve the job name from.
        /// </summary>
        [Input("jobLabel")]
        public Input<string>? JobLabel { get; set; }

        [Input("metricRelabelings")]
        private InputList<string>? _metricRelabelings;
        public InputList<string> MetricRelabelings
        {
            get => _metricRelabelings ?? (_metricRelabelings = new InputList<string>());
            set => _metricRelabelings = value;
        }

        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("namespaceSelector")]
        private InputMap<ImmutableDictionary<string, string>>? _namespaceSelector;
        public InputMap<ImmutableDictionary<string, string>> NamespaceSelector
        {
            get => _namespaceSelector ?? (_namespaceSelector = new InputMap<ImmutableDictionary<string, string>>());
            set => _namespaceSelector = value;
        }

        [Input("scrapeInterval")]
        public Input<string>? ScrapeInterval { get; set; }

        [Input("targetLabels")]
        private InputList<string>? _targetLabels;
        public InputList<string> TargetLabels
        {
            get => _targetLabels ?? (_targetLabels = new InputList<string>());
            set => _targetLabels = value;
        }

        public ControllerMetricsServiceMonitorArgs()
        {
        }
        public static new ControllerMetricsServiceMonitorArgs Empty => new ControllerMetricsServiceMonitorArgs();
    }
}
