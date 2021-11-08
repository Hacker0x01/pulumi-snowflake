// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemGetSnowflakePlatformInfo(opts?: pulumi.InvokeOptions): Promise<GetSystemGetSnowflakePlatformInfoResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("snowflake:index/getSystemGetSnowflakePlatformInfo:getSystemGetSnowflakePlatformInfo", {
    }, opts);
}

/**
 * A collection of values returned by getSystemGetSnowflakePlatformInfo.
 */
export interface GetSystemGetSnowflakePlatformInfoResult {
    /**
     * Snowflake AWS Virtual Private Cloud IDs
     */
    readonly awsVpcIds: string[];
    /**
     * Snowflake Azure Virtual Network Subnet IDs
     */
    readonly azureVnetSubnetIds: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
