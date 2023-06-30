// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allow list of IP range
 */
export class AllowList extends pulumi.CustomResource {
    /**
     * Get an existing AllowList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AllowListState, opts?: pulumi.CustomResourceOptions): AllowList {
        return new AllowList(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cockroach:index/allowList:AllowList';

    /**
     * Returns true if the given object is an instance of AllowList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AllowList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AllowList.__pulumiType;
    }

    public readonly cidrIp!: pulumi.Output<string>;
    public readonly cidrMask!: pulumi.Output<number>;
    public readonly clusterId!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly sql!: pulumi.Output<boolean>;
    public readonly ui!: pulumi.Output<boolean>;

    /**
     * Create a AllowList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AllowListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AllowListArgs | AllowListState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AllowListState | undefined;
            resourceInputs["cidrIp"] = state ? state.cidrIp : undefined;
            resourceInputs["cidrMask"] = state ? state.cidrMask : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["sql"] = state ? state.sql : undefined;
            resourceInputs["ui"] = state ? state.ui : undefined;
        } else {
            const args = argsOrState as AllowListArgs | undefined;
            if ((!args || args.cidrIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cidrIp'");
            }
            if ((!args || args.cidrMask === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cidrMask'");
            }
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.sql === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sql'");
            }
            if ((!args || args.ui === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ui'");
            }
            resourceInputs["cidrIp"] = args ? args.cidrIp : undefined;
            resourceInputs["cidrMask"] = args ? args.cidrMask : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sql"] = args ? args.sql : undefined;
            resourceInputs["ui"] = args ? args.ui : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AllowList.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AllowList resources.
 */
export interface AllowListState {
    cidrIp?: pulumi.Input<string>;
    cidrMask?: pulumi.Input<number>;
    clusterId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    sql?: pulumi.Input<boolean>;
    ui?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AllowList resource.
 */
export interface AllowListArgs {
    cidrIp: pulumi.Input<string>;
    cidrMask: pulumi.Input<number>;
    clusterId: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    sql: pulumi.Input<boolean>;
    ui: pulumi.Input<boolean>;
}