// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetesingressnginx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ReleaseStatus {
    /**
     * @return The version number of the application being deployed.
     * 
     */
    private String appVersion;
    /**
     * @return The name of the chart.
     * 
     */
    private String chart;
    /**
     * @return Name is the name of the release.
     * 
     */
    private String name;
    /**
     * @return Namespace is the kubernetes namespace of the release.
     * 
     */
    private String namespace;
    /**
     * @return Version is an int32 which represents the version of the release.
     * 
     */
    private Integer revision;
    /**
     * @return Status of the release.
     * 
     */
    private String status;
    /**
     * @return A SemVer 2 conformant version string of the chart.
     * 
     */
    private String version;

    private ReleaseStatus() {}
    /**
     * @return The version number of the application being deployed.
     * 
     */
    public String appVersion() {
        return this.appVersion;
    }
    /**
     * @return The name of the chart.
     * 
     */
    public String chart() {
        return this.chart;
    }
    /**
     * @return Name is the name of the release.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Namespace is the kubernetes namespace of the release.
     * 
     */
    public String namespace() {
        return this.namespace;
    }
    /**
     * @return Version is an int32 which represents the version of the release.
     * 
     */
    public Integer revision() {
        return this.revision;
    }
    /**
     * @return Status of the release.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return A SemVer 2 conformant version string of the chart.
     * 
     */
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ReleaseStatus defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String appVersion;
        private String chart;
        private String name;
        private String namespace;
        private Integer revision;
        private String status;
        private String version;
        public Builder() {}
        public Builder(ReleaseStatus defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.appVersion = defaults.appVersion;
    	      this.chart = defaults.chart;
    	      this.name = defaults.name;
    	      this.namespace = defaults.namespace;
    	      this.revision = defaults.revision;
    	      this.status = defaults.status;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder appVersion(String appVersion) {
            if (appVersion == null) {
              throw new MissingRequiredPropertyException("ReleaseStatus", "appVersion");
            }
            this.appVersion = appVersion;
            return this;
        }
        @CustomType.Setter
        public Builder chart(String chart) {
            if (chart == null) {
              throw new MissingRequiredPropertyException("ReleaseStatus", "chart");
            }
            this.chart = chart;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("ReleaseStatus", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(String namespace) {
            if (namespace == null) {
              throw new MissingRequiredPropertyException("ReleaseStatus", "namespace");
            }
            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder revision(Integer revision) {
            if (revision == null) {
              throw new MissingRequiredPropertyException("ReleaseStatus", "revision");
            }
            this.revision = revision;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("ReleaseStatus", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            if (version == null) {
              throw new MissingRequiredPropertyException("ReleaseStatus", "version");
            }
            this.version = version;
            return this;
        }
        public ReleaseStatus build() {
            final var _resultValue = new ReleaseStatus();
            _resultValue.appVersion = appVersion;
            _resultValue.chart = chart;
            _resultValue.name = name;
            _resultValue.namespace = namespace;
            _resultValue.revision = revision;
            _resultValue.status = status;
            _resultValue.version = version;
            return _resultValue;
        }
    }
}
