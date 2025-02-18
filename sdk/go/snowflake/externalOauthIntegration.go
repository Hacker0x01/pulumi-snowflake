// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package snowflake

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-snowflake/sdk/go/snowflake"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := snowflake.NewExternalOauthIntegration(ctx, "azure", &snowflake.ExternalOauthIntegrationArgs{
// 			AudienceUrls: pulumi.StringArray{
// 				pulumi.String("https://analysis.windows.net/powerbi/connector/Snowflake"),
// 			},
// 			Enabled: pulumi.Bool(true),
// 			Issuer:  pulumi.String("https://sts.windows.net/00000000-0000-0000-0000-000000000000"),
// 			JwsKeysUrls: pulumi.StringArray{
// 				pulumi.String("https://login.windows.net/common/discovery/keys"),
// 			},
// 			SnowflakeUserMappingAttribute: pulumi.String("LOGIN_NAME"),
// 			TokenUserMappingClaims: pulumi.StringArray{
// 				pulumi.String("upn"),
// 			},
// 			Type: pulumi.String("AZURE"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ```sh
//  $ pulumi import snowflake:index/externalOauthIntegration:ExternalOauthIntegration example name
// ```
type ExternalOauthIntegration struct {
	pulumi.CustomResourceState

	// Specifies the list of roles that the client can set as the primary role.
	AllowedRoles pulumi.StringArrayOutput `pulumi:"allowedRoles"`
	// Specifies whether the OAuth client or user can use a role that is not defined in the OAuth access token.
	AnyRoleMode pulumi.StringPtrOutput `pulumi:"anyRoleMode"`
	// Specifies additional values that can be used for the access token's audience validation on top of using the Customer's Snowflake Account URL
	AudienceUrls pulumi.StringArrayOutput `pulumi:"audienceUrls"`
	// Specifies the list of roles that a client cannot set as the primary role. Do not include ACCOUNTADMIN, ORGADMIN or SECURITYADMIN as they are already implicitly enforced and will cause in-place updates.
	BlockedRoles pulumi.StringArrayOutput `pulumi:"blockedRoles"`
	// Specifies a comment for the OAuth integration.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Date and time when the External OAUTH integration was created.
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// Specifies whether to initiate operation of the integration or suspend it.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Specifies the URL to define the OAuth 2.0 authorization server.
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// Specifies the endpoint or a list of endpoints from which to download public keys or certificates to validate an External OAuth access token. The maximum number of URLs that can be specified in the list is 3.
	JwsKeysUrls pulumi.StringArrayOutput `pulumi:"jwsKeysUrls"`
	// Specifies the name of the External Oath integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies a Base64-encoded RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers.
	RsaPublicKey pulumi.StringPtrOutput `pulumi:"rsaPublicKey"`
	// Specifies a second RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers. Used for key rotation.
	RsaPublicKey2 pulumi.StringPtrOutput `pulumi:"rsaPublicKey2"`
	// Specifies the scope delimiter in the authorization token.
	ScopeDelimiter pulumi.StringPtrOutput `pulumi:"scopeDelimiter"`
	// Indicates which Snowflake user record attribute should be used to map the access token to a Snowflake user record.
	SnowflakeUserMappingAttribute pulumi.StringOutput `pulumi:"snowflakeUserMappingAttribute"`
	// Specifies the access token claim or claims that can be used to map the access token to a Snowflake user record.
	TokenUserMappingClaims pulumi.StringArrayOutput `pulumi:"tokenUserMappingClaims"`
	// Specifies the OAuth 2.0 authorization server to be Okta, Microsoft Azure AD, Ping Identity PingFederate, or a Custom OAuth 2.0 authorization server.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExternalOauthIntegration registers a new resource with the given unique name, arguments, and options.
func NewExternalOauthIntegration(ctx *pulumi.Context,
	name string, args *ExternalOauthIntegrationArgs, opts ...pulumi.ResourceOption) (*ExternalOauthIntegration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Issuer == nil {
		return nil, errors.New("invalid value for required argument 'Issuer'")
	}
	if args.SnowflakeUserMappingAttribute == nil {
		return nil, errors.New("invalid value for required argument 'SnowflakeUserMappingAttribute'")
	}
	if args.TokenUserMappingClaims == nil {
		return nil, errors.New("invalid value for required argument 'TokenUserMappingClaims'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource ExternalOauthIntegration
	err := ctx.RegisterResource("snowflake:index/externalOauthIntegration:ExternalOauthIntegration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalOauthIntegration gets an existing ExternalOauthIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalOauthIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalOauthIntegrationState, opts ...pulumi.ResourceOption) (*ExternalOauthIntegration, error) {
	var resource ExternalOauthIntegration
	err := ctx.ReadResource("snowflake:index/externalOauthIntegration:ExternalOauthIntegration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalOauthIntegration resources.
type externalOauthIntegrationState struct {
	// Specifies the list of roles that the client can set as the primary role.
	AllowedRoles []string `pulumi:"allowedRoles"`
	// Specifies whether the OAuth client or user can use a role that is not defined in the OAuth access token.
	AnyRoleMode *string `pulumi:"anyRoleMode"`
	// Specifies additional values that can be used for the access token's audience validation on top of using the Customer's Snowflake Account URL
	AudienceUrls []string `pulumi:"audienceUrls"`
	// Specifies the list of roles that a client cannot set as the primary role. Do not include ACCOUNTADMIN, ORGADMIN or SECURITYADMIN as they are already implicitly enforced and will cause in-place updates.
	BlockedRoles []string `pulumi:"blockedRoles"`
	// Specifies a comment for the OAuth integration.
	Comment *string `pulumi:"comment"`
	// Date and time when the External OAUTH integration was created.
	CreatedOn *string `pulumi:"createdOn"`
	// Specifies whether to initiate operation of the integration or suspend it.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the URL to define the OAuth 2.0 authorization server.
	Issuer *string `pulumi:"issuer"`
	// Specifies the endpoint or a list of endpoints from which to download public keys or certificates to validate an External OAuth access token. The maximum number of URLs that can be specified in the list is 3.
	JwsKeysUrls []string `pulumi:"jwsKeysUrls"`
	// Specifies the name of the External Oath integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
	Name *string `pulumi:"name"`
	// Specifies a Base64-encoded RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers.
	RsaPublicKey *string `pulumi:"rsaPublicKey"`
	// Specifies a second RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers. Used for key rotation.
	RsaPublicKey2 *string `pulumi:"rsaPublicKey2"`
	// Specifies the scope delimiter in the authorization token.
	ScopeDelimiter *string `pulumi:"scopeDelimiter"`
	// Indicates which Snowflake user record attribute should be used to map the access token to a Snowflake user record.
	SnowflakeUserMappingAttribute *string `pulumi:"snowflakeUserMappingAttribute"`
	// Specifies the access token claim or claims that can be used to map the access token to a Snowflake user record.
	TokenUserMappingClaims []string `pulumi:"tokenUserMappingClaims"`
	// Specifies the OAuth 2.0 authorization server to be Okta, Microsoft Azure AD, Ping Identity PingFederate, or a Custom OAuth 2.0 authorization server.
	Type *string `pulumi:"type"`
}

type ExternalOauthIntegrationState struct {
	// Specifies the list of roles that the client can set as the primary role.
	AllowedRoles pulumi.StringArrayInput
	// Specifies whether the OAuth client or user can use a role that is not defined in the OAuth access token.
	AnyRoleMode pulumi.StringPtrInput
	// Specifies additional values that can be used for the access token's audience validation on top of using the Customer's Snowflake Account URL
	AudienceUrls pulumi.StringArrayInput
	// Specifies the list of roles that a client cannot set as the primary role. Do not include ACCOUNTADMIN, ORGADMIN or SECURITYADMIN as they are already implicitly enforced and will cause in-place updates.
	BlockedRoles pulumi.StringArrayInput
	// Specifies a comment for the OAuth integration.
	Comment pulumi.StringPtrInput
	// Date and time when the External OAUTH integration was created.
	CreatedOn pulumi.StringPtrInput
	// Specifies whether to initiate operation of the integration or suspend it.
	Enabled pulumi.BoolPtrInput
	// Specifies the URL to define the OAuth 2.0 authorization server.
	Issuer pulumi.StringPtrInput
	// Specifies the endpoint or a list of endpoints from which to download public keys or certificates to validate an External OAuth access token. The maximum number of URLs that can be specified in the list is 3.
	JwsKeysUrls pulumi.StringArrayInput
	// Specifies the name of the External Oath integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
	Name pulumi.StringPtrInput
	// Specifies a Base64-encoded RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers.
	RsaPublicKey pulumi.StringPtrInput
	// Specifies a second RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers. Used for key rotation.
	RsaPublicKey2 pulumi.StringPtrInput
	// Specifies the scope delimiter in the authorization token.
	ScopeDelimiter pulumi.StringPtrInput
	// Indicates which Snowflake user record attribute should be used to map the access token to a Snowflake user record.
	SnowflakeUserMappingAttribute pulumi.StringPtrInput
	// Specifies the access token claim or claims that can be used to map the access token to a Snowflake user record.
	TokenUserMappingClaims pulumi.StringArrayInput
	// Specifies the OAuth 2.0 authorization server to be Okta, Microsoft Azure AD, Ping Identity PingFederate, or a Custom OAuth 2.0 authorization server.
	Type pulumi.StringPtrInput
}

func (ExternalOauthIntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalOauthIntegrationState)(nil)).Elem()
}

type externalOauthIntegrationArgs struct {
	// Specifies the list of roles that the client can set as the primary role.
	AllowedRoles []string `pulumi:"allowedRoles"`
	// Specifies whether the OAuth client or user can use a role that is not defined in the OAuth access token.
	AnyRoleMode *string `pulumi:"anyRoleMode"`
	// Specifies additional values that can be used for the access token's audience validation on top of using the Customer's Snowflake Account URL
	AudienceUrls []string `pulumi:"audienceUrls"`
	// Specifies the list of roles that a client cannot set as the primary role. Do not include ACCOUNTADMIN, ORGADMIN or SECURITYADMIN as they are already implicitly enforced and will cause in-place updates.
	BlockedRoles []string `pulumi:"blockedRoles"`
	// Specifies a comment for the OAuth integration.
	Comment *string `pulumi:"comment"`
	// Specifies whether to initiate operation of the integration or suspend it.
	Enabled bool `pulumi:"enabled"`
	// Specifies the URL to define the OAuth 2.0 authorization server.
	Issuer string `pulumi:"issuer"`
	// Specifies the endpoint or a list of endpoints from which to download public keys or certificates to validate an External OAuth access token. The maximum number of URLs that can be specified in the list is 3.
	JwsKeysUrls []string `pulumi:"jwsKeysUrls"`
	// Specifies the name of the External Oath integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
	Name *string `pulumi:"name"`
	// Specifies a Base64-encoded RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers.
	RsaPublicKey *string `pulumi:"rsaPublicKey"`
	// Specifies a second RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers. Used for key rotation.
	RsaPublicKey2 *string `pulumi:"rsaPublicKey2"`
	// Specifies the scope delimiter in the authorization token.
	ScopeDelimiter *string `pulumi:"scopeDelimiter"`
	// Indicates which Snowflake user record attribute should be used to map the access token to a Snowflake user record.
	SnowflakeUserMappingAttribute string `pulumi:"snowflakeUserMappingAttribute"`
	// Specifies the access token claim or claims that can be used to map the access token to a Snowflake user record.
	TokenUserMappingClaims []string `pulumi:"tokenUserMappingClaims"`
	// Specifies the OAuth 2.0 authorization server to be Okta, Microsoft Azure AD, Ping Identity PingFederate, or a Custom OAuth 2.0 authorization server.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a ExternalOauthIntegration resource.
type ExternalOauthIntegrationArgs struct {
	// Specifies the list of roles that the client can set as the primary role.
	AllowedRoles pulumi.StringArrayInput
	// Specifies whether the OAuth client or user can use a role that is not defined in the OAuth access token.
	AnyRoleMode pulumi.StringPtrInput
	// Specifies additional values that can be used for the access token's audience validation on top of using the Customer's Snowflake Account URL
	AudienceUrls pulumi.StringArrayInput
	// Specifies the list of roles that a client cannot set as the primary role. Do not include ACCOUNTADMIN, ORGADMIN or SECURITYADMIN as they are already implicitly enforced and will cause in-place updates.
	BlockedRoles pulumi.StringArrayInput
	// Specifies a comment for the OAuth integration.
	Comment pulumi.StringPtrInput
	// Specifies whether to initiate operation of the integration or suspend it.
	Enabled pulumi.BoolInput
	// Specifies the URL to define the OAuth 2.0 authorization server.
	Issuer pulumi.StringInput
	// Specifies the endpoint or a list of endpoints from which to download public keys or certificates to validate an External OAuth access token. The maximum number of URLs that can be specified in the list is 3.
	JwsKeysUrls pulumi.StringArrayInput
	// Specifies the name of the External Oath integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
	Name pulumi.StringPtrInput
	// Specifies a Base64-encoded RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers.
	RsaPublicKey pulumi.StringPtrInput
	// Specifies a second RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers. Used for key rotation.
	RsaPublicKey2 pulumi.StringPtrInput
	// Specifies the scope delimiter in the authorization token.
	ScopeDelimiter pulumi.StringPtrInput
	// Indicates which Snowflake user record attribute should be used to map the access token to a Snowflake user record.
	SnowflakeUserMappingAttribute pulumi.StringInput
	// Specifies the access token claim or claims that can be used to map the access token to a Snowflake user record.
	TokenUserMappingClaims pulumi.StringArrayInput
	// Specifies the OAuth 2.0 authorization server to be Okta, Microsoft Azure AD, Ping Identity PingFederate, or a Custom OAuth 2.0 authorization server.
	Type pulumi.StringInput
}

func (ExternalOauthIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalOauthIntegrationArgs)(nil)).Elem()
}

type ExternalOauthIntegrationInput interface {
	pulumi.Input

	ToExternalOauthIntegrationOutput() ExternalOauthIntegrationOutput
	ToExternalOauthIntegrationOutputWithContext(ctx context.Context) ExternalOauthIntegrationOutput
}

func (*ExternalOauthIntegration) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalOauthIntegration)(nil)).Elem()
}

