// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EsxiNative.Inputs
{

    public sealed class UplinkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Uplink name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public UplinkArgs()
        {
        }
        public static new UplinkArgs Empty => new UplinkArgs();
    }
}
