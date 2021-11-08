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
// ```sh
//  $ pulumi import snowflake:index/managedAccount:ManagedAccount example name
// ```
type ManagedAccount struct {
	pulumi.CustomResourceState

	// Identifier, as well as login name, for the initial user in the managed account. This user serves as the account administrator for the account.
	AdminName pulumi.StringOutput `pulumi:"adminName"`
	// Password for the initial user in the managed account.
	AdminPassword pulumi.StringOutput `pulumi:"adminPassword"`
	// Cloud in which the managed account is located.
	Cloud pulumi.StringOutput `pulumi:"cloud"`
	// Specifies a comment for the managed account.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Date and time when the managed account was created.
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// Display name of the managed account.
	Locator pulumi.StringOutput `pulumi:"locator"`
	// Identifier for the managed account; must be unique for your account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Snowflake Region in which the managed account is located.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the type of managed account.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// URL for accessing the managed account, particularly through the web interface.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewManagedAccount registers a new resource with the given unique name, arguments, and options.
func NewManagedAccount(ctx *pulumi.Context,
	name string, args *ManagedAccountArgs, opts ...pulumi.ResourceOption) (*ManagedAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminName == nil {
		return nil, errors.New("invalid value for required argument 'AdminName'")
	}
	if args.AdminPassword == nil {
		return nil, errors.New("invalid value for required argument 'AdminPassword'")
	}
	var resource ManagedAccount
	err := ctx.RegisterResource("snowflake:index/managedAccount:ManagedAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedAccount gets an existing ManagedAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedAccountState, opts ...pulumi.ResourceOption) (*ManagedAccount, error) {
	var resource ManagedAccount
	err := ctx.ReadResource("snowflake:index/managedAccount:ManagedAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedAccount resources.
type managedAccountState struct {
	// Identifier, as well as login name, for the initial user in the managed account. This user serves as the account administrator for the account.
	AdminName *string `pulumi:"adminName"`
	// Password for the initial user in the managed account.
	AdminPassword *string `pulumi:"adminPassword"`
	// Cloud in which the managed account is located.
	Cloud *string `pulumi:"cloud"`
	// Specifies a comment for the managed account.
	Comment *string `pulumi:"comment"`
	// Date and time when the managed account was created.
	CreatedOn *string `pulumi:"createdOn"`
	// Display name of the managed account.
	Locator *string `pulumi:"locator"`
	// Identifier for the managed account; must be unique for your account.
	Name *string `pulumi:"name"`
	// Snowflake Region in which the managed account is located.
	Region *string `pulumi:"region"`
	// Specifies the type of managed account.
	Type *string `pulumi:"type"`
	// URL for accessing the managed account, particularly through the web interface.
	Url *string `pulumi:"url"`
}

type ManagedAccountState struct {
	// Identifier, as well as login name, for the initial user in the managed account. This user serves as the account administrator for the account.
	AdminName pulumi.StringPtrInput
	// Password for the initial user in the managed account.
	AdminPassword pulumi.StringPtrInput
	// Cloud in which the managed account is located.
	Cloud pulumi.StringPtrInput
	// Specifies a comment for the managed account.
	Comment pulumi.StringPtrInput
	// Date and time when the managed account was created.
	CreatedOn pulumi.StringPtrInput
	// Display name of the managed account.
	Locator pulumi.StringPtrInput
	// Identifier for the managed account; must be unique for your account.
	Name pulumi.StringPtrInput
	// Snowflake Region in which the managed account is located.
	Region pulumi.StringPtrInput
	// Specifies the type of managed account.
	Type pulumi.StringPtrInput
	// URL for accessing the managed account, particularly through the web interface.
	Url pulumi.StringPtrInput
}

func (ManagedAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedAccountState)(nil)).Elem()
}

type managedAccountArgs struct {
	// Identifier, as well as login name, for the initial user in the managed account. This user serves as the account administrator for the account.
	AdminName string `pulumi:"adminName"`
	// Password for the initial user in the managed account.
	AdminPassword string `pulumi:"adminPassword"`
	// Specifies a comment for the managed account.
	Comment *string `pulumi:"comment"`
	// Identifier for the managed account; must be unique for your account.
	Name *string `pulumi:"name"`
	// Specifies the type of managed account.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a ManagedAccount resource.
type ManagedAccountArgs struct {
	// Identifier, as well as login name, for the initial user in the managed account. This user serves as the account administrator for the account.
	AdminName pulumi.StringInput
	// Password for the initial user in the managed account.
	AdminPassword pulumi.StringInput
	// Specifies a comment for the managed account.
	Comment pulumi.StringPtrInput
	// Identifier for the managed account; must be unique for your account.
	Name pulumi.StringPtrInput
	// Specifies the type of managed account.
	Type pulumi.StringPtrInput
}

func (ManagedAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedAccountArgs)(nil)).Elem()
}

type ManagedAccountInput interface {
	pulumi.Input

	ToManagedAccountOutput() ManagedAccountOutput
	ToManagedAccountOutputWithContext(ctx context.Context) ManagedAccountOutput
}

func (*ManagedAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedAccount)(nil))
}

