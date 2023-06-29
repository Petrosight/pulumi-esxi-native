// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package esxi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PortGroup struct {
	pulumi.CustomResourceState

	// Forged transmits (true=Accept/false=Reject).
	ForgedTransmits pulumi.BoolPtrOutput `pulumi:"forgedTransmits"`
	// MAC address changes (true=Accept/false=Reject).
	MacChanges pulumi.BoolPtrOutput `pulumi:"macChanges"`
	// Port Group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Promiscuous mode (true=Accept/false=Reject).
	PromiscuousMode pulumi.BoolPtrOutput `pulumi:"promiscuousMode"`
	// Virtual Switch Name.
	VSwitch pulumi.StringPtrOutput `pulumi:"vSwitch"`
	// Port Group vlan id
	Vlan pulumi.IntPtrOutput `pulumi:"vlan"`
}

// NewPortGroup registers a new resource with the given unique name, arguments, and options.
func NewPortGroup(ctx *pulumi.Context,
	name string, args *PortGroupArgs, opts ...pulumi.ResourceOption) (*PortGroup, error) {
	if args == nil {
		args = &PortGroupArgs{}
	}

	var resource PortGroup
	err := ctx.RegisterResource("esxi-native:index:PortGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPortGroup gets an existing PortGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPortGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortGroupState, opts ...pulumi.ResourceOption) (*PortGroup, error) {
	var resource PortGroup
	err := ctx.ReadResource("esxi-native:index:PortGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PortGroup resources.
type portGroupState struct {
}

type PortGroupState struct {
}

func (PortGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*portGroupState)(nil)).Elem()
}

type portGroupArgs struct {
	// Forged transmits (true=Accept/false=Reject).
	ForgedTransmits *bool `pulumi:"forgedTransmits"`
	// MAC address changes (true=Accept/false=Reject).
	MacChanges *bool `pulumi:"macChanges"`
	// Virtual Switch name.
	Name *string `pulumi:"name"`
	// Promiscuous mode (true=Accept/false=Reject).
	PromiscuousMode *bool `pulumi:"promiscuousMode"`
	// Virtual Switch Name.
	VSwitch *string `pulumi:"vSwitch"`
	// Port Group vlan id
	Vlan *int `pulumi:"vlan"`
}

// The set of arguments for constructing a PortGroup resource.
type PortGroupArgs struct {
	// Forged transmits (true=Accept/false=Reject).
	ForgedTransmits pulumi.BoolPtrInput
	// MAC address changes (true=Accept/false=Reject).
	MacChanges pulumi.BoolPtrInput
	// Virtual Switch name.
	Name pulumi.StringPtrInput
	// Promiscuous mode (true=Accept/false=Reject).
	PromiscuousMode pulumi.BoolPtrInput
	// Virtual Switch Name.
	VSwitch pulumi.StringPtrInput
	// Port Group vlan id
	Vlan pulumi.IntPtrInput
}

func (PortGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portGroupArgs)(nil)).Elem()
}

type PortGroupInput interface {
	pulumi.Input

	ToPortGroupOutput() PortGroupOutput
	ToPortGroupOutputWithContext(ctx context.Context) PortGroupOutput
}

func (*PortGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**PortGroup)(nil)).Elem()
}

func (i *PortGroup) ToPortGroupOutput() PortGroupOutput {
	return i.ToPortGroupOutputWithContext(context.Background())
}

func (i *PortGroup) ToPortGroupOutputWithContext(ctx context.Context) PortGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortGroupOutput)
}

// PortGroupArrayInput is an input type that accepts PortGroupArray and PortGroupArrayOutput values.
// You can construct a concrete instance of `PortGroupArrayInput` via:
//
//	PortGroupArray{ PortGroupArgs{...} }
type PortGroupArrayInput interface {
	pulumi.Input

	ToPortGroupArrayOutput() PortGroupArrayOutput
	ToPortGroupArrayOutputWithContext(context.Context) PortGroupArrayOutput
}

type PortGroupArray []PortGroupInput

func (PortGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PortGroup)(nil)).Elem()
}

func (i PortGroupArray) ToPortGroupArrayOutput() PortGroupArrayOutput {
	return i.ToPortGroupArrayOutputWithContext(context.Background())
}

