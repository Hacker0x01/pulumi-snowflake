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
// 		_, err := snowflake.NewResourceMonitor(ctx, "monitor", &snowflake.ResourceMonitorArgs{
// 			CreditQuota:  pulumi.Int(100),
// 			EndTimestamp: pulumi.String("2021-12-07 00:00"),
// 			Frequency:    pulumi.String("DAILY"),
// 			NotifyTriggers: pulumi.IntArray{
// 				pulumi.Int(40),
// 			},
// 			StartTimestamp: pulumi.String("2020-12-07 00:00"),
// 			SuspendImmediateTriggers: pulumi.IntArray{
// 				pulumi.Int(90),
// 			},
// 			SuspendTriggers: pulumi.IntArray{
// 				pulumi.Int(50),
// 			},
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
//  $ pulumi import snowflake:index/resourceMonitor:ResourceMonitor example
// ```
type ResourceMonitor struct {
	pulumi.CustomResourceState

	// The number of credits allocated monthly to the resource monitor.
	CreditQuota pulumi.IntOutput `pulumi:"creditQuota"`
	// The date and time when the resource monitor suspends the assigned warehouses.
	EndTimestamp pulumi.StringPtrOutput `pulumi:"endTimestamp"`
	// The frequency interval at which the credit usage resets to 0. If you set a frequency for a resource monitor, you must also set START_TIMESTAMP.
	Frequency pulumi.StringOutput `pulumi:"frequency"`
	// Identifier for the resource monitor; must be unique for your account.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of percentage thresholds at which to send an alert to subscribed users.
	NotifyTriggers pulumi.IntArrayOutput `pulumi:"notifyTriggers"`
	// Specifies whether the resource monitor should be applied globally to your Snowflake account.
	SetForAccount pulumi.BoolPtrOutput `pulumi:"setForAccount"`
	// The date and time when the resource monitor starts monitoring credit usage for the assigned warehouses.
	StartTimestamp pulumi.StringOutput `pulumi:"startTimestamp"`
	// A list of percentage thresholds at which to immediately suspend all warehouses.
	SuspendImmediateTriggers pulumi.IntArrayOutput `pulumi:"suspendImmediateTriggers"`
	// A list of percentage thresholds at which to suspend all warehouses.
	SuspendTriggers pulumi.IntArrayOutput `pulumi:"suspendTriggers"`
}

