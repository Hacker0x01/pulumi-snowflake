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
// 		_, err := snowflake.NewExternalTableGrant(ctx, "grant", &snowflake.ExternalTableGrantArgs{
// 			DatabaseName:      pulumi.String("db"),
// 			ExternalTableName: pulumi.String("external_table"),
// 			OnFuture:          pulumi.Bool(false),
// 			Privilege:         pulumi.String("select"),
// 			Roles: pulumi.StringArray{
// 				pulumi.String("role1"),
// 				pulumi.String("role2"),
// 			},
// 			SchemaName: pulumi.String("schema"),
// 			Shares: pulumi.StringArray{
// 				pulumi.String("share1"),
// 				pulumi.String("share2"),
// 			},
// 			WithGrantOption: pulumi.Bool(false),
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
// # format is database name | schema name | external table name | privilege | true/false for with_grant_option
//
// ```sh
//  $ pulumi import snowflake:index/externalTableGrant:ExternalTableGrant example 'dbName|schemaName|externalTableName|SELECT|false'
// ```
type ExternalTableGrant struct {
	pulumi.CustomResourceState

	// The name of the database containing the current or future external tables on which to grant privileges.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The name of the external table on which to grant privileges immediately (only valid if onFuture is false).
	ExternalTableName pulumi.StringPtrOutput `pulumi:"externalTableName"`
	// When this is set to true and a schema*name is provided, apply this grant on all future external tables in the given schema. When this is true and no schema*name is provided apply this grant on all future external tables in the given database. The external*table*name and shares fields must be unset in order to use on_future.
	OnFuture pulumi.BoolPtrOutput `pulumi:"onFuture"`
	// The privilege to grant on the current or future external table.
	Privilege pulumi.StringPtrOutput `pulumi:"privilege"`
	// Grants privilege to these roles.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The name of the schema containing the current or future external tables on which to grant privileges.
	SchemaName pulumi.StringOutput `pulumi:"schemaName"`
	// Grants privilege to these shares (only valid if onFuture is false).
	Shares pulumi.StringArrayOutput `pulumi:"shares"`
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	WithGrantOption pulumi.BoolPtrOutput `pulumi:"withGrantOption"`
}

