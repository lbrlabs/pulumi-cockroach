// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCockroachClusterDedicated {
    private Integer diskIops;
    private String machineType;
    private Double memoryGib;
    private Integer numVirtualCpus;
    private Boolean privateNetworkVisibility;
    private Integer storageGib;

    private GetCockroachClusterDedicated() {}
    public Integer diskIops() {
        return this.diskIops;
    }
    public String machineType() {
        return this.machineType;
    }
    public Double memoryGib() {
        return this.memoryGib;
    }
    public Integer numVirtualCpus() {
        return this.numVirtualCpus;
    }
    public Boolean privateNetworkVisibility() {
        return this.privateNetworkVisibility;
    }
    public Integer storageGib() {
        return this.storageGib;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCockroachClusterDedicated defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer diskIops;
        private String machineType;
        private Double memoryGib;
        private Integer numVirtualCpus;
        private Boolean privateNetworkVisibility;
        private Integer storageGib;
        public Builder() {}
        public Builder(GetCockroachClusterDedicated defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.diskIops = defaults.diskIops;
    	      this.machineType = defaults.machineType;
    	      this.memoryGib = defaults.memoryGib;
    	      this.numVirtualCpus = defaults.numVirtualCpus;
    	      this.privateNetworkVisibility = defaults.privateNetworkVisibility;
    	      this.storageGib = defaults.storageGib;
        }

        @CustomType.Setter
        public Builder diskIops(Integer diskIops) {
            this.diskIops = Objects.requireNonNull(diskIops);
            return this;
        }
        @CustomType.Setter
        public Builder machineType(String machineType) {
            this.machineType = Objects.requireNonNull(machineType);
            return this;
        }
        @CustomType.Setter
        public Builder memoryGib(Double memoryGib) {
            this.memoryGib = Objects.requireNonNull(memoryGib);
            return this;
        }
        @CustomType.Setter
        public Builder numVirtualCpus(Integer numVirtualCpus) {
            this.numVirtualCpus = Objects.requireNonNull(numVirtualCpus);
            return this;
        }
        @CustomType.Setter
        public Builder privateNetworkVisibility(Boolean privateNetworkVisibility) {
            this.privateNetworkVisibility = Objects.requireNonNull(privateNetworkVisibility);
            return this;
        }
        @CustomType.Setter
        public Builder storageGib(Integer storageGib) {
            this.storageGib = Objects.requireNonNull(storageGib);
            return this;
        }
        public GetCockroachClusterDedicated build() {
            final var o = new GetCockroachClusterDedicated();
            o.diskIops = diskIops;
            o.machineType = machineType;
            o.memoryGib = memoryGib;
            o.numVirtualCpus = numVirtualCpus;
            o.privateNetworkVisibility = privateNetworkVisibility;
            o.storageGib = storageGib;
            return o;
        }
    }
}
