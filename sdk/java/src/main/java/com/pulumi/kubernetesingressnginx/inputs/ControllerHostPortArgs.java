// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetesingressnginx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetesingressnginx.inputs.ControllerHostPortPortsArgs;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ControllerHostPortArgs extends com.pulumi.resources.ResourceArgs {

    public static final ControllerHostPortArgs Empty = new ControllerHostPortArgs();

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="ports")
    private @Nullable Output<ControllerHostPortPortsArgs> ports;

    public Optional<Output<ControllerHostPortPortsArgs>> ports() {
        return Optional.ofNullable(this.ports);
    }

    private ControllerHostPortArgs() {}

    private ControllerHostPortArgs(ControllerHostPortArgs $) {
        this.enabled = $.enabled;
        this.ports = $.ports;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ControllerHostPortArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ControllerHostPortArgs $;

        public Builder() {
            $ = new ControllerHostPortArgs();
        }

        public Builder(ControllerHostPortArgs defaults) {
            $ = new ControllerHostPortArgs(Objects.requireNonNull(defaults));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public Builder ports(@Nullable Output<ControllerHostPortPortsArgs> ports) {
            $.ports = ports;
            return this;
        }

        public Builder ports(ControllerHostPortPortsArgs ports) {
            return ports(Output.of(ports));
        }

        public ControllerHostPortArgs build() {
            return $;
        }
    }

}
