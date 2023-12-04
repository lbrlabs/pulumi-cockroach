// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cockroach

import (
	"context"
	"reflect"

	"errors"
	"github.com/lbrlabs/pulumi-cockroach/sdk/go/cockroach/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Role grants for a single user.
type UserRoleGrants struct {
	pulumi.CustomResourceState

	Roles UserRoleGrantsRoleArrayOutput `pulumi:"roles"`
	// ID of the user to grant these roles to.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewUserRoleGrants registers a new resource with the given unique name, arguments, and options.
func NewUserRoleGrants(ctx *pulumi.Context,
	name string, args *UserRoleGrantsArgs, opts ...pulumi.ResourceOption) (*UserRoleGrants, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserRoleGrants
	err := ctx.RegisterResource("cockroach:index/userRoleGrants:UserRoleGrants", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserRoleGrants gets an existing UserRoleGrants resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserRoleGrants(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserRoleGrantsState, opts ...pulumi.ResourceOption) (*UserRoleGrants, error) {
	var resource UserRoleGrants
	err := ctx.ReadResource("cockroach:index/userRoleGrants:UserRoleGrants", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserRoleGrants resources.
type userRoleGrantsState struct {
	Roles []UserRoleGrantsRole `pulumi:"roles"`
	// ID of the user to grant these roles to.
	UserId *string `pulumi:"userId"`
}

type UserRoleGrantsState struct {
	Roles UserRoleGrantsRoleArrayInput
	// ID of the user to grant these roles to.
	UserId pulumi.StringPtrInput
}

func (UserRoleGrantsState) ElementType() reflect.Type {
	return reflect.TypeOf((*userRoleGrantsState)(nil)).Elem()
}

type userRoleGrantsArgs struct {
	Roles []UserRoleGrantsRole `pulumi:"roles"`
	// ID of the user to grant these roles to.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a UserRoleGrants resource.
type UserRoleGrantsArgs struct {
	Roles UserRoleGrantsRoleArrayInput
	// ID of the user to grant these roles to.
	UserId pulumi.StringInput
}

func (UserRoleGrantsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userRoleGrantsArgs)(nil)).Elem()
}

type UserRoleGrantsInput interface {
	pulumi.Input

	ToUserRoleGrantsOutput() UserRoleGrantsOutput
	ToUserRoleGrantsOutputWithContext(ctx context.Context) UserRoleGrantsOutput
}

func (*UserRoleGrants) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRoleGrants)(nil)).Elem()
}

func (i *UserRoleGrants) ToUserRoleGrantsOutput() UserRoleGrantsOutput {
	return i.ToUserRoleGrantsOutputWithContext(context.Background())
}

func (i *UserRoleGrants) ToUserRoleGrantsOutputWithContext(ctx context.Context) UserRoleGrantsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRoleGrantsOutput)
}

// UserRoleGrantsArrayInput is an input type that accepts UserRoleGrantsArray and UserRoleGrantsArrayOutput values.
// You can construct a concrete instance of `UserRoleGrantsArrayInput` via:
//
//	UserRoleGrantsArray{ UserRoleGrantsArgs{...} }
type UserRoleGrantsArrayInput interface {
	pulumi.Input

	ToUserRoleGrantsArrayOutput() UserRoleGrantsArrayOutput
	ToUserRoleGrantsArrayOutputWithContext(context.Context) UserRoleGrantsArrayOutput
}

type UserRoleGrantsArray []UserRoleGrantsInput

func (UserRoleGrantsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserRoleGrants)(nil)).Elem()
}

func (i UserRoleGrantsArray) ToUserRoleGrantsArrayOutput() UserRoleGrantsArrayOutput {
	return i.ToUserRoleGrantsArrayOutputWithContext(context.Background())
}

func (i UserRoleGrantsArray) ToUserRoleGrantsArrayOutputWithContext(ctx context.Context) UserRoleGrantsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRoleGrantsArrayOutput)
}