// NewExternalTableGrant registers a new resource with the given unique name, arguments, and options.
func NewExternalTableGrant(ctx *pulumi.Context,
	name string, args *ExternalTableGrantArgs, opts ...pulumi.ResourceOption) (*ExternalTableGrant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.SchemaName == nil {
		return nil, errors.New("invalid value for required argument 'SchemaName'")
	}
	var resource ExternalTableGrant
	err := ctx.RegisterResource("snowflake:index/externalTableGrant:ExternalTableGrant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalTableGrant gets an existing ExternalTableGrant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalTableGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalTableGrantState, opts ...pulumi.ResourceOption) (*ExternalTableGrant, error) {
	var resource ExternalTableGrant
	err := ctx.ReadResource("snowflake:index/externalTableGrant:ExternalTableGrant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalTableGrant resources.
type externalTableGrantState struct {
	// The name of the database containing the current or future external tables on which to grant privileges.
	DatabaseName *string `pulumi:"databaseName"`
	// The name of the external table on which to grant privileges immediately (only valid if onFuture is false).
	ExternalTableName *string `pulumi:"externalTableName"`
	// When this is set to true and a schema*name is provided, apply this grant on all future external tables in the given schema. When this is true and no schema*name is provided apply this grant on all future external tables in the given database. The external*table*name and shares fields must be unset in order to use on_future.
	OnFuture *bool `pulumi:"onFuture"`
	// The privilege to grant on the current or future external table.
	Privilege *string `pulumi:"privilege"`
	// Grants privilege to these roles.
	Roles []string `pulumi:"roles"`
	// The name of the schema containing the current or future external tables on which to grant privileges.
	SchemaName *string `pulumi:"schemaName"`
	// Grants privilege to these shares (only valid if onFuture is false).
	Shares []string `pulumi:"shares"`
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

type ExternalTableGrantState struct {
	// The name of the database containing the current or future external tables on which to grant privileges.
	DatabaseName pulumi.StringPtrInput
	// The name of the external table on which to grant privileges immediately (only valid if onFuture is false).
	ExternalTableName pulumi.StringPtrInput
	// When this is set to true and a schema*name is provided, apply this grant on all future external tables in the given schema. When this is true and no schema*name is provided apply this grant on all future external tables in the given database. The external*table*name and shares fields must be unset in order to use on_future.
	OnFuture pulumi.BoolPtrInput
	// The privilege to grant on the current or future external table.
	Privilege pulumi.StringPtrInput
	// Grants privilege to these roles.
	Roles pulumi.StringArrayInput
	// The name of the schema containing the current or future external tables on which to grant privileges.
	SchemaName pulumi.StringPtrInput
	// Grants privilege to these shares (only valid if onFuture is false).
	Shares pulumi.StringArrayInput
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	WithGrantOption pulumi.BoolPtrInput
}

func (ExternalTableGrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalTableGrantState)(nil)).Elem()
}

type externalTableGrantArgs struct {
	// The name of the database containing the current or future external tables on which to grant privileges.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the external table on which to grant privileges immediately (only valid if onFuture is false).
	ExternalTableName *string `pulumi:"externalTableName"`
	// When this is set to true and a schema*name is provided, apply this grant on all future external tables in the given schema. When this is true and no schema*name is provided apply this grant on all future external tables in the given database. The external*table*name and shares fields must be unset in order to use on_future.
	OnFuture *bool `pulumi:"onFuture"`
	// The privilege to grant on the current or future external table.
	Privilege *string `pulumi:"privilege"`
	// Grants privilege to these roles.
	Roles []string `pulumi:"roles"`
	// The name of the schema containing the current or future external tables on which to grant privileges.
	SchemaName string `pulumi:"schemaName"`
	// Grants privilege to these shares (only valid if onFuture is false).
	Shares []string `pulumi:"shares"`
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

// The set of arguments for constructing a ExternalTableGrant resource.
type ExternalTableGrantArgs struct {
	// The name of the database containing the current or future external tables on which to grant privileges.
	DatabaseName pulumi.StringInput
	// The name of the external table on which to grant privileges immediately (only valid if onFuture is false).
	ExternalTableName pulumi.StringPtrInput
	// When this is set to true and a schema*name is provided, apply this grant on all future external tables in the given schema. When this is true and no schema*name is provided apply this grant on all future external tables in the given database. The external*table*name and shares fields must be unset in order to use on_future.
	OnFuture pulumi.BoolPtrInput
	// The privilege to grant on the current or future external table.
	Privilege pulumi.StringPtrInput
	// Grants privilege to these roles.
	Roles pulumi.StringArrayInput
	// The name of the schema containing the current or future external tables on which to grant privileges.
	SchemaName pulumi.StringInput
	// Grants privilege to these shares (only valid if onFuture is false).
	Shares pulumi.StringArrayInput
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	WithGrantOption pulumi.BoolPtrInput
}

func (ExternalTableGrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalTableGrantArgs)(nil)).Elem()
}

type ExternalTableGrantInput interface {
	pulumi.Input

	ToExternalTableGrantOutput() ExternalTableGrantOutput
	ToExternalTableGrantOutputWithContext(ctx context.Context) ExternalTableGrantOutput
}

func (*ExternalTableGrant) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalTableGrant)(nil)).Elem()
}

func (i *ExternalTableGrant) ToExternalTableGrantOutput() ExternalTableGrantOutput {
	return i.ToExternalTableGrantOutputWithContext(context.Background())
}

func (i *ExternalTableGrant) ToExternalTableGrantOutputWithContext(ctx context.Context) ExternalTableGrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalTableGrantOutput)
}

// ExternalTableGrantArrayInput is an input type that accepts ExternalTableGrantArray and ExternalTableGrantArrayOutput values.
// You can construct a concrete instance of `ExternalTableGrantArrayInput` via:
//
//          ExternalTableGrantArray{ ExternalTableGrantArgs{...} }
type ExternalTableGrantArrayInput interface {
	pulumi.Input

	ToExternalTableGrantArrayOutput() ExternalTableGrantArrayOutput
	ToExternalTableGrantArrayOutputWithContext(context.Context) ExternalTableGrantArrayOutput
}

