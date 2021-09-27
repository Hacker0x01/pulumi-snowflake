// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package snowflake

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSystemGetPrivateLinkConfig(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetSystemGetPrivateLinkConfigResult, error) {
	var rv GetSystemGetPrivateLinkConfigResult
	err := ctx.Invoke("snowflake:index/getSystemGetPrivateLinkConfig:getSystemGetPrivateLinkConfig", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getSystemGetPrivateLinkConfig.
type GetSystemGetPrivateLinkConfigResult struct {
	// The name of your Snowflake account.
	AccountName string `pulumi:"accountName"`
	// The URL used to connect to Snowflake through AWS PrivateLink or Azure Private Link.
	AccountUrl string `pulumi:"accountUrl"`
	// The AWS VPCE ID for your account.
	AwsVpceId string `pulumi:"awsVpceId"`
	// The Azure Private Link Service ID for your account.
	AzurePlsId string `pulumi:"azurePlsId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The OCSP URL corresponding to your Snowflake account that uses AWS PrivateLink or Azure Private Link.
	OcspUrl string `pulumi:"ocspUrl"`
}
