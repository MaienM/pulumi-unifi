// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Unifi
{
    public static class GetNetwork
    {
        /// <summary>
        /// `unifi.Network` data source can be used to retrieve settings for a network by name or ID.
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
        ///     var lanNetwork = Unifi.GetNetwork.Invoke(new()
        ///     {
        ///         Name = "LAN",
        ///     });
        /// 
        ///     var myDevice = Unifi.IAM.GetUser.Invoke(new()
        ///     {
        ///         Mac = "01:23:45:67:89:ab",
        ///     });
        /// 
        ///     var myNetwork = Unifi.GetNetwork.Invoke(new()
        ///     {
        ///         Id = myDevice.Apply(getUserResult =&gt; getUserResult.NetworkId),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNetworkResult> InvokeAsync(GetNetworkArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkResult>("unifi:index/getNetwork:getNetwork", args ?? new GetNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// `unifi.Network` data source can be used to retrieve settings for a network by name or ID.
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
        ///     var lanNetwork = Unifi.GetNetwork.Invoke(new()
        ///     {
        ///         Name = "LAN",
        ///     });
        /// 
        ///     var myDevice = Unifi.IAM.GetUser.Invoke(new()
        ///     {
        ///         Mac = "01:23:45:67:89:ab",
        ///     });
        /// 
        ///     var myNetwork = Unifi.GetNetwork.Invoke(new()
        ///     {
        ///         Id = myDevice.Apply(getUserResult =&gt; getUserResult.NetworkId),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNetworkResult> Invoke(GetNetworkInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkResult>("unifi:index/getNetwork:getNetwork", args ?? new GetNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the network.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the network.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The name of the site to associate the network with.
        /// </summary>
        [Input("site")]
        public string? Site { get; set; }

        public GetNetworkArgs()
        {
        }
        public static new GetNetworkArgs Empty => new GetNetworkArgs();
    }

    public sealed class GetNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the network.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the site to associate the network with.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        public GetNetworkInvokeArgs()
        {
        }
        public static new GetNetworkInvokeArgs Empty => new GetNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkResult
    {
        /// <summary>
        /// IPv4 addresses for the DNS server to be returned from the DHCP server.
        /// </summary>
        public readonly ImmutableArray<string> DhcpDns;
        /// <summary>
        /// whether DHCP is enabled or not on this network.
        /// </summary>
        public readonly bool DhcpEnabled;
        /// <summary>
        /// lease time for DHCP addresses.
        /// </summary>
        public readonly int DhcpLease;
        /// <summary>
        /// The IPv4 address where the DHCP range of addresses starts.
        /// </summary>
        public readonly string DhcpStart;
        /// <summary>
        /// The IPv4 address where the DHCP range of addresses stops.
        /// </summary>
        public readonly string DhcpStop;
        /// <summary>
        /// Toggles on the DHCP boot options. will be set to true if you have dhcpd*boot*filename, and dhcpd*boot*server set.
        /// </summary>
        public readonly bool DhcpdBootEnabled;
        /// <summary>
        /// the file to PXE boot from on the dhcpd*boot*server.
        /// </summary>
        public readonly string DhcpdBootFilename;
        /// <summary>
        /// IPv4 address of a TFTP server to network boot from.
        /// </summary>
        public readonly string DhcpdBootServer;
        /// <summary>
        /// The domain name of this network.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// The ID of the network.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies whether IGMP snooping is enabled or not.
        /// </summary>
        public readonly bool IgmpSnooping;
        /// <summary>
        /// Specifies which type of IPv6 connection to use.
        /// </summary>
        public readonly string Ipv6InterfaceType;
        /// <summary>
        /// Specifies which WAN interface is used for IPv6 Prefix Delegation.
        /// </summary>
        public readonly string Ipv6PdInterface;
        /// <summary>
        /// Specifies the IPv6 Prefix ID.
        /// </summary>
        public readonly string Ipv6PdPrefixid;
        /// <summary>
        /// Specifies whether to enable router advertisements or not.
        /// </summary>
        public readonly bool Ipv6RaEnable;
        /// <summary>
        /// Specifies the static IPv6 subnet (when ipv6*interface*type is 'static').
        /// </summary>
        public readonly string Ipv6StaticSubnet;
        /// <summary>
        /// The name of the network.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The group of the network.
        /// </summary>
        public readonly string NetworkGroup;
        /// <summary>
        /// The purpose of the network. One of `corporate`, `guest`, `wan`, or `vlan-only`.
        /// </summary>
        public readonly string Purpose;
        /// <summary>
        /// The name of the site to associate the network with.
        /// </summary>
        public readonly string Site;
        /// <summary>
        /// The subnet of the network (CIDR address).
        /// </summary>
        public readonly string Subnet;
        /// <summary>
        /// The VLAN ID of the network.
        /// </summary>
        public readonly int VlanId;
        /// <summary>
        /// DNS servers IPs of the WAN.
        /// </summary>
        public readonly ImmutableArray<string> WanDns;
        /// <summary>
        /// Specifies the WAN egress quality of service.
        /// </summary>
        public readonly int WanEgressQos;
        /// <summary>
        /// The IPv4 gateway of the WAN.
        /// </summary>
        public readonly string WanGateway;
        /// <summary>
        /// The IPv4 address of the WAN.
        /// </summary>
        public readonly string WanIp;
        /// <summary>
        /// The IPv4 netmask of the WAN.
        /// </summary>
        public readonly string WanNetmask;
        /// <summary>
        /// Specifies the WAN network group. One of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
        /// </summary>
        public readonly string WanNetworkgroup;
        /// <summary>
        /// Specifies the IPV4 WAN connection type. One of either `disabled`, `static`, `dhcp`, or `pppoe`.
        /// </summary>
        public readonly string WanType;
        /// <summary>
        /// Specifies the IPV4 WAN username.
        /// </summary>
        public readonly string WanUsername;
        /// <summary>
        /// Specifies the IPV4 WAN password.
        /// </summary>
        public readonly string XWanPassword;

        [OutputConstructor]
        private GetNetworkResult(
            ImmutableArray<string> dhcpDns,

            bool dhcpEnabled,

            int dhcpLease,

            string dhcpStart,

            string dhcpStop,

            bool dhcpdBootEnabled,

            string dhcpdBootFilename,

            string dhcpdBootServer,

            string domainName,

            string id,

            bool igmpSnooping,

            string ipv6InterfaceType,

            string ipv6PdInterface,

            string ipv6PdPrefixid,

            bool ipv6RaEnable,

            string ipv6StaticSubnet,

            string name,

            string networkGroup,

            string purpose,

            string site,

            string subnet,

            int vlanId,

            ImmutableArray<string> wanDns,

            int wanEgressQos,

            string wanGateway,

            string wanIp,

            string wanNetmask,

            string wanNetworkgroup,

            string wanType,

            string wanUsername,

            string xWanPassword)
        {
            DhcpDns = dhcpDns;
            DhcpEnabled = dhcpEnabled;
            DhcpLease = dhcpLease;
            DhcpStart = dhcpStart;
            DhcpStop = dhcpStop;
            DhcpdBootEnabled = dhcpdBootEnabled;
            DhcpdBootFilename = dhcpdBootFilename;
            DhcpdBootServer = dhcpdBootServer;
            DomainName = domainName;
            Id = id;
            IgmpSnooping = igmpSnooping;
            Ipv6InterfaceType = ipv6InterfaceType;
            Ipv6PdInterface = ipv6PdInterface;
            Ipv6PdPrefixid = ipv6PdPrefixid;
            Ipv6RaEnable = ipv6RaEnable;
            Ipv6StaticSubnet = ipv6StaticSubnet;
            Name = name;
            NetworkGroup = networkGroup;
            Purpose = purpose;
            Site = site;
            Subnet = subnet;
            VlanId = vlanId;
            WanDns = wanDns;
            WanEgressQos = wanEgressQos;
            WanGateway = wanGateway;
            WanIp = wanIp;
            WanNetmask = wanNetmask;
            WanNetworkgroup = wanNetworkgroup;
            WanType = wanType;
            WanUsername = wanUsername;
            XWanPassword = xWanPassword;
        }
    }
}
