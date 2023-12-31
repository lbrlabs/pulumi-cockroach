// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach;

import com.pulumi.cockroach.PrivateEndpointTrustedOwnerArgs;
import com.pulumi.cockroach.Utilities;
import com.pulumi.cockroach.inputs.PrivateEndpointTrustedOwnerState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Private Endpoint Trusted Owner.
 * 
 */
@ResourceType(type="cockroach:index/privateEndpointTrustedOwner:PrivateEndpointTrustedOwner")
public class PrivateEndpointTrustedOwner extends com.pulumi.resources.CustomResource {
    /**
     * UUID of the cluster the private endpoint trusted owner entry belongs to.
     * 
     */
    @Export(name="clusterId", type=String.class, parameters={})
    private Output<String> clusterId;

    /**
     * @return UUID of the cluster the private endpoint trusted owner entry belongs to.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * Owner ID of the private endpoint connection in the cloud provider.
     * 
     */
    @Export(name="externalOwnerId", type=String.class, parameters={})
    private Output<String> externalOwnerId;

    /**
     * @return Owner ID of the private endpoint connection in the cloud provider.
     * 
     */
    public Output<String> externalOwnerId() {
        return this.externalOwnerId;
    }
    /**
     * UUID of the private endpoint trusted owner entry.
     * 
     */
    @Export(name="ownerId", type=String.class, parameters={})
    private Output<String> ownerId;

    /**
     * @return UUID of the private endpoint trusted owner entry.
     * 
     */
    public Output<String> ownerId() {
        return this.ownerId;
    }
    /**
     * Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PrivateEndpointTrustedOwner(String name) {
        this(name, PrivateEndpointTrustedOwnerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PrivateEndpointTrustedOwner(String name, PrivateEndpointTrustedOwnerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PrivateEndpointTrustedOwner(String name, PrivateEndpointTrustedOwnerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cockroach:index/privateEndpointTrustedOwner:PrivateEndpointTrustedOwner", name, args == null ? PrivateEndpointTrustedOwnerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PrivateEndpointTrustedOwner(String name, Output<String> id, @Nullable PrivateEndpointTrustedOwnerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cockroach:index/privateEndpointTrustedOwner:PrivateEndpointTrustedOwner", name, state, makeResourceOptions(options, id));
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
    public static PrivateEndpointTrustedOwner get(String name, Output<String> id, @Nullable PrivateEndpointTrustedOwnerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PrivateEndpointTrustedOwner(name, id, state, options);
    }
}
