# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TaskGrantArgs', 'TaskGrant']

@pulumi.input_type
class TaskGrantArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[str],
                 schema_name: pulumi.Input[str],
                 on_future: Optional[pulumi.Input[bool]] = None,
                 privilege: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 task_name: Optional[pulumi.Input[str]] = None,
                 with_grant_option: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a TaskGrant resource.
        :param pulumi.Input[str] database_name: The name of the database containing the current or future tasks on which to grant privileges.
        :param pulumi.Input[str] schema_name: The name of the schema containing the current or future tasks on which to grant privileges.
        :param pulumi.Input[bool] on_future: When this is set to true and a schema*name is provided, apply this grant on all future tasks in the given schema. When this is true and no schema*name is provided apply this grant on all future tasks in the given database. The task*name field must be unset in order to use on*future.
        :param pulumi.Input[str] privilege: The privilege to grant on the current or future task.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Grants privilege to these roles.
        :param pulumi.Input[str] task_name: The name of the task on which to grant privileges immediately (only valid if on_future is false).
        :param pulumi.Input[bool] with_grant_option: When this is set to true, allows the recipient role to grant the privileges to other roles.
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "schema_name", schema_name)
        if on_future is not None:
            pulumi.set(__self__, "on_future", on_future)
        if privilege is not None:
            pulumi.set(__self__, "privilege", privilege)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if task_name is not None:
            pulumi.set(__self__, "task_name", task_name)
        if with_grant_option is not None:
            pulumi.set(__self__, "with_grant_option", with_grant_option)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        The name of the database containing the current or future tasks on which to grant privileges.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> pulumi.Input[str]:
        """
        The name of the schema containing the current or future tasks on which to grant privileges.
        """
        return pulumi.get(self, "schema_name")

    @schema_name.setter
    def schema_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "schema_name", value)

    @property
    @pulumi.getter(name="onFuture")
    def on_future(self) -> Optional[pulumi.Input[bool]]:
        """
        When this is set to true and a schema*name is provided, apply this grant on all future tasks in the given schema. When this is true and no schema*name is provided apply this grant on all future tasks in the given database. The task*name field must be unset in order to use on*future.
        """
        return pulumi.get(self, "on_future")

    @on_future.setter
    def on_future(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "on_future", value)

    @property
    @pulumi.getter
    def privilege(self) -> Optional[pulumi.Input[str]]:
        """
        The privilege to grant on the current or future task.
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
    @pulumi.getter(name="taskName")
    def task_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the task on which to grant privileges immediately (only valid if on_future is false).
        """
        return pulumi.get(self, "task_name")

    @task_name.setter
    def task_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "task_name", value)

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
class _TaskGrantState:
    def __init__(__self__, *,
                 database_name: Optional[pulumi.Input[str]] = None,
                 on_future: Optional[pulumi.Input[bool]] = None,
                 privilege: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 schema_name: Optional[pulumi.Input[str]] = None,
                 task_name: Optional[pulumi.Input[str]] = None,
                 with_grant_option: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering TaskGrant resources.
        :param pulumi.Input[str] database_name: The name of the database containing the current or future tasks on which to grant privileges.
        :param pulumi.Input[bool] on_future: When this is set to true and a schema*name is provided, apply this grant on all future tasks in the given schema. When this is true and no schema*name is provided apply this grant on all future tasks in the given database. The task*name field must be unset in order to use on*future.
        :param pulumi.Input[str] privilege: The privilege to grant on the current or future task.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Grants privilege to these roles.
        :param pulumi.Input[str] schema_name: The name of the schema containing the current or future tasks on which to grant privileges.
        :param pulumi.Input[str] task_name: The name of the task on which to grant privileges immediately (only valid if on_future is false).
        :param pulumi.Input[bool] with_grant_option: When this is set to true, allows the recipient role to grant the privileges to other roles.
        """
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if on_future is not None:
            pulumi.set(__self__, "on_future", on_future)
        if privilege is not None:
            pulumi.set(__self__, "privilege", privilege)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if schema_name is not None:
            pulumi.set(__self__, "schema_name", schema_name)
        if task_name is not None:
            pulumi.set(__self__, "task_name", task_name)
        if with_grant_option is not None:
            pulumi.set(__self__, "with_grant_option", with_grant_option)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the database containing the current or future tasks on which to grant privileges.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="onFuture")
    def on_future(self) -> Optional[pulumi.Input[bool]]:
        """
        When this is set to true and a schema*name is provided, apply this grant on all future tasks in the given schema. When this is true and no schema*name is provided apply this grant on all future tasks in the given database. The task*name field must be unset in order to use on*future.
        """
        return pulumi.get(self, "on_future")

    @on_future.setter
    def on_future(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "on_future", value)

    @property
    @pulumi.getter
    def privilege(self) -> Optional[pulumi.Input[str]]:
        """
        The privilege to grant on the current or future task.
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
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the schema containing the current or future tasks on which to grant privileges.
        """
        return pulumi.get(self, "schema_name")

    @schema_name.setter
    def schema_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema_name", value)

    @property
    @pulumi.getter(name="taskName")
    def task_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the task on which to grant privileges immediately (only valid if on_future is false).
        """
        return pulumi.get(self, "task_name")

    @task_name.setter
    def task_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "task_name", value)

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


