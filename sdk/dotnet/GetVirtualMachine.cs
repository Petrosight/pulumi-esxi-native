// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.EsxiNative
{
    public static class GetVirtualMachine
    {
        public static Task<GetVirtualMachineResult> InvokeAsync(GetVirtualMachineArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineResult>("esxi-native:index:getVirtualMachine", args ?? new GetVirtualMachineArgs(), options.WithDefaults());

        public static Output<GetVirtualMachineResult> Invoke(GetVirtualMachineInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualMachineResult>("esxi-native:index:getVirtualMachine", args ?? new GetVirtualMachineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualMachineArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Virtual Machine Name to get details of
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetVirtualMachineArgs()
        {
        }
        public static new GetVirtualMachineArgs Empty => new GetVirtualMachineArgs();
    }

    public sealed class GetVirtualMachineInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Virtual Machine Name to get details of
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetVirtualMachineInvokeArgs()
        {
        }
        public static new GetVirtualMachineInvokeArgs Empty => new GetVirtualMachineInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualMachineResult
    {
        /// <summary>
        /// VM boot disk size. Will expand boot disk to this size.
        /// </summary>
        public readonly int? BootDiskSize;
        /// <summary>
        /// VM boot disk type. thin, zeroedthick, eagerzeroedthick
        /// </summary>
        public readonly Pulumiverse.EsxiNative.DiskType? BootDiskType;
        /// <summary>
        /// Boot type('efi' is boot uefi mode)
        /// </summary>
        public readonly Pulumiverse.EsxiNative.BootFirmwareType? BootFirmware;
        /// <summary>
        /// esxi diskstore for boot disk.
        /// </summary>
        public readonly string? DiskStore;
        /// <summary>
        /// esxi vm id.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// pass data to VM
        /// </summary>
        public readonly ImmutableArray<Outputs.KeyValuePair> Info;
        /// <summary>
        /// The IP address reported by VMWare tools.
        /// </summary>
        public readonly string? IpAddress;
        /// <summary>
        /// VM memory size.
        /// </summary>
        public readonly int? MemSize;
        /// <summary>
        /// esxi vm name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// VM network interfaces.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterface> NetworkInterfaces;
        /// <summary>
        /// VM memory size.
        /// </summary>
        public readonly string? Notes;
        /// <summary>
        /// VM number of virtual cpus.
        /// </summary>
        public readonly int? NumVCpus;
        /// <summary>
        /// VM OS type.
        /// </summary>
        public readonly string? Os;
        /// <summary>
        /// VM power state.
        /// </summary>
        public readonly string? Power;
        /// <summary>
        /// Resource pool name to place vm.
        /// </summary>
        public readonly string? ResourcePoolName;
        /// <summary>
        /// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
        /// </summary>
        public readonly int? ShutdownTimeout;
        /// <summary>
        /// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
        /// </summary>
        public readonly int? StartupTimeout;
        /// <summary>
        /// VM virtual disks.
        /// </summary>
        public readonly ImmutableArray<Outputs.VMVirtualDisk> VirtualDisks;
        /// <summary>
        /// VM Virtual HW version.
        /// </summary>
        public readonly int? VirtualHWVer;

        [OutputConstructor]
        private GetVirtualMachineResult(
            int? bootDiskSize,

            Pulumiverse.EsxiNative.DiskType? bootDiskType,

            Pulumiverse.EsxiNative.BootFirmwareType? bootFirmware,

            string? diskStore,

            string? id,

            ImmutableArray<Outputs.KeyValuePair> info,

            string? ipAddress,

            int? memSize,

            string? name,

            ImmutableArray<Outputs.NetworkInterface> networkInterfaces,

            string? notes,

            int? numVCpus,

            string? os,

            string? power,

            string? resourcePoolName,

            int? shutdownTimeout,

            int? startupTimeout,

            ImmutableArray<Outputs.VMVirtualDisk> virtualDisks,

            int? virtualHWVer)
        {
            BootDiskSize = bootDiskSize;
            BootDiskType = bootDiskType;
            BootFirmware = bootFirmware;
            DiskStore = diskStore;
            Id = id;
            Info = info;
            IpAddress = ipAddress;
            MemSize = memSize;
            Name = name;
            NetworkInterfaces = networkInterfaces;
            Notes = notes;
            NumVCpus = numVCpus;
            Os = os;
            Power = power;
            ResourcePoolName = resourcePoolName;
            ShutdownTimeout = shutdownTimeout;
            StartupTimeout = startupTimeout;
            VirtualDisks = virtualDisks;
            VirtualHWVer = virtualHWVer;
        }
    }
}
