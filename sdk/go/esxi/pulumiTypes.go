// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package esxi

import (
	"context"
	"reflect"

	"github.com/edmondshtogu/pulumi-esxi-native/sdk/v3/go/esxi/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type KeyValuePair struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// KeyValuePairInput is an input type that accepts KeyValuePairArgs and KeyValuePairOutput values.
// You can construct a concrete instance of `KeyValuePairInput` via:
//
//	KeyValuePairArgs{...}
type KeyValuePairInput interface {
	pulumi.Input

	ToKeyValuePairOutput() KeyValuePairOutput
	ToKeyValuePairOutputWithContext(context.Context) KeyValuePairOutput
}

type KeyValuePairArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (KeyValuePairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyValuePair)(nil)).Elem()
}

func (i KeyValuePairArgs) ToKeyValuePairOutput() KeyValuePairOutput {
	return i.ToKeyValuePairOutputWithContext(context.Background())
}

func (i KeyValuePairArgs) ToKeyValuePairOutputWithContext(ctx context.Context) KeyValuePairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyValuePairOutput)
}

// KeyValuePairArrayInput is an input type that accepts KeyValuePairArray and KeyValuePairArrayOutput values.
// You can construct a concrete instance of `KeyValuePairArrayInput` via:
//
//	KeyValuePairArray{ KeyValuePairArgs{...} }
type KeyValuePairArrayInput interface {
	pulumi.Input

	ToKeyValuePairArrayOutput() KeyValuePairArrayOutput
	ToKeyValuePairArrayOutputWithContext(context.Context) KeyValuePairArrayOutput
}

type KeyValuePairArray []KeyValuePairInput

func (KeyValuePairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyValuePair)(nil)).Elem()
}

func (i KeyValuePairArray) ToKeyValuePairArrayOutput() KeyValuePairArrayOutput {
	return i.ToKeyValuePairArrayOutputWithContext(context.Background())
}

func (i KeyValuePairArray) ToKeyValuePairArrayOutputWithContext(ctx context.Context) KeyValuePairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyValuePairArrayOutput)
}

type KeyValuePairOutput struct{ *pulumi.OutputState }

func (KeyValuePairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyValuePair)(nil)).Elem()
}

func (o KeyValuePairOutput) ToKeyValuePairOutput() KeyValuePairOutput {
	return o
}

func (o KeyValuePairOutput) ToKeyValuePairOutputWithContext(ctx context.Context) KeyValuePairOutput {
	return o
}

func (o KeyValuePairOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v KeyValuePair) string { return v.Key }).(pulumi.StringOutput)
}

func (o KeyValuePairOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v KeyValuePair) string { return v.Value }).(pulumi.StringOutput)
}

type KeyValuePairArrayOutput struct{ *pulumi.OutputState }

func (KeyValuePairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyValuePair)(nil)).Elem()
}

func (o KeyValuePairArrayOutput) ToKeyValuePairArrayOutput() KeyValuePairArrayOutput {
	return o
}

func (o KeyValuePairArrayOutput) ToKeyValuePairArrayOutputWithContext(ctx context.Context) KeyValuePairArrayOutput {
	return o
}

func (o KeyValuePairArrayOutput) Index(i pulumi.IntInput) KeyValuePairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyValuePair {
		return vs[0].([]KeyValuePair)[vs[1].(int)]
	}).(KeyValuePairOutput)
}

type NetworkInterface struct {
	MacAddress     *string `pulumi:"macAddress"`
	NicType        *string `pulumi:"nicType"`
	VirtualNetwork string  `pulumi:"virtualNetwork"`
}

// NetworkInterfaceInput is an input type that accepts NetworkInterfaceArgs and NetworkInterfaceOutput values.
// You can construct a concrete instance of `NetworkInterfaceInput` via:
//
//	NetworkInterfaceArgs{...}
type NetworkInterfaceInput interface {
	pulumi.Input

	ToNetworkInterfaceOutput() NetworkInterfaceOutput
	ToNetworkInterfaceOutputWithContext(context.Context) NetworkInterfaceOutput
}

type NetworkInterfaceArgs struct {
	MacAddress     pulumi.StringPtrInput `pulumi:"macAddress"`
	NicType        pulumi.StringPtrInput `pulumi:"nicType"`
	VirtualNetwork pulumi.StringInput    `pulumi:"virtualNetwork"`
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterface)(nil)).Elem()
}

func (i NetworkInterfaceArgs) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return i.ToNetworkInterfaceOutputWithContext(context.Background())
}

