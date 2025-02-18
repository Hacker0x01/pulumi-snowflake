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
// 		_, err := snowflake.NewTask(ctx, "task", &snowflake.TaskArgs{
// 			Comment:      pulumi.String("my task"),
// 			Database:     pulumi.String("db"),
// 			Schema:       pulumi.String("schema"),
// 			Warehouse:    pulumi.String("warehouse"),
// 			Schedule:     pulumi.String("10 MINUTE"),
// 			SqlStatement: pulumi.String("select * from foo;"),
// 			SessionParameters: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			UserTaskTimeoutMs: pulumi.Int(10000),
// 			After:             pulumi.String("preceding_task"),
// 			When:              pulumi.String("foo AND bar"),
// 			Enabled:           pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = snowflake.NewTask(ctx, "serverlessTask", &snowflake.TaskArgs{
// 			Comment:      pulumi.String("my serverless task"),
// 			Database:     pulumi.String("db"),
// 			Schema:       pulumi.String("schema"),
// 			Schedule:     pulumi.String("10 MINUTE"),
// 			SqlStatement: pulumi.String("select * from foo;"),
// 			SessionParameters: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			UserTaskTimeoutMs:                   pulumi.Int(10000),
// 			UserTaskManagedInitialWarehouseSize: pulumi.String("XSMALL"),
// 			After:                               pulumi.String("preceding_task"),
// 			When:                                pulumi.String("foo AND bar"),
// 			Enabled:                             pulumi.Bool(true),
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
// # format is database name | schema name | task name
//
// ```sh
//  $ pulumi import snowflake:index/task:Task example 'dbName|schemaName|taskName'
// ```
type Task struct {
	pulumi.CustomResourceState

	// Specifies the predecessor task in the same database and schema of the current task. When a run of the predecessor task finishes successfully, it triggers this task (after a brief lag). (Conflict with schedule)
	After pulumi.StringPtrOutput `pulumi:"after"`
	// Specifies a comment for the task.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The database in which to create the task.
	Database pulumi.StringOutput `pulumi:"database"`
	// Specifies if the task should be started (enabled) after creation or should remain suspended (default).
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies the name of the notification integration used for error notifications.
	ErrorIntegration pulumi.StringPtrOutput `pulumi:"errorIntegration"`
	// Specifies the identifier for the task; must be unique for the database and schema in which the task is created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The schedule for periodically running the task. This can be a cron or interval in minutes. (Conflict with after)
	Schedule pulumi.StringPtrOutput `pulumi:"schedule"`
	// The schema in which to create the task.
	Schema pulumi.StringOutput `pulumi:"schema"`
	// Specifies session parameters to set for the session when the task runs. A task supports all session parameters.
	SessionParameters pulumi.StringMapOutput `pulumi:"sessionParameters"`
	// Any single SQL statement, or a call to a stored procedure, executed when the task runs.
	SqlStatement pulumi.StringOutput `pulumi:"sqlStatement"`
	// Specifies the size of the compute resources to provision for the first run of the task, before a task history is available for Snowflake to determine an ideal size. Once a task has successfully completed a few runs, Snowflake ignores this parameter setting. (Conflicts with warehouse)
	UserTaskManagedInitialWarehouseSize pulumi.StringPtrOutput `pulumi:"userTaskManagedInitialWarehouseSize"`
	// Specifies the time limit on a single run of the task before it times out (in milliseconds).
	UserTaskTimeoutMs pulumi.IntPtrOutput `pulumi:"userTaskTimeoutMs"`
	// The warehouse the task will use. Omit this parameter to use Snowflake-managed compute resources for runs of this task. (Conflicts with user*task*managed*initial*warehouse_size)
	Warehouse pulumi.StringPtrOutput `pulumi:"warehouse"`
	// Specifies a Boolean SQL expression; multiple conditions joined with AND/OR are supported.
	When pulumi.StringPtrOutput `pulumi:"when"`
}