func (i *ExternalOauthIntegration) ToExternalOauthIntegrationOutput() ExternalOauthIntegrationOutput {
	return i.ToExternalOauthIntegrationOutputWithContext(context.Background())
}

func (i *ExternalOauthIntegration) ToExternalOauthIntegrationOutputWithContext(ctx context.Context) ExternalOauthIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalOauthIntegrationOutput)
}

// ExternalOauthIntegrationArrayInput is an input type that accepts ExternalOauthIntegrationArray and ExternalOauthIntegrationArrayOutput values.
// You can construct a concrete instance of `ExternalOauthIntegrationArrayInput` via:
//
//          ExternalOauthIntegrationArray{ ExternalOauthIntegrationArgs{...} }
type ExternalOauthIntegrationArrayInput interface {
	pulumi.Input

	ToExternalOauthIntegrationArrayOutput() ExternalOauthIntegrationArrayOutput
	ToExternalOauthIntegrationArrayOutputWithContext(context.Context) ExternalOauthIntegrationArrayOutput
}

type ExternalOauthIntegrationArray []ExternalOauthIntegrationInput

func (ExternalOauthIntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalOauthIntegration)(nil)).Elem()
}

func (i ExternalOauthIntegrationArray) ToExternalOauthIntegrationArrayOutput() ExternalOauthIntegrationArrayOutput {
	return i.ToExternalOauthIntegrationArrayOutputWithContext(context.Background())
}

