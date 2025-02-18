// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package snowflake

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ssm"
// 	"github.com/pulumi/pulumi-snowflake/sdk/go/snowflake"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		this, err := snowflake.GetCurrentAccount(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ssm.NewParameter(ctx, "snowflakeAccountUrl", &ssm.ParameterArgs{
// 			Type:  pulumi.String("String"),
// 			Value: pulumi.String(this.Url),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCurrentAccount(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetCurrentAccountResult, error) {
	var rv GetCurrentAccountResult
	err := ctx.Invoke("snowflake:index/getCurrentAccount:getCurrentAccount", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getCurrentAccount.
type GetCurrentAccountResult struct {
	// The Snowflake Account ID; as returned by CURRENT_ACCOUNT().
	Account string `pulumi:"account"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Snowflake Region; as returned by CURRENT_REGION()
	Region string `pulumi:"region"`
	// The Snowflake URL.
	Url string `pulumi:"url"`
}
