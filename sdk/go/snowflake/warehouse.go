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
// 		_, err := snowflake.NewWarehouse(ctx, "warehouse", &snowflake.WarehouseArgs{
// 			Comment:       pulumi.String("foo"),
// 			WarehouseSize: pulumi.String("small"),
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
//  $ pulumi import snowflake:index/warehouse:Warehouse example warehouseName
// ```
type Warehouse struct {
	pulumi.CustomResourceState

	// Specifies whether to automatically resume a warehouse when a SQL statement (e.g. query) is submitted to it.
	AutoResume pulumi.BoolOutput `pulumi:"autoResume"`
	// Specifies the number of seconds of inactivity after which a warehouse is automatically suspended.
	AutoSuspend pulumi.IntOutput       `pulumi:"autoSuspend"`
	Comment     pulumi.StringPtrOutput `pulumi:"comment"`
	// Specifies whether the warehouse is created initially in the ‘Suspended’ state.
	InitiallySuspended pulumi.BoolPtrOutput `pulumi:"initiallySuspended"`
	// Specifies the maximum number of server clusters for the warehouse.
	MaxClusterCount pulumi.IntOutput `pulumi:"maxClusterCount"`
	// Object parameter that specifies the concurrency level for SQL statements (i.e. queries and DML) executed by a warehouse.
	MaxConcurrencyLevel pulumi.IntPtrOutput `pulumi:"maxConcurrencyLevel"`
	// Specifies the minimum number of server clusters for the warehouse (only applies to multi-cluster warehouses).
	MinClusterCount pulumi.IntOutput `pulumi:"minClusterCount"`
	// Identifier for the virtual warehouse; must be unique for your account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of a resource monitor that is explicitly assigned to the warehouse.
	ResourceMonitor pulumi.StringOutput `pulumi:"resourceMonitor"`
	// Specifies the policy for automatically starting and shutting down clusters in a multi-cluster warehouse running in Auto-scale mode.
	ScalingPolicy pulumi.StringOutput `pulumi:"scalingPolicy"`
	// Object parameter that specifies the time, in seconds, a SQL statement (query, DDL, DML, etc.) can be queued on a warehouse before it is canceled by the system.
	StatementQueuedTimeoutInSeconds pulumi.IntPtrOutput `pulumi:"statementQueuedTimeoutInSeconds"`
	// Specifies the time, in seconds, after which a running SQL statement (query, DDL, DML, etc.) is canceled by the system
	StatementTimeoutInSeconds pulumi.IntPtrOutput `pulumi:"statementTimeoutInSeconds"`
	// Specifies whether the warehouse, after being resized, waits for all the servers to provision before executing any queued or new queries.
	WaitForProvisioning pulumi.BoolPtrOutput `pulumi:"waitForProvisioning"`
	// Specifies the size of the virtual warehouse. Larger warehouse sizes 5X-Large and 6X-Large are currently in preview and only available on Amazon Web Services (AWS).
	WarehouseSize pulumi.StringOutput `pulumi:"warehouseSize"`
}

