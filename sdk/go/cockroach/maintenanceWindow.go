// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cockroach

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Maintenance window resource for a cluster.
type MaintenanceWindow struct {
	pulumi.CustomResourceState

	// The offset duration is the duration in seconds from the beginning of each Monday (UTC) after which the maintenance window starts.
	OffsetDuration pulumi.IntOutput `pulumi:"offsetDuration"`
	// The window duration is the duration in seconds that the maintenance window will remain active for after it starts.
	WindowDuration pulumi.IntOutput `pulumi:"windowDuration"`
}

// NewMaintenanceWindow registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceWindow(ctx *pulumi.Context,
	name string, args *MaintenanceWindowArgs, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OffsetDuration == nil {
		return nil, errors.New("invalid value for required argument 'OffsetDuration'")
	}
	if args.WindowDuration == nil {
		return nil, errors.New("invalid value for required argument 'WindowDuration'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MaintenanceWindow
	err := ctx.RegisterResource("cockroach:index/maintenanceWindow:MaintenanceWindow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceWindow gets an existing MaintenanceWindow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceWindow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceWindowState, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	var resource MaintenanceWindow
	err := ctx.ReadResource("cockroach:index/maintenanceWindow:MaintenanceWindow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceWindow resources.
type maintenanceWindowState struct {
	// The offset duration is the duration in seconds from the beginning of each Monday (UTC) after which the maintenance window starts.
	OffsetDuration *int `pulumi:"offsetDuration"`
	// The window duration is the duration in seconds that the maintenance window will remain active for after it starts.
	WindowDuration *int `pulumi:"windowDuration"`
}

type MaintenanceWindowState struct {
	// The offset duration is the duration in seconds from the beginning of each Monday (UTC) after which the maintenance window starts.
	OffsetDuration pulumi.IntPtrInput
	// The window duration is the duration in seconds that the maintenance window will remain active for after it starts.
	WindowDuration pulumi.IntPtrInput
}

func (MaintenanceWindowState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowState)(nil)).Elem()
}

type maintenanceWindowArgs struct {
	// The offset duration is the duration in seconds from the beginning of each Monday (UTC) after which the maintenance window starts.
	OffsetDuration int `pulumi:"offsetDuration"`
	// The window duration is the duration in seconds that the maintenance window will remain active for after it starts.
	WindowDuration int `pulumi:"windowDuration"`
}

// The set of arguments for constructing a MaintenanceWindow resource.
type MaintenanceWindowArgs struct {
	// The offset duration is the duration in seconds from the beginning of each Monday (UTC) after which the maintenance window starts.
	OffsetDuration pulumi.IntInput
	// The window duration is the duration in seconds that the maintenance window will remain active for after it starts.
	WindowDuration pulumi.IntInput
}

func (MaintenanceWindowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowArgs)(nil)).Elem()
}

type MaintenanceWindowInput interface {
	pulumi.Input

	ToMaintenanceWindowOutput() MaintenanceWindowOutput
	ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput
}

func (*MaintenanceWindow) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (i *MaintenanceWindow) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return i.ToMaintenanceWindowOutputWithContext(context.Background())
}

func (i *MaintenanceWindow) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowOutput)
}

// MaintenanceWindowArrayInput is an input type that accepts MaintenanceWindowArray and MaintenanceWindowArrayOutput values.
// You can construct a concrete instance of `MaintenanceWindowArrayInput` via:
//
//	MaintenanceWindowArray{ MaintenanceWindowArgs{...} }
type MaintenanceWindowArrayInput interface {
	pulumi.Input

	ToMaintenanceWindowArrayOutput() MaintenanceWindowArrayOutput
	ToMaintenanceWindowArrayOutputWithContext(context.Context) MaintenanceWindowArrayOutput
}

type MaintenanceWindowArray []MaintenanceWindowInput

func (MaintenanceWindowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MaintenanceWindow)(nil)).Elem()
}

func (i MaintenanceWindowArray) ToMaintenanceWindowArrayOutput() MaintenanceWindowArrayOutput {
	return i.ToMaintenanceWindowArrayOutputWithContext(context.Background())
}

func (i MaintenanceWindowArray) ToMaintenanceWindowArrayOutputWithContext(ctx context.Context) MaintenanceWindowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowArrayOutput)
}

// MaintenanceWindowMapInput is an input type that accepts MaintenanceWindowMap and MaintenanceWindowMapOutput values.
// You can construct a concrete instance of `MaintenanceWindowMapInput` via:
//
//	MaintenanceWindowMap{ "key": MaintenanceWindowArgs{...} }
type MaintenanceWindowMapInput interface {
	pulumi.Input

	ToMaintenanceWindowMapOutput() MaintenanceWindowMapOutput
	ToMaintenanceWindowMapOutputWithContext(context.Context) MaintenanceWindowMapOutput
}

type MaintenanceWindowMap map[string]MaintenanceWindowInput

func (MaintenanceWindowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MaintenanceWindow)(nil)).Elem()
}

func (i MaintenanceWindowMap) ToMaintenanceWindowMapOutput() MaintenanceWindowMapOutput {
	return i.ToMaintenanceWindowMapOutputWithContext(context.Background())
}

func (i MaintenanceWindowMap) ToMaintenanceWindowMapOutputWithContext(ctx context.Context) MaintenanceWindowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowMapOutput)
}

type MaintenanceWindowOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return o
}

// The offset duration is the duration in seconds from the beginning of each Monday (UTC) after which the maintenance window starts.
func (o MaintenanceWindowOutput) OffsetDuration() pulumi.IntOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.IntOutput { return v.OffsetDuration }).(pulumi.IntOutput)
}

// The window duration is the duration in seconds that the maintenance window will remain active for after it starts.
func (o MaintenanceWindowOutput) WindowDuration() pulumi.IntOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.IntOutput { return v.WindowDuration }).(pulumi.IntOutput)
}

type MaintenanceWindowArrayOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowArrayOutput) ToMaintenanceWindowArrayOutput() MaintenanceWindowArrayOutput {
	return o
}

func (o MaintenanceWindowArrayOutput) ToMaintenanceWindowArrayOutputWithContext(ctx context.Context) MaintenanceWindowArrayOutput {
	return o
}

func (o MaintenanceWindowArrayOutput) Index(i pulumi.IntInput) MaintenanceWindowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MaintenanceWindow {
		return vs[0].([]*MaintenanceWindow)[vs[1].(int)]
	}).(MaintenanceWindowOutput)
}

type MaintenanceWindowMapOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowMapOutput) ToMaintenanceWindowMapOutput() MaintenanceWindowMapOutput {
	return o
}

func (o MaintenanceWindowMapOutput) ToMaintenanceWindowMapOutputWithContext(ctx context.Context) MaintenanceWindowMapOutput {
	return o
}

func (o MaintenanceWindowMapOutput) MapIndex(k pulumi.StringInput) MaintenanceWindowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MaintenanceWindow {
		return vs[0].(map[string]*MaintenanceWindow)[vs[1].(string)]
	}).(MaintenanceWindowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowInput)(nil)).Elem(), &MaintenanceWindow{})
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowArrayInput)(nil)).Elem(), MaintenanceWindowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowMapInput)(nil)).Elem(), MaintenanceWindowMap{})
	pulumi.RegisterOutputType(MaintenanceWindowOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowArrayOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowMapOutput{})
}