// NewResourceMonitor registers a new resource with the given unique name, arguments, and options.
func NewResourceMonitor(ctx *pulumi.Context,
	name string, args *ResourceMonitorArgs, opts ...pulumi.ResourceOption) (*ResourceMonitor, error) {
	if args == nil {
		args = &ResourceMonitorArgs{}
	}

	var resource ResourceMonitor
	err := ctx.RegisterResource("snowflake:index/resourceMonitor:ResourceMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceMonitor gets an existing ResourceMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceMonitorState, opts ...pulumi.ResourceOption) (*ResourceMonitor, error) {
	var resource ResourceMonitor
	err := ctx.ReadResource("snowflake:index/resourceMonitor:ResourceMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceMonitor resources.
type resourceMonitorState struct {
	// The number of credits allocated monthly to the resource monitor.
	CreditQuota *int `pulumi:"creditQuota"`
	// The date and time when the resource monitor suspends the assigned warehouses.
	EndTimestamp *string `pulumi:"endTimestamp"`
	// The frequency interval at which the credit usage resets to 0. If you set a frequency for a resource monitor, you must also set START_TIMESTAMP.
	Frequency *string `pulumi:"frequency"`
	// Identifier for the resource monitor; must be unique for your account.
	Name *string `pulumi:"name"`
	// A list of percentage thresholds at which to send an alert to subscribed users.
	NotifyTriggers []int `pulumi:"notifyTriggers"`
	// Specifies whether the resource monitor should be applied globally to your Snowflake account.
	SetForAccount *bool `pulumi:"setForAccount"`
	// The date and time when the resource monitor starts monitoring credit usage for the assigned warehouses.
	StartTimestamp *string `pulumi:"startTimestamp"`
	// A list of percentage thresholds at which to immediately suspend all warehouses.
	SuspendImmediateTriggers []int `pulumi:"suspendImmediateTriggers"`
	// A list of percentage thresholds at which to suspend all warehouses.
	SuspendTriggers []int `pulumi:"suspendTriggers"`
}

type ResourceMonitorState struct {
	// The number of credits allocated monthly to the resource monitor.
	CreditQuota pulumi.IntPtrInput
	// The date and time when the resource monitor suspends the assigned warehouses.
	EndTimestamp pulumi.StringPtrInput
	// The frequency interval at which the credit usage resets to 0. If you set a frequency for a resource monitor, you must also set START_TIMESTAMP.
	Frequency pulumi.StringPtrInput
	// Identifier for the resource monitor; must be unique for your account.
	Name pulumi.StringPtrInput
	// A list of percentage thresholds at which to send an alert to subscribed users.
	NotifyTriggers pulumi.IntArrayInput
	// Specifies whether the resource monitor should be applied globally to your Snowflake account.
	SetForAccount pulumi.BoolPtrInput
	// The date and time when the resource monitor starts monitoring credit usage for the assigned warehouses.
	StartTimestamp pulumi.StringPtrInput
	// A list of percentage thresholds at which to immediately suspend all warehouses.
	SuspendImmediateTriggers pulumi.IntArrayInput
	// A list of percentage thresholds at which to suspend all warehouses.
	SuspendTriggers pulumi.IntArrayInput
}

func (ResourceMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceMonitorState)(nil)).Elem()
}

type resourceMonitorArgs struct {
	// The number of credits allocated monthly to the resource monitor.
	CreditQuota *int `pulumi:"creditQuota"`
	// The date and time when the resource monitor suspends the assigned warehouses.
	EndTimestamp *string `pulumi:"endTimestamp"`
	// The frequency interval at which the credit usage resets to 0. If you set a frequency for a resource monitor, you must also set START_TIMESTAMP.
	Frequency *string `pulumi:"frequency"`
	// Identifier for the resource monitor; must be unique for your account.
	Name *string `pulumi:"name"`
	// A list of percentage thresholds at which to send an alert to subscribed users.
	NotifyTriggers []int `pulumi:"notifyTriggers"`
	// Specifies whether the resource monitor should be applied globally to your Snowflake account.
	SetForAccount *bool `pulumi:"setForAccount"`
	// The date and time when the resource monitor starts monitoring credit usage for the assigned warehouses.
	StartTimestamp *string `pulumi:"startTimestamp"`
	// A list of percentage thresholds at which to immediately suspend all warehouses.
	SuspendImmediateTriggers []int `pulumi:"suspendImmediateTriggers"`
	// A list of percentage thresholds at which to suspend all warehouses.
	SuspendTriggers []int `pulumi:"suspendTriggers"`
}

// The set of arguments for constructing a ResourceMonitor resource.
type ResourceMonitorArgs struct {
	// The number of credits allocated monthly to the resource monitor.
	CreditQuota pulumi.IntPtrInput
	// The date and time when the resource monitor suspends the assigned warehouses.
	EndTimestamp pulumi.StringPtrInput
	// The frequency interval at which the credit usage resets to 0. If you set a frequency for a resource monitor, you must also set START_TIMESTAMP.
	Frequency pulumi.StringPtrInput
	// Identifier for the resource monitor; must be unique for your account.
	Name pulumi.StringPtrInput
	// A list of percentage thresholds at which to send an alert to subscribed users.
	NotifyTriggers pulumi.IntArrayInput
	// Specifies whether the resource monitor should be applied globally to your Snowflake account.
	SetForAccount pulumi.BoolPtrInput
	// The date and time when the resource monitor starts monitoring credit usage for the assigned warehouses.
	StartTimestamp pulumi.StringPtrInput
	// A list of percentage thresholds at which to immediately suspend all warehouses.
	SuspendImmediateTriggers pulumi.IntArrayInput
	// A list of percentage thresholds at which to suspend all warehouses.
	SuspendTriggers pulumi.IntArrayInput
}

func (ResourceMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceMonitorArgs)(nil)).Elem()
}

