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
// # format is database name | schema name | pipe name
//
// ```sh
//  $ pulumi import snowflake:index/pipe:Pipe example 'dbName|schemaName|pipeName'
// ```
type Pipe struct {
	pulumi.CustomResourceState

	// Specifies a autoIngest param for the pipe.
	AutoIngest pulumi.BoolPtrOutput `pulumi:"autoIngest"`
	// Specifies the Amazon Resource Name (ARN) for the SNS topic for your S3 bucket.
	AwsSnsTopicArn pulumi.StringPtrOutput `pulumi:"awsSnsTopicArn"`
	// Specifies a comment for the pipe.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Specifies the copy statement for the pipe.
	CopyStatement pulumi.StringOutput `pulumi:"copyStatement"`
	// The database in which to create the pipe.
	Database pulumi.StringOutput `pulumi:"database"`
	// Specifies the name of the notification integration used for error notifications.
	ErrorIntegration pulumi.StringPtrOutput `pulumi:"errorIntegration"`
	// Specifies an integration for the pipe.
	Integration pulumi.StringPtrOutput `pulumi:"integration"`
	// Specifies the identifier for the pipe; must be unique for the database and schema in which the pipe is created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Amazon Resource Name of the Amazon SQS queue for the stage named in the DEFINITION column.
	NotificationChannel pulumi.StringOutput `pulumi:"notificationChannel"`
	// Name of the role that owns the pipe.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The schema in which to create the pipe.
	Schema pulumi.StringOutput `pulumi:"schema"`
}