func (i ExternalOauthIntegrationArray) ToExternalOauthIntegrationArrayOutputWithContext(ctx context.Context) ExternalOauthIntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalOauthIntegrationArrayOutput)
}

// ExternalOauthIntegrationMapInput is an input type that accepts ExternalOauthIntegrationMap and ExternalOauthIntegrationMapOutput values.
// You can construct a concrete instance of `ExternalOauthIntegrationMapInput` via:
//
//          ExternalOauthIntegrationMap{ "key": ExternalOauthIntegrationArgs{...} }
type ExternalOauthIntegrationMapInput interface {
	pulumi.Input

	ToExternalOauthIntegrationMapOutput() ExternalOauthIntegrationMapOutput
	ToExternalOauthIntegrationMapOutputWithContext(context.Context) ExternalOauthIntegrationMapOutput
}

type ExternalOauthIntegrationMap map[string]ExternalOauthIntegrationInput

func (ExternalOauthIntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalOauthIntegration)(nil)).Elem()
}

func (i ExternalOauthIntegrationMap) ToExternalOauthIntegrationMapOutput() ExternalOauthIntegrationMapOutput {
	return i.ToExternalOauthIntegrationMapOutputWithContext(context.Background())
}

func (i ExternalOauthIntegrationMap) ToExternalOauthIntegrationMapOutputWithContext(ctx context.Context) ExternalOauthIntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalOauthIntegrationMapOutput)
}

