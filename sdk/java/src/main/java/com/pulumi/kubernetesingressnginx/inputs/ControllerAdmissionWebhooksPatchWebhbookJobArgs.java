// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetesingressnginx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetes.core.v1.inputs.ResourceRequirementsArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ControllerAdmissionWebhooksPatchWebhbookJobArgs extends com.pulumi.resources.ResourceArgs {

    public static final ControllerAdmissionWebhooksPatchWebhbookJobArgs Empty = new ControllerAdmissionWebhooksPatchWebhbookJobArgs();

    @Import(name="resources")
    private @Nullable Output<ResourceRequirementsArgs> resources;

    public Optional<Output<ResourceRequirementsArgs>> resources() {
        return Optional.ofNullable(this.resources);
    }

    private ControllerAdmissionWebhooksPatchWebhbookJobArgs() {}

    private ControllerAdmissionWebhooksPatchWebhbookJobArgs(ControllerAdmissionWebhooksPatchWebhbookJobArgs $) {
        this.resources = $.resources;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ControllerAdmissionWebhooksPatchWebhbookJobArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ControllerAdmissionWebhooksPatchWebhbookJobArgs $;

        public Builder() {
            $ = new ControllerAdmissionWebhooksPatchWebhbookJobArgs();
        }

        public Builder(ControllerAdmissionWebhooksPatchWebhbookJobArgs defaults) {
            $ = new ControllerAdmissionWebhooksPatchWebhbookJobArgs(Objects.requireNonNull(defaults));
        }

        public Builder resources(@Nullable Output<ResourceRequirementsArgs> resources) {
            $.resources = resources;
            return this;
        }

        public Builder resources(ResourceRequirementsArgs resources) {
            return resources(Output.of(resources));
        }

        public ControllerAdmissionWebhooksPatchWebhbookJobArgs build() {
            return $;
        }
    }

}