// NewPipe registers a new resource with the given unique name, arguments, and options.
func NewPipe(ctx *pulumi.Context,
	name string, args *PipeArgs, opts ...pulumi.ResourceOption) (*Pipe, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CopyStatement == nil {
		return nil, errors.New("invalid value for required argument 'CopyStatement'")
	}
	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.Schema == nil {
		return nil, errors.New("invalid value for required argument 'Schema'")
	}
	var resource Pipe
	err := ctx.RegisterResource("snowflake:index/pipe:Pipe", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipe gets an existing Pipe resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipe(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipeState, opts ...pulumi.ResourceOption) (*Pipe, error) {
	var resource Pipe
	err := ctx.ReadResource("snowflake:index/pipe:Pipe", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pipe resources.
type pipeState struct {
	// Specifies a autoIngest param for the pipe.
	AutoIngest *bool `pulumi:"autoIngest"`
	// Specifies the Amazon Resource Name (ARN) for the SNS topic for your S3 bucket.
	AwsSnsTopicArn *string `pulumi:"awsSnsTopicArn"`
	// Specifies a comment for the pipe.
	Comment *string `pulumi:"comment"`
	// Specifies the copy statement for the pipe.
	CopyStatement *string `pulumi:"copyStatement"`
	// The database in which to create the pipe.
	Database *string `pulumi:"database"`
	// Specifies the name of the notification integration used for error notifications.
	ErrorIntegration *string `pulumi:"errorIntegration"`
	// Specifies an integration for the pipe.
	Integration *string `pulumi:"integration"`
	// Specifies the identifier for the pipe; must be unique for the database and schema in which the pipe is created.
	Name *string `pulumi:"name"`
	// Amazon Resource Name of the Amazon SQS queue for the stage named in the DEFINITION column.
	NotificationChannel *string `pulumi:"notificationChannel"`
	// Name of the role that owns the pipe.
	Owner *string `pulumi:"owner"`
	// The schema in which to create the pipe.
	Schema *string `pulumi:"schema"`
}

type PipeState struct {
	// Specifies a autoIngest param for the pipe.
	AutoIngest pulumi.BoolPtrInput
	// Specifies the Amazon Resource Name (ARN) for the SNS topic for your S3 bucket.
	AwsSnsTopicArn pulumi.StringPtrInput
	// Specifies a comment for the pipe.
	Comment pulumi.StringPtrInput
	// Specifies the copy statement for the pipe.
	CopyStatement pulumi.StringPtrInput
	// The database in which to create the pipe.
	Database pulumi.StringPtrInput
	// Specifies the name of the notification integration used for error notifications.
	ErrorIntegration pulumi.StringPtrInput
	// Specifies an integration for the pipe.
	Integration pulumi.StringPtrInput
	// Specifies the identifier for the pipe; must be unique for the database and schema in which the pipe is created.
	Name pulumi.StringPtrInput
	// Amazon Resource Name of the Amazon SQS queue for the stage named in the DEFINITION column.
	NotificationChannel pulumi.StringPtrInput
	// Name of the role that owns the pipe.
	Owner pulumi.StringPtrInput
	// The schema in which to create the pipe.
	Schema pulumi.StringPtrInput
}

func (PipeState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipeState)(nil)).Elem()
}

type pipeArgs struct {
	// Specifies a autoIngest param for the pipe.
	AutoIngest *bool `pulumi:"autoIngest"`
	// Specifies the Amazon Resource Name (ARN) for the SNS topic for your S3 bucket.
	AwsSnsTopicArn *string `pulumi:"awsSnsTopicArn"`
	// Specifies a comment for the pipe.
	Comment *string `pulumi:"comment"`
	// Specifies the copy statement for the pipe.
	CopyStatement string `pulumi:"copyStatement"`
	// The database in which to create the pipe.
	Database string `pulumi:"database"`
	// Specifies the name of the notification integration used for error notifications.
	ErrorIntegration *string `pulumi:"errorIntegration"`
	// Specifies an integration for the pipe.
	Integration *string `pulumi:"integration"`
	// Specifies the identifier for the pipe; must be unique for the database and schema in which the pipe is created.
	Name *string `pulumi:"name"`
	// The schema in which to create the pipe.
	Schema string `pulumi:"schema"`
}

// The set of arguments for constructing a Pipe resource.
type PipeArgs struct {
	// Specifies a autoIngest param for the pipe.
	AutoIngest pulumi.BoolPtrInput
	// Specifies the Amazon Resource Name (ARN) for the SNS topic for your S3 bucket.
	AwsSnsTopicArn pulumi.StringPtrInput
	// Specifies a comment for the pipe.
	Comment pulumi.StringPtrInput
	// Specifies the copy statement for the pipe.
	CopyStatement pulumi.StringInput
	// The database in which to create the pipe.
	Database pulumi.StringInput
	// Specifies the name of the notification integration used for error notifications.
	ErrorIntegration pulumi.StringPtrInput
	// Specifies an integration for the pipe.
	Integration pulumi.StringPtrInput
	// Specifies the identifier for the pipe; must be unique for the database and schema in which the pipe is created.
	Name pulumi.StringPtrInput
	// The schema in which to create the pipe.
	Schema pulumi.StringInput
}

func (PipeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipeArgs)(nil)).Elem()
}

type PipeInput interface {
	pulumi.Input

	ToPipeOutput() PipeOutput
	ToPipeOutputWithContext(ctx context.Context) PipeOutput
}

func (*Pipe) ElementType() reflect.Type {
	return reflect.TypeOf((*Pipe)(nil))
}

func (i *Pipe) ToPipeOutput() PipeOutput {
	return i.ToPipeOutputWithContext(context.Background())
}

func (i *Pipe) ToPipeOutputWithContext(ctx context.Context) PipeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipeOutput)
}

func (i *Pipe) ToPipePtrOutput() PipePtrOutput {
	return i.ToPipePtrOutputWithContext(context.Background())
}

func (i *Pipe) ToPipePtrOutputWithContext(ctx context.Context) PipePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipePtrOutput)
}

type PipePtrInput interface {
	pulumi.Input

	ToPipePtrOutput() PipePtrOutput
	ToPipePtrOutputWithContext(ctx context.Context) PipePtrOutput
}

type pipePtrType PipeArgs

func (*pipePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipe)(nil))
}

func (i *pipePtrType) ToPipePtrOutput() PipePtrOutput {
	return i.ToPipePtrOutputWithContext(context.Background())
}

