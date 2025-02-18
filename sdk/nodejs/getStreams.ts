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
 * const current = pulumi.output(snowflake.getStreams({
 *     database: "MYDB",
 *     schema: "MYSCHEMA",
 * }));
 * ```
 */
export function getStreams(args: GetStreamsArgs, opts?: pulumi.InvokeOptions): Promise<GetStreamsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("snowflake:index/getStreams:getStreams", {
        "database": args.database,
        "schema": args.schema,
    }, opts);
}

/**
 * A collection of arguments for invoking getStreams.
 */
export interface GetStreamsArgs {
    /**
     * The database from which to return the streams from.
     */
    database: string;
    /**
     * The schema from which to return the streams from.
     */
    schema: string;
}

/**
 * A collection of values returned by getStreams.
 */
export interface GetStreamsResult {
    /**
     * The database from which to return the streams from.
     */
    readonly database: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The schema from which to return the streams from.
     */
    readonly schema: string;
    /**
     * The streams in the schema
     */
    readonly streams: outputs.GetStreamsStream[];
}

export function getStreamsOutput(args: GetStreamsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStreamsResult> {
    return pulumi.output(args).apply(a => getStreams(a, opts))
}

/**
 * A collection of arguments for invoking getStreams.
 */
export interface GetStreamsOutputArgs {
    /**
     * The database from which to return the streams from.
     */
    database: pulumi.Input<string>;
    /**
     * The schema from which to return the streams from.
     */
    schema: pulumi.Input<string>;
}