type ExternalTableGrantArray []ExternalTableGrantInput

func (ExternalTableGrantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalTableGrant)(nil)).Elem()
}

func (i ExternalTableGrantArray) ToExternalTableGrantArrayOutput() ExternalTableGrantArrayOutput {
	return i.ToExternalTableGrantArrayOutputWithContext(context.Background())
}

func (i ExternalTableGrantArray) ToExternalTableGrantArrayOutputWithContext(ctx context.Context) ExternalTableGrantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalTableGrantArrayOutput)
}

// ExternalTableGrantMapInput is an input type that accepts ExternalTableGrantMap and ExternalTableGrantMapOutput values.
// You can construct a concrete instance of `ExternalTableGrantMapInput` via:
//
//          ExternalTableGrantMap{ "key": ExternalTableGrantArgs{...} }
type ExternalTableGrantMapInput interface {
	pulumi.Input

	ToExternalTableGrantMapOutput() ExternalTableGrantMapOutput
	ToExternalTableGrantMapOutputWithContext(context.Context) ExternalTableGrantMapOutput
}

type ExternalTableGrantMap map[string]ExternalTableGrantInput

func (ExternalTableGrantMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalTableGrant)(nil)).Elem()
}

func (i ExternalTableGrantMap) ToExternalTableGrantMapOutput() ExternalTableGrantMapOutput {
	return i.ToExternalTableGrantMapOutputWithContext(context.Background())
}

func (i ExternalTableGrantMap) ToExternalTableGrantMapOutputWithContext(ctx context.Context) ExternalTableGrantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalTableGrantMapOutput)
}

type ExternalTableGrantOutput struct{ *pulumi.OutputState }

func (ExternalTableGrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalTableGrant)(nil)).Elem()
}

func (o ExternalTableGrantOutput) ToExternalTableGrantOutput() ExternalTableGrantOutput {
	return o
}

func (o ExternalTableGrantOutput) ToExternalTableGrantOutputWithContext(ctx context.Context) ExternalTableGrantOutput {
	return o
}

type ExternalTableGrantArrayOutput struct{ *pulumi.OutputState }

func (ExternalTableGrantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalTableGrant)(nil)).Elem()
}

func (o ExternalTableGrantArrayOutput) ToExternalTableGrantArrayOutput() ExternalTableGrantArrayOutput {
	return o
}

func (o ExternalTableGrantArrayOutput) ToExternalTableGrantArrayOutputWithContext(ctx context.Context) ExternalTableGrantArrayOutput {
	return o
}

func (o ExternalTableGrantArrayOutput) Index(i pulumi.IntInput) ExternalTableGrantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExternalTableGrant {
		return vs[0].([]*ExternalTableGrant)[vs[1].(int)]
	}).(ExternalTableGrantOutput)
}

type ExternalTableGrantMapOutput struct{ *pulumi.OutputState }

func (ExternalTableGrantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalTableGrant)(nil)).Elem()
}

func (o ExternalTableGrantMapOutput) ToExternalTableGrantMapOutput() ExternalTableGrantMapOutput {
	return o
}

func (o ExternalTableGrantMapOutput) ToExternalTableGrantMapOutputWithContext(ctx context.Context) ExternalTableGrantMapOutput {
	return o
}

func (o ExternalTableGrantMapOutput) MapIndex(k pulumi.StringInput) ExternalTableGrantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExternalTableGrant {
		return vs[0].(map[string]*ExternalTableGrant)[vs[1].(string)]
	}).(ExternalTableGrantOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalTableGrantInput)(nil)).Elem(), &ExternalTableGrant{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalTableGrantArrayInput)(nil)).Elem(), ExternalTableGrantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalTableGrantMapInput)(nil)).Elem(), ExternalTableGrantMap{})
	pulumi.RegisterOutputType(ExternalTableGrantOutput{})
	pulumi.RegisterOutputType(ExternalTableGrantArrayOutput{})
	pulumi.RegisterOutputType(ExternalTableGrantMapOutput{})
}
