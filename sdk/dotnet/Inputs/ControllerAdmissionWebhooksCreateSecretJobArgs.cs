// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesIngressNginx.Inputs
{

    public sealed class ControllerAdmissionWebhooksCreateSecretJobArgs : global::Pulumi.ResourceArgs
    {
        [Input("resources")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs>? Resources { get; set; }

        public ControllerAdmissionWebhooksCreateSecretJobArgs()
        {
        }
        public static new ControllerAdmissionWebhooksCreateSecretJobArgs Empty => new ControllerAdmissionWebhooksCreateSecretJobArgs();
    }
}
