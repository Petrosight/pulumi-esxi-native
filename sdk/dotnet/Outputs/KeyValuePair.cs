// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EsxiNative.Outputs
{

    [OutputType]
    public sealed class KeyValuePair
    {
        public readonly string? Key;
        public readonly string? Value;

        [OutputConstructor]
        private KeyValuePair(
            string? key,

            string? value)
        {
            Key = key;
            Value = value;
        }
    }
}