type ExternalOauthIntegrationOutput struct{ *pulumi.OutputState }

func (ExternalOauthIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalOauthIntegration)(nil)).Elem()
}

func (o ExternalOauthIntegrationOutput) ToExternalOauthIntegrationOutput() ExternalOauthIntegrationOutput {
	return o
}

func (o ExternalOauthIntegrationOutput) ToExternalOauthIntegrationOutputWithContext(ctx context.Context) ExternalOauthIntegrationOutput {
	return o
}

type ExternalOauthIntegrationArrayOutput struct{ *pulumi.OutputState }

func (ExternalOauthIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalOauthIntegration)(nil)).Elem()
}

func (o ExternalOauthIntegrationArrayOutput) ToExternalOauthIntegrationArrayOutput() ExternalOauthIntegrationArrayOutput {
	return o
}

func (o ExternalOauthIntegrationArrayOutput) ToExternalOauthIntegrationArrayOutputWithContext(ctx context.Context) ExternalOauthIntegrationArrayOutput {
	return o
}

func (o ExternalOauthIntegrationArrayOutput) Index(i pulumi.IntInput) ExternalOauthIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExternalOauthIntegration {
		return vs[0].([]*ExternalOauthIntegration)[vs[1].(int)]
	}).(ExternalOauthIntegrationOutput)
}

