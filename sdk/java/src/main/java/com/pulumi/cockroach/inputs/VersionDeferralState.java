// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VersionDeferralState extends com.pulumi.resources.ResourceArgs {

    public static final VersionDeferralState Empty = new VersionDeferralState();

    /**
     * The policy for managing automated minor version upgrades. Set to FIXED*DEFERRAL to defer upgrades by 60 days or NOT*DEFERRED to apply upgrades immediately.
     * 
     */
    @Import(name="deferralPolicy")
    private @Nullable Output<String> deferralPolicy;

    /**
     * @return The policy for managing automated minor version upgrades. Set to FIXED*DEFERRAL to defer upgrades by 60 days or NOT*DEFERRED to apply upgrades immediately.
     * 
     */
    public Optional<Output<String>> deferralPolicy() {
        return Optional.ofNullable(this.deferralPolicy);
    }

    private VersionDeferralState() {}

    private VersionDeferralState(VersionDeferralState $) {
        this.deferralPolicy = $.deferralPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VersionDeferralState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VersionDeferralState $;

        public Builder() {
            $ = new VersionDeferralState();
        }

        public Builder(VersionDeferralState defaults) {
            $ = new VersionDeferralState(Objects.requireNonNull(defaults));
        }

        /**
         * @param deferralPolicy The policy for managing automated minor version upgrades. Set to FIXED*DEFERRAL to defer upgrades by 60 days or NOT*DEFERRED to apply upgrades immediately.
         * 
         * @return builder
         * 
         */
        public Builder deferralPolicy(@Nullable Output<String> deferralPolicy) {
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

        public VersionDeferralState build() {
            return $;
        }
    }

}
