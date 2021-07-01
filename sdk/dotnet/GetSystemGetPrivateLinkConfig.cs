// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Snowflake
{
    public static class GetSystemGetPrivateLinkConfig
    {
        public static Task<GetSystemGetPrivateLinkConfigResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSystemGetPrivateLinkConfigResult>("snowflake:index/getSystemGetPrivateLinkConfig:getSystemGetPrivateLinkConfig", InvokeArgs.Empty, options.WithVersion());
    }


    [OutputType]
    public sealed class GetSystemGetPrivateLinkConfigResult
    {
        /// <summary>
        /// The name of your Snowflake account.
        /// </summary>
        public readonly string AccountName;
        /// <summary>
        /// The URL used to connect to Snowflake through AWS PrivateLink or Azure Private Link.
        /// </summary>
        public readonly string AccountUrl;
        /// <summary>
        /// The AWS VPCE ID for your account.
        /// </summary>
        public readonly string AwsVpceId;
        /// <summary>
        /// The Azure Private Link Service ID for your account.
        /// </summary>
        public readonly string AzurePlsId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The OCSP URL corresponding to your Snowflake account that uses AWS PrivateLink or Azure Private Link.
        /// </summary>
        public readonly string OscpUrl;

        [OutputConstructor]
        private GetSystemGetPrivateLinkConfigResult(
            string accountName,

            string accountUrl,

            string awsVpceId,

            string azurePlsId,

            string id,

            string oscpUrl)
        {
            AccountName = accountName;
            AccountUrl = accountUrl;
            AwsVpceId = awsVpceId;
            AzurePlsId = azurePlsId;
            Id = id;
            OscpUrl = oscpUrl;
        }
    }
}