func (i NetworkInterfaceArgs) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceOutput)
}

// NetworkInterfaceArrayInput is an input type that accepts NetworkInterfaceArray and NetworkInterfaceArrayOutput values.
// You can construct a concrete instance of `NetworkInterfaceArrayInput` via:
//
//	NetworkInterfaceArray{ NetworkInterfaceArgs{...} }
type NetworkInterfaceArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput
	ToNetworkInterfaceArrayOutputWithContext(context.Context) NetworkInterfaceArrayOutput
}

type NetworkInterfaceArray []NetworkInterfaceInput

func (NetworkInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterface)(nil)).Elem()
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return i.ToNetworkInterfaceArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceArrayOutput)
}

type NetworkInterfaceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterface) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) NicType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterface) *string { return v.NicType }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) VirtualNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterface) string { return v.VirtualNetwork }).(pulumi.StringOutput)
}

type NetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterface {
		return vs[0].([]NetworkInterface)[vs[1].(int)]
	}).(NetworkInterfaceOutput)
}

type Uplink struct {
	// Uplink name.
	Name string `pulumi:"name"`
}

// UplinkInput is an input type that accepts UplinkArgs and UplinkOutput values.
// You can construct a concrete instance of `UplinkInput` via:
//
//	UplinkArgs{...}
type UplinkInput interface {
	pulumi.Input

	ToUplinkOutput() UplinkOutput
	ToUplinkOutputWithContext(context.Context) UplinkOutput
}

type UplinkArgs struct {
	// Uplink name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (UplinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Uplink)(nil)).Elem()
}

func (i UplinkArgs) ToUplinkOutput() UplinkOutput {
	return i.ToUplinkOutputWithContext(context.Background())
}

func (i UplinkArgs) ToUplinkOutputWithContext(ctx context.Context) UplinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UplinkOutput)
}

// UplinkArrayInput is an input type that accepts UplinkArray and UplinkArrayOutput values.
// You can construct a concrete instance of `UplinkArrayInput` via:
//
//	UplinkArray{ UplinkArgs{...} }
type UplinkArrayInput interface {
	pulumi.Input

	ToUplinkArrayOutput() UplinkArrayOutput
	ToUplinkArrayOutputWithContext(context.Context) UplinkArrayOutput
}

type UplinkArray []UplinkInput

func (UplinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Uplink)(nil)).Elem()
}

func (i UplinkArray) ToUplinkArrayOutput() UplinkArrayOutput {
	return i.ToUplinkArrayOutputWithContext(context.Background())
}

func (i UplinkArray) ToUplinkArrayOutputWithContext(ctx context.Context) UplinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UplinkArrayOutput)
}

type UplinkOutput struct{ *pulumi.OutputState }

func (UplinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Uplink)(nil)).Elem()
}

func (o UplinkOutput) ToUplinkOutput() UplinkOutput {
	return o
}

func (o UplinkOutput) ToUplinkOutputWithContext(ctx context.Context) UplinkOutput {
	return o
}

// Uplink name.
func (o UplinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Uplink) string { return v.Name }).(pulumi.StringOutput)
}

type UplinkArrayOutput struct{ *pulumi.OutputState }

func (UplinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Uplink)(nil)).Elem()
}

func (o UplinkArrayOutput) ToUplinkArrayOutput() UplinkArrayOutput {
	return o
}

func (o UplinkArrayOutput) ToUplinkArrayOutputWithContext(ctx context.Context) UplinkArrayOutput {
	return o
}

func (o UplinkArrayOutput) Index(i pulumi.IntInput) UplinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Uplink {
		return vs[0].([]Uplink)[vs[1].(int)]
	}).(UplinkOutput)
}

type VMVirtualDisk struct {
	// SCSI_Ctrl:SCSI_id.    Range  '0:1' to '0:15'.   SCSI_id 7 is not allowed.
	Slot          *string `pulumi:"slot"`
	VirtualDiskId string  `pulumi:"virtualDiskId"`
}

// VMVirtualDiskInput is an input type that accepts VMVirtualDiskArgs and VMVirtualDiskOutput values.
// You can construct a concrete instance of `VMVirtualDiskInput` via:
//
//	VMVirtualDiskArgs{...}
type VMVirtualDiskInput interface {
	pulumi.Input

	ToVMVirtualDiskOutput() VMVirtualDiskOutput
	ToVMVirtualDiskOutputWithContext(context.Context) VMVirtualDiskOutput
}

