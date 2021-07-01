# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'ExternalFunctionArgArgs',
    'ExternalFunctionHeaderArgs',
    'ExternalTableColumnArgs',
    'FunctionGrantArgumentArgs',
    'ProcedureGrantArgumentArgs',
    'TableColumnArgs',
]

@pulumi.input_type
class ExternalFunctionArgArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name: Argument name
        :param pulumi.Input[str] type: Argument type, e.g. VARCHAR
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Argument name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Argument type, e.g. VARCHAR
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ExternalFunctionHeaderArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name: Header name
        :param pulumi.Input[str] value: Header value
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Header name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Header value
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ExternalTableColumnArgs:
    def __init__(__self__, *,
                 as_: pulumi.Input[str],
                 name: pulumi.Input[str],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] as_: String that specifies the expression for the column. When queried, the column returns results derived from this expression.
        :param pulumi.Input[str] name: Column name
        :param pulumi.Input[str] type: Column type, e.g. VARIANT
        """
        pulumi.set(__self__, "as_", as_)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="as")
    def as_(self) -> pulumi.Input[str]:
        """
        String that specifies the expression for the column. When queried, the column returns results derived from this expression.
        """
        return pulumi.get(self, "as_")

    @as_.setter
    def as_(self, value: pulumi.Input[str]):
        pulumi.set(self, "as_", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Column name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Column type, e.g. VARIANT
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class FunctionGrantArgumentArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name: The argument name
        :param pulumi.Input[str] type: The argument type
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The argument name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The argument type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ProcedureGrantArgumentArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name: The argument name
        :param pulumi.Input[str] type: The argument type
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The argument name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The argument type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class TableColumnArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name: Column name
        :param pulumi.Input[str] type: Column type, e.g. VARIANT
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Column name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Column type, e.g. VARIANT
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


