// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetesingressnginx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KedaTriggerArgs extends com.pulumi.resources.ResourceArgs {

    public static final KedaTriggerArgs Empty = new KedaTriggerArgs();

    @Import(name="metadata")
    private @Nullable Output<Map<String,Map<String,String>>> metadata;

    public Optional<Output<Map<String,Map<String,String>>>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private KedaTriggerArgs() {}

    private KedaTriggerArgs(KedaTriggerArgs $) {
        this.metadata = $.metadata;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KedaTriggerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KedaTriggerArgs $;

        public Builder() {
            $ = new KedaTriggerArgs();
        }

        public Builder(KedaTriggerArgs defaults) {
            $ = new KedaTriggerArgs(Objects.requireNonNull(defaults));
        }

        public Builder metadata(@Nullable Output<Map<String,Map<String,String>>> metadata) {
            $.metadata = metadata;
            return this;
        }

        public Builder metadata(Map<String,Map<String,String>> metadata) {
            return metadata(Output.of(metadata));
        }

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public KedaTriggerArgs build() {
            return $;
        }
    }

}