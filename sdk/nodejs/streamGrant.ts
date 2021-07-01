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
 * const grant = new snowflake.StreamGrant("grant", {
 *     databaseName: "db",
 *     onFuture: false,
 *     privilege: "select",
 *     roles: [
 *         "role1",
 *         "role2",
 *     ],
 *     schemaName: "schema",
 *     streamName: "view",
 *     withGrantOption: false,
 * });
 * ```
 *
 * ## Import
 *
 * # format is database name | schema name | stream name | privilege | true/false for with_grant_option
 *
 * ```sh
 *  $ pulumi import snowflake:index/streamGrant:StreamGrant example 'dbName|schemaName|streamName|SELECT|false'
 * ```
 */
export class StreamGrant extends pulumi.CustomResource {
    /**
     * Get an existing StreamGrant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StreamGrantState, opts?: pulumi.CustomResourceOptions): StreamGrant {
        return new StreamGrant(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'snowflake:index/streamGrant:StreamGrant';

    /**
     * Returns true if the given object is an instance of StreamGrant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StreamGrant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StreamGrant.__pulumiType;
    }

    /**
     * The name of the database containing the current or future streams on which to grant privileges.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * When this is set to true and a schema*name is provided, apply this grant on all future streams in the given schema. When this is true and no schema*name is provided apply this grant on all future streams in the given database. The stream*name field must be unset in order to use on*future.
     */
    public readonly onFuture!: pulumi.Output<boolean | undefined>;
    /**
     * The privilege to grant on the current or future stream.
     */
    public readonly privilege!: pulumi.Output<string | undefined>;
    /**
     * Grants privilege to these roles.
     */
    public readonly roles!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the schema containing the current or future streams on which to grant privileges.
     */
    public readonly schemaName!: pulumi.Output<string>;
    /**
     * The name of the stream on which to grant privileges immediately (only valid if onFuture is false).
     */
    public readonly streamName!: pulumi.Output<string | undefined>;
    /**
     * When this is set to true, allows the recipient role to grant the privileges to other roles.
     */
    public readonly withGrantOption!: pulumi.Output<boolean | undefined>;

    /**
     * Create a StreamGrant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StreamGrantArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StreamGrantArgs | StreamGrantState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StreamGrantState | undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["onFuture"] = state ? state.onFuture : undefined;
            inputs["privilege"] = state ? state.privilege : undefined;
            inputs["roles"] = state ? state.roles : undefined;
            inputs["schemaName"] = state ? state.schemaName : undefined;
            inputs["streamName"] = state ? state.streamName : undefined;
            inputs["withGrantOption"] = state ? state.withGrantOption : undefined;
        } else {
            const args = argsOrState as StreamGrantArgs | undefined;
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.schemaName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schemaName'");
            }
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["onFuture"] = args ? args.onFuture : undefined;
            inputs["privilege"] = args ? args.privilege : undefined;
            inputs["roles"] = args ? args.roles : undefined;
            inputs["schemaName"] = args ? args.schemaName : undefined;
            inputs["streamName"] = args ? args.streamName : undefined;
            inputs["withGrantOption"] = args ? args.withGrantOption : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(StreamGrant.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StreamGrant resources.
 */
export interface StreamGrantState {
    /**
     * The name of the database containing the current or future streams on which to grant privileges.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * When this is set to true and a schema*name is provided, apply this grant on all future streams in the given schema. When this is true and no schema*name is provided apply this grant on all future streams in the given database. The stream*name field must be unset in order to use on*future.
     */
    onFuture?: pulumi.Input<boolean>;
    /**
     * The privilege to grant on the current or future stream.
     */
    privilege?: pulumi.Input<string>;
    /**
     * Grants privilege to these roles.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the schema containing the current or future streams on which to grant privileges.
     */
    schemaName?: pulumi.Input<string>;
    /**
     * The name of the stream on which to grant privileges immediately (only valid if onFuture is false).
     */
    streamName?: pulumi.Input<string>;
    /**
     * When this is set to true, allows the recipient role to grant the privileges to other roles.
     */
    withGrantOption?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a StreamGrant resource.
 */
export interface StreamGrantArgs {
    /**
     * The name of the database containing the current or future streams on which to grant privileges.
     */
    databaseName: pulumi.Input<string>;
    /**
     * When this is set to true and a schema*name is provided, apply this grant on all future streams in the given schema. When this is true and no schema*name is provided apply this grant on all future streams in the given database. The stream*name field must be unset in order to use on*future.
     */
    onFuture?: pulumi.Input<boolean>;
    /**
     * The privilege to grant on the current or future stream.
     */
    privilege?: pulumi.Input<string>;
    /**
     * Grants privilege to these roles.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the schema containing the current or future streams on which to grant privileges.
     */
    schemaName: pulumi.Input<string>;
    /**
     * The name of the stream on which to grant privileges immediately (only valid if onFuture is false).
     */
    streamName?: pulumi.Input<string>;
    /**
     * When this is set to true, allows the recipient role to grant the privileges to other roles.
     */
    withGrantOption?: pulumi.Input<boolean>;
}
