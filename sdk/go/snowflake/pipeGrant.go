// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package snowflake

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// # format is database name | schema name | pipe name | privilege | true/false for with_grant_option
//
// ```sh
//  $ pulumi import snowflake:index/pipeGrant:PipeGrant example 'dbName|schemaName|pipeName|OPERATE|false'
// ```
type PipeGrant struct {
	pulumi.CustomResourceState

	// The name of the database containing the current or future pipes on which to grant privileges.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// When this is set to true and a schema*name is provided, apply this grant on all future pipes in the given schema. When this is true and no schema*name is provided apply this grant on all future pipes in the given database. The pipe*name field must be unset in order to use on*future.
	OnFuture pulumi.BoolPtrOutput `pulumi:"onFuture"`
	// The name of the pipe on which to grant privileges immediately (only valid if onFuture is false).
	PipeName pulumi.StringPtrOutput `pulumi:"pipeName"`
	// The privilege to grant on the current or future pipe.
	Privilege pulumi.StringPtrOutput `pulumi:"privilege"`
	// Grants privilege to these roles.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The name of the schema containing the current or future pipes on which to grant privileges.
	SchemaName pulumi.StringOutput `pulumi:"schemaName"`
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	WithGrantOption pulumi.BoolPtrOutput `pulumi:"withGrantOption"`
}

// NewPipeGrant registers a new resource with the given unique name, arguments, and options.
func NewPipeGrant(ctx *pulumi.Context,
	name string, args *PipeGrantArgs, opts ...pulumi.ResourceOption) (*PipeGrant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.SchemaName == nil {
		return nil, errors.New("invalid value for required argument 'SchemaName'")
	}
	var resource PipeGrant
	err := ctx.RegisterResource("snowflake:index/pipeGrant:PipeGrant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipeGrant gets an existing PipeGrant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipeGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipeGrantState, opts ...pulumi.ResourceOption) (*PipeGrant, error) {
	var resource PipeGrant
	err := ctx.ReadResource("snowflake:index/pipeGrant:PipeGrant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PipeGrant resources.
type pipeGrantState struct {
	// The name of the database containing the current or future pipes on which to grant privileges.
	DatabaseName *string `pulumi:"databaseName"`
	// When this is set to true and a schema*name is provided, apply this grant on all future pipes in the given schema. When this is true and no schema*name is provided apply this grant on all future pipes in the given database. The pipe*name field must be unset in order to use on*future.
	OnFuture *bool `pulumi:"onFuture"`
	// The name of the pipe on which to grant privileges immediately (only valid if onFuture is false).
	PipeName *string `pulumi:"pipeName"`
	// The privilege to grant on the current or future pipe.
	Privilege *string `pulumi:"privilege"`
	// Grants privilege to these roles.
	Roles []string `pulumi:"roles"`
	// The name of the schema containing the current or future pipes on which to grant privileges.
	SchemaName *string `pulumi:"schemaName"`
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

type PipeGrantState struct {
	// The name of the database containing the current or future pipes on which to grant privileges.
	DatabaseName pulumi.StringPtrInput
	// When this is set to true and a schema*name is provided, apply this grant on all future pipes in the given schema. When this is true and no schema*name is provided apply this grant on all future pipes in the given database. The pipe*name field must be unset in order to use on*future.
	OnFuture pulumi.BoolPtrInput
	// The name of the pipe on which to grant privileges immediately (only valid if onFuture is false).
	PipeName pulumi.StringPtrInput
	// The privilege to grant on the current or future pipe.
	Privilege pulumi.StringPtrInput
	// Grants privilege to these roles.
	Roles pulumi.StringArrayInput
	// The name of the schema containing the current or future pipes on which to grant privileges.
	SchemaName pulumi.StringPtrInput
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	WithGrantOption pulumi.BoolPtrInput
}

func (PipeGrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipeGrantState)(nil)).Elem()
}

type pipeGrantArgs struct {
	// The name of the database containing the current or future pipes on which to grant privileges.
	DatabaseName string `pulumi:"databaseName"`
	// When this is set to true and a schema*name is provided, apply this grant on all future pipes in the given schema. When this is true and no schema*name is provided apply this grant on all future pipes in the given database. The pipe*name field must be unset in order to use on*future.
	OnFuture *bool `pulumi:"onFuture"`
	// The name of the pipe on which to grant privileges immediately (only valid if onFuture is false).
	PipeName *string `pulumi:"pipeName"`
	// The privilege to grant on the current or future pipe.
	Privilege *string `pulumi:"privilege"`
	// Grants privilege to these roles.
	Roles []string `pulumi:"roles"`
	// The name of the schema containing the current or future pipes on which to grant privileges.
	SchemaName string `pulumi:"schemaName"`
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

// The set of arguments for constructing a PipeGrant resource.
type PipeGrantArgs struct {
	// The name of the database containing the current or future pipes on which to grant privileges.
	DatabaseName pulumi.StringInput
	// When this is set to true and a schema*name is provided, apply this grant on all future pipes in the given schema. When this is true and no schema*name is provided apply this grant on all future pipes in the given database. The pipe*name field must be unset in order to use on*future.
	OnFuture pulumi.BoolPtrInput
	// The name of the pipe on which to grant privileges immediately (only valid if onFuture is false).
	PipeName pulumi.StringPtrInput
	// The privilege to grant on the current or future pipe.
	Privilege pulumi.StringPtrInput
	// Grants privilege to these roles.
	Roles pulumi.StringArrayInput
	// The name of the schema containing the current or future pipes on which to grant privileges.
	SchemaName pulumi.StringInput
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	WithGrantOption pulumi.BoolPtrInput
}

func (PipeGrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipeGrantArgs)(nil)).Elem()
}

type PipeGrantInput interface {
	pulumi.Input

	ToPipeGrantOutput() PipeGrantOutput
	ToPipeGrantOutputWithContext(ctx context.Context) PipeGrantOutput
}

func (*PipeGrant) ElementType() reflect.Type {
	return reflect.TypeOf((*PipeGrant)(nil))
}

func (i *PipeGrant) ToPipeGrantOutput() PipeGrantOutput {
	return i.ToPipeGrantOutputWithContext(context.Background())
}

func (i *PipeGrant) ToPipeGrantOutputWithContext(ctx context.Context) PipeGrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipeGrantOutput)
}

