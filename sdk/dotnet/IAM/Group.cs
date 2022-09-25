// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Unifi.IAM
{
    /// <summary>
    /// `unifi.iam.Group` manages a user group (called "client group" in the UI), which can be used to limit bandwidth for groups of users.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Unifi = Pulumiverse.Unifi;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var wifi = new Unifi.IAM.Group("wifi", new()
    ///     {
    ///         QosRateMaxDown = 2000,
    ///         QosRateMaxUp = 10,
    ///     });
    /// 
    ///     // 10kbps
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// # import using the ID
    /// 
    /// ```sh
    ///  $ pulumi import unifi:iam/group:Group wifi 5fe6261995fe130013456a36
    /// ```
    /// </summary>
    [UnifiResourceType("unifi:iam/group:Group")]
    public partial class Group : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the user group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The QOS maximum download rate. Defaults to `-1`.
        /// </summary>
        [Output("qosRateMaxDown")]
        public Output<int?> QosRateMaxDown { get; private set; } = null!;

        /// <summary>
        /// The QOS maximum upload rate. Defaults to `-1`.
        /// </summary>
        [Output("qosRateMaxUp")]
        public Output<int?> QosRateMaxUp { get; private set; } = null!;

        /// <summary>
        /// The name of the site to associate the user group with.
        /// </summary>
        [Output("site")]
        public Output<string> Site { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs? args = null, CustomResourceOptions? options = null)
            : base("unifi:iam/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("unifi:iam/group:Group", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the user group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The QOS maximum download rate. Defaults to `-1`.
        /// </summary>
        [Input("qosRateMaxDown")]
        public Input<int>? QosRateMaxDown { get; set; }

        /// <summary>
        /// The QOS maximum upload rate. Defaults to `-1`.
        /// </summary>
        [Input("qosRateMaxUp")]
        public Input<int>? QosRateMaxUp { get; set; }

        /// <summary>
        /// The name of the site to associate the user group with.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        public GroupArgs()
        {
        }
        public static new GroupArgs Empty => new GroupArgs();
    }

    public sealed class GroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the user group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The QOS maximum download rate. Defaults to `-1`.
        /// </summary>
        [Input("qosRateMaxDown")]
        public Input<int>? QosRateMaxDown { get; set; }

        /// <summary>
        /// The QOS maximum upload rate. Defaults to `-1`.
        /// </summary>
        [Input("qosRateMaxUp")]
        public Input<int>? QosRateMaxUp { get; set; }

        /// <summary>
        /// The name of the site to associate the user group with.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        public GroupState()
        {
        }
        public static new GroupState Empty => new GroupState();
    }
}
