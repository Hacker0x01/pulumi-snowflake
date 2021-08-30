// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Snowflake
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Snowflake = Pulumi.Snowflake;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var externalTable = new Snowflake.ExternalTable("externalTable", new Snowflake.ExternalTableArgs
    ///         {
    ///             Columns = 
    ///             {
    ///                 new Snowflake.Inputs.ExternalTableColumnArgs
    ///                 {
    ///                     Name = "id",
    ///                     Type = "int",
    ///                 },
    ///                 new Snowflake.Inputs.ExternalTableColumnArgs
    ///                 {
    ///                     Name = "data",
    ///                     Type = "text",
    ///                 },
    ///             },
    ///             Comment = "External table",
    ///             Database = "db",
    ///             Schema = "schema",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// # format is database name | schema name | external table name
    /// 
    /// ```sh
    ///  $ pulumi import snowflake:index/externalTable:ExternalTable example 'dbName|schemaName|externalTableName'
    /// ```
    /// </summary>
    [SnowflakeResourceType("snowflake:index/externalTable:ExternalTable")]
    public partial class ExternalTable : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether to automatically refresh the external table metadata once, immediately after the external table is created.
        /// </summary>
        [Output("autoRefresh")]
        public Output<bool?> AutoRefresh { get; private set; } = null!;

        /// <summary>
        /// Specifies the aws sns topic for the external table.
        /// </summary>
        [Output("awsSnsTopic")]
        public Output<string?> AwsSnsTopic { get; private set; } = null!;

        /// <summary>
        /// Definitions of a column to create in the external table. Minimum one required.
        /// </summary>
        [Output("columns")]
        public Output<ImmutableArray<Outputs.ExternalTableColumn>> Columns { get; private set; } = null!;

        /// <summary>
        /// Specifies a comment for the external table.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Specifies to retain the access permissions from the original table when an external table is recreated using the CREATE OR REPLACE TABLE variant
        /// </summary>
        [Output("copyGrants")]
        public Output<bool?> CopyGrants { get; private set; } = null!;

        /// <summary>
        /// The database in which to create the external table.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// Specifies the file format for the external table.
        /// </summary>
        [Output("fileFormat")]
        public Output<string> FileFormat { get; private set; } = null!;

        /// <summary>
        /// Specifies a location for the external table.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the identifier for the external table; must be unique for the database and schema in which the externalTable is created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name of the role that owns the external table.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// Specifies any partition columns to evaluate for the external table.
        /// </summary>
        [Output("partitionBies")]
        public Output<ImmutableArray<string>> PartitionBies { get; private set; } = null!;

        /// <summary>
        /// Specifies the file names and/or paths on the external stage to match.
        /// </summary>
        [Output("pattern")]
        public Output<string?> Pattern { get; private set; } = null!;

        /// <summary>
        /// Specifies weather to refresh when an external table is created.
        /// </summary>
        [Output("refreshOnCreate")]
        public Output<bool?> RefreshOnCreate { get; private set; } = null!;

        /// <summary>
        /// The schema in which to create the external table.
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;


        /// <summary>
        /// Create a ExternalTable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExternalTable(string name, ExternalTableArgs args, CustomResourceOptions? options = null)
            : base("snowflake:index/externalTable:ExternalTable", name, args ?? new ExternalTableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExternalTable(string name, Input<string> id, ExternalTableState? state = null, CustomResourceOptions? options = null)
            : base("snowflake:index/externalTable:ExternalTable", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ExternalTable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExternalTable Get(string name, Input<string> id, ExternalTableState? state = null, CustomResourceOptions? options = null)
        {
            return new ExternalTable(name, id, state, options);
        }
    }

    public sealed class ExternalTableArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to automatically refresh the external table metadata once, immediately after the external table is created.
        /// </summary>
        [Input("autoRefresh")]
        public Input<bool>? AutoRefresh { get; set; }

        /// <summary>
        /// Specifies the aws sns topic for the external table.
        /// </summary>
        [Input("awsSnsTopic")]
        public Input<string>? AwsSnsTopic { get; set; }

        [Input("columns", required: true)]
        private InputList<Inputs.ExternalTableColumnArgs>? _columns;

        /// <summary>
        /// Definitions of a column to create in the external table. Minimum one required.
        /// </summary>
        public InputList<Inputs.ExternalTableColumnArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.ExternalTableColumnArgs>());
            set => _columns = value;
        }

        /// <summary>
        /// Specifies a comment for the external table.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Specifies to retain the access permissions from the original table when an external table is recreated using the CREATE OR REPLACE TABLE variant
        /// </summary>
        [Input("copyGrants")]
        public Input<bool>? CopyGrants { get; set; }

        /// <summary>
        /// The database in which to create the external table.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// Specifies the file format for the external table.
        /// </summary>
        [Input("fileFormat", required: true)]
        public Input<string> FileFormat { get; set; } = null!;

        /// <summary>
        /// Specifies a location for the external table.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Specifies the identifier for the external table; must be unique for the database and schema in which the externalTable is created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("partitionBies")]
        private InputList<string>? _partitionBies;

        /// <summary>
        /// Specifies any partition columns to evaluate for the external table.
        /// </summary>
        public InputList<string> PartitionBies
        {
            get => _partitionBies ?? (_partitionBies = new InputList<string>());
            set => _partitionBies = value;
        }

        /// <summary>
        /// Specifies the file names and/or paths on the external stage to match.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// Specifies weather to refresh when an external table is created.
        /// </summary>
        [Input("refreshOnCreate")]
        public Input<bool>? RefreshOnCreate { get; set; }

        /// <summary>
        /// The schema in which to create the external table.
        /// </summary>
        [Input("schema", required: true)]
        public Input<string> Schema { get; set; } = null!;

        public ExternalTableArgs()
        {
        }
    }

    public sealed class ExternalTableState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to automatically refresh the external table metadata once, immediately after the external table is created.
        /// </summary>
        [Input("autoRefresh")]
        public Input<bool>? AutoRefresh { get; set; }

        /// <summary>
        /// Specifies the aws sns topic for the external table.
        /// </summary>
        [Input("awsSnsTopic")]
        public Input<string>? AwsSnsTopic { get; set; }

        [Input("columns")]
        private InputList<Inputs.ExternalTableColumnGetArgs>? _columns;

        /// <summary>
        /// Definitions of a column to create in the external table. Minimum one required.
        /// </summary>
        public InputList<Inputs.ExternalTableColumnGetArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.ExternalTableColumnGetArgs>());
            set => _columns = value;
        }

        /// <summary>
        /// Specifies a comment for the external table.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Specifies to retain the access permissions from the original table when an external table is recreated using the CREATE OR REPLACE TABLE variant
        /// </summary>
        [Input("copyGrants")]
        public Input<bool>? CopyGrants { get; set; }

        /// <summary>
        /// The database in which to create the external table.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Specifies the file format for the external table.
        /// </summary>
        [Input("fileFormat")]
        public Input<string>? FileFormat { get; set; }

        /// <summary>
        /// Specifies a location for the external table.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the identifier for the external table; must be unique for the database and schema in which the externalTable is created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of the role that owns the external table.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("partitionBies")]
        private InputList<string>? _partitionBies;

        /// <summary>
        /// Specifies any partition columns to evaluate for the external table.
        /// </summary>
        public InputList<string> PartitionBies
        {
            get => _partitionBies ?? (_partitionBies = new InputList<string>());
            set => _partitionBies = value;
        }

        /// <summary>
        /// Specifies the file names and/or paths on the external stage to match.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// Specifies weather to refresh when an external table is created.
        /// </summary>
        [Input("refreshOnCreate")]
        public Input<bool>? RefreshOnCreate { get; set; }

        /// <summary>
        /// The schema in which to create the external table.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        public ExternalTableState()
        {
        }
    }
}
