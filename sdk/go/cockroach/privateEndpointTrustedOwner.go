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

// Private Endpoint Trusted Owner.
type PrivateEndpointTrustedOwner struct {
	pulumi.CustomResourceState

	// UUID of the cluster the private endpoint trusted owner entry belongs to.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Owner ID of the private endpoint connection in the cloud provider.
	ExternalOwnerId pulumi.StringOutput `pulumi:"externalOwnerId"`
	// UUID of the private endpoint trusted owner entry.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointTrustedOwner registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointTrustedOwner(ctx *pulumi.Context,
	name string, args *PrivateEndpointTrustedOwnerArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointTrustedOwner, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.ExternalOwnerId == nil {
		return nil, errors.New("invalid value for required argument 'ExternalOwnerId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpointTrustedOwner
	err := ctx.RegisterResource("cockroach:index/privateEndpointTrustedOwner:PrivateEndpointTrustedOwner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointTrustedOwner gets an existing PrivateEndpointTrustedOwner resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointTrustedOwner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointTrustedOwnerState, opts ...pulumi.ResourceOption) (*PrivateEndpointTrustedOwner, error) {
	var resource PrivateEndpointTrustedOwner
	err := ctx.ReadResource("cockroach:index/privateEndpointTrustedOwner:PrivateEndpointTrustedOwner", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointTrustedOwner resources.
type privateEndpointTrustedOwnerState struct {
	// UUID of the cluster the private endpoint trusted owner entry belongs to.
	ClusterId *string `pulumi:"clusterId"`
	// Owner ID of the private endpoint connection in the cloud provider.
	ExternalOwnerId *string `pulumi:"externalOwnerId"`
	// UUID of the private endpoint trusted owner entry.
	OwnerId *string `pulumi:"ownerId"`
	// Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
	Type *string `pulumi:"type"`
}

type PrivateEndpointTrustedOwnerState struct {
	// UUID of the cluster the private endpoint trusted owner entry belongs to.
	ClusterId pulumi.StringPtrInput
	// Owner ID of the private endpoint connection in the cloud provider.
	ExternalOwnerId pulumi.StringPtrInput
	// UUID of the private endpoint trusted owner entry.
	OwnerId pulumi.StringPtrInput
	// Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
	Type pulumi.StringPtrInput
}

func (PrivateEndpointTrustedOwnerState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointTrustedOwnerState)(nil)).Elem()
}

type privateEndpointTrustedOwnerArgs struct {
	// UUID of the cluster the private endpoint trusted owner entry belongs to.
	ClusterId string `pulumi:"clusterId"`
	// Owner ID of the private endpoint connection in the cloud provider.
	ExternalOwnerId string `pulumi:"externalOwnerId"`
	// Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a PrivateEndpointTrustedOwner resource.
type PrivateEndpointTrustedOwnerArgs struct {
	// UUID of the cluster the private endpoint trusted owner entry belongs to.
	ClusterId pulumi.StringInput
	// Owner ID of the private endpoint connection in the cloud provider.
	ExternalOwnerId pulumi.StringInput
	// Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
	Type pulumi.StringInput
}

func (PrivateEndpointTrustedOwnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointTrustedOwnerArgs)(nil)).Elem()
}

type PrivateEndpointTrustedOwnerInput interface {
	pulumi.Input

	ToPrivateEndpointTrustedOwnerOutput() PrivateEndpointTrustedOwnerOutput
	ToPrivateEndpointTrustedOwnerOutputWithContext(ctx context.Context) PrivateEndpointTrustedOwnerOutput
}

func (*PrivateEndpointTrustedOwner) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointTrustedOwner)(nil)).Elem()
}

func (i *PrivateEndpointTrustedOwner) ToPrivateEndpointTrustedOwnerOutput() PrivateEndpointTrustedOwnerOutput {
	return i.ToPrivateEndpointTrustedOwnerOutputWithContext(context.Background())
}

func (i *PrivateEndpointTrustedOwner) ToPrivateEndpointTrustedOwnerOutputWithContext(ctx context.Context) PrivateEndpointTrustedOwnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointTrustedOwnerOutput)
}

// PrivateEndpointTrustedOwnerArrayInput is an input type that accepts PrivateEndpointTrustedOwnerArray and PrivateEndpointTrustedOwnerArrayOutput values.
// You can construct a concrete instance of `PrivateEndpointTrustedOwnerArrayInput` via:
//
//	PrivateEndpointTrustedOwnerArray{ PrivateEndpointTrustedOwnerArgs{...} }
type PrivateEndpointTrustedOwnerArrayInput interface {
	pulumi.Input

	ToPrivateEndpointTrustedOwnerArrayOutput() PrivateEndpointTrustedOwnerArrayOutput
	ToPrivateEndpointTrustedOwnerArrayOutputWithContext(context.Context) PrivateEndpointTrustedOwnerArrayOutput
}

type PrivateEndpointTrustedOwnerArray []PrivateEndpointTrustedOwnerInput

func (PrivateEndpointTrustedOwnerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateEndpointTrustedOwner)(nil)).Elem()
}

