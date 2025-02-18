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
// 		_, err := snowflake.NewUser(ctx, "user", &snowflake.UserArgs{
// 			Comment:            pulumi.String("A user of snowflake."),
// 			DefaultRole:        pulumi.String("role1"),
// 			DefaultWarehouse:   pulumi.String("warehouse"),
// 			Disabled:           pulumi.Bool(false),
// 			DisplayName:        pulumi.String("Snowflake User"),
// 			Email:              pulumi.String("user@snowflake.example"),
// 			FirstName:          pulumi.String("Snowflake"),
// 			LastName:           pulumi.String("User"),
// 			LoginName:          pulumi.String("snowflake_user"),
// 			MustChangePassword: pulumi.Bool(false),
// 			Password:           pulumi.String("secret"),
// 			RsaPublicKey:       pulumi.String("..."),
// 			RsaPublicKey2:      pulumi.String("..."),
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
//  $ pulumi import snowflake:index/user:User example userName
// ```
type User struct {
	pulumi.CustomResourceState

	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.
	DefaultNamespace pulumi.StringPtrOutput `pulumi:"defaultNamespace"`
	// Specifies the role that is active by default for the user’s session upon login.
	DefaultRole pulumi.StringOutput `pulumi:"defaultRole"`
	// Specifies the virtual warehouse that is active by default for the user’s session upon login.
	DefaultWarehouse pulumi.StringPtrOutput `pulumi:"defaultWarehouse"`
	Disabled         pulumi.BoolOutput      `pulumi:"disabled"`
	// Name displayed for the user in the Snowflake web interface.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Email address for the user.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// First name of the user.
	FirstName pulumi.StringPtrOutput `pulumi:"firstName"`
	// Will be true if user as an RSA key set.
	HasRsaPublicKey pulumi.BoolOutput `pulumi:"hasRsaPublicKey"`
	// Last name of the user.
	LastName pulumi.StringPtrOutput `pulumi:"lastName"`
	// The name users use to log in. If not supplied, snowflake will use name instead.
	LoginName pulumi.StringOutput `pulumi:"loginName"`
	// Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.
	MustChangePassword pulumi.BoolPtrOutput `pulumi:"mustChangePassword"`
	// Name of the user. Note that if you do not supply login*name this will be used as login*name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)
	Name pulumi.StringOutput `pulumi:"name"`
	// **WARNING:** this will put the password in the terraform state file. Use carefully.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
	RsaPublicKey pulumi.StringPtrOutput `pulumi:"rsaPublicKey"`
	// Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
	RsaPublicKey2 pulumi.StringPtrOutput `pulumi:"rsaPublicKey2"`
	// Definitions of a tag to associate with the resource.
	Tags UserTagArrayOutput `pulumi:"tags"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		args = &UserArgs{}
	}

	var resource User
	err := ctx.RegisterResource("snowflake:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("snowflake:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	Comment *string `pulumi:"comment"`
	// Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.
	DefaultNamespace *string `pulumi:"defaultNamespace"`
	// Specifies the role that is active by default for the user’s session upon login.
	DefaultRole *string `pulumi:"defaultRole"`
	// Specifies the virtual warehouse that is active by default for the user’s session upon login.
	DefaultWarehouse *string `pulumi:"defaultWarehouse"`
	Disabled         *bool   `pulumi:"disabled"`
	// Name displayed for the user in the Snowflake web interface.
	DisplayName *string `pulumi:"displayName"`
	// Email address for the user.
	Email *string `pulumi:"email"`
	// First name of the user.
	FirstName *string `pulumi:"firstName"`
	// Will be true if user as an RSA key set.
	HasRsaPublicKey *bool `pulumi:"hasRsaPublicKey"`
	// Last name of the user.
	LastName *string `pulumi:"lastName"`
	// The name users use to log in. If not supplied, snowflake will use name instead.
	LoginName *string `pulumi:"loginName"`
	// Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.
	MustChangePassword *bool `pulumi:"mustChangePassword"`
	// Name of the user. Note that if you do not supply login*name this will be used as login*name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)
	Name *string `pulumi:"name"`
	// **WARNING:** this will put the password in the terraform state file. Use carefully.
	Password *string `pulumi:"password"`
	// Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
	RsaPublicKey *string `pulumi:"rsaPublicKey"`
	// Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
	RsaPublicKey2 *string `pulumi:"rsaPublicKey2"`
	// Definitions of a tag to associate with the resource.
	Tags []UserTag `pulumi:"tags"`
}

type UserState struct {
	Comment pulumi.StringPtrInput
	// Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.
	DefaultNamespace pulumi.StringPtrInput
	// Specifies the role that is active by default for the user’s session upon login.
	DefaultRole pulumi.StringPtrInput
	// Specifies the virtual warehouse that is active by default for the user’s session upon login.
	DefaultWarehouse pulumi.StringPtrInput
	Disabled         pulumi.BoolPtrInput
	// Name displayed for the user in the Snowflake web interface.
	DisplayName pulumi.StringPtrInput
	// Email address for the user.
	Email pulumi.StringPtrInput
	// First name of the user.
	FirstName pulumi.StringPtrInput
	// Will be true if user as an RSA key set.
	HasRsaPublicKey pulumi.BoolPtrInput
	// Last name of the user.
	LastName pulumi.StringPtrInput
	// The name users use to log in. If not supplied, snowflake will use name instead.
	LoginName pulumi.StringPtrInput
	// Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.
	MustChangePassword pulumi.BoolPtrInput
	// Name of the user. Note that if you do not supply login*name this will be used as login*name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)
	Name pulumi.StringPtrInput
	// **WARNING:** this will put the password in the terraform state file. Use carefully.
	Password pulumi.StringPtrInput
	// Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
	RsaPublicKey pulumi.StringPtrInput
	// Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
	RsaPublicKey2 pulumi.StringPtrInput
	// Definitions of a tag to associate with the resource.
	Tags UserTagArrayInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	Comment *string `pulumi:"comment"`
	// Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.
	DefaultNamespace *string `pulumi:"defaultNamespace"`
	// Specifies the role that is active by default for the user’s session upon login.
	DefaultRole *string `pulumi:"defaultRole"`
	// Specifies the virtual warehouse that is active by default for the user’s session upon login.
	DefaultWarehouse *string `pulumi:"defaultWarehouse"`
	Disabled         *bool   `pulumi:"disabled"`
	// Name displayed for the user in the Snowflake web interface.
	DisplayName *string `pulumi:"displayName"`
	// Email address for the user.
	Email *string `pulumi:"email"`
	// First name of the user.
	FirstName *string `pulumi:"firstName"`
	// Last name of the user.
	LastName *string `pulumi:"lastName"`
	// The name users use to log in. If not supplied, snowflake will use name instead.
	LoginName *string `pulumi:"loginName"`
	// Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.
	MustChangePassword *bool `pulumi:"mustChangePassword"`
	// Name of the user. Note that if you do not supply login*name this will be used as login*name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)
	Name *string `pulumi:"name"`
	// **WARNING:** this will put the password in the terraform state file. Use carefully.
	Password *string `pulumi:"password"`
	// Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
	RsaPublicKey *string `pulumi:"rsaPublicKey"`
	// Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
	RsaPublicKey2 *string `pulumi:"rsaPublicKey2"`
	// Definitions of a tag to associate with the resource.
	Tags []UserTag `pulumi:"tags"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	Comment pulumi.StringPtrInput
	// Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.
	DefaultNamespace pulumi.StringPtrInput
	// Specifies the role that is active by default for the user’s session upon login.
	DefaultRole pulumi.StringPtrInput
	// Specifies the virtual warehouse that is active by default for the user’s session upon login.
	DefaultWarehouse pulumi.StringPtrInput
	Disabled         pulumi.BoolPtrInput
	// Name displayed for the user in the Snowflake web interface.
	DisplayName pulumi.StringPtrInput
	// Email address for the user.
	Email pulumi.StringPtrInput
	// First name of the user.
	FirstName pulumi.StringPtrInput
	// Last name of the user.
	LastName pulumi.StringPtrInput
	// The name users use to log in. If not supplied, snowflake will use name instead.
	LoginName pulumi.StringPtrInput
	// Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.
	MustChangePassword pulumi.BoolPtrInput
	// Name of the user. Note that if you do not supply login*name this will be used as login*name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters)
	Name pulumi.StringPtrInput
	// **WARNING:** this will put the password in the terraform state file. Use carefully.
	Password pulumi.StringPtrInput
	// Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
	RsaPublicKey pulumi.StringPtrInput
	// Specifies the user’s second RSA public key; used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
	RsaPublicKey2 pulumi.StringPtrInput
	// Definitions of a tag to associate with the resource.
	Tags UserTagArrayInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//          UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//          UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
