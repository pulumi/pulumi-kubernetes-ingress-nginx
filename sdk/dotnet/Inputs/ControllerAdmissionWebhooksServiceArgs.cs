// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesIngressNginx.Inputs
{

    public sealed class ControllerAdmissionWebhooksServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        [Input("clusterIP")]
        public Input<string>? ClusterIP { get; set; }

        [Input("externalIPs")]
        private InputList<string>? _externalIPs;
        public InputList<string> ExternalIPs
        {
            get => _externalIPs ?? (_externalIPs = new InputList<string>());
            set => _externalIPs = value;
        }

        [Input("loadBalancerIPs")]
        public Input<string>? LoadBalancerIPs { get; set; }

        [Input("loadBalancerSourceRanges")]
        private InputList<string>? _loadBalancerSourceRanges;
        public InputList<string> LoadBalancerSourceRanges
        {
            get => _loadBalancerSourceRanges ?? (_loadBalancerSourceRanges = new InputList<string>());
            set => _loadBalancerSourceRanges = value;
        }

        [Input("servicePort")]
        public Input<int>? ServicePort { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public ControllerAdmissionWebhooksServiceArgs()
        {
        }
        public static new ControllerAdmissionWebhooksServiceArgs Empty => new ControllerAdmissionWebhooksServiceArgs();
    }
}
