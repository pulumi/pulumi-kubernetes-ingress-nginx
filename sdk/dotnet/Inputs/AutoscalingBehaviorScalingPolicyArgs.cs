// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesIngressNginx.Inputs
{

    public sealed class AutoscalingBehaviorScalingPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("periodSeconds")]
        public Input<int>? PeriodSeconds { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("value")]
        public Input<int>? Value { get; set; }

        public AutoscalingBehaviorScalingPolicyArgs()
        {
        }
        public static new AutoscalingBehaviorScalingPolicyArgs Empty => new AutoscalingBehaviorScalingPolicyArgs();
    }
}