func (i *ManagedAccount) ToManagedAccountOutput() ManagedAccountOutput {
	return i.ToManagedAccountOutputWithContext(context.Background())
}

func (i *ManagedAccount) ToManagedAccountOutputWithContext(ctx context.Context) ManagedAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedAccountOutput)
}

func (i *ManagedAccount) ToManagedAccountPtrOutput() ManagedAccountPtrOutput {
	return i.ToManagedAccountPtrOutputWithContext(context.Background())
}

func (i *ManagedAccount) ToManagedAccountPtrOutputWithContext(ctx context.Context) ManagedAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedAccountPtrOutput)
}

type ManagedAccountPtrInput interface {
	pulumi.Input

	ToManagedAccountPtrOutput() ManagedAccountPtrOutput
	ToManagedAccountPtrOutputWithContext(ctx context.Context) ManagedAccountPtrOutput
}

type managedAccountPtrType ManagedAccountArgs

func (*managedAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedAccount)(nil))
}

func (i *managedAccountPtrType) ToManagedAccountPtrOutput() ManagedAccountPtrOutput {
	return i.ToManagedAccountPtrOutputWithContext(context.Background())
}

func (i *managedAccountPtrType) ToManagedAccountPtrOutputWithContext(ctx context.Context) ManagedAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedAccountPtrOutput)
}

// ManagedAccountArrayInput is an input type that accepts ManagedAccountArray and ManagedAccountArrayOutput values.
// You can construct a concrete instance of `ManagedAccountArrayInput` via:
//
//          ManagedAccountArray{ ManagedAccountArgs{...} }
type ManagedAccountArrayInput interface {
	pulumi.Input

	ToManagedAccountArrayOutput() ManagedAccountArrayOutput
	ToManagedAccountArrayOutputWithContext(context.Context) ManagedAccountArrayOutput
}

type ManagedAccountArray []ManagedAccountInput

func (ManagedAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedAccount)(nil)).Elem()
}

func (i ManagedAccountArray) ToManagedAccountArrayOutput() ManagedAccountArrayOutput {
	return i.ToManagedAccountArrayOutputWithContext(context.Background())
}

func (i ManagedAccountArray) ToManagedAccountArrayOutputWithContext(ctx context.Context) ManagedAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedAccountArrayOutput)
}

// ManagedAccountMapInput is an input type that accepts ManagedAccountMap and ManagedAccountMapOutput values.
// You can construct a concrete instance of `ManagedAccountMapInput` via:
//
//          ManagedAccountMap{ "key": ManagedAccountArgs{...} }
type ManagedAccountMapInput interface {
	pulumi.Input

	ToManagedAccountMapOutput() ManagedAccountMapOutput
	ToManagedAccountMapOutputWithContext(context.Context) ManagedAccountMapOutput
}

