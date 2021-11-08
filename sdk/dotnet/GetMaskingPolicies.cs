// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Snowflake
{
    public static class GetMaskingPolicies
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
        ///         var current = Output.Create(Snowflake.GetMaskingPolicies.InvokeAsync(new Snowflake.GetMaskingPoliciesArgs
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
        public static Task<GetMaskingPoliciesResult> InvokeAsync(GetMaskingPoliciesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMaskingPoliciesResult>("snowflake:index/getMaskingPolicies:getMaskingPolicies", args ?? new GetMaskingPoliciesArgs(), options.WithVersion());

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
        ///         var current = Output.Create(Snowflake.GetMaskingPolicies.InvokeAsync(new Snowflake.GetMaskingPoliciesArgs
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
        public static Output<GetMaskingPoliciesResult> Invoke(GetMaskingPoliciesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMaskingPoliciesResult>("snowflake:index/getMaskingPolicies:getMaskingPolicies", args ?? new GetMaskingPoliciesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetMaskingPoliciesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database from which to return the schemas from.
        /// </summary>
        [Input("database", required: true)]
        public string Database { get; set; } = null!;

        /// <summary>
        /// The schema from which to return the maskingPolicies from.
        /// </summary>
        [Input("schema", required: true)]
        public string Schema { get; set; } = null!;

        public GetMaskingPoliciesArgs()
        {
        }
    }

    public sealed class GetMaskingPoliciesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database from which to return the schemas from.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// The schema from which to return the maskingPolicies from.
        /// </summary>
        [Input("schema", required: true)]
        public Input<string> Schema { get; set; } = null!;

        public GetMaskingPoliciesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMaskingPoliciesResult
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
        /// The maskingPolicies in the schema
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMaskingPoliciesMaskingPolicyResult> MaskingPolicies;
        /// <summary>
        /// The schema from which to return the maskingPolicies from.
        /// </summary>
        public readonly string Schema;

        [OutputConstructor]
        private GetMaskingPoliciesResult(
            string database,

            string id,

            ImmutableArray<Outputs.GetMaskingPoliciesMaskingPolicyResult> maskingPolicies,

            string schema)
        {
            Database = database;
            Id = id;
            MaskingPolicies = maskingPolicies;
            Schema = schema;
        }
    }
}
