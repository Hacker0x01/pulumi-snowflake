// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package snowflake

import (
	"context"
	"reflect"

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
// 		_, err := snowflake.GetExternalFunctions(ctx, &GetExternalFunctionsArgs{
// 			Database: "MYDB",
// 			Schema:   "MYSCHEMA",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetExternalFunctions(ctx *pulumi.Context, args *GetExternalFunctionsArgs, opts ...pulumi.InvokeOption) (*GetExternalFunctionsResult, error) {
	var rv GetExternalFunctionsResult
	err := ctx.Invoke("snowflake:index/getExternalFunctions:getExternalFunctions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getExternalFunctions.
type GetExternalFunctionsArgs struct {
	// The database from which to return the schemas from.
	Database string `pulumi:"database"`
	// The schema from which to return the external functions from.
	Schema string `pulumi:"schema"`
}

// A collection of values returned by getExternalFunctions.
type GetExternalFunctionsResult struct {
	// The database from which to return the schemas from.
	Database string `pulumi:"database"`
	// The external functions in the schema
	ExternalFunctions []GetExternalFunctionsExternalFunction `pulumi:"externalFunctions"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The schema from which to return the external functions from.
	Schema string `pulumi:"schema"`
}

func GetExternalFunctionsOutput(ctx *pulumi.Context, args GetExternalFunctionsOutputArgs, opts ...pulumi.InvokeOption) GetExternalFunctionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetExternalFunctionsResult, error) {
			args := v.(GetExternalFunctionsArgs)
			r, err := GetExternalFunctions(ctx, &args, opts...)
			return *r, err
		}).(GetExternalFunctionsResultOutput)
}

// A collection of arguments for invoking getExternalFunctions.
type GetExternalFunctionsOutputArgs struct {
	// The database from which to return the schemas from.
	Database pulumi.StringInput `pulumi:"database"`
	// The schema from which to return the external functions from.
	Schema pulumi.StringInput `pulumi:"schema"`
}

func (GetExternalFunctionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExternalFunctionsArgs)(nil)).Elem()
}

// A collection of values returned by getExternalFunctions.
type GetExternalFunctionsResultOutput struct{ *pulumi.OutputState }

func (GetExternalFunctionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExternalFunctionsResult)(nil)).Elem()
}

func (o GetExternalFunctionsResultOutput) ToGetExternalFunctionsResultOutput() GetExternalFunctionsResultOutput {
	return o
}

func (o GetExternalFunctionsResultOutput) ToGetExternalFunctionsResultOutputWithContext(ctx context.Context) GetExternalFunctionsResultOutput {
	return o
}

// The database from which to return the schemas from.
func (o GetExternalFunctionsResultOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v GetExternalFunctionsResult) string { return v.Database }).(pulumi.StringOutput)
}

// The external functions in the schema
func (o GetExternalFunctionsResultOutput) ExternalFunctions() GetExternalFunctionsExternalFunctionArrayOutput {
	return o.ApplyT(func(v GetExternalFunctionsResult) []GetExternalFunctionsExternalFunction { return v.ExternalFunctions }).(GetExternalFunctionsExternalFunctionArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetExternalFunctionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetExternalFunctionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The schema from which to return the external functions from.
func (o GetExternalFunctionsResultOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v GetExternalFunctionsResult) string { return v.Schema }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetExternalFunctionsResultOutput{})
}