// NewTask registers a new resource with the given unique name, arguments, and options.
func NewTask(ctx *pulumi.Context,
	name string, args *TaskArgs, opts ...pulumi.ResourceOption) (*Task, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.Schema == nil {
		return nil, errors.New("invalid value for required argument 'Schema'")
	}
	if args.SqlStatement == nil {
		return nil, errors.New("invalid value for required argument 'SqlStatement'")
	}
	var resource Task
	err := ctx.RegisterResource("snowflake:index/task:Task", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTask gets an existing Task resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskState, opts ...pulumi.ResourceOption) (*Task, error) {
	var resource Task
	err := ctx.ReadResource("snowflake:index/task:Task", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Task resources.
type taskState struct {
	// Specifies the predecessor task in the same database and schema of the current task. When a run of the predecessor task finishes successfully, it triggers this task (after a brief lag). (Conflict with schedule)
	After *string `pulumi:"after"`
	// Specifies a comment for the task.
	Comment *string `pulumi:"comment"`
	// The database in which to create the task.
	Database *string `pulumi:"database"`
	// Specifies if the task should be started (enabled) after creation or should remain suspended (default).
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of the notification integration used for error notifications.
	ErrorIntegration *string `pulumi:"errorIntegration"`
	// Specifies the identifier for the task; must be unique for the database and schema in which the task is created.
	Name *string `pulumi:"name"`
	// The schedule for periodically running the task. This can be a cron or interval in minutes. (Conflict with after)
	Schedule *string `pulumi:"schedule"`
	// The schema in which to create the task.
	Schema *string `pulumi:"schema"`
	// Specifies session parameters to set for the session when the task runs. A task supports all session parameters.
	SessionParameters map[string]string `pulumi:"sessionParameters"`
	// Any single SQL statement, or a call to a stored procedure, executed when the task runs.
	SqlStatement *string `pulumi:"sqlStatement"`
	// Specifies the size of the compute resources to provision for the first run of the task, before a task history is available for Snowflake to determine an ideal size. Once a task has successfully completed a few runs, Snowflake ignores this parameter setting. (Conflicts with warehouse)
	UserTaskManagedInitialWarehouseSize *string `pulumi:"userTaskManagedInitialWarehouseSize"`
	// Specifies the time limit on a single run of the task before it times out (in milliseconds).
	UserTaskTimeoutMs *int `pulumi:"userTaskTimeoutMs"`
	// The warehouse the task will use. Omit this parameter to use Snowflake-managed compute resources for runs of this task. (Conflicts with user*task*managed*initial*warehouse_size)
	Warehouse *string `pulumi:"warehouse"`
	// Specifies a Boolean SQL expression; multiple conditions joined with AND/OR are supported.
	When *string `pulumi:"when"`
}

type TaskState struct {
	// Specifies the predecessor task in the same database and schema of the current task. When a run of the predecessor task finishes successfully, it triggers this task (after a brief lag). (Conflict with schedule)
	After pulumi.StringPtrInput
	// Specifies a comment for the task.
	Comment pulumi.StringPtrInput
	// The database in which to create the task.
	Database pulumi.StringPtrInput
	// Specifies if the task should be started (enabled) after creation or should remain suspended (default).
	Enabled pulumi.BoolPtrInput
	// Specifies the name of the notification integration used for error notifications.
	ErrorIntegration pulumi.StringPtrInput
	// Specifies the identifier for the task; must be unique for the database and schema in which the task is created.
	Name pulumi.StringPtrInput
	// The schedule for periodically running the task. This can be a cron or interval in minutes. (Conflict with after)
	Schedule pulumi.StringPtrInput
	// The schema in which to create the task.
	Schema pulumi.StringPtrInput
	// Specifies session parameters to set for the session when the task runs. A task supports all session parameters.
	SessionParameters pulumi.StringMapInput
	// Any single SQL statement, or a call to a stored procedure, executed when the task runs.
	SqlStatement pulumi.StringPtrInput
	// Specifies the size of the compute resources to provision for the first run of the task, before a task history is available for Snowflake to determine an ideal size. Once a task has successfully completed a few runs, Snowflake ignores this parameter setting. (Conflicts with warehouse)
	UserTaskManagedInitialWarehouseSize pulumi.StringPtrInput
	// Specifies the time limit on a single run of the task before it times out (in milliseconds).
	UserTaskTimeoutMs pulumi.IntPtrInput
	// The warehouse the task will use. Omit this parameter to use Snowflake-managed compute resources for runs of this task. (Conflicts with user*task*managed*initial*warehouse_size)
	Warehouse pulumi.StringPtrInput
	// Specifies a Boolean SQL expression; multiple conditions joined with AND/OR are supported.
	When pulumi.StringPtrInput
}

func (TaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*taskState)(nil)).Elem()
}

type taskArgs struct {
	// Specifies the predecessor task in the same database and schema of the current task. When a run of the predecessor task finishes successfully, it triggers this task (after a brief lag). (Conflict with schedule)
	After *string `pulumi:"after"`
	// Specifies a comment for the task.
	Comment *string `pulumi:"comment"`
	// The database in which to create the task.
	Database string `pulumi:"database"`
	// Specifies if the task should be started (enabled) after creation or should remain suspended (default).
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of the notification integration used for error notifications.
	ErrorIntegration *string `pulumi:"errorIntegration"`
	// Specifies the identifier for the task; must be unique for the database and schema in which the task is created.
	Name *string `pulumi:"name"`
	// The schedule for periodically running the task. This can be a cron or interval in minutes. (Conflict with after)
	Schedule *string `pulumi:"schedule"`
	// The schema in which to create the task.
	Schema string `pulumi:"schema"`
	// Specifies session parameters to set for the session when the task runs. A task supports all session parameters.
	SessionParameters map[string]string `pulumi:"sessionParameters"`
	// Any single SQL statement, or a call to a stored procedure, executed when the task runs.
	SqlStatement string `pulumi:"sqlStatement"`
	// Specifies the size of the compute resources to provision for the first run of the task, before a task history is available for Snowflake to determine an ideal size. Once a task has successfully completed a few runs, Snowflake ignores this parameter setting. (Conflicts with warehouse)
	UserTaskManagedInitialWarehouseSize *string `pulumi:"userTaskManagedInitialWarehouseSize"`
	// Specifies the time limit on a single run of the task before it times out (in milliseconds).
	UserTaskTimeoutMs *int `pulumi:"userTaskTimeoutMs"`
	// The warehouse the task will use. Omit this parameter to use Snowflake-managed compute resources for runs of this task. (Conflicts with user*task*managed*initial*warehouse_size)
	Warehouse *string `pulumi:"warehouse"`
	// Specifies a Boolean SQL expression; multiple conditions joined with AND/OR are supported.
	When *string `pulumi:"when"`
}

// The set of arguments for constructing a Task resource.
type TaskArgs struct {
	// Specifies the predecessor task in the same database and schema of the current task. When a run of the predecessor task finishes successfully, it triggers this task (after a brief lag). (Conflict with schedule)
	After pulumi.StringPtrInput
	// Specifies a comment for the task.
	Comment pulumi.StringPtrInput
	// The database in which to create the task.
	Database pulumi.StringInput
	// Specifies if the task should be started (enabled) after creation or should remain suspended (default).
	Enabled pulumi.BoolPtrInput
	// Specifies the name of the notification integration used for error notifications.
	ErrorIntegration pulumi.StringPtrInput
	// Specifies the identifier for the task; must be unique for the database and schema in which the task is created.
	Name pulumi.StringPtrInput
	// The schedule for periodically running the task. This can be a cron or interval in minutes. (Conflict with after)
	Schedule pulumi.StringPtrInput
	// The schema in which to create the task.
	Schema pulumi.StringInput
	// Specifies session parameters to set for the session when the task runs. A task supports all session parameters.
	SessionParameters pulumi.StringMapInput
	// Any single SQL statement, or a call to a stored procedure, executed when the task runs.
	SqlStatement pulumi.StringInput
	// Specifies the size of the compute resources to provision for the first run of the task, before a task history is available for Snowflake to determine an ideal size. Once a task has successfully completed a few runs, Snowflake ignores this parameter setting. (Conflicts with warehouse)
	UserTaskManagedInitialWarehouseSize pulumi.StringPtrInput
	// Specifies the time limit on a single run of the task before it times out (in milliseconds).
	UserTaskTimeoutMs pulumi.IntPtrInput
	// The warehouse the task will use. Omit this parameter to use Snowflake-managed compute resources for runs of this task. (Conflicts with user*task*managed*initial*warehouse_size)
	Warehouse pulumi.StringPtrInput
	// Specifies a Boolean SQL expression; multiple conditions joined with AND/OR are supported.
	When pulumi.StringPtrInput
}

func (TaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taskArgs)(nil)).Elem()
}

