// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Snowflake.Outputs
{

    [OutputType]
    public sealed class TableColumn
    {
        /// <summary>
        /// Column name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Whether this column can contain null values. **Note**: Depending on your Snowflake version, the default value will not suffice if this column is used in a primary key constraint.
        /// </summary>
        public readonly bool? Nullable;
        /// <summary>
        /// Column type, e.g. VARIANT
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private TableColumn(
            string name,

            bool? nullable,

            string type)
        {
            Name = name;
            Nullable = nullable;
            Type = type;
        }
    }
}