func (i *pipePtrType) ToPipePtrOutputWithContext(ctx context.Context) PipePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipePtrOutput)
}

// PipeArrayInput is an input type that accepts PipeArray and PipeArrayOutput values.
// You can construct a concrete instance of `PipeArrayInput` via:
//
//          PipeArray{ PipeArgs{...} }
type PipeArrayInput interface {
	pulumi.Input

	ToPipeArrayOutput() PipeArrayOutput
	ToPipeArrayOutputWithContext(context.Context) PipeArrayOutput
}

type PipeArray []PipeInput

func (PipeArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Pipe)(nil))
}

func (i PipeArray) ToPipeArrayOutput() PipeArrayOutput {
	return i.ToPipeArrayOutputWithContext(context.Background())
}

func (i PipeArray) ToPipeArrayOutputWithContext(ctx context.Context) PipeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipeArrayOutput)
}

// PipeMapInput is an input type that accepts PipeMap and PipeMapOutput values.
// You can construct a concrete instance of `PipeMapInput` via:
//
//          PipeMap{ "key": PipeArgs{...} }
type PipeMapInput interface {
	pulumi.Input

	ToPipeMapOutput() PipeMapOutput
	ToPipeMapOutputWithContext(context.Context) PipeMapOutput
}

type PipeMap map[string]PipeInput

func (PipeMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Pipe)(nil))
}

func (i PipeMap) ToPipeMapOutput() PipeMapOutput {
	return i.ToPipeMapOutputWithContext(context.Background())
}

func (i PipeMap) ToPipeMapOutputWithContext(ctx context.Context) PipeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipeMapOutput)
}

type PipeOutput struct {
	*pulumi.OutputState
}

func (PipeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pipe)(nil))
}

func (o PipeOutput) ToPipeOutput() PipeOutput {
	return o
}

func (o PipeOutput) ToPipeOutputWithContext(ctx context.Context) PipeOutput {
	return o
}

func (o PipeOutput) ToPipePtrOutput() PipePtrOutput {
	return o.ToPipePtrOutputWithContext(context.Background())
}

func (o PipeOutput) ToPipePtrOutputWithContext(ctx context.Context) PipePtrOutput {
	return o.ApplyT(func(v Pipe) *Pipe {
		return &v
	}).(PipePtrOutput)
}

type PipePtrOutput struct {
	*pulumi.OutputState
}

func (PipePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipe)(nil))
}

func (o PipePtrOutput) ToPipePtrOutput() PipePtrOutput {
	return o
}

func (o PipePtrOutput) ToPipePtrOutputWithContext(ctx context.Context) PipePtrOutput {
	return o
}

type PipeArrayOutput struct{ *pulumi.OutputState }

func (PipeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Pipe)(nil))
}

func (o PipeArrayOutput) ToPipeArrayOutput() PipeArrayOutput {
	return o
}

func (o PipeArrayOutput) ToPipeArrayOutputWithContext(ctx context.Context) PipeArrayOutput {
	return o
}

func (o PipeArrayOutput) Index(i pulumi.IntInput) PipeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Pipe {
		return vs[0].([]Pipe)[vs[1].(int)]
	}).(PipeOutput)
}

type PipeMapOutput struct{ *pulumi.OutputState }

func (PipeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Pipe)(nil))
}

func (o PipeMapOutput) ToPipeMapOutput() PipeMapOutput {
	return o
}

func (o PipeMapOutput) ToPipeMapOutputWithContext(ctx context.Context) PipeMapOutput {
	return o
}

func (o PipeMapOutput) MapIndex(k pulumi.StringInput) PipeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Pipe {
		return vs[0].(map[string]Pipe)[vs[1].(string)]
	}).(PipeOutput)
}

func init() {
	pulumi.RegisterOutputType(PipeOutput{})
	pulumi.RegisterOutputType(PipePtrOutput{})
	pulumi.RegisterOutputType(PipeArrayOutput{})
	pulumi.RegisterOutputType(PipeMapOutput{})
}