func (i *PipeGrant) ToPipeGrantPtrOutput() PipeGrantPtrOutput {
	return i.ToPipeGrantPtrOutputWithContext(context.Background())
}

func (i *PipeGrant) ToPipeGrantPtrOutputWithContext(ctx context.Context) PipeGrantPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipeGrantPtrOutput)
}

type PipeGrantPtrInput interface {
	pulumi.Input

	ToPipeGrantPtrOutput() PipeGrantPtrOutput
	ToPipeGrantPtrOutputWithContext(ctx context.Context) PipeGrantPtrOutput
}

type pipeGrantPtrType PipeGrantArgs

func (*pipeGrantPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipeGrant)(nil))
}

func (i *pipeGrantPtrType) ToPipeGrantPtrOutput() PipeGrantPtrOutput {
	return i.ToPipeGrantPtrOutputWithContext(context.Background())
}

func (i *pipeGrantPtrType) ToPipeGrantPtrOutputWithContext(ctx context.Context) PipeGrantPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipeGrantPtrOutput)
}

// PipeGrantArrayInput is an input type that accepts PipeGrantArray and PipeGrantArrayOutput values.
// You can construct a concrete instance of `PipeGrantArrayInput` via:
//
//          PipeGrantArray{ PipeGrantArgs{...} }
type PipeGrantArrayInput interface {
	pulumi.Input

	ToPipeGrantArrayOutput() PipeGrantArrayOutput
	ToPipeGrantArrayOutputWithContext(context.Context) PipeGrantArrayOutput
}

type PipeGrantArray []PipeGrantInput

func (PipeGrantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PipeGrant)(nil)).Elem()
}

func (i PipeGrantArray) ToPipeGrantArrayOutput() PipeGrantArrayOutput {
	return i.ToPipeGrantArrayOutputWithContext(context.Background())
}

func (i PipeGrantArray) ToPipeGrantArrayOutputWithContext(ctx context.Context) PipeGrantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipeGrantArrayOutput)
}