type ManagedAccountMap map[string]ManagedAccountInput

func (ManagedAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedAccount)(nil)).Elem()
}

func (i ManagedAccountMap) ToManagedAccountMapOutput() ManagedAccountMapOutput {
	return i.ToManagedAccountMapOutputWithContext(context.Background())
}

func (i ManagedAccountMap) ToManagedAccountMapOutputWithContext(ctx context.Context) ManagedAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedAccountMapOutput)
}

type ManagedAccountOutput struct{ *pulumi.OutputState }

func (ManagedAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedAccount)(nil))
}

func (o ManagedAccountOutput) ToManagedAccountOutput() ManagedAccountOutput {
	return o
}

func (o ManagedAccountOutput) ToManagedAccountOutputWithContext(ctx context.Context) ManagedAccountOutput {
	return o
}

func (o ManagedAccountOutput) ToManagedAccountPtrOutput() ManagedAccountPtrOutput {
	return o.ToManagedAccountPtrOutputWithContext(context.Background())
}

func (o ManagedAccountOutput) ToManagedAccountPtrOutputWithContext(ctx context.Context) ManagedAccountPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedAccount) *ManagedAccount {
		return &v
	}).(ManagedAccountPtrOutput)
}

type ManagedAccountPtrOutput struct{ *pulumi.OutputState }

func (ManagedAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedAccount)(nil))
}

func (o ManagedAccountPtrOutput) ToManagedAccountPtrOutput() ManagedAccountPtrOutput {
	return o
}

func (o ManagedAccountPtrOutput) ToManagedAccountPtrOutputWithContext(ctx context.Context) ManagedAccountPtrOutput {
	return o
}

func (o ManagedAccountPtrOutput) Elem() ManagedAccountOutput {
	return o.ApplyT(func(v *ManagedAccount) ManagedAccount {
		if v != nil {
			return *v
		}
		var ret ManagedAccount
		return ret
	}).(ManagedAccountOutput)
}

type ManagedAccountArrayOutput struct{ *pulumi.OutputState }

func (ManagedAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedAccount)(nil))
}

func (o ManagedAccountArrayOutput) ToManagedAccountArrayOutput() ManagedAccountArrayOutput {
	return o
}

func (o ManagedAccountArrayOutput) ToManagedAccountArrayOutputWithContext(ctx context.Context) ManagedAccountArrayOutput {
	return o
}

func (o ManagedAccountArrayOutput) Index(i pulumi.IntInput) ManagedAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedAccount {
		return vs[0].([]ManagedAccount)[vs[1].(int)]
	}).(ManagedAccountOutput)
}

type ManagedAccountMapOutput struct{ *pulumi.OutputState }

func (ManagedAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedAccount)(nil))
}

func (o ManagedAccountMapOutput) ToManagedAccountMapOutput() ManagedAccountMapOutput {
	return o
}

func (o ManagedAccountMapOutput) ToManagedAccountMapOutputWithContext(ctx context.Context) ManagedAccountMapOutput {
	return o
}

func (o ManagedAccountMapOutput) MapIndex(k pulumi.StringInput) ManagedAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedAccount {
		return vs[0].(map[string]ManagedAccount)[vs[1].(string)]
	}).(ManagedAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedAccountInput)(nil)).Elem(), &ManagedAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedAccountPtrInput)(nil)).Elem(), &ManagedAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedAccountArrayInput)(nil)).Elem(), ManagedAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedAccountMapInput)(nil)).Elem(), ManagedAccountMap{})
	pulumi.RegisterOutputType(ManagedAccountOutput{})
	pulumi.RegisterOutputType(ManagedAccountPtrOutput{})
	pulumi.RegisterOutputType(ManagedAccountArrayOutput{})
	pulumi.RegisterOutputType(ManagedAccountMapOutput{})
}
