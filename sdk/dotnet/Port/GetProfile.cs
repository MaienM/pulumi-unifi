// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Unifi.Port
{
    public static class GetProfile
    {
        /// <summary>
        /// `unifi.port.Profile` data source can be used to retrieve the ID for a port profile by name.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Unifi = Pulumi.Unifi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var all = Unifi.Port.GetProfile.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProfileResult> InvokeAsync(GetProfileArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProfileResult>("unifi:port/getProfile:getProfile", args ?? new GetProfileArgs(), options.WithDefaults());

        /// <summary>
        /// `unifi.port.Profile` data source can be used to retrieve the ID for a port profile by name.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Unifi = Pulumi.Unifi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var all = Unifi.Port.GetProfile.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProfileResult> Invoke(GetProfileInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProfileResult>("unifi:port/getProfile:getProfile", args ?? new GetProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the port profile to look up. Defaults to `All`.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The name of the site the port profile is associated with.
        /// </summary>
        [Input("site")]
        public string? Site { get; set; }

        public GetProfileArgs()
        {
        }
        public static new GetProfileArgs Empty => new GetProfileArgs();
    }

    public sealed class GetProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the port profile to look up. Defaults to `All`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the site the port profile is associated with.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        public GetProfileInvokeArgs()
        {
        }
        public static new GetProfileInvokeArgs Empty => new GetProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetProfileResult
    {
        /// <summary>
        /// The ID of this port profile.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the port profile to look up. Defaults to `All`.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The name of the site the port profile is associated with.
        /// </summary>
        public readonly string Site;

        [OutputConstructor]
        private GetProfileResult(
            string id,

            string? name,

            string site)
        {
            Id = id;
            Name = name;
            Site = site;
        }
    }
}
