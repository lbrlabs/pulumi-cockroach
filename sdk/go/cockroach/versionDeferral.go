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

// Configure minor version upgrade deferral for a cluster.
type VersionDeferral struct {
	pulumi.CustomResourceState

	// The policy for managing automated minor version upgrades. Set to FIXED*DEFERRAL to defer upgrades by 60 days or NOT*DEFERRED to apply upgrades immediately.
	DeferralPolicy pulumi.StringOutput `pulumi:"deferralPolicy"`
}

// NewVersionDeferral registers a new resource with the given unique name, arguments, and options.
func NewVersionDeferral(ctx *pulumi.Context,
	name string, args *VersionDeferralArgs, opts ...pulumi.ResourceOption) (*VersionDeferral, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeferralPolicy == nil {
		return nil, errors.New("invalid value for required argument 'DeferralPolicy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VersionDeferral
	err := ctx.RegisterResource("cockroach:index/versionDeferral:VersionDeferral", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVersionDeferral gets an existing VersionDeferral resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVersionDeferral(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VersionDeferralState, opts ...pulumi.ResourceOption) (*VersionDeferral, error) {
	var resource VersionDeferral
	err := ctx.ReadResource("cockroach:index/versionDeferral:VersionDeferral", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VersionDeferral resources.
type versionDeferralState struct {
	// The policy for managing automated minor version upgrades. Set to FIXED*DEFERRAL to defer upgrades by 60 days or NOT*DEFERRED to apply upgrades immediately.
	DeferralPolicy *string `pulumi:"deferralPolicy"`
}

type VersionDeferralState struct {
	// The policy for managing automated minor version upgrades. Set to FIXED*DEFERRAL to defer upgrades by 60 days or NOT*DEFERRED to apply upgrades immediately.
	DeferralPolicy pulumi.StringPtrInput
}

func (VersionDeferralState) ElementType() reflect.Type {
	return reflect.TypeOf((*versionDeferralState)(nil)).Elem()
}

type versionDeferralArgs struct {
	// The policy for managing automated minor version upgrades. Set to FIXED*DEFERRAL to defer upgrades by 60 days or NOT*DEFERRED to apply upgrades immediately.
	DeferralPolicy string `pulumi:"deferralPolicy"`
}

// The set of arguments for constructing a VersionDeferral resource.
type VersionDeferralArgs struct {
	// The policy for managing automated minor version upgrades. Set to FIXED*DEFERRAL to defer upgrades by 60 days or NOT*DEFERRED to apply upgrades immediately.
	DeferralPolicy pulumi.StringInput
}

func (VersionDeferralArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*versionDeferralArgs)(nil)).Elem()
}

type VersionDeferralInput interface {
	pulumi.Input

	ToVersionDeferralOutput() VersionDeferralOutput
	ToVersionDeferralOutputWithContext(ctx context.Context) VersionDeferralOutput
}

func (*VersionDeferral) ElementType() reflect.Type {
	return reflect.TypeOf((**VersionDeferral)(nil)).Elem()
}

func (i *VersionDeferral) ToVersionDeferralOutput() VersionDeferralOutput {
	return i.ToVersionDeferralOutputWithContext(context.Background())
}

func (i *VersionDeferral) ToVersionDeferralOutputWithContext(ctx context.Context) VersionDeferralOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VersionDeferralOutput)
}

// VersionDeferralArrayInput is an input type that accepts VersionDeferralArray and VersionDeferralArrayOutput values.
// You can construct a concrete instance of `VersionDeferralArrayInput` via:
//
//	VersionDeferralArray{ VersionDeferralArgs{...} }
type VersionDeferralArrayInput interface {
	pulumi.Input

	ToVersionDeferralArrayOutput() VersionDeferralArrayOutput
	ToVersionDeferralArrayOutputWithContext(context.Context) VersionDeferralArrayOutput
}

type VersionDeferralArray []VersionDeferralInput

func (VersionDeferralArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VersionDeferral)(nil)).Elem()
}