func (i PrivateEndpointTrustedOwnerArray) ToPrivateEndpointTrustedOwnerArrayOutput() PrivateEndpointTrustedOwnerArrayOutput {
	return i.ToPrivateEndpointTrustedOwnerArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointTrustedOwnerArray) ToPrivateEndpointTrustedOwnerArrayOutputWithContext(ctx context.Context) PrivateEndpointTrustedOwnerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointTrustedOwnerArrayOutput)
}

// PrivateEndpointTrustedOwnerMapInput is an input type that accepts PrivateEndpointTrustedOwnerMap and PrivateEndpointTrustedOwnerMapOutput values.
// You can construct a concrete instance of `PrivateEndpointTrustedOwnerMapInput` via:
//
//	PrivateEndpointTrustedOwnerMap{ "key": PrivateEndpointTrustedOwnerArgs{...} }
type PrivateEndpointTrustedOwnerMapInput interface {
	pulumi.Input

	ToPrivateEndpointTrustedOwnerMapOutput() PrivateEndpointTrustedOwnerMapOutput
	ToPrivateEndpointTrustedOwnerMapOutputWithContext(context.Context) PrivateEndpointTrustedOwnerMapOutput
}

type PrivateEndpointTrustedOwnerMap map[string]PrivateEndpointTrustedOwnerInput

func (PrivateEndpointTrustedOwnerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateEndpointTrustedOwner)(nil)).Elem()
}

func (i PrivateEndpointTrustedOwnerMap) ToPrivateEndpointTrustedOwnerMapOutput() PrivateEndpointTrustedOwnerMapOutput {
	return i.ToPrivateEndpointTrustedOwnerMapOutputWithContext(context.Background())
}

func (i PrivateEndpointTrustedOwnerMap) ToPrivateEndpointTrustedOwnerMapOutputWithContext(ctx context.Context) PrivateEndpointTrustedOwnerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointTrustedOwnerMapOutput)
}

type PrivateEndpointTrustedOwnerOutput struct{ *pulumi.OutputState }

func (PrivateEndpointTrustedOwnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointTrustedOwner)(nil)).Elem()
}

func (o PrivateEndpointTrustedOwnerOutput) ToPrivateEndpointTrustedOwnerOutput() PrivateEndpointTrustedOwnerOutput {
	return o
}

func (o PrivateEndpointTrustedOwnerOutput) ToPrivateEndpointTrustedOwnerOutputWithContext(ctx context.Context) PrivateEndpointTrustedOwnerOutput {
	return o
}

// UUID of the cluster the private endpoint trusted owner entry belongs to.
func (o PrivateEndpointTrustedOwnerOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointTrustedOwner) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Owner ID of the private endpoint connection in the cloud provider.
func (o PrivateEndpointTrustedOwnerOutput) ExternalOwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointTrustedOwner) pulumi.StringOutput { return v.ExternalOwnerId }).(pulumi.StringOutput)
}

// UUID of the private endpoint trusted owner entry.
func (o PrivateEndpointTrustedOwnerOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointTrustedOwner) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
func (o PrivateEndpointTrustedOwnerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointTrustedOwner) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointTrustedOwnerArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointTrustedOwnerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateEndpointTrustedOwner)(nil)).Elem()
}

func (o PrivateEndpointTrustedOwnerArrayOutput) ToPrivateEndpointTrustedOwnerArrayOutput() PrivateEndpointTrustedOwnerArrayOutput {
	return o
}

func (o PrivateEndpointTrustedOwnerArrayOutput) ToPrivateEndpointTrustedOwnerArrayOutputWithContext(ctx context.Context) PrivateEndpointTrustedOwnerArrayOutput {
	return o
}

func (o PrivateEndpointTrustedOwnerArrayOutput) Index(i pulumi.IntInput) PrivateEndpointTrustedOwnerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivateEndpointTrustedOwner {
		return vs[0].([]*PrivateEndpointTrustedOwner)[vs[1].(int)]
	}).(PrivateEndpointTrustedOwnerOutput)
}

type PrivateEndpointTrustedOwnerMapOutput struct{ *pulumi.OutputState }

func (PrivateEndpointTrustedOwnerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateEndpointTrustedOwner)(nil)).Elem()
}

func (o PrivateEndpointTrustedOwnerMapOutput) ToPrivateEndpointTrustedOwnerMapOutput() PrivateEndpointTrustedOwnerMapOutput {
	return o
}

func (o PrivateEndpointTrustedOwnerMapOutput) ToPrivateEndpointTrustedOwnerMapOutputWithContext(ctx context.Context) PrivateEndpointTrustedOwnerMapOutput {
	return o
}

func (o PrivateEndpointTrustedOwnerMapOutput) MapIndex(k pulumi.StringInput) PrivateEndpointTrustedOwnerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivateEndpointTrustedOwner {
		return vs[0].(map[string]*PrivateEndpointTrustedOwner)[vs[1].(string)]
	}).(PrivateEndpointTrustedOwnerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointTrustedOwnerInput)(nil)).Elem(), &PrivateEndpointTrustedOwner{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointTrustedOwnerArrayInput)(nil)).Elem(), PrivateEndpointTrustedOwnerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointTrustedOwnerMapInput)(nil)).Elem(), PrivateEndpointTrustedOwnerMap{})
	pulumi.RegisterOutputType(PrivateEndpointTrustedOwnerOutput{})
	pulumi.RegisterOutputType(PrivateEndpointTrustedOwnerArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointTrustedOwnerMapOutput{})
}
