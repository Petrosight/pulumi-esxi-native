package esxi

type KeyValuePair struct {
	Key   string
	Value string
}

type NetworkInterface struct {
	MacAddress     string
	NicType        string
	VirtualNetwork string
}

type PortGroup struct {
	// Forged transmits (true=Accept/false=Reject).
	ForgedTransmits bool
	// MAC address changes (true=Accept/false=Reject).
	MacChanges bool
	// Port Group Virtual Switch.
	VSwitch string
	// Port Group Virtual LAN.
	Vlan int
	// Virtual Switch name.
	Name string
	// Virtual Switch number of ports. (1-4096)
	Ports int
	// Promiscuous mode (true=Accept/false=Reject).
	PromiscuousMode bool
}

type ResourcePool struct {
	// CPU maximum (in MHz).
	CpuMax int
	// CPU minimum (in MHz).
	CpuMin int
	// Can pool borrow CPU resources from parent?
	CpuMinExpandable string
	// CPU shares (low/normal/high/<custom>).
	CpuShares string
	// Memory maximum (in MB).
	MemMax int
	// Memory minimum (in MB).
	MemMin int
	// Can pool borrow memory resources from parent?
	MemMinExpandable string
	// Memory shares (low/normal/high/<custom>).
	MemShares string
	// Resource Pool Name
	Name string
}

type Uplink struct {
	// Uplink name.
	Name string
}

type VMVirtualDisk struct {
	// SCSI_Ctrl:SCSI_id.    Range  '0:1' to '0:15'.   SCSI_id 7 is not allowed.
	Slot          string
	VirtualDiskId string
}

type VirtualDisk struct {
	// Disk directory.
	Directory string
	// Disk Store.
	DiskStore string
	// Virtual Disk type.
	DiskType string
	// Virtual Disk Name.
	Name string
	// Virtual Disk size in GB.
	Size int
}

type VirtualMachine struct {
	// VM boot disk size. Will expand boot disk to this size.
	BootDiskSize string
	// VM boot disk type.
	BootDiskType string
	// Boot type('efi' is boot uefi mode)
	BootFirmware string
	// Source vm path on esxi host to clone.
	CloneFromVirtualMachine string
	// esxi DiskStore for boot disk.
	DiskStore string
	// pass data to VM
	Id string
	// pass data to VM
	Info []KeyValuePair
	// The IP address reported by VMWare tools.
	IpAddress string
	// VM memory size.
	MemSize string
	// esxi vm name.
	Name string
	// VM network interfaces.
	NetworkInterfaces []NetworkInterface
	// VM memory size.
	Notes string
	// VM number of virtual cpus.
	NumVCpus string
	// VM OS type.
	Os string
	// Path on esxi host of ovf files.
	OvfHostPathSource string
	// VM OVF properties.
	OvfProperties []KeyValuePair
	// The amount of time, in seconds, to wait for the guest to boot and run ovfProperties. (0-6000)
	OvfPropertiesTimer int
	// Local path to source ovf files.
	OvfSourceLocalPath string
	// VM power state.
	Power string
	// Resource pool name to place vm.
	ResourcePoolName string
	// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
	ShutdownTimeout int
	// The amount of vm uptime, in seconds, to wait for an available IP address on this virtual machine.
	StartupTimeout int
	// VM virtual disks.
	VirtualDisks []VMVirtualDisk
	// VM Virtual HW version.
	VirtualHWVer string
}

type VirtualSwitch struct {
	// Forged transmits (true=Accept/false=Reject).
	ForgedTransmits bool
	// Virtual Switch Link Discovery Mode.
	LinkDiscoveryMode string
	// MAC address changes (true=Accept/false=Reject).
	MacChanges bool
	// Virtual Switch mtu. (1280-9000)
	Mtu int
	// Virtual Switch name.
	Name string
	// Virtual Switch number of ports. (1-4096)
	Ports int
	// Promiscuous mode (true=Accept/false=Reject).
	PromiscuousMode bool
	// Uplink configuration.
	Uplinks []Uplink
}
