# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetFileFormatsResult',
    'AwaitableGetFileFormatsResult',
    'get_file_formats',
]

@pulumi.output_type
class GetFileFormatsResult:
    """
    A collection of values returned by getFileFormats.
    """
    def __init__(__self__, database=None, file_formats=None, id=None, schema=None):
        if database and not isinstance(database, str):
            raise TypeError("Expected argument 'database' to be a str")
        pulumi.set(__self__, "database", database)
        if file_formats and not isinstance(file_formats, list):
            raise TypeError("Expected argument 'file_formats' to be a list")
        pulumi.set(__self__, "file_formats", file_formats)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if schema and not isinstance(schema, str):
            raise TypeError("Expected argument 'schema' to be a str")
        pulumi.set(__self__, "schema", schema)

    @property
    @pulumi.getter
    def database(self) -> str:
        """
        The database from which to return the schemas from.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="fileFormats")
    def file_formats(self) -> Sequence['outputs.GetFileFormatsFileFormatResult']:
        """
        The file formats in the schema
        """
        return pulumi.get(self, "file_formats")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def schema(self) -> str:
        """
        The schema from which to return the file formats from.
        """
        return pulumi.get(self, "schema")


class AwaitableGetFileFormatsResult(GetFileFormatsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileFormatsResult(
            database=self.database,
            file_formats=self.file_formats,
            id=self.id,
            schema=self.schema)


def get_file_formats(database: Optional[str] = None,
                     schema: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFileFormatsResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_snowflake as snowflake

    current = snowflake.get_file_formats(database="MYDB",
        schema="MYSCHEMA")
    ```


    :param str database: The database from which to return the schemas from.
    :param str schema: The schema from which to return the file formats from.
    """
    __args__ = dict()
    __args__['database'] = database
    __args__['schema'] = schema
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('snowflake:index/getFileFormats:getFileFormats', __args__, opts=opts, typ=GetFileFormatsResult).value

    return AwaitableGetFileFormatsResult(
        database=__ret__.database,
        file_formats=__ret__.file_formats,
        id=__ret__.id,
        schema=__ret__.schema)
