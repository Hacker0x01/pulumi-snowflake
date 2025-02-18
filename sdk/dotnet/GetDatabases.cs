// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Snowflake
{
    public static class GetDatabases
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
        ///         var @this = Output.Create(Snowflake.GetDatabases.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatabasesResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabasesResult>("snowflake:index/getDatabases:getDatabases", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetDatabasesResult
    {
        /// <summary>
        /// Snowflake databases
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabasesDatabaseResult> Databases;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDatabasesResult(
            ImmutableArray<Outputs.GetDatabasesDatabaseResult> databases,

            string id)
        {
            Databases = databases;
            Id = id;
        }
    }
}
