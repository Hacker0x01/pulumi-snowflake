# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApiIntegrationArgs', 'ApiIntegration']

@pulumi.input_type
class ApiIntegrationArgs:
    def __init__(__self__, *,
                 api_allowed_prefixes: pulumi.Input[Sequence[pulumi.Input[str]]],
                 api_provider: pulumi.Input[str],
                 api_aws_role_arn: Optional[pulumi.Input[str]] = None,
                 api_blocked_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 azure_ad_application_id: Optional[pulumi.Input[str]] = None,
                 azure_tenant_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApiIntegration resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_allowed_prefixes: Explicitly limits external functions that use the integration to reference one or more HTTPS proxy service endpoints and resources within those proxies.
        :param pulumi.Input[str] api_provider: Specifies the HTTPS proxy service type.
        :param pulumi.Input[str] api_aws_role_arn: ARN of a cloud platform role.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_blocked_prefixes: Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
        :param pulumi.Input[str] azure_ad_application_id: The 'Application (client) id' of the Azure AD app for your remote service.
        :param pulumi.Input[str] azure_tenant_id: Specifies the ID for your Office 365 tenant that all Azure API Management instances belong to.
        :param pulumi.Input[bool] enabled: Specifies whether this API integration is enabled or disabled. If the API integration is disabled, any external function that relies on it will not work.
        :param pulumi.Input[str] name: Specifies the name of the API integration. This name follows the rules for Object Identifiers. The name should be unique among api integrations in your account.
        """
        pulumi.set(__self__, "api_allowed_prefixes", api_allowed_prefixes)
        pulumi.set(__self__, "api_provider", api_provider)
        if api_aws_role_arn is not None:
            pulumi.set(__self__, "api_aws_role_arn", api_aws_role_arn)
        if api_blocked_prefixes is not None:
            pulumi.set(__self__, "api_blocked_prefixes", api_blocked_prefixes)
        if azure_ad_application_id is not None:
            pulumi.set(__self__, "azure_ad_application_id", azure_ad_application_id)
        if azure_tenant_id is not None:
            pulumi.set(__self__, "azure_tenant_id", azure_tenant_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="apiAllowedPrefixes")
    def api_allowed_prefixes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Explicitly limits external functions that use the integration to reference one or more HTTPS proxy service endpoints and resources within those proxies.
        """
        return pulumi.get(self, "api_allowed_prefixes")

    @api_allowed_prefixes.setter
    def api_allowed_prefixes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "api_allowed_prefixes", value)

    @property
    @pulumi.getter(name="apiProvider")
    def api_provider(self) -> pulumi.Input[str]:
        """
        Specifies the HTTPS proxy service type.
        """
        return pulumi.get(self, "api_provider")

    @api_provider.setter
    def api_provider(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_provider", value)

    @property
    @pulumi.getter(name="apiAwsRoleArn")
    def api_aws_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of a cloud platform role.
        """
        return pulumi.get(self, "api_aws_role_arn")

    @api_aws_role_arn.setter
    def api_aws_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_aws_role_arn", value)

    @property
    @pulumi.getter(name="apiBlockedPrefixes")
    def api_blocked_prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
        """
        return pulumi.get(self, "api_blocked_prefixes")

    @api_blocked_prefixes.setter
    def api_blocked_prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "api_blocked_prefixes", value)

    @property
    @pulumi.getter(name="azureAdApplicationId")
    def azure_ad_application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The 'Application (client) id' of the Azure AD app for your remote service.
        """
        return pulumi.get(self, "azure_ad_application_id")

    @azure_ad_application_id.setter
    def azure_ad_application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_ad_application_id", value)

    @property
    @pulumi.getter(name="azureTenantId")
    def azure_tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID for your Office 365 tenant that all Azure API Management instances belong to.
        """
        return pulumi.get(self, "azure_tenant_id")

    @azure_tenant_id.setter
    def azure_tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_tenant_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether this API integration is enabled or disabled. If the API integration is disabled, any external function that relies on it will not work.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the API integration. This name follows the rules for Object Identifiers. The name should be unique among api integrations in your account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ApiIntegrationState:
    def __init__(__self__, *,
                 api_allowed_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 api_aws_external_id: Optional[pulumi.Input[str]] = None,
                 api_aws_iam_user_arn: Optional[pulumi.Input[str]] = None,
                 api_aws_role_arn: Optional[pulumi.Input[str]] = None,
                 api_blocked_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 api_provider: Optional[pulumi.Input[str]] = None,
                 azure_ad_application_id: Optional[pulumi.Input[str]] = None,
                 azure_consent_url: Optional[pulumi.Input[str]] = None,
                 azure_multi_tenant_app_name: Optional[pulumi.Input[str]] = None,
                 azure_tenant_id: Optional[pulumi.Input[str]] = None,
                 created_on: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApiIntegration resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_allowed_prefixes: Explicitly limits external functions that use the integration to reference one or more HTTPS proxy service endpoints and resources within those proxies.
        :param pulumi.Input[str] api_aws_external_id: The external ID that Snowflake will use when assuming the AWS role.
        :param pulumi.Input[str] api_aws_iam_user_arn: The Snowflake user that will attempt to assume the AWS role.
        :param pulumi.Input[str] api_aws_role_arn: ARN of a cloud platform role.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_blocked_prefixes: Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
        :param pulumi.Input[str] api_provider: Specifies the HTTPS proxy service type.
        :param pulumi.Input[str] azure_ad_application_id: The 'Application (client) id' of the Azure AD app for your remote service.
        :param pulumi.Input[str] azure_tenant_id: Specifies the ID for your Office 365 tenant that all Azure API Management instances belong to.
        :param pulumi.Input[str] created_on: Date and time when the API integration was created.
        :param pulumi.Input[bool] enabled: Specifies whether this API integration is enabled or disabled. If the API integration is disabled, any external function that relies on it will not work.
        :param pulumi.Input[str] name: Specifies the name of the API integration. This name follows the rules for Object Identifiers. The name should be unique among api integrations in your account.
        """
        if api_allowed_prefixes is not None:
            pulumi.set(__self__, "api_allowed_prefixes", api_allowed_prefixes)
        if api_aws_external_id is not None:
            pulumi.set(__self__, "api_aws_external_id", api_aws_external_id)
        if api_aws_iam_user_arn is not None:
            pulumi.set(__self__, "api_aws_iam_user_arn", api_aws_iam_user_arn)
        if api_aws_role_arn is not None:
            pulumi.set(__self__, "api_aws_role_arn", api_aws_role_arn)
        if api_blocked_prefixes is not None:
            pulumi.set(__self__, "api_blocked_prefixes", api_blocked_prefixes)
        if api_provider is not None:
            pulumi.set(__self__, "api_provider", api_provider)
        if azure_ad_application_id is not None:
            pulumi.set(__self__, "azure_ad_application_id", azure_ad_application_id)
        if azure_consent_url is not None:
            pulumi.set(__self__, "azure_consent_url", azure_consent_url)
        if azure_multi_tenant_app_name is not None:
            pulumi.set(__self__, "azure_multi_tenant_app_name", azure_multi_tenant_app_name)
        if azure_tenant_id is not None:
            pulumi.set(__self__, "azure_tenant_id", azure_tenant_id)
        if created_on is not None:
            pulumi.set(__self__, "created_on", created_on)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="apiAllowedPrefixes")
    def api_allowed_prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Explicitly limits external functions that use the integration to reference one or more HTTPS proxy service endpoints and resources within those proxies.
        """
        return pulumi.get(self, "api_allowed_prefixes")

    @api_allowed_prefixes.setter
    def api_allowed_prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "api_allowed_prefixes", value)

    @property
    @pulumi.getter(name="apiAwsExternalId")
    def api_aws_external_id(self) -> Optional[pulumi.Input[str]]:
        """
        The external ID that Snowflake will use when assuming the AWS role.
        """
        return pulumi.get(self, "api_aws_external_id")

    @api_aws_external_id.setter
    def api_aws_external_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_aws_external_id", value)

    @property
    @pulumi.getter(name="apiAwsIamUserArn")
    def api_aws_iam_user_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Snowflake user that will attempt to assume the AWS role.
        """
        return pulumi.get(self, "api_aws_iam_user_arn")

    @api_aws_iam_user_arn.setter
    def api_aws_iam_user_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_aws_iam_user_arn", value)

    @property
    @pulumi.getter(name="apiAwsRoleArn")
    def api_aws_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of a cloud platform role.
        """
        return pulumi.get(self, "api_aws_role_arn")

    @api_aws_role_arn.setter
    def api_aws_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_aws_role_arn", value)

    @property
    @pulumi.getter(name="apiBlockedPrefixes")
    def api_blocked_prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
        """
        return pulumi.get(self, "api_blocked_prefixes")

    @api_blocked_prefixes.setter
    def api_blocked_prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "api_blocked_prefixes", value)

    @property
    @pulumi.getter(name="apiProvider")
    def api_provider(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the HTTPS proxy service type.
        """
        return pulumi.get(self, "api_provider")

    @api_provider.setter
    def api_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_provider", value)

    @property
    @pulumi.getter(name="azureAdApplicationId")
    def azure_ad_application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The 'Application (client) id' of the Azure AD app for your remote service.
        """
        return pulumi.get(self, "azure_ad_application_id")

    @azure_ad_application_id.setter
    def azure_ad_application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_ad_application_id", value)

    @property
    @pulumi.getter(name="azureConsentUrl")
    def azure_consent_url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "azure_consent_url")

    @azure_consent_url.setter
    def azure_consent_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_consent_url", value)

    @property
    @pulumi.getter(name="azureMultiTenantAppName")
    def azure_multi_tenant_app_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "azure_multi_tenant_app_name")

    @azure_multi_tenant_app_name.setter
    def azure_multi_tenant_app_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_multi_tenant_app_name", value)

    @property
    @pulumi.getter(name="azureTenantId")
    def azure_tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID for your Office 365 tenant that all Azure API Management instances belong to.
        """
        return pulumi.get(self, "azure_tenant_id")

    @azure_tenant_id.setter
    def azure_tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "azure_tenant_id", value)

    @property
    @pulumi.getter(name="createdOn")
    def created_on(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time when the API integration was created.
        """
        return pulumi.get(self, "created_on")

    @created_on.setter
    def created_on(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_on", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether this API integration is enabled or disabled. If the API integration is disabled, any external function that relies on it will not work.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the API integration. This name follows the rules for Object Identifiers. The name should be unique among api integrations in your account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class ApiIntegration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_allowed_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 api_aws_role_arn: Optional[pulumi.Input[str]] = None,
                 api_blocked_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 api_provider: Optional[pulumi.Input[str]] = None,
                 azure_ad_application_id: Optional[pulumi.Input[str]] = None,
                 azure_tenant_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_snowflake as snowflake

        api_integration = snowflake.ApiIntegration("apiIntegration",
            api_allowed_prefixes=["https://123456.execute-api.us-west-2.amazonaws.com/prod/"],
            api_aws_role_arn="arn:aws:iam::000000000001:/role/test",
            api_provider="aws_api_gateway",
            enabled=True)
        ```

        ## Import

        ```sh
         $ pulumi import snowflake:index/apiIntegration:ApiIntegration example name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_allowed_prefixes: Explicitly limits external functions that use the integration to reference one or more HTTPS proxy service endpoints and resources within those proxies.
        :param pulumi.Input[str] api_aws_role_arn: ARN of a cloud platform role.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_blocked_prefixes: Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
        :param pulumi.Input[str] api_provider: Specifies the HTTPS proxy service type.
        :param pulumi.Input[str] azure_ad_application_id: The 'Application (client) id' of the Azure AD app for your remote service.
        :param pulumi.Input[str] azure_tenant_id: Specifies the ID for your Office 365 tenant that all Azure API Management instances belong to.
        :param pulumi.Input[bool] enabled: Specifies whether this API integration is enabled or disabled. If the API integration is disabled, any external function that relies on it will not work.
        :param pulumi.Input[str] name: Specifies the name of the API integration. This name follows the rules for Object Identifiers. The name should be unique among api integrations in your account.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApiIntegrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_snowflake as snowflake

        api_integration = snowflake.ApiIntegration("apiIntegration",
            api_allowed_prefixes=["https://123456.execute-api.us-west-2.amazonaws.com/prod/"],
            api_aws_role_arn="arn:aws:iam::000000000001:/role/test",
            api_provider="aws_api_gateway",
            enabled=True)
        ```

        ## Import

        ```sh
         $ pulumi import snowflake:index/apiIntegration:ApiIntegration example name
        ```

        :param str resource_name: The name of the resource.
        :param ApiIntegrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApiIntegrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_allowed_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 api_aws_role_arn: Optional[pulumi.Input[str]] = None,
                 api_blocked_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 api_provider: Optional[pulumi.Input[str]] = None,
                 azure_ad_application_id: Optional[pulumi.Input[str]] = None,
                 azure_tenant_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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
            __props__ = ApiIntegrationArgs.__new__(ApiIntegrationArgs)

            if api_allowed_prefixes is None and not opts.urn:
                raise TypeError("Missing required property 'api_allowed_prefixes'")
            __props__.__dict__["api_allowed_prefixes"] = api_allowed_prefixes
            __props__.__dict__["api_aws_role_arn"] = api_aws_role_arn
            __props__.__dict__["api_blocked_prefixes"] = api_blocked_prefixes
            if api_provider is None and not opts.urn:
                raise TypeError("Missing required property 'api_provider'")
            __props__.__dict__["api_provider"] = api_provider
            __props__.__dict__["azure_ad_application_id"] = azure_ad_application_id
            __props__.__dict__["azure_tenant_id"] = azure_tenant_id
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["api_aws_external_id"] = None
            __props__.__dict__["api_aws_iam_user_arn"] = None
            __props__.__dict__["azure_consent_url"] = None
            __props__.__dict__["azure_multi_tenant_app_name"] = None
            __props__.__dict__["created_on"] = None
        super(ApiIntegration, __self__).__init__(
            'snowflake:index/apiIntegration:ApiIntegration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_allowed_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            api_aws_external_id: Optional[pulumi.Input[str]] = None,
            api_aws_iam_user_arn: Optional[pulumi.Input[str]] = None,
            api_aws_role_arn: Optional[pulumi.Input[str]] = None,
            api_blocked_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            api_provider: Optional[pulumi.Input[str]] = None,
            azure_ad_application_id: Optional[pulumi.Input[str]] = None,
            azure_consent_url: Optional[pulumi.Input[str]] = None,
            azure_multi_tenant_app_name: Optional[pulumi.Input[str]] = None,
            azure_tenant_id: Optional[pulumi.Input[str]] = None,
            created_on: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'ApiIntegration':
        """
        Get an existing ApiIntegration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_allowed_prefixes: Explicitly limits external functions that use the integration to reference one or more HTTPS proxy service endpoints and resources within those proxies.
        :param pulumi.Input[str] api_aws_external_id: The external ID that Snowflake will use when assuming the AWS role.
        :param pulumi.Input[str] api_aws_iam_user_arn: The Snowflake user that will attempt to assume the AWS role.
        :param pulumi.Input[str] api_aws_role_arn: ARN of a cloud platform role.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_blocked_prefixes: Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
        :param pulumi.Input[str] api_provider: Specifies the HTTPS proxy service type.
        :param pulumi.Input[str] azure_ad_application_id: The 'Application (client) id' of the Azure AD app for your remote service.
        :param pulumi.Input[str] azure_tenant_id: Specifies the ID for your Office 365 tenant that all Azure API Management instances belong to.
        :param pulumi.Input[str] created_on: Date and time when the API integration was created.
        :param pulumi.Input[bool] enabled: Specifies whether this API integration is enabled or disabled. If the API integration is disabled, any external function that relies on it will not work.
        :param pulumi.Input[str] name: Specifies the name of the API integration. This name follows the rules for Object Identifiers. The name should be unique among api integrations in your account.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApiIntegrationState.__new__(_ApiIntegrationState)

        __props__.__dict__["api_allowed_prefixes"] = api_allowed_prefixes
        __props__.__dict__["api_aws_external_id"] = api_aws_external_id
        __props__.__dict__["api_aws_iam_user_arn"] = api_aws_iam_user_arn
        __props__.__dict__["api_aws_role_arn"] = api_aws_role_arn
        __props__.__dict__["api_blocked_prefixes"] = api_blocked_prefixes
        __props__.__dict__["api_provider"] = api_provider
        __props__.__dict__["azure_ad_application_id"] = azure_ad_application_id
        __props__.__dict__["azure_consent_url"] = azure_consent_url
        __props__.__dict__["azure_multi_tenant_app_name"] = azure_multi_tenant_app_name
        __props__.__dict__["azure_tenant_id"] = azure_tenant_id
        __props__.__dict__["created_on"] = created_on
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["name"] = name
        return ApiIntegration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiAllowedPrefixes")
    def api_allowed_prefixes(self) -> pulumi.Output[Sequence[str]]:
        """
        Explicitly limits external functions that use the integration to reference one or more HTTPS proxy service endpoints and resources within those proxies.
        """
        return pulumi.get(self, "api_allowed_prefixes")

    @property
    @pulumi.getter(name="apiAwsExternalId")
    def api_aws_external_id(self) -> pulumi.Output[str]:
        """
        The external ID that Snowflake will use when assuming the AWS role.
        """
        return pulumi.get(self, "api_aws_external_id")

    @property
    @pulumi.getter(name="apiAwsIamUserArn")
    def api_aws_iam_user_arn(self) -> pulumi.Output[str]:
        """
        The Snowflake user that will attempt to assume the AWS role.
        """
        return pulumi.get(self, "api_aws_iam_user_arn")

    @property
    @pulumi.getter(name="apiAwsRoleArn")
    def api_aws_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        ARN of a cloud platform role.
        """
        return pulumi.get(self, "api_aws_role_arn")

    @property
    @pulumi.getter(name="apiBlockedPrefixes")
    def api_blocked_prefixes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
        """
        return pulumi.get(self, "api_blocked_prefixes")

    @property
    @pulumi.getter(name="apiProvider")
    def api_provider(self) -> pulumi.Output[str]:
        """
        Specifies the HTTPS proxy service type.
        """
        return pulumi.get(self, "api_provider")

    @property
    @pulumi.getter(name="azureAdApplicationId")
    def azure_ad_application_id(self) -> pulumi.Output[Optional[str]]:
        """
        The 'Application (client) id' of the Azure AD app for your remote service.
        """
        return pulumi.get(self, "azure_ad_application_id")

    @property
    @pulumi.getter(name="azureConsentUrl")
    def azure_consent_url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "azure_consent_url")

    @property
    @pulumi.getter(name="azureMultiTenantAppName")
    def azure_multi_tenant_app_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "azure_multi_tenant_app_name")

    @property
    @pulumi.getter(name="azureTenantId")
    def azure_tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the ID for your Office 365 tenant that all Azure API Management instances belong to.
        """
        return pulumi.get(self, "azure_tenant_id")

    @property
    @pulumi.getter(name="createdOn")
    def created_on(self) -> pulumi.Output[str]:
        """
        Date and time when the API integration was created.
        """
        return pulumi.get(self, "created_on")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether this API integration is enabled or disabled. If the API integration is disabled, any external function that relies on it will not work.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the API integration. This name follows the rules for Object Identifiers. The name should be unique among api integrations in your account.
        """
        return pulumi.get(self, "name")

