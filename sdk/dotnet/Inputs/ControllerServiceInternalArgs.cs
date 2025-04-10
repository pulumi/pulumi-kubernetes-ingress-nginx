// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesIngressNginx.Inputs
{

    public sealed class ControllerServiceInternalArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Set external traffic policy to: "Local" to preserve source IP on providers supporting it. Ref: https://kubernetes.io/docs/tutorials/services/source-ip/#source-ip-for-services-with-typeloadbalancer
        /// </summary>
        [Input("externalTrafficPolicy")]
        public Input<string>? ExternalTrafficPolicy { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("loadBalancerIPs")]
        public Input<string>? LoadBalancerIPs { get; set; }

        [Input("loadBalancerSourceRanges")]
        private InputList<string>? _loadBalancerSourceRanges;

        /// <summary>
        /// Restrict access For LoadBalancer service. Defaults to 0.0.0.0/0.
        /// </summary>
        public InputList<string> LoadBalancerSourceRanges
        {
            get => _loadBalancerSourceRanges ?? (_loadBalancerSourceRanges = new InputList<string>());
            set => _loadBalancerSourceRanges = value;
        }

        public ControllerServiceInternalArgs()
        {
        }
        public static new ControllerServiceInternalArgs Empty => new ControllerServiceInternalArgs();
    }
}
