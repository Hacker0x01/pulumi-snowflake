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
 * const current = pulumi.output(snowflake.getExternalTables({
 *     database: "MYDB",
 *     schema: "MYSCHEMA",
 * }));
 * ```
 */
export function getExternalTables(args: GetExternalTablesArgs, opts?: pulumi.InvokeOptions): Promise<GetExternalTablesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("snowflake:index/getExternalTables:getExternalTables", {
        "database": args.database,
        "schema": args.schema,
    }, opts);
}

/**
 * A collection of arguments for invoking getExternalTables.
 */
export interface GetExternalTablesArgs {
    /**
     * The database from which to return the schemas from.
     */
    database: string;
    /**
     * The schema from which to return the external tables from.
     */
    schema: string;
}

/**
 * A collection of values returned by getExternalTables.
 */
export interface GetExternalTablesResult {
    /**
     * The database from which to return the schemas from.
     */
    readonly database: string;
    /**
     * The external tables in the schema
     */
    readonly externalTables: outputs.GetExternalTablesExternalTable[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The schema from which to return the external tables from.
     */
    readonly schema: string;
}
