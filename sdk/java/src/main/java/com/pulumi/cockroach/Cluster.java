// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach;

import com.pulumi.cockroach.ClusterArgs;
import com.pulumi.cockroach.Utilities;
import com.pulumi.cockroach.inputs.ClusterState;
import com.pulumi.cockroach.outputs.ClusterDedicated;
import com.pulumi.cockroach.outputs.ClusterRegion;
import com.pulumi.cockroach.outputs.ClusterServerless;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * CockroachDB Cloud cluster. Can be Dedicated or Serverless.
 * 
 */
@ResourceType(type="cockroach:index/cluster:Cluster")
public class Cluster extends com.pulumi.resources.CustomResource {
    /**
     * The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
     * 
     */
    @Export(name="accountId", type=String.class, parameters={})
    private Output<String> accountId;

    /**
     * @return The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
     * 
     */
    public Output<String> accountId() {
        return this.accountId;
    }
    /**
     * Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
     * 
     */
    @Export(name="cloudProvider", type=String.class, parameters={})
    private Output<String> cloudProvider;

    /**
     * @return Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
     * 
     */
    public Output<String> cloudProvider() {
        return this.cloudProvider;
    }
    @Export(name="clusterId", type=String.class, parameters={})
    private Output<String> clusterId;

    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * Major version of CockroachDB running on the cluster.
     * 
     */
    @Export(name="cockroachVersion", type=String.class, parameters={})
    private Output<String> cockroachVersion;

    /**
     * @return Major version of CockroachDB running on the cluster.
     * 
     */
    public Output<String> cockroachVersion() {
        return this.cockroachVersion;
    }
    /**
     * ID of the user who created the cluster.
     * 
     */
    @Export(name="creatorId", type=String.class, parameters={})
    private Output<String> creatorId;

    /**
     * @return ID of the user who created the cluster.
     * 
     */
    public Output<String> creatorId() {
        return this.creatorId;
    }
    @Export(name="dedicated", type=ClusterDedicated.class, parameters={})
    private Output</* @Nullable */ ClusterDedicated> dedicated;

    public Output<Optional<ClusterDedicated>> dedicated() {
        return Codegen.optional(this.dedicated);
    }
    /**
     * Name of the cluster.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the cluster.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Describes the current long-running operation, if any.
     * 
     */
    @Export(name="operationStatus", type=String.class, parameters={})
    private Output<String> operationStatus;

    /**
     * @return Describes the current long-running operation, if any.
     * 
     */
    public Output<String> operationStatus() {
        return this.operationStatus;
    }
    /**
     * The ID of the cluster&#39;s parent folder. &#39;root&#39; is used for a cluster at the root level.
     * 
     */
    @Export(name="parentId", type=String.class, parameters={})
    private Output<String> parentId;

    /**
     * @return The ID of the cluster&#39;s parent folder. &#39;root&#39; is used for a cluster at the root level.
     * 
     */
    public Output<String> parentId() {
        return this.parentId;
    }
    /**
     * Denotes cluster deployment type: &#39;DEDICATED&#39; or &#39;SERVERLESS&#39;.
     * 
     */
    @Export(name="plan", type=String.class, parameters={})
    private Output<String> plan;

    /**
     * @return Denotes cluster deployment type: &#39;DEDICATED&#39; or &#39;SERVERLESS&#39;.
     * 
     */
    public Output<String> plan() {
        return this.plan;
    }
    @Export(name="regions", type=List.class, parameters={ClusterRegion.class})
    private Output<List<ClusterRegion>> regions;

    public Output<List<ClusterRegion>> regions() {
        return this.regions;
    }
    @Export(name="serverless", type=ClusterServerless.class, parameters={})
    private Output</* @Nullable */ ClusterServerless> serverless;

    public Output<Optional<ClusterServerless>> serverless() {
        return Codegen.optional(this.serverless);
    }
    /**
     * Describes whether the cluster is being created, updated, deleted, etc.
     * 
     */
    @Export(name="state", type=String.class, parameters={})
    private Output<String> state;

    /**
     * @return Describes whether the cluster is being created, updated, deleted, etc.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Describes the status of any in-progress CockroachDB upgrade or rollback.
     * 
     */
    @Export(name="upgradeStatus", type=String.class, parameters={})
    private Output<String> upgradeStatus;

    /**
     * @return Describes the status of any in-progress CockroachDB upgrade or rollback.
     * 
     */
    public Output<String> upgradeStatus() {
        return this.upgradeStatus;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Cluster(String name) {
        this(name, ClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Cluster(String name, ClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Cluster(String name, ClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cockroach:index/cluster:Cluster", name, args == null ? ClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Cluster(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cockroach:index/cluster:Cluster", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Cluster get(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Cluster(name, id, state, options);
    }
}
