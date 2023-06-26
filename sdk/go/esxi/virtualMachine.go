// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package esxi

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachine struct {
	pulumi.CustomResourceState

	// VM boot disk size. Will expand boot disk to this size.
	BootDiskSize pulumi.StringPtrOutput `pulumi:"bootDiskSize"`
	// VM boot disk type. thin, zeroedthick, eagerzeroedthick
	BootDiskType DiskTypePtrOutput `pulumi:"bootDiskType"`
	// Boot type('efi' is boot uefi mode)
	BootFirmware BootFirmwareTypePtrOutput `pulumi:"bootFirmware"`
	// esxi diskstore for boot disk.
	DiskStore pulumi.StringOutput `pulumi:"diskStore"`
	// pass data to VM
	Info KeyValuePairArrayOutput `pulumi:"info"`
	// The IP address reported by VMWare tools.
	IpAddress pulumi.StringPtrOutput `pulumi:"ipAddress"`
	// VM memory size.
	MemSize pulumi.StringOutput `pulumi:"memSize"`
	// esxi vm name.
	Name pulumi.StringOutput `pulumi:"name"`
	// VM network interfaces.
	NetworkInterfaces NetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	// VM memory size.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// VM number of virtual cpus.
	NumVCpus pulumi.StringOutput `pulumi:"numVCpus"`
	// VM OS type.
	Os pulumi.StringOutput `pulumi:"os"`
	// VM power state.
	Power pulumi.StringPtrOutput `pulumi:"power"`
	// Resource pool name to place vm.
	ResourcePoolName pulumi.StringOutput `pulumi:"resourcePoolName"`
	// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
	ShutdownTimeout pulumi.IntPtrOutput `pulumi:"shutdownTimeout"`
	// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
	StartupTimeout pulumi.IntPtrOutput `pulumi:"startupTimeout"`
	// VM virtual disks.
	VirtualDisks VMVirtualDiskArrayOutput `pulumi:"virtualDisks"`
	// VM Virtual HW version.
	VirtualHWVer pulumi.StringPtrOutput `pulumi:"virtualHWVer"`
}

// NewVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskStore == nil {
		return nil, errors.New("invalid value for required argument 'DiskStore'")
	}
	if args.MemSize == nil {
		return nil, errors.New("invalid value for required argument 'MemSize'")
	}
	if args.NumVCpus == nil {
		return nil, errors.New("invalid value for required argument 'NumVCpus'")
	}
	if args.Os == nil {
		return nil, errors.New("invalid value for required argument 'Os'")
	}
	if isZero(args.BootDiskType) {
		args.BootDiskType = DiskType("thin")
	}
	if isZero(args.BootFirmware) {
		args.BootFirmware = BootFirmwareType("bios")
	}
	if isZero(args.OvfPropertiesTimer) {
		args.OvfPropertiesTimer = pulumi.IntPtr(6000)
	}
	if isZero(args.ResourcePoolName) {
		args.ResourcePoolName = pulumi.String("/")
	}
	if isZero(args.ShutdownTimeout) {
		args.ShutdownTimeout = pulumi.IntPtr(600)
	}
	if isZero(args.StartupTimeout) {
		args.StartupTimeout = pulumi.IntPtr(600)
	}
	var resource VirtualMachine
	err := ctx.RegisterResource("esxi-native:index:VirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachine gets an existing VirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	var resource VirtualMachine
	err := ctx.ReadResource("esxi-native:index:VirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachine resources.
type virtualMachineState struct {
}

type VirtualMachineState struct {
}

func (VirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineState)(nil)).Elem()
}

type virtualMachineArgs struct {
	// VM boot disk type. thin, zeroedthick, eagerzeroedthick
	BootDiskType *DiskType `pulumi:"bootDiskType"`
	// Boot type('efi' is boot uefi mode)
	BootFirmware *BootFirmwareType `pulumi:"bootFirmware"`
	// Source vm path on esxi host to clone.
	CloneFromVirtualMachine *string `pulumi:"cloneFromVirtualMachine"`
	// esxi diskstore for boot disk.
	DiskStore string `pulumi:"diskStore"`
	// pass data to VM
	Info []KeyValuePair `pulumi:"info"`
	// VM memory size.
	MemSize string `pulumi:"memSize"`
	// esxi vm name.
	Name *string `pulumi:"name"`
	// VM network interfaces.
	NetworkInterfaces []NetworkInterface `pulumi:"networkInterfaces"`
	// VM memory size.
	Notes *string `pulumi:"notes"`
	// VM number of virtual cpus.
	NumVCpus int `pulumi:"numVCpus"`
	// VM OS type.
	Os string `pulumi:"os"`
	// Path on esxi host of ovf files.
	OvfHostPathSource *string `pulumi:"ovfHostPathSource"`
	// VM OVF properties.
	OvfProperties []KeyValuePair `pulumi:"ovfProperties"`
	// The amount of time, in seconds, to wait for the guest to boot and run ovfProperties. (0-6000)
	OvfPropertiesTimer *int `pulumi:"ovfPropertiesTimer"`
	// Local path to source ovf files.
	OvfSourceLocalPath *string `pulumi:"ovfSourceLocalPath"`
	// VM power state.
	Power *string `pulumi:"power"`
	// Resource pool name to place vm.
	ResourcePoolName string `pulumi:"resourcePoolName"`
	// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine. (0-600)
	ShutdownTimeout *int `pulumi:"shutdownTimeout"`
	// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine. (0-600)
	StartupTimeout *int `pulumi:"startupTimeout"`
	// VM virtual disks.
	VirtualDisks []VMVirtualDisk `pulumi:"virtualDisks"`
	// VM Virtual HW version.
	VirtualHWVer *string `pulumi:"virtualHWVer"`
}

