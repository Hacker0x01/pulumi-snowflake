# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NotificationIntegrationArgs', 'NotificationIntegration']

@pulumi.input_type
class NotificationIntegrationArgs:
    def __init__(__self__, *,
                 aws_sqs_arn: Optional[pulumi.Input[str]] = None,
                 aws_sqs_external_id: Optional[pulumi.Input[str]] = None,
                 aws_sqs_role_arn: Optional[pulumi.Input[str]] = None,
                 azure_storage_queue_primary_uri: Optional[pulumi.Input[str]] = None,
                 azure_tenant_id: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 gcp_pubsub_subscription_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_provider: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NotificationIntegration resource.
        :param pulumi.Input[str] aws_sqs_arn: AWS SQS queue ARN for notification integration to connect to
        :param pulumi.Input[str] aws_sqs_external_id: The external ID that Snowflake will use when assuming the AWS role
        :param pulumi.Input[str] aws_sqs_role_arn: AWS IAM role ARN for notification integration to assume
        :param pulumi.Input[str] azure_storage_queue_primary_uri: The queue ID for the Azure Queue Storage queue created for Event Grid notifications
        :param pulumi.Input[str] azure_tenant_id: The ID of the Azure Active Directory tenant used for identity management
        :param pulumi.Input[str] direction: Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
        :param pulumi.Input[str] gcp_pubsub_subscription_name: The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
        :param pulumi.Input[str] notification_provider: The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS_SQS)
        :param pulumi.Input[str] type: A type of integration
        """
        if aws_sqs_arn is not None:
            pulumi.set(__self__, "aws_sqs_arn", aws_sqs_arn)
        if aws_sqs_external_id is not None:
            pulumi.set(__self__, "aws_sqs_external_id", aws_sqs_external_id)
        if aws_sqs_role_arn is not None:
            pulumi.set(__self__, "aws_sqs_role_arn", aws_sqs_role_arn)
        if azure_storage_queue_primary_uri is not None:
            pulumi.set(__self__, "azure_storage_queue_primary_uri", azure_storage_queue_primary_uri)
        if azure_tenant_id is not None:
            pulumi.set(__self__, "azure_tenant_id", azure_tenant_id)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if gcp_pubsub_subscription_name is not None:
            pulumi.set(__self__, "gcp_pubsub_subscription_name", gcp_pubsub_subscription_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_provider is not None:
            pulumi.set(__self__, "notification_provider", notification_provider)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="awsSqsArn")
    def aws_sqs_arn(self) -> Optional[pulumi.Input[str]]:
        """
        AWS SQS queue ARN for notification integration to connect to
        """
        return pulumi.get(self, "aws_sqs_arn")

    @aws_sqs_arn.setter
    def aws_sqs_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_sqs_arn", value)

    @property
    @pulumi.getter(name="awsSqsExternalId")
    def aws_sqs_external_id(self) -> Optional[pulumi.Input[str]]:
        """
        The external ID that Snowflake will use when assuming the AWS role
        """
        return pulumi.get(self, "aws_sqs_external_id")

    @aws_sqs_external_id.setter
    def aws_sqs_external_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_sqs_external_id", value)

    @property
    @pulumi.getter(name="awsSqsRoleArn")
    def aws_sqs_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        AWS IAM role ARN for notification integration to assume
        """
        return pulumi.get(self, "aws_sqs_role_arn")

    @aws_sqs_role_arn.setter
    def aws_sqs_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_sqs_role_arn", value)

    @property
    @pulumi.getter(name="azureStorageQueuePrimaryUri")
    def azure_storage_queue_primary_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The queue ID for the Azure Queue Storage queue created for Event Grid notifications
        """
        return pulumi.get(self, "azure_storage_queue_primary_uri")

    @azure_storage_queue_primary_uri.setter
    def azure_storage_queue_primary_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_storage_queue_primary_uri", value)

    @property
    @pulumi.getter(name="azureTenantId")
    def azure_tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Azure Active Directory tenant used for identity management
        """
        return pulumi.get(self, "azure_tenant_id")

    @azure_tenant_id.setter
    def azure_tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_tenant_id", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="gcpPubsubSubscriptionName")
    def gcp_pubsub_subscription_name(self) -> Optional[pulumi.Input[str]]:
        """
        The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
        """
        return pulumi.get(self, "gcp_pubsub_subscription_name")

    @gcp_pubsub_subscription_name.setter
    def gcp_pubsub_subscription_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gcp_pubsub_subscription_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationProvider")
    def notification_provider(self) -> Optional[pulumi.Input[str]]:
        """
        The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS_SQS)
        """
        return pulumi.get(self, "notification_provider")

    @notification_provider.setter
    def notification_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_provider", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        A type of integration
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _NotificationIntegrationState:
    def __init__(__self__, *,
                 aws_sqs_arn: Optional[pulumi.Input[str]] = None,
                 aws_sqs_external_id: Optional[pulumi.Input[str]] = None,
                 aws_sqs_role_arn: Optional[pulumi.Input[str]] = None,
                 azure_storage_queue_primary_uri: Optional[pulumi.Input[str]] = None,
                 azure_tenant_id: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 created_on: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 gcp_pubsub_subscription_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_provider: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NotificationIntegration resources.
        :param pulumi.Input[str] aws_sqs_arn: AWS SQS queue ARN for notification integration to connect to
        :param pulumi.Input[str] aws_sqs_external_id: The external ID that Snowflake will use when assuming the AWS role
        :param pulumi.Input[str] aws_sqs_role_arn: AWS IAM role ARN for notification integration to assume
        :param pulumi.Input[str] azure_storage_queue_primary_uri: The queue ID for the Azure Queue Storage queue created for Event Grid notifications
        :param pulumi.Input[str] azure_tenant_id: The ID of the Azure Active Directory tenant used for identity management
        :param pulumi.Input[str] created_on: Date and time when the notification integration was created.
        :param pulumi.Input[str] direction: Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
        :param pulumi.Input[str] gcp_pubsub_subscription_name: The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
        :param pulumi.Input[str] notification_provider: The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS_SQS)
        :param pulumi.Input[str] type: A type of integration
        """
        if aws_sqs_arn is not None:
            pulumi.set(__self__, "aws_sqs_arn", aws_sqs_arn)
        if aws_sqs_external_id is not None:
            pulumi.set(__self__, "aws_sqs_external_id", aws_sqs_external_id)
        if aws_sqs_role_arn is not None:
            pulumi.set(__self__, "aws_sqs_role_arn", aws_sqs_role_arn)
        if azure_storage_queue_primary_uri is not None:
            pulumi.set(__self__, "azure_storage_queue_primary_uri", azure_storage_queue_primary_uri)
        if azure_tenant_id is not None:
            pulumi.set(__self__, "azure_tenant_id", azure_tenant_id)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if created_on is not None:
            pulumi.set(__self__, "created_on", created_on)
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if gcp_pubsub_subscription_name is not None:
            pulumi.set(__self__, "gcp_pubsub_subscription_name", gcp_pubsub_subscription_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_provider is not None:
            pulumi.set(__self__, "notification_provider", notification_provider)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="awsSqsArn")
    def aws_sqs_arn(self) -> Optional[pulumi.Input[str]]:
        """
        AWS SQS queue ARN for notification integration to connect to
        """
        return pulumi.get(self, "aws_sqs_arn")

    @aws_sqs_arn.setter
    def aws_sqs_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_sqs_arn", value)

    @property
    @pulumi.getter(name="awsSqsExternalId")
    def aws_sqs_external_id(self) -> Optional[pulumi.Input[str]]:
        """
        The external ID that Snowflake will use when assuming the AWS role
        """
        return pulumi.get(self, "aws_sqs_external_id")

    @aws_sqs_external_id.setter
    def aws_sqs_external_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_sqs_external_id", value)

    @property
    @pulumi.getter(name="awsSqsRoleArn")
    def aws_sqs_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        AWS IAM role ARN for notification integration to assume
        """
        return pulumi.get(self, "aws_sqs_role_arn")

    @aws_sqs_role_arn.setter
    def aws_sqs_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_sqs_role_arn", value)

    @property
    @pulumi.getter(name="azureStorageQueuePrimaryUri")
    def azure_storage_queue_primary_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The queue ID for the Azure Queue Storage queue created for Event Grid notifications
        """
        return pulumi.get(self, "azure_storage_queue_primary_uri")

    @azure_storage_queue_primary_uri.setter
    def azure_storage_queue_primary_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_storage_queue_primary_uri", value)

    @property
    @pulumi.getter(name="azureTenantId")
    def azure_tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Azure Active Directory tenant used for identity management
        """
        return pulumi.get(self, "azure_tenant_id")

    @azure_tenant_id.setter
    def azure_tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_tenant_id", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="createdOn")
    def created_on(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time when the notification integration was created.
        """
        return pulumi.get(self, "created_on")

    @created_on.setter
    def created_on(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_on", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="gcpPubsubSubscriptionName")
    def gcp_pubsub_subscription_name(self) -> Optional[pulumi.Input[str]]:
        """
        The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
        """
        return pulumi.get(self, "gcp_pubsub_subscription_name")

    @gcp_pubsub_subscription_name.setter
    def gcp_pubsub_subscription_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gcp_pubsub_subscription_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationProvider")
    def notification_provider(self) -> Optional[pulumi.Input[str]]:
        """
        The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS_SQS)
        """
        return pulumi.get(self, "notification_provider")

    @notification_provider.setter
    def notification_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_provider", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        A type of integration
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class NotificationIntegration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_sqs_arn: Optional[pulumi.Input[str]] = None,
                 aws_sqs_external_id: Optional[pulumi.Input[str]] = None,
                 aws_sqs_role_arn: Optional[pulumi.Input[str]] = None,
                 azure_storage_queue_primary_uri: Optional[pulumi.Input[str]] = None,
                 azure_tenant_id: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 gcp_pubsub_subscription_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_provider: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a NotificationIntegration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_sqs_arn: AWS SQS queue ARN for notification integration to connect to
        :param pulumi.Input[str] aws_sqs_external_id: The external ID that Snowflake will use when assuming the AWS role
        :param pulumi.Input[str] aws_sqs_role_arn: AWS IAM role ARN for notification integration to assume
        :param pulumi.Input[str] azure_storage_queue_primary_uri: The queue ID for the Azure Queue Storage queue created for Event Grid notifications
        :param pulumi.Input[str] azure_tenant_id: The ID of the Azure Active Directory tenant used for identity management
        :param pulumi.Input[str] direction: Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
        :param pulumi.Input[str] gcp_pubsub_subscription_name: The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
        :param pulumi.Input[str] notification_provider: The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS_SQS)
        :param pulumi.Input[str] type: A type of integration
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NotificationIntegrationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NotificationIntegration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NotificationIntegrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NotificationIntegrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_sqs_arn: Optional[pulumi.Input[str]] = None,
                 aws_sqs_external_id: Optional[pulumi.Input[str]] = None,
                 aws_sqs_role_arn: Optional[pulumi.Input[str]] = None,
                 azure_storage_queue_primary_uri: Optional[pulumi.Input[str]] = None,
                 azure_tenant_id: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 gcp_pubsub_subscription_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_provider: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NotificationIntegrationArgs.__new__(NotificationIntegrationArgs)

            __props__.__dict__["aws_sqs_arn"] = aws_sqs_arn
            __props__.__dict__["aws_sqs_external_id"] = aws_sqs_external_id
            __props__.__dict__["aws_sqs_role_arn"] = aws_sqs_role_arn
            __props__.__dict__["azure_storage_queue_primary_uri"] = azure_storage_queue_primary_uri
            __props__.__dict__["azure_tenant_id"] = azure_tenant_id
            __props__.__dict__["comment"] = comment
            __props__.__dict__["direction"] = direction
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["gcp_pubsub_subscription_name"] = gcp_pubsub_subscription_name
            __props__.__dict__["name"] = name
            __props__.__dict__["notification_provider"] = notification_provider
            __props__.__dict__["type"] = type
            __props__.__dict__["created_on"] = None
        super(NotificationIntegration, __self__).__init__(
            'snowflake:index/notificationIntegration:NotificationIntegration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aws_sqs_arn: Optional[pulumi.Input[str]] = None,
            aws_sqs_external_id: Optional[pulumi.Input[str]] = None,
            aws_sqs_role_arn: Optional[pulumi.Input[str]] = None,
            azure_storage_queue_primary_uri: Optional[pulumi.Input[str]] = None,
            azure_tenant_id: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            created_on: Optional[pulumi.Input[str]] = None,
            direction: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            gcp_pubsub_subscription_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notification_provider: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'NotificationIntegration':
        """
        Get an existing NotificationIntegration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_sqs_arn: AWS SQS queue ARN for notification integration to connect to
        :param pulumi.Input[str] aws_sqs_external_id: The external ID that Snowflake will use when assuming the AWS role
        :param pulumi.Input[str] aws_sqs_role_arn: AWS IAM role ARN for notification integration to assume
        :param pulumi.Input[str] azure_storage_queue_primary_uri: The queue ID for the Azure Queue Storage queue created for Event Grid notifications
        :param pulumi.Input[str] azure_tenant_id: The ID of the Azure Active Directory tenant used for identity management
        :param pulumi.Input[str] created_on: Date and time when the notification integration was created.
        :param pulumi.Input[str] direction: Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
        :param pulumi.Input[str] gcp_pubsub_subscription_name: The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
        :param pulumi.Input[str] notification_provider: The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS_SQS)
        :param pulumi.Input[str] type: A type of integration
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NotificationIntegrationState.__new__(_NotificationIntegrationState)

        __props__.__dict__["aws_sqs_arn"] = aws_sqs_arn
        __props__.__dict__["aws_sqs_external_id"] = aws_sqs_external_id
        __props__.__dict__["aws_sqs_role_arn"] = aws_sqs_role_arn
        __props__.__dict__["azure_storage_queue_primary_uri"] = azure_storage_queue_primary_uri
        __props__.__dict__["azure_tenant_id"] = azure_tenant_id
        __props__.__dict__["comment"] = comment
        __props__.__dict__["created_on"] = created_on
        __props__.__dict__["direction"] = direction
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["gcp_pubsub_subscription_name"] = gcp_pubsub_subscription_name
        __props__.__dict__["name"] = name
        __props__.__dict__["notification_provider"] = notification_provider
        __props__.__dict__["type"] = type
        return NotificationIntegration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsSqsArn")
    def aws_sqs_arn(self) -> pulumi.Output[Optional[str]]:
        """
        AWS SQS queue ARN for notification integration to connect to
        """
        return pulumi.get(self, "aws_sqs_arn")

    @property
    @pulumi.getter(name="awsSqsExternalId")
    def aws_sqs_external_id(self) -> pulumi.Output[Optional[str]]:
        """
        The external ID that Snowflake will use when assuming the AWS role
        """
        return pulumi.get(self, "aws_sqs_external_id")

    @property
    @pulumi.getter(name="awsSqsRoleArn")
    def aws_sqs_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        AWS IAM role ARN for notification integration to assume
        """
        return pulumi.get(self, "aws_sqs_role_arn")

    @property
    @pulumi.getter(name="azureStorageQueuePrimaryUri")
    def azure_storage_queue_primary_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The queue ID for the Azure Queue Storage queue created for Event Grid notifications
        """
        return pulumi.get(self, "azure_storage_queue_primary_uri")

    @property
    @pulumi.getter(name="azureTenantId")
    def azure_tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Azure Active Directory tenant used for identity management
        """
        return pulumi.get(self, "azure_tenant_id")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="createdOn")
    def created_on(self) -> pulumi.Output[str]:
        """
        Date and time when the notification integration was created.
        """
        return pulumi.get(self, "created_on")

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Output[Optional[str]]:
        """
        Direction of the cloud messaging with respect to Snowflake (required only for error notifications)
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="gcpPubsubSubscriptionName")
    def gcp_pubsub_subscription_name(self) -> pulumi.Output[Optional[str]]:
        """
        The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
        """
        return pulumi.get(self, "gcp_pubsub_subscription_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationProvider")
    def notification_provider(self) -> pulumi.Output[Optional[str]]:
        """
        The third-party cloud message queuing service (e.g. AZURE*STORAGE*QUEUE, AWS_SQS)
        """
        return pulumi.get(self, "notification_provider")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        A type of integration
        """
        return pulumi.get(self, "type")