type ExternalOauthIntegrationMapOutput struct{ *pulumi.OutputState }

func (ExternalOauthIntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalOauthIntegration)(nil)).Elem()
}

func (o ExternalOauthIntegrationMapOutput) ToExternalOauthIntegrationMapOutput() ExternalOauthIntegrationMapOutput {
	return o
}

func (o ExternalOauthIntegrationMapOutput) ToExternalOauthIntegrationMapOutputWithContext(ctx context.Context) ExternalOauthIntegrationMapOutput {
	return o
}

func (o ExternalOauthIntegrationMapOutput) MapIndex(k pulumi.StringInput) ExternalOauthIntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExternalOauthIntegration {
		return vs[0].(map[string]*ExternalOauthIntegration)[vs[1].(string)]
	}).(ExternalOauthIntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalOauthIntegrationInput)(nil)).Elem(), &ExternalOauthIntegration{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalOauthIntegrationArrayInput)(nil)).Elem(), ExternalOauthIntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalOauthIntegrationMapInput)(nil)).Elem(), ExternalOauthIntegrationMap{})
	pulumi.RegisterOutputType(ExternalOauthIntegrationOutput{})
	pulumi.RegisterOutputType(ExternalOauthIntegrationArrayOutput{})
	pulumi.RegisterOutputType(ExternalOauthIntegrationMapOutput{})
}
