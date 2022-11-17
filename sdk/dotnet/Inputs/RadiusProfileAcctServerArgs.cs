// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Unifi.Inputs
{

    public sealed class RadiusProfileAcctServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address of accounting service server.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// Port of accounting service. Defaults to `1813`.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("xsecret", required: true)]
        private Input<string>? _xsecret;

        /// <summary>
        /// RADIUS secret.
        /// </summary>
        public Input<string>? Xsecret
        {
            get => _xsecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _xsecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public RadiusProfileAcctServerArgs()
        {
        }
        public static new RadiusProfileAcctServerArgs Empty => new RadiusProfileAcctServerArgs();
    }
}
