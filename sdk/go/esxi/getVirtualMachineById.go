// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package esxi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualMachineById(ctx *pulumi.Context, args *GetVirtualMachineByIdArgs, opts ...pulumi.InvokeOption) (*GetVirtualMachineByIdResult, error) {
	var rv GetVirtualMachineByIdResult
	err := ctx.Invoke("esxi-native:index:getVirtualMachineById", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type GetVirtualMachineByIdArgs struct {
	// Virtual Machine Id to get details of
	Id string `pulumi:"id"`
}

type GetVirtualMachineByIdResult struct {
	// VM boot disk size. Will expand boot disk to this size.
	BootDiskSize *int `pulumi:"bootDiskSize"`
	// VM boot disk type. thin, zeroedthick, eagerzeroedthick
	BootDiskType *DiskType `pulumi:"bootDiskType"`
	// Boot type('efi' is boot uefi mode)
	BootFirmware *BootFirmwareType `pulumi:"bootFirmware"`
	// esxi diskstore for boot disk.
	DiskStore *string `pulumi:"diskStore"`
	// esxi vm id.
	Id *string `pulumi:"id"`
	// pass data to VM
	Info []KeyValuePair `pulumi:"info"`
	// The IP address reported by VMWare tools.
	IpAddress *string `pulumi:"ipAddress"`
	// VM memory size.
	MemSize *int `pulumi:"memSize"`
	// esxi vm name.
	Name *string `pulumi:"name"`
	// VM network interfaces.
	NetworkInterfaces []NetworkInterface `pulumi:"networkInterfaces"`
	// VM memory size.
	Notes *string `pulumi:"notes"`
	// VM number of virtual cpus.
	NumVCpus *int `pulumi:"numVCpus"`
	// VM OS type.
	Os *string `pulumi:"os"`
	// VM power state.
	Power *string `pulumi:"power"`
	// Resource pool name to place vm.
	ResourcePoolName *string `pulumi:"resourcePoolName"`
	// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
	ShutdownTimeout *int `pulumi:"shutdownTimeout"`
	// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
	StartupTimeout *int `pulumi:"startupTimeout"`
	// VM virtual disks.
	VirtualDisks []VMVirtualDisk `pulumi:"virtualDisks"`
	// VM Virtual HW version.
	VirtualHWVer *int `pulumi:"virtualHWVer"`
}

// Defaults sets the appropriate defaults for GetVirtualMachineByIdResult
func (val *GetVirtualMachineByIdResult) Defaults() *GetVirtualMachineByIdResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ShutdownTimeout) {
		shutdownTimeout_ := 600
		tmp.ShutdownTimeout = &shutdownTimeout_
	}
	if isZero(tmp.StartupTimeout) {
		startupTimeout_ := 600
		tmp.StartupTimeout = &startupTimeout_
	}
	return &tmp
}

func GetVirtualMachineByIdOutput(ctx *pulumi.Context, args GetVirtualMachineByIdOutputArgs, opts ...pulumi.InvokeOption) GetVirtualMachineByIdResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualMachineByIdResult, error) {
			args := v.(GetVirtualMachineByIdArgs)
			r, err := GetVirtualMachineById(ctx, &args, opts...)
			var s GetVirtualMachineByIdResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualMachineByIdResultOutput)
}

type GetVirtualMachineByIdOutputArgs struct {
	// Virtual Machine Id to get details of
	Id pulumi.StringInput `pulumi:"id"`
}

func (GetVirtualMachineByIdOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualMachineByIdArgs)(nil)).Elem()
}

type GetVirtualMachineByIdResultOutput struct{ *pulumi.OutputState }

func (GetVirtualMachineByIdResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualMachineByIdResult)(nil)).Elem()
}

func (o GetVirtualMachineByIdResultOutput) ToGetVirtualMachineByIdResultOutput() GetVirtualMachineByIdResultOutput {
	return o
}

func (o GetVirtualMachineByIdResultOutput) ToGetVirtualMachineByIdResultOutputWithContext(ctx context.Context) GetVirtualMachineByIdResultOutput {
	return o
}

// VM boot disk size. Will expand boot disk to this size.
func (o GetVirtualMachineByIdResultOutput) BootDiskSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *int { return v.BootDiskSize }).(pulumi.IntPtrOutput)
}

// VM boot disk type. thin, zeroedthick, eagerzeroedthick
func (o GetVirtualMachineByIdResultOutput) BootDiskType() DiskTypePtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *DiskType { return v.BootDiskType }).(DiskTypePtrOutput)
}

// Boot type('efi' is boot uefi mode)
func (o GetVirtualMachineByIdResultOutput) BootFirmware() BootFirmwareTypePtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *BootFirmwareType { return v.BootFirmware }).(BootFirmwareTypePtrOutput)
}

// esxi diskstore for boot disk.
func (o GetVirtualMachineByIdResultOutput) DiskStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *string { return v.DiskStore }).(pulumi.StringPtrOutput)
}

// esxi vm id.
func (o GetVirtualMachineByIdResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// pass data to VM
func (o GetVirtualMachineByIdResultOutput) Info() KeyValuePairArrayOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) []KeyValuePair { return v.Info }).(KeyValuePairArrayOutput)
}

// The IP address reported by VMWare tools.
func (o GetVirtualMachineByIdResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// VM memory size.
func (o GetVirtualMachineByIdResultOutput) MemSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *int { return v.MemSize }).(pulumi.IntPtrOutput)
}

// esxi vm name.
func (o GetVirtualMachineByIdResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// VM network interfaces.
func (o GetVirtualMachineByIdResultOutput) NetworkInterfaces() NetworkInterfaceArrayOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) []NetworkInterface { return v.NetworkInterfaces }).(NetworkInterfaceArrayOutput)
}

// VM memory size.
func (o GetVirtualMachineByIdResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// VM number of virtual cpus.
func (o GetVirtualMachineByIdResultOutput) NumVCpus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *int { return v.NumVCpus }).(pulumi.IntPtrOutput)
}

// VM OS type.
func (o GetVirtualMachineByIdResultOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *string { return v.Os }).(pulumi.StringPtrOutput)
}

// VM power state.
func (o GetVirtualMachineByIdResultOutput) Power() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *string { return v.Power }).(pulumi.StringPtrOutput)
}

// Resource pool name to place vm.
func (o GetVirtualMachineByIdResultOutput) ResourcePoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *string { return v.ResourcePoolName }).(pulumi.StringPtrOutput)
}

// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
func (o GetVirtualMachineByIdResultOutput) ShutdownTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *int { return v.ShutdownTimeout }).(pulumi.IntPtrOutput)
}

// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
func (o GetVirtualMachineByIdResultOutput) StartupTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *int { return v.StartupTimeout }).(pulumi.IntPtrOutput)
}

// VM virtual disks.
func (o GetVirtualMachineByIdResultOutput) VirtualDisks() VMVirtualDiskArrayOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) []VMVirtualDisk { return v.VirtualDisks }).(VMVirtualDiskArrayOutput)
}

// VM Virtual HW version.
func (o GetVirtualMachineByIdResultOutput) VirtualHWVer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetVirtualMachineByIdResult) *int { return v.VirtualHWVer }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualMachineByIdResultOutput{})
}