// NewWarehouse registers a new resource with the given unique name, arguments, and options.
func NewWarehouse(ctx *pulumi.Context,
	name string, args *WarehouseArgs, opts ...pulumi.ResourceOption) (*Warehouse, error) {
	if args == nil {
		args = &WarehouseArgs{}
	}

	var resource Warehouse
	err := ctx.RegisterResource("snowflake:index/warehouse:Warehouse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWarehouse gets an existing Warehouse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWarehouse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WarehouseState, opts ...pulumi.ResourceOption) (*Warehouse, error) {
	var resource Warehouse
	err := ctx.ReadResource("snowflake:index/warehouse:Warehouse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Warehouse resources.
type warehouseState struct {
	// Specifies whether to automatically resume a warehouse when a SQL statement (e.g. query) is submitted to it.
	AutoResume *bool `pulumi:"autoResume"`
	// Specifies the number of seconds of inactivity after which a warehouse is automatically suspended.
	AutoSuspend *int    `pulumi:"autoSuspend"`
	Comment     *string `pulumi:"comment"`
	// Specifies whether the warehouse is created initially in the ‘Suspended’ state.
	InitiallySuspended *bool `pulumi:"initiallySuspended"`
	// Specifies the maximum number of server clusters for the warehouse.
	MaxClusterCount *int `pulumi:"maxClusterCount"`
	// Object parameter that specifies the concurrency level for SQL statements (i.e. queries and DML) executed by a warehouse.
	MaxConcurrencyLevel *int `pulumi:"maxConcurrencyLevel"`
	// Specifies the minimum number of server clusters for the warehouse (only applies to multi-cluster warehouses).
	MinClusterCount *int `pulumi:"minClusterCount"`
	// Identifier for the virtual warehouse; must be unique for your account.
	Name *string `pulumi:"name"`
	// Specifies the name of a resource monitor that is explicitly assigned to the warehouse.
	ResourceMonitor *string `pulumi:"resourceMonitor"`
	// Specifies the policy for automatically starting and shutting down clusters in a multi-cluster warehouse running in Auto-scale mode.
	ScalingPolicy *string `pulumi:"scalingPolicy"`
	// Object parameter that specifies the time, in seconds, a SQL statement (query, DDL, DML, etc.) can be queued on a warehouse before it is canceled by the system.
	StatementQueuedTimeoutInSeconds *int `pulumi:"statementQueuedTimeoutInSeconds"`
	// Specifies the time, in seconds, after which a running SQL statement (query, DDL, DML, etc.) is canceled by the system
	StatementTimeoutInSeconds *int `pulumi:"statementTimeoutInSeconds"`
	// Specifies whether the warehouse, after being resized, waits for all the servers to provision before executing any queued or new queries.
	WaitForProvisioning *bool `pulumi:"waitForProvisioning"`
	// Specifies the size of the virtual warehouse. Larger warehouse sizes 5X-Large and 6X-Large are currently in preview and only available on Amazon Web Services (AWS).
	WarehouseSize *string `pulumi:"warehouseSize"`
}

type WarehouseState struct {
	// Specifies whether to automatically resume a warehouse when a SQL statement (e.g. query) is submitted to it.
	AutoResume pulumi.BoolPtrInput
	// Specifies the number of seconds of inactivity after which a warehouse is automatically suspended.
	AutoSuspend pulumi.IntPtrInput
	Comment     pulumi.StringPtrInput
	// Specifies whether the warehouse is created initially in the ‘Suspended’ state.
	InitiallySuspended pulumi.BoolPtrInput
	// Specifies the maximum number of server clusters for the warehouse.
	MaxClusterCount pulumi.IntPtrInput
	// Object parameter that specifies the concurrency level for SQL statements (i.e. queries and DML) executed by a warehouse.
	MaxConcurrencyLevel pulumi.IntPtrInput
	// Specifies the minimum number of server clusters for the warehouse (only applies to multi-cluster warehouses).
	MinClusterCount pulumi.IntPtrInput
	// Identifier for the virtual warehouse; must be unique for your account.
	Name pulumi.StringPtrInput
	// Specifies the name of a resource monitor that is explicitly assigned to the warehouse.
	ResourceMonitor pulumi.StringPtrInput
	// Specifies the policy for automatically starting and shutting down clusters in a multi-cluster warehouse running in Auto-scale mode.
	ScalingPolicy pulumi.StringPtrInput
	// Object parameter that specifies the time, in seconds, a SQL statement (query, DDL, DML, etc.) can be queued on a warehouse before it is canceled by the system.
	StatementQueuedTimeoutInSeconds pulumi.IntPtrInput
	// Specifies the time, in seconds, after which a running SQL statement (query, DDL, DML, etc.) is canceled by the system
	StatementTimeoutInSeconds pulumi.IntPtrInput
	// Specifies whether the warehouse, after being resized, waits for all the servers to provision before executing any queued or new queries.
	WaitForProvisioning pulumi.BoolPtrInput
	// Specifies the size of the virtual warehouse. Larger warehouse sizes 5X-Large and 6X-Large are currently in preview and only available on Amazon Web Services (AWS).
	WarehouseSize pulumi.StringPtrInput
}

func (WarehouseState) ElementType() reflect.Type {
	return reflect.TypeOf((*warehouseState)(nil)).Elem()
}

type warehouseArgs struct {
	// Specifies whether to automatically resume a warehouse when a SQL statement (e.g. query) is submitted to it.
	AutoResume *bool `pulumi:"autoResume"`
	// Specifies the number of seconds of inactivity after which a warehouse is automatically suspended.
	AutoSuspend *int    `pulumi:"autoSuspend"`
	Comment     *string `pulumi:"comment"`
	// Specifies whether the warehouse is created initially in the ‘Suspended’ state.
	InitiallySuspended *bool `pulumi:"initiallySuspended"`
	// Specifies the maximum number of server clusters for the warehouse.
	MaxClusterCount *int `pulumi:"maxClusterCount"`
	// Object parameter that specifies the concurrency level for SQL statements (i.e. queries and DML) executed by a warehouse.
	MaxConcurrencyLevel *int `pulumi:"maxConcurrencyLevel"`
	// Specifies the minimum number of server clusters for the warehouse (only applies to multi-cluster warehouses).
	MinClusterCount *int `pulumi:"minClusterCount"`
	// Identifier for the virtual warehouse; must be unique for your account.
	Name *string `pulumi:"name"`
	// Specifies the name of a resource monitor that is explicitly assigned to the warehouse.
	ResourceMonitor *string `pulumi:"resourceMonitor"`
	// Specifies the policy for automatically starting and shutting down clusters in a multi-cluster warehouse running in Auto-scale mode.
	ScalingPolicy *string `pulumi:"scalingPolicy"`
	// Object parameter that specifies the time, in seconds, a SQL statement (query, DDL, DML, etc.) can be queued on a warehouse before it is canceled by the system.
	StatementQueuedTimeoutInSeconds *int `pulumi:"statementQueuedTimeoutInSeconds"`
	// Specifies the time, in seconds, after which a running SQL statement (query, DDL, DML, etc.) is canceled by the system
	StatementTimeoutInSeconds *int `pulumi:"statementTimeoutInSeconds"`
	// Specifies whether the warehouse, after being resized, waits for all the servers to provision before executing any queued or new queries.
	WaitForProvisioning *bool `pulumi:"waitForProvisioning"`
	// Specifies the size of the virtual warehouse. Larger warehouse sizes 5X-Large and 6X-Large are currently in preview and only available on Amazon Web Services (AWS).
	WarehouseSize *string `pulumi:"warehouseSize"`
}

// The set of arguments for constructing a Warehouse resource.
type WarehouseArgs struct {
	// Specifies whether to automatically resume a warehouse when a SQL statement (e.g. query) is submitted to it.
	AutoResume pulumi.BoolPtrInput
	// Specifies the number of seconds of inactivity after which a warehouse is automatically suspended.
	AutoSuspend pulumi.IntPtrInput
	Comment     pulumi.StringPtrInput
	// Specifies whether the warehouse is created initially in the ‘Suspended’ state.
	InitiallySuspended pulumi.BoolPtrInput
	// Specifies the maximum number of server clusters for the warehouse.
	MaxClusterCount pulumi.IntPtrInput
	// Object parameter that specifies the concurrency level for SQL statements (i.e. queries and DML) executed by a warehouse.
	MaxConcurrencyLevel pulumi.IntPtrInput
	// Specifies the minimum number of server clusters for the warehouse (only applies to multi-cluster warehouses).
	MinClusterCount pulumi.IntPtrInput
	// Identifier for the virtual warehouse; must be unique for your account.
	Name pulumi.StringPtrInput
	// Specifies the name of a resource monitor that is explicitly assigned to the warehouse.
	ResourceMonitor pulumi.StringPtrInput
	// Specifies the policy for automatically starting and shutting down clusters in a multi-cluster warehouse running in Auto-scale mode.
	ScalingPolicy pulumi.StringPtrInput
	// Object parameter that specifies the time, in seconds, a SQL statement (query, DDL, DML, etc.) can be queued on a warehouse before it is canceled by the system.
	StatementQueuedTimeoutInSeconds pulumi.IntPtrInput
	// Specifies the time, in seconds, after which a running SQL statement (query, DDL, DML, etc.) is canceled by the system
	StatementTimeoutInSeconds pulumi.IntPtrInput
	// Specifies whether the warehouse, after being resized, waits for all the servers to provision before executing any queued or new queries.
	WaitForProvisioning pulumi.BoolPtrInput
	// Specifies the size of the virtual warehouse. Larger warehouse sizes 5X-Large and 6X-Large are currently in preview and only available on Amazon Web Services (AWS).
	WarehouseSize pulumi.StringPtrInput
}

func (WarehouseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*warehouseArgs)(nil)).Elem()
}

type WarehouseInput interface {
	pulumi.Input

	ToWarehouseOutput() WarehouseOutput
	ToWarehouseOutputWithContext(ctx context.Context) WarehouseOutput
}

func (*Warehouse) ElementType() reflect.Type {
	return reflect.TypeOf((*Warehouse)(nil))
}

func (i *Warehouse) ToWarehouseOutput() WarehouseOutput {
	return i.ToWarehouseOutputWithContext(context.Background())
}

func (i *Warehouse) ToWarehouseOutputWithContext(ctx context.Context) WarehouseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WarehouseOutput)
}

