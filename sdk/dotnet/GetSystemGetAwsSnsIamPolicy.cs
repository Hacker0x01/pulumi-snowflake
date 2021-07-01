// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Snowflake
{
    public static class GetSystemGetAwsSnsIamPolicy
    {
        public static Task<GetSystemGetAwsSnsIamPolicyResult> InvokeAsync(GetSystemGetAwsSnsIamPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSystemGetAwsSnsIamPolicyResult>("snowflake:index/getSystemGetAwsSnsIamPolicy:getSystemGetAwsSnsIamPolicy", args ?? new GetSystemGetAwsSnsIamPolicyArgs(), options.WithVersion());
    }


    public sealed class GetSystemGetAwsSnsIamPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the SNS topic for your S3 bucket
        /// </summary>
        [Input("awsSnsTopicArn", required: true)]
        public string AwsSnsTopicArn { get; set; } = null!;

        public GetSystemGetAwsSnsIamPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSystemGetAwsSnsIamPolicyResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the SNS topic for your S3 bucket
        /// </summary>
        public readonly string AwsSnsTopicArn;
        /// <summary>
        /// IAM policy for Snowflake’s SQS queue to subscribe to this topic
        /// </summary>
        public readonly string AwsSnsTopicPolicyJson;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSystemGetAwsSnsIamPolicyResult(
            string awsSnsTopicArn,

            string awsSnsTopicPolicyJson,

            string id)
        {
            AwsSnsTopicArn = awsSnsTopicArn;
            AwsSnsTopicPolicyJson = awsSnsTopicPolicyJson;
            Id = id;
        }
    }
}