// UserRoleGrantsMapInput is an input type that accepts UserRoleGrantsMap and UserRoleGrantsMapOutput values.
// You can construct a concrete instance of `UserRoleGrantsMapInput` via:
//
//	UserRoleGrantsMap{ "key": UserRoleGrantsArgs{...} }
type UserRoleGrantsMapInput interface {
	pulumi.Input

	ToUserRoleGrantsMapOutput() UserRoleGrantsMapOutput
	ToUserRoleGrantsMapOutputWithContext(context.Context) UserRoleGrantsMapOutput
}

type UserRoleGrantsMap map[string]UserRoleGrantsInput

func (UserRoleGrantsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserRoleGrants)(nil)).Elem()
}

func (i UserRoleGrantsMap) ToUserRoleGrantsMapOutput() UserRoleGrantsMapOutput {
	return i.ToUserRoleGrantsMapOutputWithContext(context.Background())
}

func (i UserRoleGrantsMap) ToUserRoleGrantsMapOutputWithContext(ctx context.Context) UserRoleGrantsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRoleGrantsMapOutput)
}

type UserRoleGrantsOutput struct{ *pulumi.OutputState }

func (UserRoleGrantsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRoleGrants)(nil)).Elem()
}

func (o UserRoleGrantsOutput) ToUserRoleGrantsOutput() UserRoleGrantsOutput {
	return o
}

func (o UserRoleGrantsOutput) ToUserRoleGrantsOutputWithContext(ctx context.Context) UserRoleGrantsOutput {
	return o
}

func (o UserRoleGrantsOutput) Roles() UserRoleGrantsRoleArrayOutput {
	return o.ApplyT(func(v *UserRoleGrants) UserRoleGrantsRoleArrayOutput { return v.Roles }).(UserRoleGrantsRoleArrayOutput)
}

// ID of the user to grant these roles to.
func (o UserRoleGrantsOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRoleGrants) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type UserRoleGrantsArrayOutput struct{ *pulumi.OutputState }

func (UserRoleGrantsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserRoleGrants)(nil)).Elem()
}

func (o UserRoleGrantsArrayOutput) ToUserRoleGrantsArrayOutput() UserRoleGrantsArrayOutput {
	return o
}

func (o UserRoleGrantsArrayOutput) ToUserRoleGrantsArrayOutputWithContext(ctx context.Context) UserRoleGrantsArrayOutput {
	return o
}

func (o UserRoleGrantsArrayOutput) Index(i pulumi.IntInput) UserRoleGrantsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserRoleGrants {
		return vs[0].([]*UserRoleGrants)[vs[1].(int)]
	}).(UserRoleGrantsOutput)
}

type UserRoleGrantsMapOutput struct{ *pulumi.OutputState }

func (UserRoleGrantsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserRoleGrants)(nil)).Elem()
}

func (o UserRoleGrantsMapOutput) ToUserRoleGrantsMapOutput() UserRoleGrantsMapOutput {
	return o
}

func (o UserRoleGrantsMapOutput) ToUserRoleGrantsMapOutputWithContext(ctx context.Context) UserRoleGrantsMapOutput {
	return o
}

func (o UserRoleGrantsMapOutput) MapIndex(k pulumi.StringInput) UserRoleGrantsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserRoleGrants {
		return vs[0].(map[string]*UserRoleGrants)[vs[1].(string)]
	}).(UserRoleGrantsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserRoleGrantsInput)(nil)).Elem(), &UserRoleGrants{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserRoleGrantsArrayInput)(nil)).Elem(), UserRoleGrantsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserRoleGrantsMapInput)(nil)).Elem(), UserRoleGrantsMap{})
	pulumi.RegisterOutputType(UserRoleGrantsOutput{})
	pulumi.RegisterOutputType(UserRoleGrantsArrayOutput{})
	pulumi.RegisterOutputType(UserRoleGrantsMapOutput{})
}
