// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package esxi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourcePool struct {
	pulumi.CustomResourceState

	// CPU maximum (in MHz).
	CpuMax pulumi.IntPtrOutput `pulumi:"cpuMax"`
	// CPU minimum (in MHz).
	CpuMin pulumi.IntPtrOutput `pulumi:"cpuMin"`
	// Can pool borrow CPU resources from parent?
	CpuMinExpandable pulumi.StringPtrOutput `pulumi:"cpuMinExpandable"`
	// CPU shares (low/normal/high/<custom>).
	CpuShares pulumi.StringPtrOutput `pulumi:"cpuShares"`
	// Memory maximum (in MB).
	MemMax pulumi.IntPtrOutput `pulumi:"memMax"`
	// Memory minimum (in MB).
	MemMin pulumi.IntPtrOutput `pulumi:"memMin"`
	// Can pool borrow memory resources from parent?
	MemMinExpandable pulumi.StringPtrOutput `pulumi:"memMinExpandable"`
	// Memory shares (low/normal/high/<custom>).
	MemShares pulumi.StringPtrOutput `pulumi:"memShares"`
	// Resource Pool Name
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewResourcePool registers a new resource with the given unique name, arguments, and options.
func NewResourcePool(ctx *pulumi.Context,
	name string, args *ResourcePoolArgs, opts ...pulumi.ResourceOption) (*ResourcePool, error) {
	if args == nil {
		args = &ResourcePoolArgs{}
	}

	var resource ResourcePool
	err := ctx.RegisterResource("esxi-native:index:ResourcePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcePool gets an existing ResourcePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePoolState, opts ...pulumi.ResourceOption) (*ResourcePool, error) {
	var resource ResourcePool
	err := ctx.ReadResource("esxi-native:index:ResourcePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcePool resources.
type resourcePoolState struct {
}

type ResourcePoolState struct {
}

func (ResourcePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePoolState)(nil)).Elem()
}

type resourcePoolArgs struct {
	// CPU maximum (in MHz).
	CpuMax *int `pulumi:"cpuMax"`
	// CPU minimum (in MHz).
	CpuMin *int `pulumi:"cpuMin"`
	// Can pool borrow CPU resources from parent?
	CpuMinExpandable *string `pulumi:"cpuMinExpandable"`
	// CPU shares (low/normal/high/<custom>).
	CpuShares *string `pulumi:"cpuShares"`
	// Memory maximum (in MB).
	MemMax *int `pulumi:"memMax"`
	// Memory minimum (in MB).
	MemMin *int `pulumi:"memMin"`
	// Can pool borrow memory resources from parent?
	MemMinExpandable *string `pulumi:"memMinExpandable"`
	// Memory shares (low/normal/high/<custom>).
	MemShares *string `pulumi:"memShares"`
	// Resource Pool Name
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ResourcePool resource.
type ResourcePoolArgs struct {
	// CPU maximum (in MHz).
	CpuMax pulumi.IntPtrInput
	// CPU minimum (in MHz).
	CpuMin pulumi.IntPtrInput
	// Can pool borrow CPU resources from parent?
	CpuMinExpandable pulumi.StringPtrInput
	// CPU shares (low/normal/high/<custom>).
	CpuShares pulumi.StringPtrInput
	// Memory maximum (in MB).
	MemMax pulumi.IntPtrInput
	// Memory minimum (in MB).
	MemMin pulumi.IntPtrInput
	// Can pool borrow memory resources from parent?
	MemMinExpandable pulumi.StringPtrInput
	// Memory shares (low/normal/high/<custom>).
	MemShares pulumi.StringPtrInput
	// Resource Pool Name
	Name pulumi.StringPtrInput
}

func (ResourcePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePoolArgs)(nil)).Elem()
}

type ResourcePoolInput interface {
	pulumi.Input

	ToResourcePoolOutput() ResourcePoolOutput
	ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput
}

func (*ResourcePool) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePool)(nil)).Elem()
}

func (i *ResourcePool) ToResourcePoolOutput() ResourcePoolOutput {
	return i.ToResourcePoolOutputWithContext(context.Background())
}

func (i *ResourcePool) ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolOutput)
}

// ResourcePoolArrayInput is an input type that accepts ResourcePoolArray and ResourcePoolArrayOutput values.
// You can construct a concrete instance of `ResourcePoolArrayInput` via:
//
//          ResourcePoolArray{ ResourcePoolArgs{...} }
type ResourcePoolArrayInput interface {
	pulumi.Input

	ToResourcePoolArrayOutput() ResourcePoolArrayOutput
	ToResourcePoolArrayOutputWithContext(context.Context) ResourcePoolArrayOutput
}

type ResourcePoolArray []ResourcePoolInput

func (ResourcePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcePool)(nil)).Elem()
}