// The set of arguments for constructing a VirtualMachine resource.
type VirtualMachineArgs struct {
	// VM boot disk type. thin, zeroedthick, eagerzeroedthick
	BootDiskType DiskTypePtrInput
	// Boot type('efi' is boot uefi mode)
	BootFirmware BootFirmwareTypePtrInput
	// Source vm path on esxi host to clone.
	CloneFromVirtualMachine pulumi.StringPtrInput
	// esxi diskstore for boot disk.
	DiskStore pulumi.StringInput
	// pass data to VM
	Info KeyValuePairArrayInput
	// VM memory size.
	MemSize pulumi.StringInput
	// esxi vm name.
	Name pulumi.StringPtrInput
	// VM network interfaces.
	NetworkInterfaces NetworkInterfaceArrayInput
	// VM memory size.
	Notes pulumi.StringPtrInput
	// VM number of virtual cpus.
	NumVCpus pulumi.IntInput
	// VM OS type.
	Os pulumi.StringInput
	// Path on esxi host of ovf files.
	OvfHostPathSource pulumi.StringPtrInput
	// VM OVF properties.
	OvfProperties KeyValuePairArrayInput
	// The amount of time, in seconds, to wait for the guest to boot and run ovfProperties. (0-6000)
	OvfPropertiesTimer pulumi.IntPtrInput
	// Local path to source ovf files.
	OvfSourceLocalPath pulumi.StringPtrInput
	// VM power state.
	Power pulumi.StringPtrInput
	// Resource pool name to place vm.
	ResourcePoolName pulumi.StringInput
	// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine. (0-600)
	ShutdownTimeout pulumi.IntPtrInput
	// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine. (0-600)
	StartupTimeout pulumi.IntPtrInput
	// VM virtual disks.
	VirtualDisks VMVirtualDiskArrayInput
	// VM Virtual HW version.
	VirtualHWVer pulumi.StringPtrInput
}

func (VirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineArgs)(nil)).Elem()
}

type VirtualMachineInput interface {
	pulumi.Input

	ToVirtualMachineOutput() VirtualMachineOutput
	ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput
}

func (*VirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachine)(nil)).Elem()
}

func (i *VirtualMachine) ToVirtualMachineOutput() VirtualMachineOutput {
	return i.ToVirtualMachineOutputWithContext(context.Background())
}

func (i *VirtualMachine) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineOutput)
}

// VirtualMachineArrayInput is an input type that accepts VirtualMachineArray and VirtualMachineArrayOutput values.
// You can construct a concrete instance of `VirtualMachineArrayInput` via:
//
//          VirtualMachineArray{ VirtualMachineArgs{...} }
type VirtualMachineArrayInput interface {
	pulumi.Input

	ToVirtualMachineArrayOutput() VirtualMachineArrayOutput
	ToVirtualMachineArrayOutputWithContext(context.Context) VirtualMachineArrayOutput
}

type VirtualMachineArray []VirtualMachineInput

func (VirtualMachineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualMachine)(nil)).Elem()
}

func (i VirtualMachineArray) ToVirtualMachineArrayOutput() VirtualMachineArrayOutput {
	return i.ToVirtualMachineArrayOutputWithContext(context.Background())
}

func (i VirtualMachineArray) ToVirtualMachineArrayOutputWithContext(ctx context.Context) VirtualMachineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineArrayOutput)
}

// VirtualMachineMapInput is an input type that accepts VirtualMachineMap and VirtualMachineMapOutput values.
// You can construct a concrete instance of `VirtualMachineMapInput` via:
//
//          VirtualMachineMap{ "key": VirtualMachineArgs{...} }
type VirtualMachineMapInput interface {
	pulumi.Input

	ToVirtualMachineMapOutput() VirtualMachineMapOutput
	ToVirtualMachineMapOutputWithContext(context.Context) VirtualMachineMapOutput
}

type VirtualMachineMap map[string]VirtualMachineInput

func (VirtualMachineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualMachine)(nil)).Elem()
}

