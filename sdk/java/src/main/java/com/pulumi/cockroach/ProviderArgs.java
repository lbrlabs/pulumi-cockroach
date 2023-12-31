// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * apikey to access cockroach cloud
     * 
     */
    @Import(name="apikey")
    private @Nullable Output<String> apikey;

    /**
     * @return apikey to access cockroach cloud
     * 
     */
    public Optional<Output<String>> apikey() {
        return Optional.ofNullable(this.apikey);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.apikey = $.apikey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apikey apikey to access cockroach cloud
         * 
         * @return builder
         * 
         */
        public Builder apikey(@Nullable Output<String> apikey) {
            $.apikey = apikey;
            return this;
        }

        /**
         * @param apikey apikey to access cockroach cloud
         * 
         * @return builder
         * 
         */
        public Builder apikey(String apikey) {
            return apikey(Output.of(apikey));
        }

        public ProviderArgs build() {
            return $;
        }
    }

}