class TaskGrant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 on_future: Optional[pulumi.Input[bool]] = None,
                 privilege: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 schema_name: Optional[pulumi.Input[str]] = None,
                 task_name: Optional[pulumi.Input[str]] = None,
                 with_grant_option: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## Import

        # format is database name | schema name | task name | privilege | true/false for with_grant_option

        ```sh
         $ pulumi import snowflake:index/taskGrant:TaskGrant example 'dbName|schemaName|taskName|OPERATE|false'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: The name of the database containing the current or future tasks on which to grant privileges.
        :param pulumi.Input[bool] on_future: When this is set to true and a schema*name is provided, apply this grant on all future tasks in the given schema. When this is true and no schema*name is provided apply this grant on all future tasks in the given database. The task*name field must be unset in order to use on*future.
        :param pulumi.Input[str] privilege: The privilege to grant on the current or future task.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Grants privilege to these roles.
        :param pulumi.Input[str] schema_name: The name of the schema containing the current or future tasks on which to grant privileges.
        :param pulumi.Input[str] task_name: The name of the task on which to grant privileges immediately (only valid if on_future is false).
        :param pulumi.Input[bool] with_grant_option: When this is set to true, allows the recipient role to grant the privileges to other roles.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TaskGrantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        # format is database name | schema name | task name | privilege | true/false for with_grant_option

        ```sh
         $ pulumi import snowflake:index/taskGrant:TaskGrant example 'dbName|schemaName|taskName|OPERATE|false'
        ```

        :param str resource_name: The name of the resource.
        :param TaskGrantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TaskGrantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 on_future: Optional[pulumi.Input[bool]] = None,
                 privilege: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 schema_name: Optional[pulumi.Input[str]] = None,
                 task_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = TaskGrantArgs.__new__(TaskGrantArgs)

            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            __props__.__dict__["on_future"] = on_future
            __props__.__dict__["privilege"] = privilege
            __props__.__dict__["roles"] = roles
            if schema_name is None and not opts.urn:
                raise TypeError("Missing required property 'schema_name'")
            __props__.__dict__["schema_name"] = schema_name
            __props__.__dict__["task_name"] = task_name
            __props__.__dict__["with_grant_option"] = with_grant_option
        super(TaskGrant, __self__).__init__(
            'snowflake:index/taskGrant:TaskGrant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            on_future: Optional[pulumi.Input[bool]] = None,
            privilege: Optional[pulumi.Input[str]] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            schema_name: Optional[pulumi.Input[str]] = None,
            task_name: Optional[pulumi.Input[str]] = None,
            with_grant_option: Optional[pulumi.Input[bool]] = None) -> 'TaskGrant':
        """
        Get an existing TaskGrant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: The name of the database containing the current or future tasks on which to grant privileges.
        :param pulumi.Input[bool] on_future: When this is set to true and a schema*name is provided, apply this grant on all future tasks in the given schema. When this is true and no schema*name is provided apply this grant on all future tasks in the given database. The task*name field must be unset in order to use on*future.
        :param pulumi.Input[str] privilege: The privilege to grant on the current or future task.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: Grants privilege to these roles.
        :param pulumi.Input[str] schema_name: The name of the schema containing the current or future tasks on which to grant privileges.
        :param pulumi.Input[str] task_name: The name of the task on which to grant privileges immediately (only valid if on_future is false).
        :param pulumi.Input[bool] with_grant_option: When this is set to true, allows the recipient role to grant the privileges to other roles.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TaskGrantState.__new__(_TaskGrantState)

        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["on_future"] = on_future
        __props__.__dict__["privilege"] = privilege
        __props__.__dict__["roles"] = roles
        __props__.__dict__["schema_name"] = schema_name
        __props__.__dict__["task_name"] = task_name
        __props__.__dict__["with_grant_option"] = with_grant_option
        return TaskGrant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        """
        The name of the database containing the current or future tasks on which to grant privileges.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="onFuture")
    def on_future(self) -> pulumi.Output[Optional[bool]]:
        """
        When this is set to true and a schema*name is provided, apply this grant on all future tasks in the given schema. When this is true and no schema*name is provided apply this grant on all future tasks in the given database. The task*name field must be unset in order to use on*future.
        """
        return pulumi.get(self, "on_future")

    @property
    @pulumi.getter
    def privilege(self) -> pulumi.Output[Optional[str]]:
        """
        The privilege to grant on the current or future task.
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
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> pulumi.Output[str]:
        """
        The name of the schema containing the current or future tasks on which to grant privileges.
        """
        return pulumi.get(self, "schema_name")

    @property
    @pulumi.getter(name="taskName")
    def task_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the task on which to grant privileges immediately (only valid if on_future is false).
        """
        return pulumi.get(self, "task_name")

    @property
    @pulumi.getter(name="withGrantOption")
    def with_grant_option(self) -> pulumi.Output[Optional[bool]]:
        """
        When this is set to true, allows the recipient role to grant the privileges to other roles.
        """
        return pulumi.get(self, "with_grant_option")

