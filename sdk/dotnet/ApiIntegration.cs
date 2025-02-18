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
    ///         var apiIntegration = new Snowflake.ApiIntegration("apiIntegration", new Snowflake.ApiIntegrationArgs
    ///         {
    ///             ApiAllowedPrefixes = 
    ///             {
    ///                 "https://123456.execute-api.us-west-2.amazonaws.com/prod/",
    ///             },
    ///             ApiAwsRoleArn = "arn:aws:iam::000000000001:/role/test",
    ///             ApiProvider = "aws_api_gateway",
    ///             Enabled = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import snowflake:index/apiIntegration:ApiIntegration example name
    /// ```
    /// </summary>
    [SnowflakeResourceType("snowflake:index/apiIntegration:ApiIntegration")]
    public partial class ApiIntegration : Pulumi.CustomResource
    {
        /// <summary>
        /// Explicitly limits external functions that use the integration to reference one or more HTTPS proxy service endpoints and resources within those proxies.
        /// </summary>
        [Output("apiAllowedPrefixes")]
        public Output<ImmutableArray<string>> ApiAllowedPrefixes { get; private set; } = null!;

        /// <summary>
        /// The external ID that Snowflake will use when assuming the AWS role.
        /// </summary>
        [Output("apiAwsExternalId")]
        public Output<string> ApiAwsExternalId { get; private set; } = null!;

        /// <summary>
        /// The Snowflake user that will attempt to assume the AWS role.
        /// </summary>
        [Output("apiAwsIamUserArn")]
        public Output<string> ApiAwsIamUserArn { get; private set; } = null!;

        /// <summary>
        /// ARN of a cloud platform role.
        /// </summary>
        [Output("apiAwsRoleArn")]
        public Output<string?> ApiAwsRoleArn { get; private set; } = null!;

        /// <summary>
        /// Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
        /// </summary>
        [Output("apiBlockedPrefixes")]
        public Output<ImmutableArray<string>> ApiBlockedPrefixes { get; private set; } = null!;

        /// <summary>
        /// Specifies the HTTPS proxy service type.
        /// </summary>
        [Output("apiProvider")]
        public Output<string> ApiProvider { get; private set; } = null!;

        /// <summary>
        /// The 'Application (client) id' of the Azure AD app for your remote service.
        /// </summary>
        [Output("azureAdApplicationId")]
        public Output<string?> AzureAdApplicationId { get; private set; } = null!;

        [Output("azureConsentUrl")]
        public Output<string> AzureConsentUrl { get; private set; } = null!;

        [Output("azureMultiTenantAppName")]
        public Output<string> AzureMultiTenantAppName { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID for your Office 365 tenant that all Azure API Management instances belong to.
        /// </summary>
        [Output("azureTenantId")]
        public Output<string?> AzureTenantId { get; private set; } = null!;

        /// <summary>
        /// Date and time when the API integration was created.
        /// </summary>
        [Output("createdOn")]
        public Output<string> CreatedOn { get; private set; } = null!;

        /// <summary>
        /// Specifies whether this API integration is enabled or disabled. If the API integration is disabled, any external function that relies on it will not work.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the API integration. This name follows the rules for Object Identifiers. The name should be unique among api integrations in your account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ApiIntegration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiIntegration(string name, ApiIntegrationArgs args, CustomResourceOptions? options = null)
            : base("snowflake:index/apiIntegration:ApiIntegration", name, args ?? new ApiIntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiIntegration(string name, Input<string> id, ApiIntegrationState? state = null, CustomResourceOptions? options = null)
            : base("snowflake:index/apiIntegration:ApiIntegration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApiIntegration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiIntegration Get(string name, Input<string> id, ApiIntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new ApiIntegration(name, id, state, options);
        }
    }

    public sealed class ApiIntegrationArgs : Pulumi.ResourceArgs
    {
        [Input("apiAllowedPrefixes", required: true)]
        private InputList<string>? _apiAllowedPrefixes;

        /// <summary>
        /// Explicitly limits external functions that use the integration to reference one or more HTTPS proxy service endpoints and resources within those proxies.
        /// </summary>
        public InputList<string> ApiAllowedPrefixes
        {
            get => _apiAllowedPrefixes ?? (_apiAllowedPrefixes = new InputList<string>());
            set => _apiAllowedPrefixes = value;
        }

        /// <summary>
        /// ARN of a cloud platform role.
        /// </summary>
        [Input("apiAwsRoleArn")]
        public Input<string>? ApiAwsRoleArn { get; set; }

        [Input("apiBlockedPrefixes")]
        private InputList<string>? _apiBlockedPrefixes;

        /// <summary>
        /// Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
        /// </summary>
        public InputList<string> ApiBlockedPrefixes
        {
            get => _apiBlockedPrefixes ?? (_apiBlockedPrefixes = new InputList<string>());
            set => _apiBlockedPrefixes = value;
        }

        /// <summary>
        /// Specifies the HTTPS proxy service type.
        /// </summary>
        [Input("apiProvider", required: true)]
        public Input<string> ApiProvider { get; set; } = null!;

        /// <summary>
        /// The 'Application (client) id' of the Azure AD app for your remote service.
        /// </summary>
        [Input("azureAdApplicationId")]
        public Input<string>? AzureAdApplicationId { get; set; }

        /// <summary>
        /// Specifies the ID for your Office 365 tenant that all Azure API Management instances belong to.
        /// </summary>
        [Input("azureTenantId")]
        public Input<string>? AzureTenantId { get; set; }

        /// <summary>
        /// Specifies whether this API integration is enabled or disabled. If the API integration is disabled, any external function that relies on it will not work.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Specifies the name of the API integration. This name follows the rules for Object Identifiers. The name should be unique among api integrations in your account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ApiIntegrationArgs()
        {
        }
    }

    public sealed class ApiIntegrationState : Pulumi.ResourceArgs
    {
        [Input("apiAllowedPrefixes")]
        private InputList<string>? _apiAllowedPrefixes;

        /// <summary>
        /// Explicitly limits external functions that use the integration to reference one or more HTTPS proxy service endpoints and resources within those proxies.
        /// </summary>
        public InputList<string> ApiAllowedPrefixes
        {
            get => _apiAllowedPrefixes ?? (_apiAllowedPrefixes = new InputList<string>());
            set => _apiAllowedPrefixes = value;
        }

        /// <summary>
        /// The external ID that Snowflake will use when assuming the AWS role.
        /// </summary>
        [Input("apiAwsExternalId")]
        public Input<string>? ApiAwsExternalId { get; set; }

        /// <summary>
        /// The Snowflake user that will attempt to assume the AWS role.
        /// </summary>
        [Input("apiAwsIamUserArn")]
        public Input<string>? ApiAwsIamUserArn { get; set; }

        /// <summary>
        /// ARN of a cloud platform role.
        /// </summary>
        [Input("apiAwsRoleArn")]
        public Input<string>? ApiAwsRoleArn { get; set; }

        [Input("apiBlockedPrefixes")]
        private InputList<string>? _apiBlockedPrefixes;

        /// <summary>
        /// Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
        /// </summary>
        public InputList<string> ApiBlockedPrefixes
        {
            get => _apiBlockedPrefixes ?? (_apiBlockedPrefixes = new InputList<string>());
            set => _apiBlockedPrefixes = value;
        }

        /// <summary>
        /// Specifies the HTTPS proxy service type.
        /// </summary>
        [Input("apiProvider")]
        public Input<string>? ApiProvider { get; set; }

        /// <summary>
        /// The 'Application (client) id' of the Azure AD app for your remote service.
        /// </summary>
        [Input("azureAdApplicationId")]
        public Input<string>? AzureAdApplicationId { get; set; }

        [Input("azureConsentUrl")]
        public Input<string>? AzureConsentUrl { get; set; }

        [Input("azureMultiTenantAppName")]
        public Input<string>? AzureMultiTenantAppName { get; set; }

        /// <summary>
        /// Specifies the ID for your Office 365 tenant that all Azure API Management instances belong to.
        /// </summary>
        [Input("azureTenantId")]
        public Input<string>? AzureTenantId { get; set; }

        /// <summary>
        /// Date and time when the API integration was created.
        /// </summary>
        [Input("createdOn")]
        public Input<string>? CreatedOn { get; set; }

        /// <summary>
        /// Specifies whether this API integration is enabled or disabled. If the API integration is disabled, any external function that relies on it will not work.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Specifies the name of the API integration. This name follows the rules for Object Identifiers. The name should be unique among api integrations in your account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ApiIntegrationState()
        {
        }
    }
}
