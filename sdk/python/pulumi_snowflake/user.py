# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_namespace: Optional[pulumi.Input[str]] = None,
                 default_role: Optional[pulumi.Input[str]] = None,
                 default_warehouse: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 login_name: Optional[pulumi.Input[str]] = None,
                 must_change_password: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 rsa_public_key: Optional[pulumi.Input[str]] = None,
                 rsa_public_key2: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[str] default_namespace: Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.
        :param pulumi.Input[str] default_role: Specifies the role that is active by default for the user’s session upon login.
        :param pulumi.Input[str] default_warehouse: Specifies the virtual warehouse that is active by default for the user’s session upon login.
        :param pulumi.Input[str] display_name: Name displayed for the user in the Snowflake web interface.
        :param pulumi.Input[str] email: Email address for the user.
        :param pulumi.Input[str] first_name: First name of the user.
        :param pulumi.Input[str] last_name: Last name of the user.
        :param pulumi.Input[str] login_name: The name users use to log in. If not supplied, snowflake will use name instead.
        :param pulumi.Input[bool] must_change_password: Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.
        :param pulumi.Input[str] name: Name of the user. Note that if you do not supply login*name this will be used as login*name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)
        :param pulumi.Input[str] password: **WARNING:** this will put the password in the terraform state file. Use carefully.
        :param pulumi.Input[str] rsa_public_key: Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
        :param pulumi.Input[str] rsa_public_key2: Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if default_namespace is not None:
            pulumi.set(__self__, "default_namespace", default_namespace)
        if default_role is not None:
            pulumi.set(__self__, "default_role", default_role)
        if default_warehouse is not None:
            pulumi.set(__self__, "default_warehouse", default_warehouse)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if first_name is not None:
            pulumi.set(__self__, "first_name", first_name)
        if last_name is not None:
            pulumi.set(__self__, "last_name", last_name)
        if login_name is not None:
            pulumi.set(__self__, "login_name", login_name)
        if must_change_password is not None:
            pulumi.set(__self__, "must_change_password", must_change_password)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if rsa_public_key is not None:
            pulumi.set(__self__, "rsa_public_key", rsa_public_key)
        if rsa_public_key2 is not None:
            pulumi.set(__self__, "rsa_public_key2", rsa_public_key2)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="defaultNamespace")
    def default_namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.
        """
        return pulumi.get(self, "default_namespace")

    @default_namespace.setter
    def default_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_namespace", value)

    @property
    @pulumi.getter(name="defaultRole")
    def default_role(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the role that is active by default for the user’s session upon login.
        """
        return pulumi.get(self, "default_role")

    @default_role.setter
    def default_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_role", value)

    @property
    @pulumi.getter(name="defaultWarehouse")
    def default_warehouse(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the virtual warehouse that is active by default for the user’s session upon login.
        """
        return pulumi.get(self, "default_warehouse")

    @default_warehouse.setter
    def default_warehouse(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_warehouse", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name displayed for the user in the Snowflake web interface.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Email address for the user.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> Optional[pulumi.Input[str]]:
        """
        First name of the user.
        """
        return pulumi.get(self, "first_name")

    @first_name.setter
    def first_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "first_name", value)

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> Optional[pulumi.Input[str]]:
        """
        Last name of the user.
        """
        return pulumi.get(self, "last_name")

    @last_name.setter
    def last_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_name", value)

    @property
    @pulumi.getter(name="loginName")
    def login_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name users use to log in. If not supplied, snowflake will use name instead.
        """
        return pulumi.get(self, "login_name")

    @login_name.setter
    def login_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login_name", value)

    @property
    @pulumi.getter(name="mustChangePassword")
    def must_change_password(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.
        """
        return pulumi.get(self, "must_change_password")

    @must_change_password.setter
    def must_change_password(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "must_change_password", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the user. Note that if you do not supply login*name this will be used as login*name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        **WARNING:** this will put the password in the terraform state file. Use carefully.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="rsaPublicKey")
    def rsa_public_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
        """
        return pulumi.get(self, "rsa_public_key")

    @rsa_public_key.setter
    def rsa_public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rsa_public_key", value)

    @property
    @pulumi.getter(name="rsaPublicKey2")
    def rsa_public_key2(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
        """
        return pulumi.get(self, "rsa_public_key2")

    @rsa_public_key2.setter
    def rsa_public_key2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rsa_public_key2", value)


@pulumi.input_type
class _UserState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_namespace: Optional[pulumi.Input[str]] = None,
                 default_role: Optional[pulumi.Input[str]] = None,
                 default_warehouse: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 has_rsa_public_key: Optional[pulumi.Input[bool]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 login_name: Optional[pulumi.Input[str]] = None,
                 must_change_password: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 rsa_public_key: Optional[pulumi.Input[str]] = None,
                 rsa_public_key2: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[str] default_namespace: Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.
        :param pulumi.Input[str] default_role: Specifies the role that is active by default for the user’s session upon login.
        :param pulumi.Input[str] default_warehouse: Specifies the virtual warehouse that is active by default for the user’s session upon login.
        :param pulumi.Input[str] display_name: Name displayed for the user in the Snowflake web interface.
        :param pulumi.Input[str] email: Email address for the user.
        :param pulumi.Input[str] first_name: First name of the user.
        :param pulumi.Input[bool] has_rsa_public_key: Will be true if user as an RSA key set.
        :param pulumi.Input[str] last_name: Last name of the user.
        :param pulumi.Input[str] login_name: The name users use to log in. If not supplied, snowflake will use name instead.
        :param pulumi.Input[bool] must_change_password: Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.
        :param pulumi.Input[str] name: Name of the user. Note that if you do not supply login*name this will be used as login*name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)
        :param pulumi.Input[str] password: **WARNING:** this will put the password in the terraform state file. Use carefully.
        :param pulumi.Input[str] rsa_public_key: Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
        :param pulumi.Input[str] rsa_public_key2: Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if default_namespace is not None:
            pulumi.set(__self__, "default_namespace", default_namespace)
        if default_role is not None:
            pulumi.set(__self__, "default_role", default_role)
        if default_warehouse is not None:
            pulumi.set(__self__, "default_warehouse", default_warehouse)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if first_name is not None:
            pulumi.set(__self__, "first_name", first_name)
        if has_rsa_public_key is not None:
            pulumi.set(__self__, "has_rsa_public_key", has_rsa_public_key)
        if last_name is not None:
            pulumi.set(__self__, "last_name", last_name)
        if login_name is not None:
            pulumi.set(__self__, "login_name", login_name)
        if must_change_password is not None:
            pulumi.set(__self__, "must_change_password", must_change_password)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if rsa_public_key is not None:
            pulumi.set(__self__, "rsa_public_key", rsa_public_key)
        if rsa_public_key2 is not None:
            pulumi.set(__self__, "rsa_public_key2", rsa_public_key2)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="defaultNamespace")
    def default_namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.
        """
        return pulumi.get(self, "default_namespace")

    @default_namespace.setter
    def default_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_namespace", value)

    @property
    @pulumi.getter(name="defaultRole")
    def default_role(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the role that is active by default for the user’s session upon login.
        """
        return pulumi.get(self, "default_role")

    @default_role.setter
    def default_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_role", value)

    @property
    @pulumi.getter(name="defaultWarehouse")
    def default_warehouse(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the virtual warehouse that is active by default for the user’s session upon login.
        """
        return pulumi.get(self, "default_warehouse")

    @default_warehouse.setter
    def default_warehouse(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_warehouse", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name displayed for the user in the Snowflake web interface.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Email address for the user.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> Optional[pulumi.Input[str]]:
        """
        First name of the user.
        """
        return pulumi.get(self, "first_name")

    @first_name.setter
    def first_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "first_name", value)

    @property
    @pulumi.getter(name="hasRsaPublicKey")
    def has_rsa_public_key(self) -> Optional[pulumi.Input[bool]]:
        """
        Will be true if user as an RSA key set.
        """
        return pulumi.get(self, "has_rsa_public_key")

    @has_rsa_public_key.setter
    def has_rsa_public_key(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "has_rsa_public_key", value)

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> Optional[pulumi.Input[str]]:
        """
        Last name of the user.
        """
        return pulumi.get(self, "last_name")

    @last_name.setter
    def last_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_name", value)

    @property
    @pulumi.getter(name="loginName")
    def login_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name users use to log in. If not supplied, snowflake will use name instead.
        """
        return pulumi.get(self, "login_name")

    @login_name.setter
    def login_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login_name", value)

    @property
    @pulumi.getter(name="mustChangePassword")
    def must_change_password(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.
        """
        return pulumi.get(self, "must_change_password")

    @must_change_password.setter
    def must_change_password(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "must_change_password", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the user. Note that if you do not supply login*name this will be used as login*name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        **WARNING:** this will put the password in the terraform state file. Use carefully.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="rsaPublicKey")
    def rsa_public_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
        """
        return pulumi.get(self, "rsa_public_key")

    @rsa_public_key.setter
    def rsa_public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rsa_public_key", value)

    @property
    @pulumi.getter(name="rsaPublicKey2")
    def rsa_public_key2(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
        """
        return pulumi.get(self, "rsa_public_key2")

    @rsa_public_key2.setter
    def rsa_public_key2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rsa_public_key2", value)


class User(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_namespace: Optional[pulumi.Input[str]] = None,
                 default_role: Optional[pulumi.Input[str]] = None,
                 default_warehouse: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 login_name: Optional[pulumi.Input[str]] = None,
                 must_change_password: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 rsa_public_key: Optional[pulumi.Input[str]] = None,
                 rsa_public_key2: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_snowflake as snowflake

        user = snowflake.User("user",
            comment="A user of snowflake.",
            default_role="role1",
            default_warehouse="warehouse",
            disabled=False,
            display_name="Snowflake User",
            email="user@snowflake.example",
            first_name="Snowflake",
            last_name="User",
            login_name="snowflake_user",
            must_change_password=False,
            password="secret",
            rsa_public_key="...",
            rsa_public_key2="...")
        ```

        ## Import

        ```sh
         $ pulumi import snowflake:index/user:User example userName
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_namespace: Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.
        :param pulumi.Input[str] default_role: Specifies the role that is active by default for the user’s session upon login.
        :param pulumi.Input[str] default_warehouse: Specifies the virtual warehouse that is active by default for the user’s session upon login.
        :param pulumi.Input[str] display_name: Name displayed for the user in the Snowflake web interface.
        :param pulumi.Input[str] email: Email address for the user.
        :param pulumi.Input[str] first_name: First name of the user.
        :param pulumi.Input[str] last_name: Last name of the user.
        :param pulumi.Input[str] login_name: The name users use to log in. If not supplied, snowflake will use name instead.
        :param pulumi.Input[bool] must_change_password: Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.
        :param pulumi.Input[str] name: Name of the user. Note that if you do not supply login*name this will be used as login*name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)
        :param pulumi.Input[str] password: **WARNING:** this will put the password in the terraform state file. Use carefully.
        :param pulumi.Input[str] rsa_public_key: Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
        :param pulumi.Input[str] rsa_public_key2: Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[UserArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_snowflake as snowflake

        user = snowflake.User("user",
            comment="A user of snowflake.",
            default_role="role1",
            default_warehouse="warehouse",
            disabled=False,
            display_name="Snowflake User",
            email="user@snowflake.example",
            first_name="Snowflake",
            last_name="User",
            login_name="snowflake_user",
            must_change_password=False,
            password="secret",
            rsa_public_key="...",
            rsa_public_key2="...")
        ```

        ## Import

        ```sh
         $ pulumi import snowflake:index/user:User example userName
        ```

        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_namespace: Optional[pulumi.Input[str]] = None,
                 default_role: Optional[pulumi.Input[str]] = None,
                 default_warehouse: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 login_name: Optional[pulumi.Input[str]] = None,
                 must_change_password: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 rsa_public_key: Optional[pulumi.Input[str]] = None,
                 rsa_public_key2: Optional[pulumi.Input[str]] = None,
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
            __props__ = UserArgs.__new__(UserArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["default_namespace"] = default_namespace
            __props__.__dict__["default_role"] = default_role
            __props__.__dict__["default_warehouse"] = default_warehouse
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["email"] = email
            __props__.__dict__["first_name"] = first_name
            __props__.__dict__["last_name"] = last_name
            __props__.__dict__["login_name"] = login_name
            __props__.__dict__["must_change_password"] = must_change_password
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = password
            __props__.__dict__["rsa_public_key"] = rsa_public_key
            __props__.__dict__["rsa_public_key2"] = rsa_public_key2
            __props__.__dict__["has_rsa_public_key"] = None
        super(User, __self__).__init__(
            'snowflake:index/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            default_namespace: Optional[pulumi.Input[str]] = None,
            default_role: Optional[pulumi.Input[str]] = None,
            default_warehouse: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            email: Optional[pulumi.Input[str]] = None,
            first_name: Optional[pulumi.Input[str]] = None,
            has_rsa_public_key: Optional[pulumi.Input[bool]] = None,
            last_name: Optional[pulumi.Input[str]] = None,
            login_name: Optional[pulumi.Input[str]] = None,
            must_change_password: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            rsa_public_key: Optional[pulumi.Input[str]] = None,
            rsa_public_key2: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_namespace: Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.
        :param pulumi.Input[str] default_role: Specifies the role that is active by default for the user’s session upon login.
        :param pulumi.Input[str] default_warehouse: Specifies the virtual warehouse that is active by default for the user’s session upon login.
        :param pulumi.Input[str] display_name: Name displayed for the user in the Snowflake web interface.
        :param pulumi.Input[str] email: Email address for the user.
        :param pulumi.Input[str] first_name: First name of the user.
        :param pulumi.Input[bool] has_rsa_public_key: Will be true if user as an RSA key set.
        :param pulumi.Input[str] last_name: Last name of the user.
        :param pulumi.Input[str] login_name: The name users use to log in. If not supplied, snowflake will use name instead.
        :param pulumi.Input[bool] must_change_password: Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.
        :param pulumi.Input[str] name: Name of the user. Note that if you do not supply login*name this will be used as login*name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)
        :param pulumi.Input[str] password: **WARNING:** this will put the password in the terraform state file. Use carefully.
        :param pulumi.Input[str] rsa_public_key: Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
        :param pulumi.Input[str] rsa_public_key2: Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["default_namespace"] = default_namespace
        __props__.__dict__["default_role"] = default_role
        __props__.__dict__["default_warehouse"] = default_warehouse
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["email"] = email
        __props__.__dict__["first_name"] = first_name
        __props__.__dict__["has_rsa_public_key"] = has_rsa_public_key
        __props__.__dict__["last_name"] = last_name
        __props__.__dict__["login_name"] = login_name
        __props__.__dict__["must_change_password"] = must_change_password
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["rsa_public_key"] = rsa_public_key
        __props__.__dict__["rsa_public_key2"] = rsa_public_key2
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="defaultNamespace")
    def default_namespace(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.
        """
        return pulumi.get(self, "default_namespace")

    @property
    @pulumi.getter(name="defaultRole")
    def default_role(self) -> pulumi.Output[str]:
        """
        Specifies the role that is active by default for the user’s session upon login.
        """
        return pulumi.get(self, "default_role")

    @property
    @pulumi.getter(name="defaultWarehouse")
    def default_warehouse(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the virtual warehouse that is active by default for the user’s session upon login.
        """
        return pulumi.get(self, "default_warehouse")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Name displayed for the user in the Snowflake web interface.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[Optional[str]]:
        """
        Email address for the user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> pulumi.Output[Optional[str]]:
        """
        First name of the user.
        """
        return pulumi.get(self, "first_name")

    @property
    @pulumi.getter(name="hasRsaPublicKey")
    def has_rsa_public_key(self) -> pulumi.Output[bool]:
        """
        Will be true if user as an RSA key set.
        """
        return pulumi.get(self, "has_rsa_public_key")

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> pulumi.Output[Optional[str]]:
        """
        Last name of the user.
        """
        return pulumi.get(self, "last_name")

    @property
    @pulumi.getter(name="loginName")
    def login_name(self) -> pulumi.Output[str]:
        """
        The name users use to log in. If not supplied, snowflake will use name instead.
        """
        return pulumi.get(self, "login_name")

    @property
    @pulumi.getter(name="mustChangePassword")
    def must_change_password(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.
        """
        return pulumi.get(self, "must_change_password")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the user. Note that if you do not supply login*name this will be used as login*name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        **WARNING:** this will put the password in the terraform state file. Use carefully.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="rsaPublicKey")
    def rsa_public_key(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
        """
        return pulumi.get(self, "rsa_public_key")

    @property
    @pulumi.getter(name="rsaPublicKey2")
    def rsa_public_key2(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
        """
        return pulumi.get(self, "rsa_public_key2")

