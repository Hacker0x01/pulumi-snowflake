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
 * const integration = new snowflake.NotificationIntegration("integration", {
 *     awsSqsArn: "...",
 *     awsSqsRoleArn: "...",
 *     azureStorageQueuePrimaryUri: "...",
 *     azureTenantId: "...",
 *     comment: "A notification integration.",
 *     direction: "OUTBOUND",
 *     enabled: true,
 *     notificationProvider: "AWS_SQS",
 *     type: "QUEUE",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import snowflake:index/notificationIntegration:NotificationIntegration example name
 * ```
 */
export class NotificationIntegration extends pulumi.CustomResource {
    /**
     * Get an existing NotificationIntegration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationIntegrationState, opts?: pulumi.CustomResourceOptions): NotificationIntegration {
        return new NotificationIntegration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'snowflake:index/notificationIntegration:NotificationIntegration';

    /**
     * Returns true if the given object is an instance of NotificationIntegration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationIntegration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationIntegration.__pulumiType;
    }

    /**
     * AWS SQS queue ARN for notification integration to connect to
     */
    public readonly awsSqsArn!: pulumi.Output<string | undefined>;
    /**
     * The external ID that Snowflake will use when assuming the AWS role
     */
    public /*out*/ readonly awsSqsExternalId!: pulumi.Output<string>;
    /**
     * The Snowflake user that will attempt to assume the AWS role.
     */
    public /*out*/ readonly awsSqsIamUserArn!: pulumi.Output<string>;
    /**
     * AWS IAM role ARN for notification integration to assume
     */
    public readonly awsSqsRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The queue ID for the Azure Queue Storage queue created for Event Grid notifications
     */
    public readonly azureStorageQueuePrimaryUri!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Azure Active Directory tenant used for identity management
     */
    public readonly azureTenantId!: pulumi.Output<string | undefined>;
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Date and time when the notification integration was created.
     */
    public /*out*/ readonly createdOn!: pulumi.Output<string>;
    /**
     * Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
     */
    public readonly direction!: pulumi.Output<string | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
     */
    public readonly gcpPubsubSubscriptionName!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS_SQS)
     */
    public readonly notificationProvider!: pulumi.Output<string | undefined>;
    /**
     * A type of integration
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a NotificationIntegration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NotificationIntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationIntegrationArgs | NotificationIntegrationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotificationIntegrationState | undefined;
            inputs["awsSqsArn"] = state ? state.awsSqsArn : undefined;
            inputs["awsSqsExternalId"] = state ? state.awsSqsExternalId : undefined;
            inputs["awsSqsIamUserArn"] = state ? state.awsSqsIamUserArn : undefined;
            inputs["awsSqsRoleArn"] = state ? state.awsSqsRoleArn : undefined;
            inputs["azureStorageQueuePrimaryUri"] = state ? state.azureStorageQueuePrimaryUri : undefined;
            inputs["azureTenantId"] = state ? state.azureTenantId : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["createdOn"] = state ? state.createdOn : undefined;
            inputs["direction"] = state ? state.direction : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["gcpPubsubSubscriptionName"] = state ? state.gcpPubsubSubscriptionName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notificationProvider"] = state ? state.notificationProvider : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as NotificationIntegrationArgs | undefined;
            inputs["awsSqsArn"] = args ? args.awsSqsArn : undefined;
            inputs["awsSqsRoleArn"] = args ? args.awsSqsRoleArn : undefined;
            inputs["azureStorageQueuePrimaryUri"] = args ? args.azureStorageQueuePrimaryUri : undefined;
            inputs["azureTenantId"] = args ? args.azureTenantId : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["direction"] = args ? args.direction : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["gcpPubsubSubscriptionName"] = args ? args.gcpPubsubSubscriptionName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationProvider"] = args ? args.notificationProvider : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["awsSqsExternalId"] = undefined /*out*/;
            inputs["awsSqsIamUserArn"] = undefined /*out*/;
            inputs["createdOn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NotificationIntegration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NotificationIntegration resources.
 */
export interface NotificationIntegrationState {
    /**
     * AWS SQS queue ARN for notification integration to connect to
     */
    awsSqsArn?: pulumi.Input<string>;
    /**
     * The external ID that Snowflake will use when assuming the AWS role
     */
    awsSqsExternalId?: pulumi.Input<string>;
    /**
     * The Snowflake user that will attempt to assume the AWS role.
     */
    awsSqsIamUserArn?: pulumi.Input<string>;
    /**
     * AWS IAM role ARN for notification integration to assume
     */
    awsSqsRoleArn?: pulumi.Input<string>;
    /**
     * The queue ID for the Azure Queue Storage queue created for Event Grid notifications
     */
    azureStorageQueuePrimaryUri?: pulumi.Input<string>;
    /**
     * The ID of the Azure Active Directory tenant used for identity management
     */
    azureTenantId?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    /**
     * Date and time when the notification integration was created.
     */
    createdOn?: pulumi.Input<string>;
    /**
     * Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
     */
    direction?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    /**
     * The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
     */
    gcpPubsubSubscriptionName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS_SQS)
     */
    notificationProvider?: pulumi.Input<string>;
    /**
     * A type of integration
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NotificationIntegration resource.
 */
export interface NotificationIntegrationArgs {
    /**
     * AWS SQS queue ARN for notification integration to connect to
     */
    awsSqsArn?: pulumi.Input<string>;
    /**
     * AWS IAM role ARN for notification integration to assume
     */
    awsSqsRoleArn?: pulumi.Input<string>;
    /**
     * The queue ID for the Azure Queue Storage queue created for Event Grid notifications
     */
    azureStorageQueuePrimaryUri?: pulumi.Input<string>;
    /**
     * The ID of the Azure Active Directory tenant used for identity management
     */
    azureTenantId?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    /**
     * Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
     */
    direction?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    /**
     * The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
     */
    gcpPubsubSubscriptionName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS_SQS)
     */
    notificationProvider?: pulumi.Input<string>;
    /**
     * A type of integration
     */
    type?: pulumi.Input<string>;
}
