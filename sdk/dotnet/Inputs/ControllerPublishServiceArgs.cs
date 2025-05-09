// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesIngressNginx.Inputs
{

    public sealed class ControllerPublishServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Allows overriding of the publish service to bind to. Must be &lt;namespace&gt;/&lt;service_name&gt;.
        /// </summary>
        [Input("pathOverride")]
        public Input<string>? PathOverride { get; set; }

        public ControllerPublishServiceArgs()
        {
        }
        public static new ControllerPublishServiceArgs Empty => new ControllerPublishServiceArgs();
    }
}
