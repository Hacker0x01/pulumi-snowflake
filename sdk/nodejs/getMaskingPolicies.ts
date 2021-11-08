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
 * const current = pulumi.output(snowflake.getMaskingPolicies({
 *     database: "MYDB",
 *     schema: "MYSCHEMA",
 * }));
 * ```
 */
export function getMaskingPolicies(args: GetMaskingPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetMaskingPoliciesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("snowflake:index/getMaskingPolicies:getMaskingPolicies", {
        "database": args.database,
        "schema": args.schema,
    }, opts);
}

/**
 * A collection of arguments for invoking getMaskingPolicies.
 */
export interface GetMaskingPoliciesArgs {
    /**
     * The database from which to return the schemas from.
     */
    database: string;
    /**
     * The schema from which to return the maskingPolicies from.
     */
    schema: string;
}

/**
 * A collection of values returned by getMaskingPolicies.
 */
export interface GetMaskingPoliciesResult {
    /**
     * The database from which to return the schemas from.
     */
    readonly database: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The maskingPolicies in the schema
     */
    readonly maskingPolicies: outputs.GetMaskingPoliciesMaskingPolicy[];
    /**
     * The schema from which to return the maskingPolicies from.
     */
    readonly schema: string;
}

export function getMaskingPoliciesOutput(args: GetMaskingPoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMaskingPoliciesResult> {
    return pulumi.output(args).apply(a => getMaskingPolicies(a, opts))
}

/**
 * A collection of arguments for invoking getMaskingPolicies.
 */
export interface GetMaskingPoliciesOutputArgs {
    /**
     * The database from which to return the schemas from.
     */
    database: pulumi.Input<string>;
    /**
     * The schema from which to return the maskingPolicies from.
     */
    schema: pulumi.Input<string>;
}
