// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach.inputs;

import com.pulumi.cockroach.inputs.ClusterServerlessUsageLimitsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterServerlessArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterServerlessArgs Empty = new ClusterServerlessArgs();

    @Import(name="routingId")
    private @Nullable Output<String> routingId;

    public Optional<Output<String>> routingId() {
        return Optional.ofNullable(this.routingId);
    }

    @Import(name="spendLimit")
    private @Nullable Output<Integer> spendLimit;

    public Optional<Output<Integer>> spendLimit() {
        return Optional.ofNullable(this.spendLimit);
    }

    @Import(name="usageLimits")
    private @Nullable Output<ClusterServerlessUsageLimitsArgs> usageLimits;

    public Optional<Output<ClusterServerlessUsageLimitsArgs>> usageLimits() {
        return Optional.ofNullable(this.usageLimits);
    }

    private ClusterServerlessArgs() {}

    private ClusterServerlessArgs(ClusterServerlessArgs $) {
        this.routingId = $.routingId;
        this.spendLimit = $.spendLimit;
        this.usageLimits = $.usageLimits;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterServerlessArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterServerlessArgs $;

        public Builder() {
            $ = new ClusterServerlessArgs();
        }

        public Builder(ClusterServerlessArgs defaults) {
            $ = new ClusterServerlessArgs(Objects.requireNonNull(defaults));
        }

        public Builder routingId(@Nullable Output<String> routingId) {
            $.routingId = routingId;
            return this;
        }

        public Builder routingId(String routingId) {
            return routingId(Output.of(routingId));
        }

        public Builder spendLimit(@Nullable Output<Integer> spendLimit) {
            $.spendLimit = spendLimit;
            return this;
        }

        public Builder spendLimit(Integer spendLimit) {
            return spendLimit(Output.of(spendLimit));
        }

        public Builder usageLimits(@Nullable Output<ClusterServerlessUsageLimitsArgs> usageLimits) {
            $.usageLimits = usageLimits;
            return this;
        }

        public Builder usageLimits(ClusterServerlessUsageLimitsArgs usageLimits) {
            return usageLimits(Output.of(usageLimits));
        }

        public ClusterServerlessArgs build() {
            return $;
        }
    }

}