type ResourceMonitorInput interface {
	pulumi.Input

	ToResourceMonitorOutput() ResourceMonitorOutput
	ToResourceMonitorOutputWithContext(ctx context.Context) ResourceMonitorOutput
}

func (*ResourceMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceMonitor)(nil))
}

func (i *ResourceMonitor) ToResourceMonitorOutput() ResourceMonitorOutput {
	return i.ToResourceMonitorOutputWithContext(context.Background())
}

func (i *ResourceMonitor) ToResourceMonitorOutputWithContext(ctx context.Context) ResourceMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMonitorOutput)
}

func (i *ResourceMonitor) ToResourceMonitorPtrOutput() ResourceMonitorPtrOutput {
	return i.ToResourceMonitorPtrOutputWithContext(context.Background())
}

func (i *ResourceMonitor) ToResourceMonitorPtrOutputWithContext(ctx context.Context) ResourceMonitorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMonitorPtrOutput)
}

type ResourceMonitorPtrInput interface {
	pulumi.Input

	ToResourceMonitorPtrOutput() ResourceMonitorPtrOutput
	ToResourceMonitorPtrOutputWithContext(ctx context.Context) ResourceMonitorPtrOutput
}

type resourceMonitorPtrType ResourceMonitorArgs

func (*resourceMonitorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceMonitor)(nil))
}

func (i *resourceMonitorPtrType) ToResourceMonitorPtrOutput() ResourceMonitorPtrOutput {
	return i.ToResourceMonitorPtrOutputWithContext(context.Background())
}

func (i *resourceMonitorPtrType) ToResourceMonitorPtrOutputWithContext(ctx context.Context) ResourceMonitorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMonitorPtrOutput)
}

// ResourceMonitorArrayInput is an input type that accepts ResourceMonitorArray and ResourceMonitorArrayOutput values.
// You can construct a concrete instance of `ResourceMonitorArrayInput` via:
//
//          ResourceMonitorArray{ ResourceMonitorArgs{...} }
type ResourceMonitorArrayInput interface {
	pulumi.Input

	ToResourceMonitorArrayOutput() ResourceMonitorArrayOutput
	ToResourceMonitorArrayOutputWithContext(context.Context) ResourceMonitorArrayOutput
}

type ResourceMonitorArray []ResourceMonitorInput

func (ResourceMonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceMonitor)(nil)).Elem()
}

func (i ResourceMonitorArray) ToResourceMonitorArrayOutput() ResourceMonitorArrayOutput {
	return i.ToResourceMonitorArrayOutputWithContext(context.Background())
}

func (i ResourceMonitorArray) ToResourceMonitorArrayOutputWithContext(ctx context.Context) ResourceMonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMonitorArrayOutput)
}

// ResourceMonitorMapInput is an input type that accepts ResourceMonitorMap and ResourceMonitorMapOutput values.
// You can construct a concrete instance of `ResourceMonitorMapInput` via:
//
//          ResourceMonitorMap{ "key": ResourceMonitorArgs{...} }
type ResourceMonitorMapInput interface {
	pulumi.Input

	ToResourceMonitorMapOutput() ResourceMonitorMapOutput
	ToResourceMonitorMapOutputWithContext(context.Context) ResourceMonitorMapOutput
}

type ResourceMonitorMap map[string]ResourceMonitorInput

func (ResourceMonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceMonitor)(nil)).Elem()
}

