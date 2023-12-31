// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class VersionDeferralArgs extends com.pulumi.resources.ResourceArgs {

    public static final VersionDeferralArgs Empty = new VersionDeferralArgs();

    /**
     * The policy for managing automated minor version upgrades. Set to FIXED*DEFERRAL to defer upgrades by 60 days or NOT*DEFERRED to apply upgrades immediately.
     * 
     */
    @Import(name="deferralPolicy", required=true)
    private Output<String> deferralPolicy;

    /**
     * @return The policy for managing automated minor version upgrades. Set to FIXED*DEFERRAL to defer upgrades by 60 days or NOT*DEFERRED to apply upgrades immediately.
     * 
     */
    public Output<String> deferralPolicy() {
        return this.deferralPolicy;
    }

    private VersionDeferralArgs() {}

    private VersionDeferralArgs(VersionDeferralArgs $) {
        this.deferralPolicy = $.deferralPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VersionDeferralArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VersionDeferralArgs $;

        public Builder() {
            $ = new VersionDeferralArgs();
        }

        public Builder(VersionDeferralArgs defaults) {
            $ = new VersionDeferralArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param deferralPolicy The policy for managing automated minor version upgrades. Set to FIXED*DEFERRAL to defer upgrades by 60 days or NOT*DEFERRED to apply upgrades immediately.
         * 
         * @return builder
         * 
         */
        public Builder deferralPolicy(Output<String> deferralPolicy) {
            $.deferralPolicy = deferralPolicy;
            return this;
        }

        /**
         * @param deferralPolicy The policy for managing automated minor version upgrades. Set to FIXED*DEFERRAL to defer upgrades by 60 days or NOT*DEFERRED to apply upgrades immediately.
         * 
         * @return builder
         * 
         */
        public Builder deferralPolicy(String deferralPolicy) {
            return deferralPolicy(Output.of(deferralPolicy));
        }

        public VersionDeferralArgs build() {
            $.deferralPolicy = Objects.requireNonNull($.deferralPolicy, "expected parameter 'deferralPolicy' to be non-null");
            return $;
        }
    }

}
