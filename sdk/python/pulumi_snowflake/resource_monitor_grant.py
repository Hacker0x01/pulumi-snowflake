# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ResourceMonitorGrantArgs', 'ResourceMonitorGrant']

@pulumi.input_type
class ResourceMonitorGrantArgs:
    def __init__(__self__, *,
                 monitor_name: pulumi.Input[str],
                 privilege: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 with_grant_option: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ResourceMonitorGrant resource.
        :param pulumi.Input[str] monitor_name: Identifier for the resource monitor; must be unique for your account.
        :param pulumi.Input[str] privilege: The privilege to grant on the resource monitor.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Grants privilege to these roles.
        :param pulumi.Input[bool] with_grant_option: When this is set to true, allows the recipient role to grant the privileges to other roles.
        """
        pulumi.set(__self__, "monitor_name", monitor_name)
        if privilege is not None:
            pulumi.set(__self__, "privilege", privilege)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if with_grant_option is not None:
            pulumi.set(__self__, "with_grant_option", with_grant_option)

    @property
    @pulumi.getter(name="monitorName")
    def monitor_name(self) -> pulumi.Input[str]:
        """
        Identifier for the resource monitor; must be unique for your account.
        """
        return pulumi.get(self, "monitor_name")

    @monitor_name.setter
    def monitor_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "monitor_name", value)

    @property
    @pulumi.getter
    def privilege(self) -> Optional[pulumi.Input[str]]:
        """
        The privilege to grant on the resource monitor.
        """
        return pulumi.get(self, "privilege")

    @privilege.setter
    def privilege(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "privilege", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Grants privilege to these roles.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter(name="withGrantOption")
    def with_grant_option(self) -> Optional[pulumi.Input[bool]]:
        """
        When this is set to true, allows the recipient role to grant the privileges to other roles.
        """
        return pulumi.get(self, "with_grant_option")

    @with_grant_option.setter
    def with_grant_option(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "with_grant_option", value)


@pulumi.input_type
class _ResourceMonitorGrantState:
    def __init__(__self__, *,
                 monitor_name: Optional[pulumi.Input[str]] = None,
                 privilege: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 with_grant_option: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering ResourceMonitorGrant resources.
        :param pulumi.Input[str] monitor_name: Identifier for the resource monitor; must be unique for your account.
        :param pulumi.Input[str] privilege: The privilege to grant on the resource monitor.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Grants privilege to these roles.
        :param pulumi.Input[bool] with_grant_option: When this is set to true, allows the recipient role to grant the privileges to other roles.
        """
        if monitor_name is not None:
            pulumi.set(__self__, "monitor_name", monitor_name)
        if privilege is not None:
            pulumi.set(__self__, "privilege", privilege)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if with_grant_option is not None:
            pulumi.set(__self__, "with_grant_option", with_grant_option)

    @property
    @pulumi.getter(name="monitorName")
    def monitor_name(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier for the resource monitor; must be unique for your account.
        """
        return pulumi.get(self, "monitor_name")

    @monitor_name.setter
    def monitor_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "monitor_name", value)

    @property
    @pulumi.getter
    def privilege(self) -> Optional[pulumi.Input[str]]:
        """
        The privilege to grant on the resource monitor.
        """
        return pulumi.get(self, "privilege")

    @privilege.setter
    def privilege(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "privilege", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Grants privilege to these roles.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter(name="withGrantOption")
    def with_grant_option(self) -> Optional[pulumi.Input[bool]]:
        """
        When this is set to true, allows the recipient role to grant the privileges to other roles.
        """
        return pulumi.get(self, "with_grant_option")

    @with_grant_option.setter
    def with_grant_option(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "with_grant_option", value)


class ResourceMonitorGrant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 monitor_name: Optional[pulumi.Input[str]] = None,
                 privilege: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 with_grant_option: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## Import

        ```sh
         $ pulumi import snowflake:index/resourceMonitorGrant:ResourceMonitorGrant example name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] monitor_name: Identifier for the resource monitor; must be unique for your account.
        :param pulumi.Input[str] privilege: The privilege to grant on the resource monitor.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Grants privilege to these roles.
        :param pulumi.Input[bool] with_grant_option: When this is set to true, allows the recipient role to grant the privileges to other roles.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourceMonitorGrantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        ```sh
         $ pulumi import snowflake:index/resourceMonitorGrant:ResourceMonitorGrant example name
        ```

        :param str resource_name: The name of the resource.
        :param ResourceMonitorGrantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceMonitorGrantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 monitor_name: Optional[pulumi.Input[str]] = None,
                 privilege: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 with_grant_option: Optional[pulumi.Input[bool]] = None,
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
            __props__ = ResourceMonitorGrantArgs.__new__(ResourceMonitorGrantArgs)

            if monitor_name is None and not opts.urn:
                raise TypeError("Missing required property 'monitor_name'")
            __props__.__dict__["monitor_name"] = monitor_name
            __props__.__dict__["privilege"] = privilege
            __props__.__dict__["roles"] = roles
            __props__.__dict__["with_grant_option"] = with_grant_option
        super(ResourceMonitorGrant, __self__).__init__(
            'snowflake:index/resourceMonitorGrant:ResourceMonitorGrant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            monitor_name: Optional[pulumi.Input[str]] = None,
            privilege: Optional[pulumi.Input[str]] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            with_grant_option: Optional[pulumi.Input[bool]] = None) -> 'ResourceMonitorGrant':
        """
        Get an existing ResourceMonitorGrant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] monitor_name: Identifier for the resource monitor; must be unique for your account.
        :param pulumi.Input[str] privilege: The privilege to grant on the resource monitor.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Grants privilege to these roles.
        :param pulumi.Input[bool] with_grant_option: When this is set to true, allows the recipient role to grant the privileges to other roles.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResourceMonitorGrantState.__new__(_ResourceMonitorGrantState)

        __props__.__dict__["monitor_name"] = monitor_name
        __props__.__dict__["privilege"] = privilege
        __props__.__dict__["roles"] = roles
        __props__.__dict__["with_grant_option"] = with_grant_option
        return ResourceMonitorGrant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="monitorName")
    def monitor_name(self) -> pulumi.Output[str]:
        """
        Identifier for the resource monitor; must be unique for your account.
        """
        return pulumi.get(self, "monitor_name")

    @property
    @pulumi.getter
    def privilege(self) -> pulumi.Output[Optional[str]]:
        """
        The privilege to grant on the resource monitor.
        """
        return pulumi.get(self, "privilege")

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Grants privilege to these roles.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="withGrantOption")
    def with_grant_option(self) -> pulumi.Output[Optional[bool]]:
        """
        When this is set to true, allows the recipient role to grant the privileges to other roles.
        """
        return pulumi.get(self, "with_grant_option")