func (i VersionDeferralArray) ToVersionDeferralArrayOutput() VersionDeferralArrayOutput {
	return i.ToVersionDeferralArrayOutputWithContext(context.Background())
}

func (i VersionDeferralArray) ToVersionDeferralArrayOutputWithContext(ctx context.Context) VersionDeferralArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VersionDeferralArrayOutput)
}

// VersionDeferralMapInput is an input type that accepts VersionDeferralMap and VersionDeferralMapOutput values.
// You can construct a concrete instance of `VersionDeferralMapInput` via:
//
//	VersionDeferralMap{ "key": VersionDeferralArgs{...} }
type VersionDeferralMapInput interface {
	pulumi.Input

	ToVersionDeferralMapOutput() VersionDeferralMapOutput
	ToVersionDeferralMapOutputWithContext(context.Context) VersionDeferralMapOutput
}

type VersionDeferralMap map[string]VersionDeferralInput

func (VersionDeferralMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VersionDeferral)(nil)).Elem()
}

func (i VersionDeferralMap) ToVersionDeferralMapOutput() VersionDeferralMapOutput {
	return i.ToVersionDeferralMapOutputWithContext(context.Background())
}

func (i VersionDeferralMap) ToVersionDeferralMapOutputWithContext(ctx context.Context) VersionDeferralMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VersionDeferralMapOutput)
}

type VersionDeferralOutput struct{ *pulumi.OutputState }

func (VersionDeferralOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VersionDeferral)(nil)).Elem()
}

func (o VersionDeferralOutput) ToVersionDeferralOutput() VersionDeferralOutput {
	return o
}

func (o VersionDeferralOutput) ToVersionDeferralOutputWithContext(ctx context.Context) VersionDeferralOutput {
	return o
}

// The policy for managing automated minor version upgrades. Set to FIXED*DEFERRAL to defer upgrades by 60 days or NOT*DEFERRED to apply upgrades immediately.
func (o VersionDeferralOutput) DeferralPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *VersionDeferral) pulumi.StringOutput { return v.DeferralPolicy }).(pulumi.StringOutput)
}

type VersionDeferralArrayOutput struct{ *pulumi.OutputState }

func (VersionDeferralArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VersionDeferral)(nil)).Elem()
}

func (o VersionDeferralArrayOutput) ToVersionDeferralArrayOutput() VersionDeferralArrayOutput {
	return o
}

func (o VersionDeferralArrayOutput) ToVersionDeferralArrayOutputWithContext(ctx context.Context) VersionDeferralArrayOutput {
	return o
}

func (o VersionDeferralArrayOutput) Index(i pulumi.IntInput) VersionDeferralOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VersionDeferral {
		return vs[0].([]*VersionDeferral)[vs[1].(int)]
	}).(VersionDeferralOutput)
}

type VersionDeferralMapOutput struct{ *pulumi.OutputState }

func (VersionDeferralMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VersionDeferral)(nil)).Elem()
}

func (o VersionDeferralMapOutput) ToVersionDeferralMapOutput() VersionDeferralMapOutput {
	return o
}

func (o VersionDeferralMapOutput) ToVersionDeferralMapOutputWithContext(ctx context.Context) VersionDeferralMapOutput {
	return o
}

func (o VersionDeferralMapOutput) MapIndex(k pulumi.StringInput) VersionDeferralOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VersionDeferral {
		return vs[0].(map[string]*VersionDeferral)[vs[1].(string)]
	}).(VersionDeferralOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VersionDeferralInput)(nil)).Elem(), &VersionDeferral{})
	pulumi.RegisterInputType(reflect.TypeOf((*VersionDeferralArrayInput)(nil)).Elem(), VersionDeferralArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VersionDeferralMapInput)(nil)).Elem(), VersionDeferralMap{})
	pulumi.RegisterOutputType(VersionDeferralOutput{})
	pulumi.RegisterOutputType(VersionDeferralArrayOutput{})
	pulumi.RegisterOutputType(VersionDeferralMapOutput{})
}