type VMVirtualDiskArgs struct {
	// SCSI_Ctrl:SCSI_id.    Range  '0:1' to '0:15'.   SCSI_id 7 is not allowed.
	Slot          pulumi.StringPtrInput `pulumi:"slot"`
	VirtualDiskId pulumi.StringInput    `pulumi:"virtualDiskId"`
}

func (VMVirtualDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMVirtualDisk)(nil)).Elem()
}

func (i VMVirtualDiskArgs) ToVMVirtualDiskOutput() VMVirtualDiskOutput {
	return i.ToVMVirtualDiskOutputWithContext(context.Background())
}

func (i VMVirtualDiskArgs) ToVMVirtualDiskOutputWithContext(ctx context.Context) VMVirtualDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMVirtualDiskOutput)
}

// VMVirtualDiskArrayInput is an input type that accepts VMVirtualDiskArray and VMVirtualDiskArrayOutput values.
// You can construct a concrete instance of `VMVirtualDiskArrayInput` via:
//
//	VMVirtualDiskArray{ VMVirtualDiskArgs{...} }
type VMVirtualDiskArrayInput interface {
	pulumi.Input

	ToVMVirtualDiskArrayOutput() VMVirtualDiskArrayOutput
	ToVMVirtualDiskArrayOutputWithContext(context.Context) VMVirtualDiskArrayOutput
}

type VMVirtualDiskArray []VMVirtualDiskInput

func (VMVirtualDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMVirtualDisk)(nil)).Elem()
}

func (i VMVirtualDiskArray) ToVMVirtualDiskArrayOutput() VMVirtualDiskArrayOutput {
	return i.ToVMVirtualDiskArrayOutputWithContext(context.Background())
}

func (i VMVirtualDiskArray) ToVMVirtualDiskArrayOutputWithContext(ctx context.Context) VMVirtualDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMVirtualDiskArrayOutput)
}

type VMVirtualDiskOutput struct{ *pulumi.OutputState }

func (VMVirtualDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMVirtualDisk)(nil)).Elem()
}

func (o VMVirtualDiskOutput) ToVMVirtualDiskOutput() VMVirtualDiskOutput {
	return o
}

func (o VMVirtualDiskOutput) ToVMVirtualDiskOutputWithContext(ctx context.Context) VMVirtualDiskOutput {
	return o
}

// SCSI_Ctrl:SCSI_id.    Range  '0:1' to '0:15'.   SCSI_id 7 is not allowed.
func (o VMVirtualDiskOutput) Slot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMVirtualDisk) *string { return v.Slot }).(pulumi.StringPtrOutput)
}

func (o VMVirtualDiskOutput) VirtualDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v VMVirtualDisk) string { return v.VirtualDiskId }).(pulumi.StringOutput)
}

type VMVirtualDiskArrayOutput struct{ *pulumi.OutputState }

func (VMVirtualDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMVirtualDisk)(nil)).Elem()
}

func (o VMVirtualDiskArrayOutput) ToVMVirtualDiskArrayOutput() VMVirtualDiskArrayOutput {
	return o
}

func (o VMVirtualDiskArrayOutput) ToVMVirtualDiskArrayOutputWithContext(ctx context.Context) VMVirtualDiskArrayOutput {
	return o
}

func (o VMVirtualDiskArrayOutput) Index(i pulumi.IntInput) VMVirtualDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMVirtualDisk {
		return vs[0].([]VMVirtualDisk)[vs[1].(int)]
	}).(VMVirtualDiskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyValuePairInput)(nil)).Elem(), KeyValuePairArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyValuePairArrayInput)(nil)).Elem(), KeyValuePairArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceInput)(nil)).Elem(), NetworkInterfaceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceArrayInput)(nil)).Elem(), NetworkInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UplinkInput)(nil)).Elem(), UplinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UplinkArrayInput)(nil)).Elem(), UplinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMVirtualDiskInput)(nil)).Elem(), VMVirtualDiskArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMVirtualDiskArrayInput)(nil)).Elem(), VMVirtualDiskArray{})
	pulumi.RegisterOutputType(KeyValuePairOutput{})
	pulumi.RegisterOutputType(KeyValuePairArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceArrayOutput{})
	pulumi.RegisterOutputType(UplinkOutput{})
	pulumi.RegisterOutputType(UplinkArrayOutput{})
	pulumi.RegisterOutputType(VMVirtualDiskOutput{})
	pulumi.RegisterOutputType(VMVirtualDiskArrayOutput{})
}
