// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Snowflake
{
    public static class GetRowAccessPolicies
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Snowflake = Pulumi.Snowflake;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var current = Output.Create(Snowflake.GetRowAccessPolicies.InvokeAsync(new Snowflake.GetRowAccessPoliciesArgs
        ///         {
        ///             Database = "MYDB",
        ///             Schema = "MYSCHEMA",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRowAccessPoliciesResult> InvokeAsync(GetRowAccessPoliciesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRowAccessPoliciesResult>("snowflake:index/getRowAccessPolicies:getRowAccessPolicies", args ?? new GetRowAccessPoliciesArgs(), options.WithVersion());
    }


    public sealed class GetRowAccessPoliciesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database from which to return the schemas from.
        /// </summary>
        [Input("database", required: true)]
        public string Database { get; set; } = null!;

        /// <summary>
        /// The schema from which to return the row access policyfrom.
        /// </summary>
        [Input("schema", required: true)]
        public string Schema { get; set; } = null!;

        public GetRowAccessPoliciesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRowAccessPoliciesResult
    {
        /// <summary>
        /// The database from which to return the schemas from.
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The row access policy in the schema
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRowAccessPoliciesRowAccessPolicyResult> RowAccessPolicies;
        /// <summary>
        /// The schema from which to return the row access policyfrom.
        /// </summary>
        public readonly string Schema;

        [OutputConstructor]
        private GetRowAccessPoliciesResult(
            string database,

            string id,

            ImmutableArray<Outputs.GetRowAccessPoliciesRowAccessPolicyResult> rowAccessPolicies,

            string schema)
        {
            Database = database;
            Id = id;
            RowAccessPolicies = rowAccessPolicies;
            Schema = schema;
        }
    }
}
