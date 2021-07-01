# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ShareArgs', 'Share']

@pulumi.input_type
class ShareArgs:
    def __init__(__self__, *,
                 accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Share resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] accounts: A list of accounts to be added to the share.
        :param pulumi.Input[str] comment: Specifies a comment for the managed account.
        :param pulumi.Input[str] name: Specifies the identifier for the share; must be unique for the account in which the share is created.
        """
        if accounts is not None:
            pulumi.set(__self__, "accounts", accounts)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of accounts to be added to the share.
        """
        return pulumi.get(self, "accounts")

    @accounts.setter
    def accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "accounts", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a comment for the managed account.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the identifier for the share; must be unique for the account in which the share is created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ShareState:
    def __init__(__self__, *,
                 accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Share resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] accounts: A list of accounts to be added to the share.
        :param pulumi.Input[str] comment: Specifies a comment for the managed account.
        :param pulumi.Input[str] name: Specifies the identifier for the share; must be unique for the account in which the share is created.
        """
        if accounts is not None:
            pulumi.set(__self__, "accounts", accounts)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of accounts to be added to the share.
        """
        return pulumi.get(self, "accounts")

    @accounts.setter
    def accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "accounts", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a comment for the managed account.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the identifier for the share; must be unique for the account in which the share is created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class Share(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        ```sh
         $ pulumi import snowflake:index/share:Share example name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] accounts: A list of accounts to be added to the share.
        :param pulumi.Input[str] comment: Specifies a comment for the managed account.
        :param pulumi.Input[str] name: Specifies the identifier for the share; must be unique for the account in which the share is created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ShareArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        ```sh
         $ pulumi import snowflake:index/share:Share example name
        ```

        :param str resource_name: The name of the resource.
        :param ShareArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ShareArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
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
            __props__ = ShareArgs.__new__(ShareArgs)

            __props__.__dict__["accounts"] = accounts
            __props__.__dict__["comment"] = comment
            __props__.__dict__["name"] = name
        super(Share, __self__).__init__(
            'snowflake:index/share:Share',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'Share':
        """
        Get an existing Share resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] accounts: A list of accounts to be added to the share.
        :param pulumi.Input[str] comment: Specifies a comment for the managed account.
        :param pulumi.Input[str] name: Specifies the identifier for the share; must be unique for the account in which the share is created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ShareState.__new__(_ShareState)

        __props__.__dict__["accounts"] = accounts
        __props__.__dict__["comment"] = comment
        __props__.__dict__["name"] = name
        return Share(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accounts(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of accounts to be added to the share.
        """
        return pulumi.get(self, "accounts")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a comment for the managed account.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the identifier for the share; must be unique for the account in which the share is created.
        """
        return pulumi.get(self, "name")