type TaskInput interface {
	pulumi.Input

	ToTaskOutput() TaskOutput
	ToTaskOutputWithContext(ctx context.Context) TaskOutput
}

func (*Task) ElementType() reflect.Type {
	return reflect.TypeOf((**Task)(nil)).Elem()
}

func (i *Task) ToTaskOutput() TaskOutput {
	return i.ToTaskOutputWithContext(context.Background())
}

func (i *Task) ToTaskOutputWithContext(ctx context.Context) TaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskOutput)
}

// TaskArrayInput is an input type that accepts TaskArray and TaskArrayOutput values.
// You can construct a concrete instance of `TaskArrayInput` via:
//
//          TaskArray{ TaskArgs{...} }
type TaskArrayInput interface {
	pulumi.Input

	ToTaskArrayOutput() TaskArrayOutput
	ToTaskArrayOutputWithContext(context.Context) TaskArrayOutput
}

type TaskArray []TaskInput

func (TaskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Task)(nil)).Elem()
}

func (i TaskArray) ToTaskArrayOutput() TaskArrayOutput {
	return i.ToTaskArrayOutputWithContext(context.Background())
}

func (i TaskArray) ToTaskArrayOutputWithContext(ctx context.Context) TaskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskArrayOutput)
}

// TaskMapInput is an input type that accepts TaskMap and TaskMapOutput values.
// You can construct a concrete instance of `TaskMapInput` via:
//
//          TaskMap{ "key": TaskArgs{...} }
type TaskMapInput interface {
	pulumi.Input

	ToTaskMapOutput() TaskMapOutput
	ToTaskMapOutputWithContext(context.Context) TaskMapOutput
}

