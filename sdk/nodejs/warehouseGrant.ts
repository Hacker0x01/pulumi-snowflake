// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as snowflake from "@pulumi/snowflake";
 *
 * const grant = new snowflake.WarehouseGrant("grant", {
 *     privilege: "MODIFY",
 *     roles: ["role1"],
 *     warehouseName: "wh",
 *     withGrantOption: false,
 * });
 * ```
 *
 * ## Import
 *
 * # format is warehouse name | | | privilege | true/false for with_grant_option
 *
 * ```sh
 *  $ pulumi import snowflake:index/warehouseGrant:WarehouseGrant example 'warehouseName|||MODIFY|true'
 * ```
 */
export class WarehouseGrant extends pulumi.CustomResource {
    /**
     * Get an existing WarehouseGrant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WarehouseGrantState, opts?: pulumi.CustomResourceOptions): WarehouseGrant {
        return new WarehouseGrant(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'snowflake:index/warehouseGrant:WarehouseGrant';

    /**
     * Returns true if the given object is an instance of WarehouseGrant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WarehouseGrant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WarehouseGrant.__pulumiType;
    }

    /**
     * The privilege to grant on the warehouse.
     */
    public readonly privilege!: pulumi.Output<string | undefined>;
    /**
     * Grants privilege to these roles.
     */
    public readonly roles!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the warehouse on which to grant privileges.
     */
    public readonly warehouseName!: pulumi.Output<string>;
    /**
     * When this is set to true, allows the recipient role to grant the privileges to other roles.
     */
    public readonly withGrantOption!: pulumi.Output<boolean | undefined>;

    /**
     * Create a WarehouseGrant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WarehouseGrantArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WarehouseGrantArgs | WarehouseGrantState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WarehouseGrantState | undefined;
            resourceInputs["privilege"] = state ? state.privilege : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["warehouseName"] = state ? state.warehouseName : undefined;
            resourceInputs["withGrantOption"] = state ? state.withGrantOption : undefined;
        } else {
            const args = argsOrState as WarehouseGrantArgs | undefined;
            if ((!args || args.warehouseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'warehouseName'");
            }
            resourceInputs["privilege"] = args ? args.privilege : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["warehouseName"] = args ? args.warehouseName : undefined;
            resourceInputs["withGrantOption"] = args ? args.withGrantOption : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WarehouseGrant.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WarehouseGrant resources.
 */
export interface WarehouseGrantState {
    /**
     * The privilege to grant on the warehouse.
     */
    privilege?: pulumi.Input<string>;
    /**
     * Grants privilege to these roles.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the warehouse on which to grant privileges.
     */
    warehouseName?: pulumi.Input<string>;
    /**
     * When this is set to true, allows the recipient role to grant the privileges to other roles.
     */
    withGrantOption?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a WarehouseGrant resource.
 */
export interface WarehouseGrantArgs {
    /**
     * The privilege to grant on the warehouse.
     */
    privilege?: pulumi.Input<string>;
    /**
     * Grants privilege to these roles.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the warehouse on which to grant privileges.
     */
    warehouseName: pulumi.Input<string>;
    /**
     * When this is set to true, allows the recipient role to grant the privileges to other roles.
     */
    withGrantOption?: pulumi.Input<boolean>;
}