// PipeGrantMapInput is an input type that accepts PipeGrantMap and PipeGrantMapOutput values.
// You can construct a concrete instance of `PipeGrantMapInput` via:
//
//          PipeGrantMap{ "key": PipeGrantArgs{...} }
type PipeGrantMapInput interface {
	pulumi.Input

	ToPipeGrantMapOutput() PipeGrantMapOutput
	ToPipeGrantMapOutputWithContext(context.Context) PipeGrantMapOutput
}

type PipeGrantMap map[string]PipeGrantInput

func (PipeGrantMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PipeGrant)(nil)).Elem()
}

func (i PipeGrantMap) ToPipeGrantMapOutput() PipeGrantMapOutput {
	return i.ToPipeGrantMapOutputWithContext(context.Background())
}

func (i PipeGrantMap) ToPipeGrantMapOutputWithContext(ctx context.Context) PipeGrantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipeGrantMapOutput)
}

type PipeGrantOutput struct{ *pulumi.OutputState }

func (PipeGrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipeGrant)(nil))
}

func (o PipeGrantOutput) ToPipeGrantOutput() PipeGrantOutput {
	return o
}

func (o PipeGrantOutput) ToPipeGrantOutputWithContext(ctx context.Context) PipeGrantOutput {
	return o
}

func (o PipeGrantOutput) ToPipeGrantPtrOutput() PipeGrantPtrOutput {
	return o.ToPipeGrantPtrOutputWithContext(context.Background())
}

func (o PipeGrantOutput) ToPipeGrantPtrOutputWithContext(ctx context.Context) PipeGrantPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipeGrant) *PipeGrant {
		return &v
	}).(PipeGrantPtrOutput)
}

type PipeGrantPtrOutput struct{ *pulumi.OutputState }

func (PipeGrantPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipeGrant)(nil))
}

func (o PipeGrantPtrOutput) ToPipeGrantPtrOutput() PipeGrantPtrOutput {
	return o
}

func (o PipeGrantPtrOutput) ToPipeGrantPtrOutputWithContext(ctx context.Context) PipeGrantPtrOutput {
	return o
}

func (o PipeGrantPtrOutput) Elem() PipeGrantOutput {
	return o.ApplyT(func(v *PipeGrant) PipeGrant {
		if v != nil {
			return *v
		}
		var ret PipeGrant
		return ret
	}).(PipeGrantOutput)
}

type PipeGrantArrayOutput struct{ *pulumi.OutputState }

func (PipeGrantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PipeGrant)(nil))
}

func (o PipeGrantArrayOutput) ToPipeGrantArrayOutput() PipeGrantArrayOutput {
	return o
}

func (o PipeGrantArrayOutput) ToPipeGrantArrayOutputWithContext(ctx context.Context) PipeGrantArrayOutput {
	return o
}

func (o PipeGrantArrayOutput) Index(i pulumi.IntInput) PipeGrantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PipeGrant {
		return vs[0].([]PipeGrant)[vs[1].(int)]
	}).(PipeGrantOutput)
}

type PipeGrantMapOutput struct{ *pulumi.OutputState }

func (PipeGrantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PipeGrant)(nil))
}

func (o PipeGrantMapOutput) ToPipeGrantMapOutput() PipeGrantMapOutput {
	return o
}

func (o PipeGrantMapOutput) ToPipeGrantMapOutputWithContext(ctx context.Context) PipeGrantMapOutput {
	return o
}

func (o PipeGrantMapOutput) MapIndex(k pulumi.StringInput) PipeGrantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PipeGrant {
		return vs[0].(map[string]PipeGrant)[vs[1].(string)]
	}).(PipeGrantOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PipeGrantInput)(nil)).Elem(), &PipeGrant{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipeGrantPtrInput)(nil)).Elem(), &PipeGrant{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipeGrantArrayInput)(nil)).Elem(), PipeGrantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipeGrantMapInput)(nil)).Elem(), PipeGrantMap{})
	pulumi.RegisterOutputType(PipeGrantOutput{})
	pulumi.RegisterOutputType(PipeGrantPtrOutput{})
	pulumi.RegisterOutputType(PipeGrantArrayOutput{})
	pulumi.RegisterOutputType(PipeGrantMapOutput{})
}