type TaskMap map[string]TaskInput

func (TaskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Task)(nil)).Elem()
}

func (i TaskMap) ToTaskMapOutput() TaskMapOutput {
	return i.ToTaskMapOutputWithContext(context.Background())
}

func (i TaskMap) ToTaskMapOutputWithContext(ctx context.Context) TaskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskMapOutput)
}

type TaskOutput struct{ *pulumi.OutputState }

func (TaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Task)(nil)).Elem()
}

func (o TaskOutput) ToTaskOutput() TaskOutput {
	return o
}

func (o TaskOutput) ToTaskOutputWithContext(ctx context.Context) TaskOutput {
	return o
}

type TaskArrayOutput struct{ *pulumi.OutputState }

func (TaskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Task)(nil)).Elem()
}

func (o TaskArrayOutput) ToTaskArrayOutput() TaskArrayOutput {
	return o
}

func (o TaskArrayOutput) ToTaskArrayOutputWithContext(ctx context.Context) TaskArrayOutput {
	return o
}

func (o TaskArrayOutput) Index(i pulumi.IntInput) TaskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Task {
		return vs[0].([]*Task)[vs[1].(int)]
	}).(TaskOutput)
}

type TaskMapOutput struct{ *pulumi.OutputState }

func (TaskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Task)(nil)).Elem()
}

func (o TaskMapOutput) ToTaskMapOutput() TaskMapOutput {
	return o
}

func (o TaskMapOutput) ToTaskMapOutputWithContext(ctx context.Context) TaskMapOutput {
	return o
}

func (o TaskMapOutput) MapIndex(k pulumi.StringInput) TaskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Task {
		return vs[0].(map[string]*Task)[vs[1].(string)]
	}).(TaskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaskInput)(nil)).Elem(), &Task{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskArrayInput)(nil)).Elem(), TaskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskMapInput)(nil)).Elem(), TaskMap{})
	pulumi.RegisterOutputType(TaskOutput{})
	pulumi.RegisterOutputType(TaskArrayOutput{})
	pulumi.RegisterOutputType(TaskMapOutput{})
}