func (i *Warehouse) ToWarehousePtrOutput() WarehousePtrOutput {
	return i.ToWarehousePtrOutputWithContext(context.Background())
}

func (i *Warehouse) ToWarehousePtrOutputWithContext(ctx context.Context) WarehousePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WarehousePtrOutput)
}

type WarehousePtrInput interface {
	pulumi.Input

	ToWarehousePtrOutput() WarehousePtrOutput
	ToWarehousePtrOutputWithContext(ctx context.Context) WarehousePtrOutput
}

type warehousePtrType WarehouseArgs

func (*warehousePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Warehouse)(nil))
}

func (i *warehousePtrType) ToWarehousePtrOutput() WarehousePtrOutput {
	return i.ToWarehousePtrOutputWithContext(context.Background())
}

func (i *warehousePtrType) ToWarehousePtrOutputWithContext(ctx context.Context) WarehousePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WarehousePtrOutput)
}

// WarehouseArrayInput is an input type that accepts WarehouseArray and WarehouseArrayOutput values.
// You can construct a concrete instance of `WarehouseArrayInput` via:
//
//          WarehouseArray{ WarehouseArgs{...} }
type WarehouseArrayInput interface {
	pulumi.Input

	ToWarehouseArrayOutput() WarehouseArrayOutput
	ToWarehouseArrayOutputWithContext(context.Context) WarehouseArrayOutput
}