func (i VirtualMachineMap) ToVirtualMachineMapOutput() VirtualMachineMapOutput {
	return i.ToVirtualMachineMapOutputWithContext(context.Background())
}

func (i VirtualMachineMap) ToVirtualMachineMapOutputWithContext(ctx context.Context) VirtualMachineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineMapOutput)
}

type VirtualMachineOutput struct{ *pulumi.OutputState }

func (VirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachine)(nil)).Elem()
}

func (o VirtualMachineOutput) ToVirtualMachineOutput() VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return o
}

// VM boot disk size. Will expand boot disk to this size.
func (o VirtualMachineOutput) BootDiskSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.BootDiskSize }).(pulumi.StringPtrOutput)
}

// VM boot disk type. thin, zeroedthick, eagerzeroedthick
func (o VirtualMachineOutput) BootDiskType() DiskTypePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) DiskTypePtrOutput { return v.BootDiskType }).(DiskTypePtrOutput)
}

// Boot type('efi' is boot uefi mode)
func (o VirtualMachineOutput) BootFirmware() BootFirmwareTypePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) BootFirmwareTypePtrOutput { return v.BootFirmware }).(BootFirmwareTypePtrOutput)
}

// esxi diskstore for boot disk.
func (o VirtualMachineOutput) DiskStore() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.DiskStore }).(pulumi.StringOutput)
}

// pass data to VM
func (o VirtualMachineOutput) Info() KeyValuePairArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) KeyValuePairArrayOutput { return v.Info }).(KeyValuePairArrayOutput)
}

// The IP address reported by VMWare tools.
func (o VirtualMachineOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// VM memory size.
func (o VirtualMachineOutput) MemSize() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.MemSize }).(pulumi.StringOutput)
}

// esxi vm name.
func (o VirtualMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// VM network interfaces.
func (o VirtualMachineOutput) NetworkInterfaces() NetworkInterfaceArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) NetworkInterfaceArrayOutput { return v.NetworkInterfaces }).(NetworkInterfaceArrayOutput)
}

// VM memory size.
func (o VirtualMachineOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// VM number of virtual cpus.
func (o VirtualMachineOutput) NumVCpus() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.NumVCpus }).(pulumi.StringOutput)
}

// VM OS type.
func (o VirtualMachineOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Os }).(pulumi.StringOutput)
}

// VM power state.
func (o VirtualMachineOutput) Power() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Power }).(pulumi.StringPtrOutput)
}

// Resource pool name to place vm.
func (o VirtualMachineOutput) ResourcePoolName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.ResourcePoolName }).(pulumi.StringOutput)
}

// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
func (o VirtualMachineOutput) ShutdownTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntPtrOutput { return v.ShutdownTimeout }).(pulumi.IntPtrOutput)
}

// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
func (o VirtualMachineOutput) StartupTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntPtrOutput { return v.StartupTimeout }).(pulumi.IntPtrOutput)
}

// VM virtual disks.
func (o VirtualMachineOutput) VirtualDisks() VMVirtualDiskArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) VMVirtualDiskArrayOutput { return v.VirtualDisks }).(VMVirtualDiskArrayOutput)
}

// VM Virtual HW version.
func (o VirtualMachineOutput) VirtualHWVer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.VirtualHWVer }).(pulumi.StringPtrOutput)
}

type VirtualMachineArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualMachine)(nil)).Elem()
}

func (o VirtualMachineArrayOutput) ToVirtualMachineArrayOutput() VirtualMachineArrayOutput {
	return o
}

func (o VirtualMachineArrayOutput) ToVirtualMachineArrayOutputWithContext(ctx context.Context) VirtualMachineArrayOutput {
	return o
}

func (o VirtualMachineArrayOutput) Index(i pulumi.IntInput) VirtualMachineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualMachine {
		return vs[0].([]*VirtualMachine)[vs[1].(int)]
	}).(VirtualMachineOutput)
}

type VirtualMachineMapOutput struct{ *pulumi.OutputState }

func (VirtualMachineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualMachine)(nil)).Elem()
}

func (o VirtualMachineMapOutput) ToVirtualMachineMapOutput() VirtualMachineMapOutput {
	return o
}

func (o VirtualMachineMapOutput) ToVirtualMachineMapOutputWithContext(ctx context.Context) VirtualMachineMapOutput {
	return o
}

func (o VirtualMachineMapOutput) MapIndex(k pulumi.StringInput) VirtualMachineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualMachine {
		return vs[0].(map[string]*VirtualMachine)[vs[1].(string)]
	}).(VirtualMachineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineInput)(nil)).Elem(), &VirtualMachine{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineArrayInput)(nil)).Elem(), VirtualMachineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineMapInput)(nil)).Elem(), VirtualMachineMap{})
	pulumi.RegisterOutputType(VirtualMachineOutput{})
	pulumi.RegisterOutputType(VirtualMachineArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineMapOutput{})
}