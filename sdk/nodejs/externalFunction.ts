// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as snowflake from "@pulumi/snowflake";
 *
 * const testExtFunc = new snowflake.ExternalFunction("test_ext_func", {
 *     apiIntegration: "api_integration_name",
 *     args: [
 *         {
 *             name: "arg1",
 *             type: "varchar",
 *         },
 *         {
 *             name: "arg2",
 *             type: "varchar",
 *         },
 *     ],
 *     database: "my_test_db",
 *     returnBehavior: "IMMUTABLE",
 *     returnType: "varchar",
 *     schema: "my_test_schema",
 *     urlOfProxyAndResource: "https://123456.execute-api.us-west-2.amazonaws.com/prod/test_func",
 * });
 * ```
 *
 * ## Import
 *
 * # format is database name | schema name | external function name | <list of function arg types, separated with '-'>
 *
 * ```sh
 *  $ pulumi import snowflake:index/externalFunction:ExternalFunction example 'dbName|schemaName|externalFunctionName|varchar-varchar-varchar'
 * ```
 */
export class ExternalFunction extends pulumi.CustomResource {
    /**
     * Get an existing ExternalFunction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExternalFunctionState, opts?: pulumi.CustomResourceOptions): ExternalFunction {
        return new ExternalFunction(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'snowflake:index/externalFunction:ExternalFunction';

    /**
     * Returns true if the given object is an instance of ExternalFunction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExternalFunction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExternalFunction.__pulumiType;
    }

    /**
     * The name of the API integration object that should be used to authenticate the call to the proxy service.
     */
    public readonly apiIntegration!: pulumi.Output<string>;
    /**
     * Specifies the arguments/inputs for the external function. These should correspond to the arguments that the remote service expects.
     */
    public readonly args!: pulumi.Output<outputs.ExternalFunctionArg[] | undefined>;
    /**
     * A description of the external function.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * If specified, the JSON payload is compressed when sent from Snowflake to the proxy service, and when sent back from the proxy service to Snowflake.
     */
    public readonly compression!: pulumi.Output<string | undefined>;
    /**
     * Binds Snowflake context function results to HTTP headers.
     */
    public readonly contextHeaders!: pulumi.Output<string[] | undefined>;
    /**
     * Date and time when the external function was created.
     */
    public /*out*/ readonly createdOn!: pulumi.Output<string>;
    /**
     * The database in which to create the external function.
     */
    public readonly database!: pulumi.Output<string>;
    /**
     * Allows users to specify key-value metadata that is sent with every request as HTTP headers.
     */
    public readonly headers!: pulumi.Output<outputs.ExternalFunctionHeader[] | undefined>;
    /**
     * This specifies the maximum number of rows in each batch sent to the proxy service.
     */
    public readonly maxBatchRows!: pulumi.Output<number | undefined>;
    /**
     * Specifies the identifier for the external function. The identifier can contain the schema name and database name, as well as the function name. The function's signature (name and argument data types) must be unique within the schema.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the behavior of the external function when called with null inputs.
     */
    public readonly nullInputBehavior!: pulumi.Output<string | undefined>;
    /**
     * Specifies the behavior of the function when returning results
     */
    public readonly returnBehavior!: pulumi.Output<string>;
    /**
     * Indicates whether the function can return NULL values or must return only NON-NULL values.
     */
    public readonly returnNullAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the data type returned by the external function.
     */
    public readonly returnType!: pulumi.Output<string>;
    /**
     * The schema in which to create the external function.
     */
    public readonly schema!: pulumi.Output<string>;
    /**
     * This is the invocation URL of the proxy service and resource through which Snowflake calls the remote service.
     */
    public readonly urlOfProxyAndResource!: pulumi.Output<string>;

    /**
     * Create a ExternalFunction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExternalFunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExternalFunctionArgs | ExternalFunctionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExternalFunctionState | undefined;
            inputs["apiIntegration"] = state ? state.apiIntegration : undefined;
            inputs["args"] = state ? state.args : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["compression"] = state ? state.compression : undefined;
            inputs["contextHeaders"] = state ? state.contextHeaders : undefined;
            inputs["createdOn"] = state ? state.createdOn : undefined;
            inputs["database"] = state ? state.database : undefined;
            inputs["headers"] = state ? state.headers : undefined;
            inputs["maxBatchRows"] = state ? state.maxBatchRows : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nullInputBehavior"] = state ? state.nullInputBehavior : undefined;
            inputs["returnBehavior"] = state ? state.returnBehavior : undefined;
            inputs["returnNullAllowed"] = state ? state.returnNullAllowed : undefined;
            inputs["returnType"] = state ? state.returnType : undefined;
            inputs["schema"] = state ? state.schema : undefined;
            inputs["urlOfProxyAndResource"] = state ? state.urlOfProxyAndResource : undefined;
        } else {
            const args = argsOrState as ExternalFunctionArgs | undefined;
            if ((!args || args.apiIntegration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiIntegration'");
            }
            if ((!args || args.database === undefined) && !opts.urn) {
                throw new Error("Missing required property 'database'");
            }
            if ((!args || args.returnBehavior === undefined) && !opts.urn) {
                throw new Error("Missing required property 'returnBehavior'");
            }
            if ((!args || args.returnType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'returnType'");
            }
            if ((!args || args.schema === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schema'");
            }
            if ((!args || args.urlOfProxyAndResource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'urlOfProxyAndResource'");
            }
            inputs["apiIntegration"] = args ? args.apiIntegration : undefined;
            inputs["args"] = args ? args.args : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["compression"] = args ? args.compression : undefined;
            inputs["contextHeaders"] = args ? args.contextHeaders : undefined;
            inputs["database"] = args ? args.database : undefined;
            inputs["headers"] = args ? args.headers : undefined;
            inputs["maxBatchRows"] = args ? args.maxBatchRows : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nullInputBehavior"] = args ? args.nullInputBehavior : undefined;
            inputs["returnBehavior"] = args ? args.returnBehavior : undefined;
            inputs["returnNullAllowed"] = args ? args.returnNullAllowed : undefined;
            inputs["returnType"] = args ? args.returnType : undefined;
            inputs["schema"] = args ? args.schema : undefined;
            inputs["urlOfProxyAndResource"] = args ? args.urlOfProxyAndResource : undefined;
            inputs["createdOn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ExternalFunction.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExternalFunction resources.
 */
export interface ExternalFunctionState {
    /**
     * The name of the API integration object that should be used to authenticate the call to the proxy service.
     */
    apiIntegration?: pulumi.Input<string>;
    /**
     * Specifies the arguments/inputs for the external function. These should correspond to the arguments that the remote service expects.
     */
    args?: pulumi.Input<pulumi.Input<inputs.ExternalFunctionArg>[]>;
    /**
     * A description of the external function.
     */
    comment?: pulumi.Input<string>;
    /**
     * If specified, the JSON payload is compressed when sent from Snowflake to the proxy service, and when sent back from the proxy service to Snowflake.
     */
    compression?: pulumi.Input<string>;
    /**
     * Binds Snowflake context function results to HTTP headers.
     */
    contextHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Date and time when the external function was created.
     */
    createdOn?: pulumi.Input<string>;
    /**
     * The database in which to create the external function.
     */
    database?: pulumi.Input<string>;
    /**
     * Allows users to specify key-value metadata that is sent with every request as HTTP headers.
     */
    headers?: pulumi.Input<pulumi.Input<inputs.ExternalFunctionHeader>[]>;
    /**
     * This specifies the maximum number of rows in each batch sent to the proxy service.
     */
    maxBatchRows?: pulumi.Input<number>;
    /**
     * Specifies the identifier for the external function. The identifier can contain the schema name and database name, as well as the function name. The function's signature (name and argument data types) must be unique within the schema.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the behavior of the external function when called with null inputs.
     */
    nullInputBehavior?: pulumi.Input<string>;
    /**
     * Specifies the behavior of the function when returning results
     */
    returnBehavior?: pulumi.Input<string>;
    /**
     * Indicates whether the function can return NULL values or must return only NON-NULL values.
     */
    returnNullAllowed?: pulumi.Input<boolean>;
    /**
     * Specifies the data type returned by the external function.
     */
    returnType?: pulumi.Input<string>;
    /**
     * The schema in which to create the external function.
     */
    schema?: pulumi.Input<string>;
    /**
     * This is the invocation URL of the proxy service and resource through which Snowflake calls the remote service.
     */
    urlOfProxyAndResource?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ExternalFunction resource.
 */
export interface ExternalFunctionArgs {
    /**
     * The name of the API integration object that should be used to authenticate the call to the proxy service.
     */
    apiIntegration: pulumi.Input<string>;
    /**
     * Specifies the arguments/inputs for the external function. These should correspond to the arguments that the remote service expects.
     */
    args?: pulumi.Input<pulumi.Input<inputs.ExternalFunctionArg>[]>;
    /**
     * A description of the external function.
     */
    comment?: pulumi.Input<string>;
    /**
     * If specified, the JSON payload is compressed when sent from Snowflake to the proxy service, and when sent back from the proxy service to Snowflake.
     */
    compression?: pulumi.Input<string>;
    /**
     * Binds Snowflake context function results to HTTP headers.
     */
    contextHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The database in which to create the external function.
     */
    database: pulumi.Input<string>;
    /**
     * Allows users to specify key-value metadata that is sent with every request as HTTP headers.
     */
    headers?: pulumi.Input<pulumi.Input<inputs.ExternalFunctionHeader>[]>;
    /**
     * This specifies the maximum number of rows in each batch sent to the proxy service.
     */
    maxBatchRows?: pulumi.Input<number>;
    /**
     * Specifies the identifier for the external function. The identifier can contain the schema name and database name, as well as the function name. The function's signature (name and argument data types) must be unique within the schema.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the behavior of the external function when called with null inputs.
     */
    nullInputBehavior?: pulumi.Input<string>;
    /**
     * Specifies the behavior of the function when returning results
     */
    returnBehavior: pulumi.Input<string>;
    /**
     * Indicates whether the function can return NULL values or must return only NON-NULL values.
     */
    returnNullAllowed?: pulumi.Input<boolean>;
    /**
     * Specifies the data type returned by the external function.
     */
    returnType: pulumi.Input<string>;
    /**
     * The schema in which to create the external function.
     */
    schema: pulumi.Input<string>;
    /**
     * This is the invocation URL of the proxy service and resource through which Snowflake calls the remote service.
     */
    urlOfProxyAndResource: pulumi.Input<string>;
}
