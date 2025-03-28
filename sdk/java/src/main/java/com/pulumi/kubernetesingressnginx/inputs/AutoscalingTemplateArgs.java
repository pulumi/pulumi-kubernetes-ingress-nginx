// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetesingressnginx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetesingressnginx.inputs.AutoscalingTemplatePodsArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AutoscalingTemplateArgs extends com.pulumi.resources.ResourceArgs {

    public static final AutoscalingTemplateArgs Empty = new AutoscalingTemplateArgs();

    @Import(name="pods")
    private @Nullable Output<AutoscalingTemplatePodsArgs> pods;

    public Optional<Output<AutoscalingTemplatePodsArgs>> pods() {
        return Optional.ofNullable(this.pods);
    }

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private AutoscalingTemplateArgs() {}

    private AutoscalingTemplateArgs(AutoscalingTemplateArgs $) {
        this.pods = $.pods;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AutoscalingTemplateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AutoscalingTemplateArgs $;

        public Builder() {
            $ = new AutoscalingTemplateArgs();
        }

        public Builder(AutoscalingTemplateArgs defaults) {
            $ = new AutoscalingTemplateArgs(Objects.requireNonNull(defaults));
        }

        public Builder pods(@Nullable Output<AutoscalingTemplatePodsArgs> pods) {
            $.pods = pods;
            return this;
        }

        public Builder pods(AutoscalingTemplatePodsArgs pods) {
            return pods(Output.of(pods));
        }

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public AutoscalingTemplateArgs build() {
            return $;
        }
    }

}
