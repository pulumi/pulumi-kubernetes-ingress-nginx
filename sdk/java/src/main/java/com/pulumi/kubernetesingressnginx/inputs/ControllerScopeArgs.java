// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetesingressnginx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ControllerScopeArgs extends com.pulumi.resources.ResourceArgs {

    public static final ControllerScopeArgs Empty = new ControllerScopeArgs();

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    private ControllerScopeArgs() {}

    private ControllerScopeArgs(ControllerScopeArgs $) {
        this.enabled = $.enabled;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ControllerScopeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ControllerScopeArgs $;

        public Builder() {
            $ = new ControllerScopeArgs();
        }

        public Builder(ControllerScopeArgs defaults) {
            $ = new ControllerScopeArgs(Objects.requireNonNull(defaults));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public ControllerScopeArgs build() {
            return $;
        }
    }

}
