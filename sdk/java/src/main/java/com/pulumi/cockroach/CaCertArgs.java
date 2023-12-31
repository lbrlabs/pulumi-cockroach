// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class CaCertArgs extends com.pulumi.resources.ResourceArgs {

    public static final CaCertArgs Empty = new CaCertArgs();

    /**
     * Cluster ID.
     * 
     */
    @Import(name="clusterId", required=true)
    private Output<String> clusterId;

    /**
     * @return Cluster ID.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }

    /**
     * X509 certificate in PEM format.
     * 
     */
    @Import(name="x509PemCert", required=true)
    private Output<String> x509PemCert;

    /**
     * @return X509 certificate in PEM format.
     * 
     */
    public Output<String> x509PemCert() {
        return this.x509PemCert;
    }

    private CaCertArgs() {}

    private CaCertArgs(CaCertArgs $) {
        this.clusterId = $.clusterId;
        this.x509PemCert = $.x509PemCert;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CaCertArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CaCertArgs $;

        public Builder() {
            $ = new CaCertArgs();
        }

        public Builder(CaCertArgs defaults) {
            $ = new CaCertArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param x509PemCert X509 certificate in PEM format.
         * 
         * @return builder
         * 
         */
        public Builder x509PemCert(Output<String> x509PemCert) {
            $.x509PemCert = x509PemCert;
            return this;
        }

        /**
         * @param x509PemCert X509 certificate in PEM format.
         * 
         * @return builder
         * 
         */
        public Builder x509PemCert(String x509PemCert) {
            return x509PemCert(Output.of(x509PemCert));
        }

        public CaCertArgs build() {
            $.clusterId = Objects.requireNonNull($.clusterId, "expected parameter 'clusterId' to be non-null");
            $.x509PemCert = Objects.requireNonNull($.x509PemCert, "expected parameter 'x509PemCert' to be non-null");
            return $;
        }
    }

}