func (i ResourcePoolArray) ToResourcePoolArrayOutput() ResourcePoolArrayOutput {
	return i.ToResourcePoolArrayOutputWithContext(context.Background())
}

func (i ResourcePoolArray) ToResourcePoolArrayOutputWithContext(ctx context.Context) ResourcePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolArrayOutput)
}

// ResourcePoolMapInput is an input type that accepts ResourcePoolMap and ResourcePoolMapOutput values.
// You can construct a concrete instance of `ResourcePoolMapInput` via:
//
//          ResourcePoolMap{ "key": ResourcePoolArgs{...} }
type ResourcePoolMapInput interface {
	pulumi.Input

	ToResourcePoolMapOutput() ResourcePoolMapOutput
	ToResourcePoolMapOutputWithContext(context.Context) ResourcePoolMapOutput
}

type ResourcePoolMap map[string]ResourcePoolInput

func (ResourcePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcePool)(nil)).Elem()
}

func (i ResourcePoolMap) ToResourcePoolMapOutput() ResourcePoolMapOutput {
	return i.ToResourcePoolMapOutputWithContext(context.Background())
}

func (i ResourcePoolMap) ToResourcePoolMapOutputWithContext(ctx context.Context) ResourcePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolMapOutput)
}

type ResourcePoolOutput struct{ *pulumi.OutputState }

func (ResourcePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePool)(nil)).Elem()
}

func (o ResourcePoolOutput) ToResourcePoolOutput() ResourcePoolOutput {
	return o
}

func (o ResourcePoolOutput) ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput {
	return o
}

// CPU maximum (in MHz).
func (o ResourcePoolOutput) CpuMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.IntPtrOutput { return v.CpuMax }).(pulumi.IntPtrOutput)
}

// CPU minimum (in MHz).
func (o ResourcePoolOutput) CpuMin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.IntPtrOutput { return v.CpuMin }).(pulumi.IntPtrOutput)
}

// Can pool borrow CPU resources from parent?
func (o ResourcePoolOutput) CpuMinExpandable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringPtrOutput { return v.CpuMinExpandable }).(pulumi.StringPtrOutput)
}

// CPU shares (low/normal/high/<custom>).
func (o ResourcePoolOutput) CpuShares() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringPtrOutput { return v.CpuShares }).(pulumi.StringPtrOutput)
}

// Memory maximum (in MB).
func (o ResourcePoolOutput) MemMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.IntPtrOutput { return v.MemMax }).(pulumi.IntPtrOutput)
}

// Memory minimum (in MB).
func (o ResourcePoolOutput) MemMin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.IntPtrOutput { return v.MemMin }).(pulumi.IntPtrOutput)
}

// Can pool borrow memory resources from parent?
func (o ResourcePoolOutput) MemMinExpandable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringPtrOutput { return v.MemMinExpandable }).(pulumi.StringPtrOutput)
}

// Memory shares (low/normal/high/<custom>).
func (o ResourcePoolOutput) MemShares() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringPtrOutput { return v.MemShares }).(pulumi.StringPtrOutput)
}

// Resource Pool Name
func (o ResourcePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type ResourcePoolArrayOutput struct{ *pulumi.OutputState }

func (ResourcePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcePool)(nil)).Elem()
}

func (o ResourcePoolArrayOutput) ToResourcePoolArrayOutput() ResourcePoolArrayOutput {
	return o
}

func (o ResourcePoolArrayOutput) ToResourcePoolArrayOutputWithContext(ctx context.Context) ResourcePoolArrayOutput {
	return o
}

func (o ResourcePoolArrayOutput) Index(i pulumi.IntInput) ResourcePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourcePool {
		return vs[0].([]*ResourcePool)[vs[1].(int)]
	}).(ResourcePoolOutput)
}

type ResourcePoolMapOutput struct{ *pulumi.OutputState }

func (ResourcePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcePool)(nil)).Elem()
}

func (o ResourcePoolMapOutput) ToResourcePoolMapOutput() ResourcePoolMapOutput {
	return o
}

func (o ResourcePoolMapOutput) ToResourcePoolMapOutputWithContext(ctx context.Context) ResourcePoolMapOutput {
	return o
}

func (o ResourcePoolMapOutput) MapIndex(k pulumi.StringInput) ResourcePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourcePool {
		return vs[0].(map[string]*ResourcePool)[vs[1].(string)]
	}).(ResourcePoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePoolInput)(nil)).Elem(), &ResourcePool{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePoolArrayInput)(nil)).Elem(), ResourcePoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePoolMapInput)(nil)).Elem(), ResourcePoolMap{})
	pulumi.RegisterOutputType(ResourcePoolOutput{})
	pulumi.RegisterOutputType(ResourcePoolArrayOutput{})
	pulumi.RegisterOutputType(ResourcePoolMapOutput{})
}