// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetesingressnginx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ControllerServiceAccountArgs extends com.pulumi.resources.ResourceArgs {

    public static final ControllerServiceAccountArgs Empty = new ControllerServiceAccountArgs();

    @Import(name="automountServiceAccountToken")
    private @Nullable Output<Boolean> automountServiceAccountToken;

    public Optional<Output<Boolean>> automountServiceAccountToken() {
        return Optional.ofNullable(this.automountServiceAccountToken);
    }

    @Import(name="create")
    private @Nullable Output<Boolean> create;

    public Optional<Output<Boolean>> create() {
        return Optional.ofNullable(this.create);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private ControllerServiceAccountArgs() {}

    private ControllerServiceAccountArgs(ControllerServiceAccountArgs $) {
        this.automountServiceAccountToken = $.automountServiceAccountToken;
        this.create = $.create;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ControllerServiceAccountArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ControllerServiceAccountArgs $;

        public Builder() {
            $ = new ControllerServiceAccountArgs();
        }

        public Builder(ControllerServiceAccountArgs defaults) {
            $ = new ControllerServiceAccountArgs(Objects.requireNonNull(defaults));
        }

        public Builder automountServiceAccountToken(@Nullable Output<Boolean> automountServiceAccountToken) {
            $.automountServiceAccountToken = automountServiceAccountToken;
            return this;
        }

        public Builder automountServiceAccountToken(Boolean automountServiceAccountToken) {
            return automountServiceAccountToken(Output.of(automountServiceAccountToken));
        }

        public Builder create(@Nullable Output<Boolean> create) {
            $.create = create;
            return this;
        }

        public Builder create(Boolean create) {
            return create(Output.of(create));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public ControllerServiceAccountArgs build() {
            return $;
        }
    }

}