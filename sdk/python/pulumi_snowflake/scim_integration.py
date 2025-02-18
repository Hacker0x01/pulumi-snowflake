# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ScimIntegrationArgs', 'ScimIntegration']

@pulumi.input_type
class ScimIntegrationArgs:
    def __init__(__self__, *,
                 provisioner_role: pulumi.Input[str],
                 scim_client: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 network_policy: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ScimIntegration resource.
        :param pulumi.Input[str] provisioner_role: Specify the SCIM role in Snowflake that owns any users and roles that are imported from the identity provider into Snowflake using SCIM.
        :param pulumi.Input[str] scim_client: Specifies the client type for the scim integration
        :param pulumi.Input[str] name: Specifies the name of the SCIM integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
        :param pulumi.Input[str] network_policy: Specifies an existing network policy active for your account. The network policy restricts the list of user IP addresses when exchanging an authorization code for an access or refresh token and when using a refresh token to obtain a new access token. If this parameter is not set, the network policy for the account (if any) is used instead.
        """
        pulumi.set(__self__, "provisioner_role", provisioner_role)
        pulumi.set(__self__, "scim_client", scim_client)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_policy is not None:
            pulumi.set(__self__, "network_policy", network_policy)

    @property
    @pulumi.getter(name="provisionerRole")
    def provisioner_role(self) -> pulumi.Input[str]:
        """
        Specify the SCIM role in Snowflake that owns any users and roles that are imported from the identity provider into Snowflake using SCIM.
        """
        return pulumi.get(self, "provisioner_role")

    @provisioner_role.setter
    def provisioner_role(self, value: pulumi.Input[str]):
        pulumi.set(self, "provisioner_role", value)

    @property
    @pulumi.getter(name="scimClient")
    def scim_client(self) -> pulumi.Input[str]:
        """
        Specifies the client type for the scim integration
        """
        return pulumi.get(self, "scim_client")

    @scim_client.setter
    def scim_client(self, value: pulumi.Input[str]):
        pulumi.set(self, "scim_client", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the SCIM integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkPolicy")
    def network_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies an existing network policy active for your account. The network policy restricts the list of user IP addresses when exchanging an authorization code for an access or refresh token and when using a refresh token to obtain a new access token. If this parameter is not set, the network policy for the account (if any) is used instead.
        """
        return pulumi.get(self, "network_policy")

    @network_policy.setter
    def network_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_policy", value)


@pulumi.input_type
class _ScimIntegrationState:
    def __init__(__self__, *,
                 created_on: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_policy: Optional[pulumi.Input[str]] = None,
                 provisioner_role: Optional[pulumi.Input[str]] = None,
                 scim_client: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ScimIntegration resources.
        :param pulumi.Input[str] created_on: Date and time when the SCIM integration was created.
        :param pulumi.Input[str] name: Specifies the name of the SCIM integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
        :param pulumi.Input[str] network_policy: Specifies an existing network policy active for your account. The network policy restricts the list of user IP addresses when exchanging an authorization code for an access or refresh token and when using a refresh token to obtain a new access token. If this parameter is not set, the network policy for the account (if any) is used instead.
        :param pulumi.Input[str] provisioner_role: Specify the SCIM role in Snowflake that owns any users and roles that are imported from the identity provider into Snowflake using SCIM.
        :param pulumi.Input[str] scim_client: Specifies the client type for the scim integration
        """
        if created_on is not None:
            pulumi.set(__self__, "created_on", created_on)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_policy is not None:
            pulumi.set(__self__, "network_policy", network_policy)
        if provisioner_role is not None:
            pulumi.set(__self__, "provisioner_role", provisioner_role)
        if scim_client is not None:
            pulumi.set(__self__, "scim_client", scim_client)

    @property
    @pulumi.getter(name="createdOn")
    def created_on(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time when the SCIM integration was created.
        """
        return pulumi.get(self, "created_on")

    @created_on.setter
    def created_on(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_on", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the SCIM integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkPolicy")
    def network_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies an existing network policy active for your account. The network policy restricts the list of user IP addresses when exchanging an authorization code for an access or refresh token and when using a refresh token to obtain a new access token. If this parameter is not set, the network policy for the account (if any) is used instead.
        """
        return pulumi.get(self, "network_policy")

    @network_policy.setter
    def network_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_policy", value)

    @property
    @pulumi.getter(name="provisionerRole")
    def provisioner_role(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the SCIM role in Snowflake that owns any users and roles that are imported from the identity provider into Snowflake using SCIM.
        """
        return pulumi.get(self, "provisioner_role")

    @provisioner_role.setter
    def provisioner_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provisioner_role", value)

    @property
    @pulumi.getter(name="scimClient")
    def scim_client(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the client type for the scim integration
        """
        return pulumi.get(self, "scim_client")

    @scim_client.setter
    def scim_client(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scim_client", value)


class ScimIntegration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_policy: Optional[pulumi.Input[str]] = None,
                 provisioner_role: Optional[pulumi.Input[str]] = None,
                 scim_client: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_snowflake as snowflake

        aad = snowflake.ScimIntegration("aad",
            network_policy="AAD_NETWORK_POLICY",
            provisioner_role="AAD_PROVISIONER",
            scim_client="AZURE")
        ```

        ## Import

        ```sh
         $ pulumi import snowflake:index/scimIntegration:ScimIntegration example name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Specifies the name of the SCIM integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
        :param pulumi.Input[str] network_policy: Specifies an existing network policy active for your account. The network policy restricts the list of user IP addresses when exchanging an authorization code for an access or refresh token and when using a refresh token to obtain a new access token. If this parameter is not set, the network policy for the account (if any) is used instead.
        :param pulumi.Input[str] provisioner_role: Specify the SCIM role in Snowflake that owns any users and roles that are imported from the identity provider into Snowflake using SCIM.
        :param pulumi.Input[str] scim_client: Specifies the client type for the scim integration
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScimIntegrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_snowflake as snowflake

        aad = snowflake.ScimIntegration("aad",
            network_policy="AAD_NETWORK_POLICY",
            provisioner_role="AAD_PROVISIONER",
            scim_client="AZURE")
        ```

        ## Import

        ```sh
         $ pulumi import snowflake:index/scimIntegration:ScimIntegration example name
        ```

        :param str resource_name: The name of the resource.
        :param ScimIntegrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScimIntegrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_policy: Optional[pulumi.Input[str]] = None,
                 provisioner_role: Optional[pulumi.Input[str]] = None,
                 scim_client: Optional[pulumi.Input[str]] = None,
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
            __props__ = ScimIntegrationArgs.__new__(ScimIntegrationArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["network_policy"] = network_policy
            if provisioner_role is None and not opts.urn:
                raise TypeError("Missing required property 'provisioner_role'")
            __props__.__dict__["provisioner_role"] = provisioner_role
            if scim_client is None and not opts.urn:
                raise TypeError("Missing required property 'scim_client'")
            __props__.__dict__["scim_client"] = scim_client
            __props__.__dict__["created_on"] = None
        super(ScimIntegration, __self__).__init__(
            'snowflake:index/scimIntegration:ScimIntegration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_on: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_policy: Optional[pulumi.Input[str]] = None,
            provisioner_role: Optional[pulumi.Input[str]] = None,
            scim_client: Optional[pulumi.Input[str]] = None) -> 'ScimIntegration':
        """
        Get an existing ScimIntegration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_on: Date and time when the SCIM integration was created.
        :param pulumi.Input[str] name: Specifies the name of the SCIM integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
        :param pulumi.Input[str] network_policy: Specifies an existing network policy active for your account. The network policy restricts the list of user IP addresses when exchanging an authorization code for an access or refresh token and when using a refresh token to obtain a new access token. If this parameter is not set, the network policy for the account (if any) is used instead.
        :param pulumi.Input[str] provisioner_role: Specify the SCIM role in Snowflake that owns any users and roles that are imported from the identity provider into Snowflake using SCIM.
        :param pulumi.Input[str] scim_client: Specifies the client type for the scim integration
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ScimIntegrationState.__new__(_ScimIntegrationState)

        __props__.__dict__["created_on"] = created_on
        __props__.__dict__["name"] = name
        __props__.__dict__["network_policy"] = network_policy
        __props__.__dict__["provisioner_role"] = provisioner_role
        __props__.__dict__["scim_client"] = scim_client
        return ScimIntegration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdOn")
    def created_on(self) -> pulumi.Output[str]:
        """
        Date and time when the SCIM integration was created.
        """
        return pulumi.get(self, "created_on")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the SCIM integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkPolicy")
    def network_policy(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies an existing network policy active for your account. The network policy restricts the list of user IP addresses when exchanging an authorization code for an access or refresh token and when using a refresh token to obtain a new access token. If this parameter is not set, the network policy for the account (if any) is used instead.
        """
        return pulumi.get(self, "network_policy")

    @property
    @pulumi.getter(name="provisionerRole")
    def provisioner_role(self) -> pulumi.Output[str]:
        """
        Specify the SCIM role in Snowflake that owns any users and roles that are imported from the identity provider into Snowflake using SCIM.
        """
        return pulumi.get(self, "provisioner_role")

    @property
    @pulumi.getter(name="scimClient")
    def scim_client(self) -> pulumi.Output[str]:
        """
        Specifies the client type for the scim integration
        """
        return pulumi.get(self, "scim_client")

