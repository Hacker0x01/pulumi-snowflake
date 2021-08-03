// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getViews(args: GetViewsArgs, opts?: pulumi.InvokeOptions): Promise<GetViewsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("snowflake:index/getViews:getViews", {
        "database": args.database,
        "schema": args.schema,
    }, opts);
}

/**
 * A collection of arguments for invoking getViews.
 */
export interface GetViewsArgs {
    database: string;
    schema: string;
}

/**
 * A collection of values returned by getViews.
 */
export interface GetViewsResult {
    readonly database: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly schema: string;
    readonly views: outputs.GetViewsView[];
}
