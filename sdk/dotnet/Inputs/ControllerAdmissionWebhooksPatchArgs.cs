// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesIngressNginx.Inputs
{

    public sealed class ControllerAdmissionWebhooksPatchArgs : global::Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("image")]
        public Input<Inputs.ControllerImageArgs>? Image { get; set; }

        [Input("nodeSelector")]
        private InputMap<string>? _nodeSelector;
        public InputMap<string> NodeSelector
        {
            get => _nodeSelector ?? (_nodeSelector = new InputMap<string>());
            set => _nodeSelector = value;
        }

        [Input("podAnnotations")]
        private InputMap<string>? _podAnnotations;
        public InputMap<string> PodAnnotations
        {
            get => _podAnnotations ?? (_podAnnotations = new InputMap<string>());
            set => _podAnnotations = value;
        }

        /// <summary>
        /// Provide a priority class name to the webhook patching job.
        /// </summary>
        [Input("priorityClassName")]
        public Input<string>? PriorityClassName { get; set; }

        [Input("runAsUser")]
        public Input<int>? RunAsUser { get; set; }

        [Input("tolerations")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TolerationArgs>? _tolerations;
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TolerationArgs> Tolerations
        {
            get => _tolerations ?? (_tolerations = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TolerationArgs>());
            set => _tolerations = value;
        }

        public ControllerAdmissionWebhooksPatchArgs()
        {
        }
        public static new ControllerAdmissionWebhooksPatchArgs Empty => new ControllerAdmissionWebhooksPatchArgs();
    }
}
