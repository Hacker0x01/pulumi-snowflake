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
    ///         var grant = new Snowflake.RowAccessPolicyGrant("grant", new Snowflake.RowAccessPolicyGrantArgs
    ///         {
    ///             DatabaseName = "db",
    ///             Privilege = "APPLY",
    ///             Roles = 
    ///             {
    ///                 "role1",
    ///                 "role2",
    ///             },
    ///             RowAccessPolicyName = "row_access_policy",
    ///             SchemaName = "schema",
    ///             WithGrantOption = false,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// # format is database name | schema name | row access policy name | privilege | true/false for with_grant_option
    /// 
    /// ```sh
    ///  $ pulumi import snowflake:index/rowAccessPolicyGrant:RowAccessPolicyGrant example 'dbName|schemaName|rowAccessPolicyName|SELECT|false'
    /// ```
    /// </summary>
    [SnowflakeResourceType("snowflake:index/rowAccessPolicyGrant:RowAccessPolicyGrant")]
    public partial class RowAccessPolicyGrant : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the database containing the row access policy on which to grant privileges.
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// The privilege to grant on the row access policy.
        /// </summary>
        [Output("privilege")]
        public Output<string?> Privilege { get; private set; } = null!;

        /// <summary>
        /// Grants privilege to these roles.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// The name of the row access policy on which to grant privileges immediately.
        /// </summary>
        [Output("rowAccessPolicyName")]
        public Output<string> RowAccessPolicyName { get; private set; } = null!;

        /// <summary>
        /// The name of the schema containing the row access policy on which to grant privileges.
        /// </summary>
        [Output("schemaName")]
        public Output<string> SchemaName { get; private set; } = null!;

        /// <summary>
        /// When this is set to true, allows the recipient role to grant the privileges to other roles.
        /// </summary>
        [Output("withGrantOption")]
        public Output<bool?> WithGrantOption { get; private set; } = null!;


        /// <summary>
        /// Create a RowAccessPolicyGrant resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RowAccessPolicyGrant(string name, RowAccessPolicyGrantArgs args, CustomResourceOptions? options = null)
            : base("snowflake:index/rowAccessPolicyGrant:RowAccessPolicyGrant", name, args ?? new RowAccessPolicyGrantArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RowAccessPolicyGrant(string name, Input<string> id, RowAccessPolicyGrantState? state = null, CustomResourceOptions? options = null)
            : base("snowflake:index/rowAccessPolicyGrant:RowAccessPolicyGrant", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RowAccessPolicyGrant resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RowAccessPolicyGrant Get(string name, Input<string> id, RowAccessPolicyGrantState? state = null, CustomResourceOptions? options = null)
        {
            return new RowAccessPolicyGrant(name, id, state, options);
        }
    }

    public sealed class RowAccessPolicyGrantArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the database containing the row access policy on which to grant privileges.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The privilege to grant on the row access policy.
        /// </summary>
        [Input("privilege")]
        public Input<string>? Privilege { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// Grants privilege to these roles.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// The name of the row access policy on which to grant privileges immediately.
        /// </summary>
        [Input("rowAccessPolicyName", required: true)]
        public Input<string> RowAccessPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the schema containing the row access policy on which to grant privileges.
        /// </summary>
        [Input("schemaName", required: true)]
        public Input<string> SchemaName { get; set; } = null!;

        /// <summary>
        /// When this is set to true, allows the recipient role to grant the privileges to other roles.
        /// </summary>
        [Input("withGrantOption")]
        public Input<bool>? WithGrantOption { get; set; }

        public RowAccessPolicyGrantArgs()
        {
        }
    }

    public sealed class RowAccessPolicyGrantState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the database containing the row access policy on which to grant privileges.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// The privilege to grant on the row access policy.
        /// </summary>
        [Input("privilege")]
        public Input<string>? Privilege { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// Grants privilege to these roles.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// The name of the row access policy on which to grant privileges immediately.
        /// </summary>
        [Input("rowAccessPolicyName")]
        public Input<string>? RowAccessPolicyName { get; set; }

        /// <summary>
        /// The name of the schema containing the row access policy on which to grant privileges.
        /// </summary>
        [Input("schemaName")]
        public Input<string>? SchemaName { get; set; }

        /// <summary>
        /// When this is set to true, allows the recipient role to grant the privileges to other roles.
        /// </summary>
        [Input("withGrantOption")]
        public Input<bool>? WithGrantOption { get; set; }

        public RowAccessPolicyGrantState()
        {
        }
    }
}