func (i PortGroupArray) ToPortGroupArrayOutputWithContext(ctx context.Context) PortGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortGroupArrayOutput)
}

// PortGroupMapInput is an input type that accepts PortGroupMap and PortGroupMapOutput values.
// You can construct a concrete instance of `PortGroupMapInput` via:
//
//	PortGroupMap{ "key": PortGroupArgs{...} }
type PortGroupMapInput interface {
	pulumi.Input

	ToPortGroupMapOutput() PortGroupMapOutput
	ToPortGroupMapOutputWithContext(context.Context) PortGroupMapOutput
}

type PortGroupMap map[string]PortGroupInput

func (PortGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PortGroup)(nil)).Elem()
}

func (i PortGroupMap) ToPortGroupMapOutput() PortGroupMapOutput {
	return i.ToPortGroupMapOutputWithContext(context.Background())
}

func (i PortGroupMap) ToPortGroupMapOutputWithContext(ctx context.Context) PortGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortGroupMapOutput)
}

type PortGroupOutput struct{ *pulumi.OutputState }

func (PortGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PortGroup)(nil)).Elem()
}

func (o PortGroupOutput) ToPortGroupOutput() PortGroupOutput {
	return o
}

func (o PortGroupOutput) ToPortGroupOutputWithContext(ctx context.Context) PortGroupOutput {
	return o
}

// Forged transmits (true=Accept/false=Reject).
func (o PortGroupOutput) ForgedTransmits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PortGroup) pulumi.BoolPtrOutput { return v.ForgedTransmits }).(pulumi.BoolPtrOutput)
}

// MAC address changes (true=Accept/false=Reject).
func (o PortGroupOutput) MacChanges() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PortGroup) pulumi.BoolPtrOutput { return v.MacChanges }).(pulumi.BoolPtrOutput)
}

// Port Group name.
func (o PortGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PortGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Promiscuous mode (true=Accept/false=Reject).
func (o PortGroupOutput) PromiscuousMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PortGroup) pulumi.BoolPtrOutput { return v.PromiscuousMode }).(pulumi.BoolPtrOutput)
}

// Virtual Switch Name.
func (o PortGroupOutput) VSwitch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PortGroup) pulumi.StringPtrOutput { return v.VSwitch }).(pulumi.StringPtrOutput)
}

// Port Group vlan id
func (o PortGroupOutput) Vlan() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PortGroup) pulumi.IntPtrOutput { return v.Vlan }).(pulumi.IntPtrOutput)
}

type PortGroupArrayOutput struct{ *pulumi.OutputState }

func (PortGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PortGroup)(nil)).Elem()
}

func (o PortGroupArrayOutput) ToPortGroupArrayOutput() PortGroupArrayOutput {
	return o
}

func (o PortGroupArrayOutput) ToPortGroupArrayOutputWithContext(ctx context.Context) PortGroupArrayOutput {
	return o
}

func (o PortGroupArrayOutput) Index(i pulumi.IntInput) PortGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PortGroup {
		return vs[0].([]*PortGroup)[vs[1].(int)]
	}).(PortGroupOutput)
}

type PortGroupMapOutput struct{ *pulumi.OutputState }

func (PortGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PortGroup)(nil)).Elem()
}

func (o PortGroupMapOutput) ToPortGroupMapOutput() PortGroupMapOutput {
	return o
}

func (o PortGroupMapOutput) ToPortGroupMapOutputWithContext(ctx context.Context) PortGroupMapOutput {
	return o
}

func (o PortGroupMapOutput) MapIndex(k pulumi.StringInput) PortGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PortGroup {
		return vs[0].(map[string]*PortGroup)[vs[1].(string)]
	}).(PortGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PortGroupInput)(nil)).Elem(), &PortGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortGroupArrayInput)(nil)).Elem(), PortGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortGroupMapInput)(nil)).Elem(), PortGroupMap{})
	pulumi.RegisterOutputType(PortGroupOutput{})
	pulumi.RegisterOutputType(PortGroupArrayOutput{})
	pulumi.RegisterOutputType(PortGroupMapOutput{})
}