type WarehouseArray []WarehouseInput

func (WarehouseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Warehouse)(nil)).Elem()
}

func (i WarehouseArray) ToWarehouseArrayOutput() WarehouseArrayOutput {
	return i.ToWarehouseArrayOutputWithContext(context.Background())
}

func (i WarehouseArray) ToWarehouseArrayOutputWithContext(ctx context.Context) WarehouseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WarehouseArrayOutput)
}

// WarehouseMapInput is an input type that accepts WarehouseMap and WarehouseMapOutput values.
// You can construct a concrete instance of `WarehouseMapInput` via:
//
//          WarehouseMap{ "key": WarehouseArgs{...} }
type WarehouseMapInput interface {
	pulumi.Input

	ToWarehouseMapOutput() WarehouseMapOutput
	ToWarehouseMapOutputWithContext(context.Context) WarehouseMapOutput
}

type WarehouseMap map[string]WarehouseInput

func (WarehouseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Warehouse)(nil)).Elem()
}

func (i WarehouseMap) ToWarehouseMapOutput() WarehouseMapOutput {
	return i.ToWarehouseMapOutputWithContext(context.Background())
}

func (i WarehouseMap) ToWarehouseMapOutputWithContext(ctx context.Context) WarehouseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WarehouseMapOutput)
}

type WarehouseOutput struct{ *pulumi.OutputState }

func (WarehouseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Warehouse)(nil))
}

func (o WarehouseOutput) ToWarehouseOutput() WarehouseOutput {
	return o
}

func (o WarehouseOutput) ToWarehouseOutputWithContext(ctx context.Context) WarehouseOutput {
	return o
}

func (o WarehouseOutput) ToWarehousePtrOutput() WarehousePtrOutput {
	return o.ToWarehousePtrOutputWithContext(context.Background())
}

func (o WarehouseOutput) ToWarehousePtrOutputWithContext(ctx context.Context) WarehousePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Warehouse) *Warehouse {
		return &v
	}).(WarehousePtrOutput)
}

type WarehousePtrOutput struct{ *pulumi.OutputState }

func (WarehousePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Warehouse)(nil))
}

func (o WarehousePtrOutput) ToWarehousePtrOutput() WarehousePtrOutput {
	return o
}

func (o WarehousePtrOutput) ToWarehousePtrOutputWithContext(ctx context.Context) WarehousePtrOutput {
	return o
}

func (o WarehousePtrOutput) Elem() WarehouseOutput {
	return o.ApplyT(func(v *Warehouse) Warehouse {
		if v != nil {
			return *v
		}
		var ret Warehouse
		return ret
	}).(WarehouseOutput)
}

type WarehouseArrayOutput struct{ *pulumi.OutputState }

func (WarehouseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Warehouse)(nil))
}

func (o WarehouseArrayOutput) ToWarehouseArrayOutput() WarehouseArrayOutput {
	return o
}

func (o WarehouseArrayOutput) ToWarehouseArrayOutputWithContext(ctx context.Context) WarehouseArrayOutput {
	return o
}

func (o WarehouseArrayOutput) Index(i pulumi.IntInput) WarehouseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Warehouse {
		return vs[0].([]Warehouse)[vs[1].(int)]
	}).(WarehouseOutput)
}

type WarehouseMapOutput struct{ *pulumi.OutputState }

func (WarehouseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Warehouse)(nil))
}

func (o WarehouseMapOutput) ToWarehouseMapOutput() WarehouseMapOutput {
	return o
}

func (o WarehouseMapOutput) ToWarehouseMapOutputWithContext(ctx context.Context) WarehouseMapOutput {
	return o
}

func (o WarehouseMapOutput) MapIndex(k pulumi.StringInput) WarehouseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Warehouse {
		return vs[0].(map[string]Warehouse)[vs[1].(string)]
	}).(WarehouseOutput)
}

func init() {
	pulumi.RegisterOutputType(WarehouseOutput{})
	pulumi.RegisterOutputType(WarehousePtrOutput{})
	pulumi.RegisterOutputType(WarehouseArrayOutput{})
	pulumi.RegisterOutputType(WarehouseMapOutput{})
}