func (i ResourceMonitorMap) ToResourceMonitorMapOutput() ResourceMonitorMapOutput {
	return i.ToResourceMonitorMapOutputWithContext(context.Background())
}

func (i ResourceMonitorMap) ToResourceMonitorMapOutputWithContext(ctx context.Context) ResourceMonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMonitorMapOutput)
}

type ResourceMonitorOutput struct{ *pulumi.OutputState }

func (ResourceMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceMonitor)(nil))
}

func (o ResourceMonitorOutput) ToResourceMonitorOutput() ResourceMonitorOutput {
	return o
}

func (o ResourceMonitorOutput) ToResourceMonitorOutputWithContext(ctx context.Context) ResourceMonitorOutput {
	return o
}

func (o ResourceMonitorOutput) ToResourceMonitorPtrOutput() ResourceMonitorPtrOutput {
	return o.ToResourceMonitorPtrOutputWithContext(context.Background())
}

func (o ResourceMonitorOutput) ToResourceMonitorPtrOutputWithContext(ctx context.Context) ResourceMonitorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceMonitor) *ResourceMonitor {
		return &v
	}).(ResourceMonitorPtrOutput)
}

type ResourceMonitorPtrOutput struct{ *pulumi.OutputState }

func (ResourceMonitorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceMonitor)(nil))
}

func (o ResourceMonitorPtrOutput) ToResourceMonitorPtrOutput() ResourceMonitorPtrOutput {
	return o
}

func (o ResourceMonitorPtrOutput) ToResourceMonitorPtrOutputWithContext(ctx context.Context) ResourceMonitorPtrOutput {
	return o
}

func (o ResourceMonitorPtrOutput) Elem() ResourceMonitorOutput {
	return o.ApplyT(func(v *ResourceMonitor) ResourceMonitor {
		if v != nil {
			return *v
		}
		var ret ResourceMonitor
		return ret
	}).(ResourceMonitorOutput)
}

type ResourceMonitorArrayOutput struct{ *pulumi.OutputState }

func (ResourceMonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceMonitor)(nil))
}

func (o ResourceMonitorArrayOutput) ToResourceMonitorArrayOutput() ResourceMonitorArrayOutput {
	return o
}

func (o ResourceMonitorArrayOutput) ToResourceMonitorArrayOutputWithContext(ctx context.Context) ResourceMonitorArrayOutput {
	return o
}

func (o ResourceMonitorArrayOutput) Index(i pulumi.IntInput) ResourceMonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceMonitor {
		return vs[0].([]ResourceMonitor)[vs[1].(int)]
	}).(ResourceMonitorOutput)
}

type ResourceMonitorMapOutput struct{ *pulumi.OutputState }

func (ResourceMonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceMonitor)(nil))
}

func (o ResourceMonitorMapOutput) ToResourceMonitorMapOutput() ResourceMonitorMapOutput {
	return o
}

func (o ResourceMonitorMapOutput) ToResourceMonitorMapOutputWithContext(ctx context.Context) ResourceMonitorMapOutput {
	return o
}

func (o ResourceMonitorMapOutput) MapIndex(k pulumi.StringInput) ResourceMonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceMonitor {
		return vs[0].(map[string]ResourceMonitor)[vs[1].(string)]
	}).(ResourceMonitorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceMonitorInput)(nil)).Elem(), &ResourceMonitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceMonitorPtrInput)(nil)).Elem(), &ResourceMonitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceMonitorArrayInput)(nil)).Elem(), ResourceMonitorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceMonitorMapInput)(nil)).Elem(), ResourceMonitorMap{})
	pulumi.RegisterOutputType(ResourceMonitorOutput{})
	pulumi.RegisterOutputType(ResourceMonitorPtrOutput{})
	pulumi.RegisterOutputType(ResourceMonitorArrayOutput{})
	pulumi.RegisterOutputType(ResourceMonitorMapOutput{})
}
