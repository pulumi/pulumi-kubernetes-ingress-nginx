// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetesingressnginx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ControllerPodSecurityPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final ControllerPodSecurityPolicyArgs Empty = new ControllerPodSecurityPolicyArgs();

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    private ControllerPodSecurityPolicyArgs() {}

    private ControllerPodSecurityPolicyArgs(ControllerPodSecurityPolicyArgs $) {
        this.enabled = $.enabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ControllerPodSecurityPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ControllerPodSecurityPolicyArgs $;

        public Builder() {
            $ = new ControllerPodSecurityPolicyArgs();
        }

        public Builder(ControllerPodSecurityPolicyArgs defaults) {
            $ = new ControllerPodSecurityPolicyArgs(Objects.requireNonNull(defaults));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public ControllerPodSecurityPolicyArgs build() {
            return $;
        }
    }

}
