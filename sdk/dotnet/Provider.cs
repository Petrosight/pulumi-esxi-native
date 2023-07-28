// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EsxiNative
{
    /// <summary>
    /// The provider type for the ESXi native package. By default, resources use package-wide configuration settings, however an explicit `Provider` instance may be created and passed during resource construction to achieve fine-grained programmatic control over provider settings. See the [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [EsxiNativeResourceType("pulumi:providers:esxi-native")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// ESXi Host Name config
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// ESXi Password config
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// ESXi Host SSH Port config
        /// </summary>
        [Output("sshPort")]
        public Output<string?> SshPort { get; private set; } = null!;

        /// <summary>
        /// ESXi Host SSL Port config
        /// </summary>
        [Output("sslPort")]
        public Output<string?> SslPort { get; private set; } = null!;

        /// <summary>
        /// ESXi Username config
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs args, CustomResourceOptions? options = null)
            : base("esxi-native", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
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
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ESXi Host Name config
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// ESXi Password config
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// ESXi Host SSH Port config
        /// </summary>
        [Input("sshPort")]
        public Input<string>? SshPort { get; set; }

        /// <summary>
        /// ESXi Host SSL Port config
        /// </summary>
        [Input("sslPort")]
        public Input<string>? SslPort { get; set; }

        /// <summary>
        /// ESXi Username config
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ProviderArgs()
        {
            SshPort = "22";
            SslPort = "443";
            Username = "root";
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
