// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EsxiNative.Inputs
{

    public sealed class VMVirtualDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SCSI_Ctrl:SCSI_id.    Range  '0:1' to '0:15'.   SCSI_id 7 is not allowed.
        /// </summary>
        [Input("slot")]
        public Input<string>? Slot { get; set; }

        [Input("virtualDiskId")]
        public Input<string>? VirtualDiskId { get; set; }

        public VMVirtualDiskArgs()
        {
        }
        public static new VMVirtualDiskArgs Empty => new VMVirtualDiskArgs();
    }
}
