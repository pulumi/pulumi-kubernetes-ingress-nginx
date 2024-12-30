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


public final class ControllerImageArgs extends com.pulumi.resources.ResourceArgs {

    public static final ControllerImageArgs Empty = new ControllerImageArgs();

    @Import(name="allowPrivilegeEscalation")
    private @Nullable Output<Boolean> allowPrivilegeEscalation;

    public Optional<Output<Boolean>> allowPrivilegeEscalation() {
        return Optional.ofNullable(this.allowPrivilegeEscalation);
    }

    @Import(name="digest")
    private @Nullable Output<String> digest;

    public Optional<Output<String>> digest() {
        return Optional.ofNullable(this.digest);
    }

    @Import(name="image")
    private @Nullable Output<String> image;

    public Optional<Output<String>> image() {
        return Optional.ofNullable(this.image);
    }

    @Import(name="pullPolicy")
    private @Nullable Output<String> pullPolicy;

    public Optional<Output<String>> pullPolicy() {
        return Optional.ofNullable(this.pullPolicy);
    }

    @Import(name="readOnlyRootFilesystem")
    private @Nullable Output<Boolean> readOnlyRootFilesystem;

    public Optional<Output<Boolean>> readOnlyRootFilesystem() {
        return Optional.ofNullable(this.readOnlyRootFilesystem);
    }

    @Import(name="registry")
    private @Nullable Output<String> registry;

    public Optional<Output<String>> registry() {
        return Optional.ofNullable(this.registry);
    }

    /**
     * for backwards compatibility consider setting the full image url via the repository value below use *either* current default registry/image or repository format or installing will fail.
     * 
     */
    @Import(name="repository")
    private @Nullable Output<String> repository;

    /**
     * @return for backwards compatibility consider setting the full image url via the repository value below use *either* current default registry/image or repository format or installing will fail.
     * 
     */
    public Optional<Output<String>> repository() {
        return Optional.ofNullable(this.repository);
    }

    @Import(name="runAsNonRoot")
    private @Nullable Output<Boolean> runAsNonRoot;

    public Optional<Output<Boolean>> runAsNonRoot() {
        return Optional.ofNullable(this.runAsNonRoot);
    }

    @Import(name="runAsUser")
    private @Nullable Output<String> runAsUser;

    public Optional<Output<String>> runAsUser() {
        return Optional.ofNullable(this.runAsUser);
    }

    @Import(name="tag")
    private @Nullable Output<String> tag;

    public Optional<Output<String>> tag() {
        return Optional.ofNullable(this.tag);
    }

    private ControllerImageArgs() {}

    private ControllerImageArgs(ControllerImageArgs $) {
        this.allowPrivilegeEscalation = $.allowPrivilegeEscalation;
        this.digest = $.digest;
        this.image = $.image;
        this.pullPolicy = $.pullPolicy;
        this.readOnlyRootFilesystem = $.readOnlyRootFilesystem;
        this.registry = $.registry;
        this.repository = $.repository;
        this.runAsNonRoot = $.runAsNonRoot;
        this.runAsUser = $.runAsUser;
        this.tag = $.tag;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ControllerImageArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ControllerImageArgs $;

        public Builder() {
            $ = new ControllerImageArgs();
        }

        public Builder(ControllerImageArgs defaults) {
            $ = new ControllerImageArgs(Objects.requireNonNull(defaults));
        }

        public Builder allowPrivilegeEscalation(@Nullable Output<Boolean> allowPrivilegeEscalation) {
            $.allowPrivilegeEscalation = allowPrivilegeEscalation;
            return this;
        }

        public Builder allowPrivilegeEscalation(Boolean allowPrivilegeEscalation) {
            return allowPrivilegeEscalation(Output.of(allowPrivilegeEscalation));
        }

        public Builder digest(@Nullable Output<String> digest) {
            $.digest = digest;
            return this;
        }

        public Builder digest(String digest) {
            return digest(Output.of(digest));
        }

        public Builder image(@Nullable Output<String> image) {
            $.image = image;
            return this;
        }

        public Builder image(String image) {
            return image(Output.of(image));
        }

        public Builder pullPolicy(@Nullable Output<String> pullPolicy) {
            $.pullPolicy = pullPolicy;
            return this;
        }

        public Builder pullPolicy(String pullPolicy) {
            return pullPolicy(Output.of(pullPolicy));
        }

        public Builder readOnlyRootFilesystem(@Nullable Output<Boolean> readOnlyRootFilesystem) {
            $.readOnlyRootFilesystem = readOnlyRootFilesystem;
            return this;
        }

        public Builder readOnlyRootFilesystem(Boolean readOnlyRootFilesystem) {
            return readOnlyRootFilesystem(Output.of(readOnlyRootFilesystem));
        }

        public Builder registry(@Nullable Output<String> registry) {
            $.registry = registry;
            return this;
        }

        public Builder registry(String registry) {
            return registry(Output.of(registry));
        }

        /**
         * @param repository for backwards compatibility consider setting the full image url via the repository value below use *either* current default registry/image or repository format or installing will fail.
         * 
         * @return builder
         * 
         */
        public Builder repository(@Nullable Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository for backwards compatibility consider setting the full image url via the repository value below use *either* current default registry/image or repository format or installing will fail.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public Builder runAsNonRoot(@Nullable Output<Boolean> runAsNonRoot) {
            $.runAsNonRoot = runAsNonRoot;
            return this;
        }

        public Builder runAsNonRoot(Boolean runAsNonRoot) {
            return runAsNonRoot(Output.of(runAsNonRoot));
        }

        public Builder runAsUser(@Nullable Output<String> runAsUser) {
            $.runAsUser = runAsUser;
            return this;
        }

        public Builder runAsUser(String runAsUser) {
            return runAsUser(Output.of(runAsUser));
        }

        public Builder tag(@Nullable Output<String> tag) {
            $.tag = tag;
            return this;
        }

        public Builder tag(String tag) {
            return tag(Output.of(tag));
        }

        public ControllerImageArgs build() {
            return $;
        }
    }

}
