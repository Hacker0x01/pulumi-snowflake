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
 * const monitor = new snowflake.ResourceMonitor("monitor", {
 *     creditQuota: 100,
 *     endTimestamp: "2021-12-07 00:00",
 *     frequency: "DAILY",
 *     notifyTriggers: [40],
 *     startTimestamp: "2020-12-07 00:00",
 *     suspendImmediateTriggers: [90],
 *     suspendTriggers: [50],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import snowflake:index/resourceMonitor:ResourceMonitor example
 * ```
 */
export class ResourceMonitor extends pulumi.CustomResource {
    /**
     * Get an existing ResourceMonitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceMonitorState, opts?: pulumi.CustomResourceOptions): ResourceMonitor {
        return new ResourceMonitor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'snowflake:index/resourceMonitor:ResourceMonitor';

    /**
     * Returns true if the given object is an instance of ResourceMonitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceMonitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceMonitor.__pulumiType;
    }

    /**
     * The number of credits allocated monthly to the resource monitor.
     */
    public readonly creditQuota!: pulumi.Output<number>;
    /**
     * The date and time when the resource monitor suspends the assigned warehouses.
     */
    public readonly endTimestamp!: pulumi.Output<string | undefined>;
    /**
     * The frequency interval at which the credit usage resets to 0. If you set a frequency for a resource monitor, you must also set START_TIMESTAMP.
     */
    public readonly frequency!: pulumi.Output<string>;
    /**
     * Identifier for the resource monitor; must be unique for your account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of percentage thresholds at which to send an alert to subscribed users.
     */
    public readonly notifyTriggers!: pulumi.Output<number[] | undefined>;
    /**
     * Specifies whether the resource monitor should be applied globally to your Snowflake account.
     */
    public readonly setForAccount!: pulumi.Output<boolean | undefined>;
    /**
     * The date and time when the resource monitor starts monitoring credit usage for the assigned warehouses.
     */
    public readonly startTimestamp!: pulumi.Output<string>;
    /**
     * A list of percentage thresholds at which to immediately suspend all warehouses.
     */
    public readonly suspendImmediateTriggers!: pulumi.Output<number[] | undefined>;
    /**
     * A list of percentage thresholds at which to suspend all warehouses.
     */
    public readonly suspendTriggers!: pulumi.Output<number[] | undefined>;
    /**
     * A list of warehouses to apply the resource monitor to.
     */
    public readonly warehouses!: pulumi.Output<string[] | undefined>;

    /**
     * Create a ResourceMonitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResourceMonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceMonitorArgs | ResourceMonitorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResourceMonitorState | undefined;
            resourceInputs["creditQuota"] = state ? state.creditQuota : undefined;
            resourceInputs["endTimestamp"] = state ? state.endTimestamp : undefined;
            resourceInputs["frequency"] = state ? state.frequency : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notifyTriggers"] = state ? state.notifyTriggers : undefined;
            resourceInputs["setForAccount"] = state ? state.setForAccount : undefined;
            resourceInputs["startTimestamp"] = state ? state.startTimestamp : undefined;
            resourceInputs["suspendImmediateTriggers"] = state ? state.suspendImmediateTriggers : undefined;
            resourceInputs["suspendTriggers"] = state ? state.suspendTriggers : undefined;
            resourceInputs["warehouses"] = state ? state.warehouses : undefined;
        } else {
            const args = argsOrState as ResourceMonitorArgs | undefined;
            resourceInputs["creditQuota"] = args ? args.creditQuota : undefined;
            resourceInputs["endTimestamp"] = args ? args.endTimestamp : undefined;
            resourceInputs["frequency"] = args ? args.frequency : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notifyTriggers"] = args ? args.notifyTriggers : undefined;
            resourceInputs["setForAccount"] = args ? args.setForAccount : undefined;
            resourceInputs["startTimestamp"] = args ? args.startTimestamp : undefined;
            resourceInputs["suspendImmediateTriggers"] = args ? args.suspendImmediateTriggers : undefined;
            resourceInputs["suspendTriggers"] = args ? args.suspendTriggers : undefined;
            resourceInputs["warehouses"] = args ? args.warehouses : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResourceMonitor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourceMonitor resources.
 */
export interface ResourceMonitorState {
    /**
     * The number of credits allocated monthly to the resource monitor.
     */
    creditQuota?: pulumi.Input<number>;
    /**
     * The date and time when the resource monitor suspends the assigned warehouses.
     */
    endTimestamp?: pulumi.Input<string>;
    /**
     * The frequency interval at which the credit usage resets to 0. If you set a frequency for a resource monitor, you must also set START_TIMESTAMP.
     */
    frequency?: pulumi.Input<string>;
    /**
     * Identifier for the resource monitor; must be unique for your account.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of percentage thresholds at which to send an alert to subscribed users.
     */
    notifyTriggers?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Specifies whether the resource monitor should be applied globally to your Snowflake account.
     */
    setForAccount?: pulumi.Input<boolean>;
    /**
     * The date and time when the resource monitor starts monitoring credit usage for the assigned warehouses.
     */
    startTimestamp?: pulumi.Input<string>;
    /**
     * A list of percentage thresholds at which to immediately suspend all warehouses.
     */
    suspendImmediateTriggers?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * A list of percentage thresholds at which to suspend all warehouses.
     */
    suspendTriggers?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * A list of warehouses to apply the resource monitor to.
     */
    warehouses?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ResourceMonitor resource.
 */
export interface ResourceMonitorArgs {
    /**
     * The number of credits allocated monthly to the resource monitor.
     */
    creditQuota?: pulumi.Input<number>;
    /**
     * The date and time when the resource monitor suspends the assigned warehouses.
     */
    endTimestamp?: pulumi.Input<string>;
    /**
     * The frequency interval at which the credit usage resets to 0. If you set a frequency for a resource monitor, you must also set START_TIMESTAMP.
     */
    frequency?: pulumi.Input<string>;
    /**
     * Identifier for the resource monitor; must be unique for your account.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of percentage thresholds at which to send an alert to subscribed users.
     */
    notifyTriggers?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Specifies whether the resource monitor should be applied globally to your Snowflake account.
     */
    setForAccount?: pulumi.Input<boolean>;
    /**
     * The date and time when the resource monitor starts monitoring credit usage for the assigned warehouses.
     */
    startTimestamp?: pulumi.Input<string>;
    /**
     * A list of percentage thresholds at which to immediately suspend all warehouses.
     */
    suspendImmediateTriggers?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * A list of percentage thresholds at which to suspend all warehouses.
     */
    suspendTriggers?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * A list of warehouses to apply the resource monitor to.
     */
    warehouses?: pulumi.Input<pulumi.Input<string>[]>;
}
