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
// 		_, err := snowflake.GetStages(ctx, &GetStagesArgs{
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
func GetStages(ctx *pulumi.Context, args *GetStagesArgs, opts ...pulumi.InvokeOption) (*GetStagesResult, error) {
	var rv GetStagesResult
	err := ctx.Invoke("snowflake:index/getStages:getStages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStages.
type GetStagesArgs struct {
	// The database from which to return the schemas from.
	Database string `pulumi:"database"`
	// The schema from which to return the stages from.
	Schema string `pulumi:"schema"`
}

// A collection of values returned by getStages.
type GetStagesResult struct {
	// The database from which to return the schemas from.
	Database string `pulumi:"database"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The schema from which to return the stages from.
	Schema string `pulumi:"schema"`
	// The stages in the schema
	Stages []GetStagesStage `pulumi:"stages"`
}

func GetStagesOutput(ctx *pulumi.Context, args GetStagesOutputArgs, opts ...pulumi.InvokeOption) GetStagesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStagesResult, error) {
			args := v.(GetStagesArgs)
			r, err := GetStages(ctx, &args, opts...)
			return *r, err
		}).(GetStagesResultOutput)
}

// A collection of arguments for invoking getStages.
type GetStagesOutputArgs struct {
	// The database from which to return the schemas from.
	Database pulumi.StringInput `pulumi:"database"`
	// The schema from which to return the stages from.
	Schema pulumi.StringInput `pulumi:"schema"`
}

func (GetStagesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStagesArgs)(nil)).Elem()
}

// A collection of values returned by getStages.
type GetStagesResultOutput struct{ *pulumi.OutputState }

func (GetStagesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStagesResult)(nil)).Elem()
}

func (o GetStagesResultOutput) ToGetStagesResultOutput() GetStagesResultOutput {
	return o
}

func (o GetStagesResultOutput) ToGetStagesResultOutputWithContext(ctx context.Context) GetStagesResultOutput {
	return o
}

// The database from which to return the schemas from.
func (o GetStagesResultOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v GetStagesResult) string { return v.Database }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStagesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStagesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The schema from which to return the stages from.
func (o GetStagesResultOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v GetStagesResult) string { return v.Schema }).(pulumi.StringOutput)
}

// The stages in the schema
func (o GetStagesResultOutput) Stages() GetStagesStageArrayOutput {
	return o.ApplyT(func(v GetStagesResult) []GetStagesStage { return v.Stages }).(GetStagesStageArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStagesResultOutput{})
}
