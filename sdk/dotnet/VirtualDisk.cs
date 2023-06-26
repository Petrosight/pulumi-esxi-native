// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EsxiNative
{
    [EsxiNativeResourceType("esxi-native:index:VirtualDisk")]
    public partial class VirtualDisk : Pulumi.CustomResource
    {
        /// <summary>
        /// Disk directory.
        /// </summary>
        [Output("directory")]
        public Output<string> Directory { get; private set; } = null!;

        /// <summary>
        /// Disk Store.
        /// </summary>
        [Output("diskStore")]
        public Output<string> DiskStore { get; private set; } = null!;

        /// <summary>
        /// Virtual Disk type. (thin, zeroedthick or eagerzeroedthick)
        /// </summary>
        [Output("diskType")]
        public Output<Pulumi.EsxiNative.DiskType> DiskType { get; private set; } = null!;

        /// <summary>
        /// Virtual Disk Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Virtual Disk size in GB.
        /// </summary>
        [Output("size")]
        public Output<int?> Size { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualDisk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualDisk(string name, VirtualDiskArgs args, CustomResourceOptions? options = null)
            : base("esxi-native:index:VirtualDisk", name, args ?? new VirtualDiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualDisk(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("esxi-native:index:VirtualDisk", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualDisk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualDisk Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualDisk(name, id, options);
        }
    }

    public sealed class VirtualDiskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disk directory.
        /// </summary>
        [Input("directory", required: true)]
        public Input<string> Directory { get; set; } = null!;

        /// <summary>
        /// Disk Store.
        /// </summary>
        [Input("diskStore", required: true)]
        public Input<string> DiskStore { get; set; } = null!;

        /// <summary>
        /// Virtual Disk type. (thin, zeroedthick or eagerzeroedthick)
        /// </summary>
        [Input("diskType", required: true)]
        public Input<Pulumi.EsxiNative.DiskType> DiskType { get; set; } = null!;

        /// <summary>
        /// Virtual Disk Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Virtual Disk size in GB.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        public VirtualDiskArgs()
        {
            Size = 1;
        }
    }
}