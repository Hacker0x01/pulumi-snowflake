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
 * const grant = new snowflake.MaterializedViewGrant("grant", {
 *     databaseName: "db",
 *     materializedViewName: "materialized_view",
 *     onFuture: false,
 *     privilege: "select",
 *     roles: [
 *         "role1",
 *         "role2",
 *     ],
 *     schemaName: "schema",
 *     shares: [
 *         "share1",
 *         "share2",
 *     ],
 *     withGrantOption: false,
 * });
 * ```
 *
 * ## Import
 *
 * # format is database name | schema name | materialized view name | privilege | true/false for with_grant_option
 *
 * ```sh
 *  $ pulumi import snowflake:index/materializedViewGrant:MaterializedViewGrant example 'dbName|schemaName|materializedViewName|SELECT|false'
 * ```
 */
export class MaterializedViewGrant extends pulumi.CustomResource {
    /**
     * Get an existing MaterializedViewGrant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MaterializedViewGrantState, opts?: pulumi.CustomResourceOptions): MaterializedViewGrant {
        return new MaterializedViewGrant(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'snowflake:index/materializedViewGrant:MaterializedViewGrant';

    /**
     * Returns true if the given object is an instance of MaterializedViewGrant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MaterializedViewGrant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MaterializedViewGrant.__pulumiType;
    }

    /**
     * The name of the database containing the current or future materialized views on which to grant privileges.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * The name of the materialized view on which to grant privileges immediately (only valid if onFuture is false).
     */
    public readonly materializedViewName!: pulumi.Output<string | undefined>;
    /**
     * When this is set to true and a schema*name is provided, apply this grant on all future materialized views in the given schema. When this is true and no schema*name is provided apply this grant on all future materialized views in the given database. The materialized*view*name and shares fields must be unset in order to use on_future.
     */
    public readonly onFuture!: pulumi.Output<boolean | undefined>;
    /**
     * The privilege to grant on the current or future materialized view view.
     */
    public readonly privilege!: pulumi.Output<string | undefined>;
    /**
     * Grants privilege to these roles.
     */
    public readonly roles!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the schema containing the current or future materialized views on which to grant privileges.
     */
    public readonly schemaName!: pulumi.Output<string>;
    /**
     * Grants privilege to these shares (only valid if onFuture is false).
     */
    public readonly shares!: pulumi.Output<string[] | undefined>;
    /**
     * When this is set to true, allows the recipient role to grant the privileges to other roles.
     */
    public readonly withGrantOption!: pulumi.Output<boolean | undefined>;

    /**
     * Create a MaterializedViewGrant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MaterializedViewGrantArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MaterializedViewGrantArgs | MaterializedViewGrantState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MaterializedViewGrantState | undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["materializedViewName"] = state ? state.materializedViewName : undefined;
            inputs["onFuture"] = state ? state.onFuture : undefined;
            inputs["privilege"] = state ? state.privilege : undefined;
            inputs["roles"] = state ? state.roles : undefined;
            inputs["schemaName"] = state ? state.schemaName : undefined;
            inputs["shares"] = state ? state.shares : undefined;
            inputs["withGrantOption"] = state ? state.withGrantOption : undefined;
        } else {
            const args = argsOrState as MaterializedViewGrantArgs | undefined;
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.schemaName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schemaName'");
            }
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["materializedViewName"] = args ? args.materializedViewName : undefined;
            inputs["onFuture"] = args ? args.onFuture : undefined;
            inputs["privilege"] = args ? args.privilege : undefined;
            inputs["roles"] = args ? args.roles : undefined;
            inputs["schemaName"] = args ? args.schemaName : undefined;
            inputs["shares"] = args ? args.shares : undefined;
            inputs["withGrantOption"] = args ? args.withGrantOption : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MaterializedViewGrant.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MaterializedViewGrant resources.
 */
export interface MaterializedViewGrantState {
    /**
     * The name of the database containing the current or future materialized views on which to grant privileges.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * The name of the materialized view on which to grant privileges immediately (only valid if onFuture is false).
     */
    materializedViewName?: pulumi.Input<string>;
    /**
     * When this is set to true and a schema*name is provided, apply this grant on all future materialized views in the given schema. When this is true and no schema*name is provided apply this grant on all future materialized views in the given database. The materialized*view*name and shares fields must be unset in order to use on_future.
     */
    onFuture?: pulumi.Input<boolean>;
    /**
     * The privilege to grant on the current or future materialized view view.
     */
    privilege?: pulumi.Input<string>;
    /**
     * Grants privilege to these roles.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the schema containing the current or future materialized views on which to grant privileges.
     */
    schemaName?: pulumi.Input<string>;
    /**
     * Grants privilege to these shares (only valid if onFuture is false).
     */
    shares?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When this is set to true, allows the recipient role to grant the privileges to other roles.
     */
    withGrantOption?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a MaterializedViewGrant resource.
 */
export interface MaterializedViewGrantArgs {
    /**
     * The name of the database containing the current or future materialized views on which to grant privileges.
     */
    databaseName: pulumi.Input<string>;
    /**
     * The name of the materialized view on which to grant privileges immediately (only valid if onFuture is false).
     */
    materializedViewName?: pulumi.Input<string>;
    /**
     * When this is set to true and a schema*name is provided, apply this grant on all future materialized views in the given schema. When this is true and no schema*name is provided apply this grant on all future materialized views in the given database. The materialized*view*name and shares fields must be unset in order to use on_future.
     */
    onFuture?: pulumi.Input<boolean>;
    /**
     * The privilege to grant on the current or future materialized view view.
     */
    privilege?: pulumi.Input<string>;
    /**
     * Grants privilege to these roles.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the schema containing the current or future materialized views on which to grant privileges.
     */
    schemaName: pulumi.Input<string>;
    /**
     * Grants privilege to these shares (only valid if onFuture is false).
     */
    shares?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When this is set to true, allows the recipient role to grant the privileges to other roles.
     */
    withGrantOption?: pulumi.Input<boolean>;
}
