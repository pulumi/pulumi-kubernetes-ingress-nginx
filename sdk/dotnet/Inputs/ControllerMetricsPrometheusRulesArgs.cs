// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesIngressNginx.Inputs
{

    public sealed class ControllerMetricsPrometheusRulesArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalLabels")]
        private InputMap<ImmutableDictionary<string, string>>? _additionalLabels;
        public InputMap<ImmutableDictionary<string, string>> AdditionalLabels
        {
            get => _additionalLabels ?? (_additionalLabels = new InputMap<ImmutableDictionary<string, string>>());
            set => _additionalLabels = value;
        }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("rules")]
        private InputList<ImmutableDictionary<string, string>>? _rules;
        public InputList<ImmutableDictionary<string, string>> Rules
        {
            get => _rules ?? (_rules = new InputList<ImmutableDictionary<string, string>>());
            set => _rules = value;
        }

        public ControllerMetricsPrometheusRulesArgs()
        {
        }
        public static new ControllerMetricsPrometheusRulesArgs Empty => new ControllerMetricsPrometheusRulesArgs();
    }
}
