// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package unifi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `Network` data source can be used to retrieve settings for a network by name or ID.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi"
//	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/iam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := unifi.LookupNetwork(ctx, &unifi.LookupNetworkArgs{
//				Name: pulumi.StringRef("Default"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			myDevice, err := iam.LookupUser(ctx, &iam.LookupUserArgs{
//				Mac: "01:23:45:67:89:ab",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = unifi.LookupNetwork(ctx, &unifi.LookupNetworkArgs{
//				Id: pulumi.StringRef(myDevice.NetworkId),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNetwork(ctx *pulumi.Context, args *LookupNetworkArgs, opts ...pulumi.InvokeOption) (*LookupNetworkResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupNetworkResult
	err := ctx.Invoke("unifi:index/getNetwork:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkArgs struct {
	// The ID of the network.
	Id *string `pulumi:"id"`
	// The name of the network.
	Name *string `pulumi:"name"`
	// The name of the site to associate the network with.
	Site *string `pulumi:"site"`
}

// A collection of values returned by getNetwork.
type LookupNetworkResult struct {
	// IPv4 addresses for the DNS server to be returned from the DHCP server.
	DhcpDns []string `pulumi:"dhcpDns"`
	// whether DHCP is enabled or not on this network.
	DhcpEnabled bool `pulumi:"dhcpEnabled"`
	// lease time for DHCP addresses.
	DhcpLease int `pulumi:"dhcpLease"`
	// The IPv4 address where the DHCP range of addresses starts.
	DhcpStart string `pulumi:"dhcpStart"`
	// The IPv4 address where the DHCP range of addresses stops.
	DhcpStop string `pulumi:"dhcpStop"`
	// Specifies the IPv6 addresses for the DNS server to be returned from the DHCP server. Used if `dhcpV6DnsAuto` is set to `false`.
	DhcpV6Dns []string `pulumi:"dhcpV6Dns"`
	// Specifies DNS source to propagate. If set `false` the entries in `dhcpV6Dns` are used, the upstream entries otherwise
	DhcpV6DnsAuto bool `pulumi:"dhcpV6DnsAuto"`
	// Enable stateful DHCPv6 for static configuration.
	DhcpV6Enabled bool `pulumi:"dhcpV6Enabled"`
	// Specifies the lease time for DHCPv6 addresses.
	DhcpV6Lease int `pulumi:"dhcpV6Lease"`
	// Start address of the DHCPv6 range. Used in static DHCPv6 configuration.
	DhcpV6Start string `pulumi:"dhcpV6Start"`
	// End address of the DHCPv6 range. Used in static DHCPv6 configuration.
	DhcpV6Stop string `pulumi:"dhcpV6Stop"`
	// Toggles on the DHCP boot options. will be set to true if you have dhcpd*boot*filename, and dhcpd*boot*server set.
	DhcpdBootEnabled bool `pulumi:"dhcpdBootEnabled"`
	// the file to PXE boot from on the dhcpd*boot*server.
	DhcpdBootFilename string `pulumi:"dhcpdBootFilename"`
	// IPv4 address of a TFTP server to network boot from.
	DhcpdBootServer string `pulumi:"dhcpdBootServer"`
	// The domain name of this network.
	DomainName string `pulumi:"domainName"`
	// The ID of the network.
	Id string `pulumi:"id"`
	// Specifies whether IGMP snooping is enabled or not.
	IgmpSnooping bool `pulumi:"igmpSnooping"`
	// Specifies which type of IPv6 connection to use. Must be one of either `static`, `pd`, or `none`.
	Ipv6InterfaceType string `pulumi:"ipv6InterfaceType"`
	// Specifies which WAN interface to use for IPv6 PD. Must be one of either `wan` or `wan2`.
	Ipv6PdInterface string `pulumi:"ipv6PdInterface"`
	// Specifies the IPv6 Prefix ID.
	Ipv6PdPrefixid string `pulumi:"ipv6PdPrefixid"`
	// Start address of the DHCPv6 range. Used if `ipv6InterfaceType` is set to `pd`.
	Ipv6PdStart string `pulumi:"ipv6PdStart"`
	// End address of the DHCPv6 range. Used if `ipv6InterfaceType` is set to `pd`.
	Ipv6PdStop string `pulumi:"ipv6PdStop"`
	// Specifies whether to enable router advertisements or not.
	Ipv6RaEnable bool `pulumi:"ipv6RaEnable"`
	// Lifetime in which the address can be used. Address becomes deprecated afterwards. Must be lower than or equal to `ipv6RaValidLifetime`
	Ipv6RaPreferredLifetime int `pulumi:"ipv6RaPreferredLifetime"`
	// IPv6 router advertisement priority. Must be one of either `high`, `medium`, or `low`
	Ipv6RaPriority string `pulumi:"ipv6RaPriority"`
	// Total lifetime in which the address can be used. Must be equal to or greater than `ipv6RaPreferredLifetime`.
	Ipv6RaValidLifetime int `pulumi:"ipv6RaValidLifetime"`
	// Specifies the static IPv6 subnet (when ipv6*interface*type is 'static').
	Ipv6StaticSubnet string `pulumi:"ipv6StaticSubnet"`
	// Specifies whether Multicast DNS (mDNS) is enabled or not on the network (Controller >=v7).
	MulticastDns bool `pulumi:"multicastDns"`
	// The name of the network.
	Name string `pulumi:"name"`
	// The group of the network.
	NetworkGroup string `pulumi:"networkGroup"`
	// The purpose of the network. One of `corporate`, `guest`, `wan`, or `vlan-only`.
	Purpose string `pulumi:"purpose"`
	// The name of the site to associate the network with.
	Site string `pulumi:"site"`
	// The subnet of the network (CIDR address).
	Subnet string `pulumi:"subnet"`
	// The VLAN ID of the network.
	VlanId int `pulumi:"vlanId"`
	// Specifies the IPv6 prefix size to request from ISP. Must be a number between 48 and 64.
	WanDhcpV6PdSize int `pulumi:"wanDhcpV6PdSize"`
	// DNS servers IPs of the WAN.
	WanDns []string `pulumi:"wanDns"`
	// Specifies the WAN egress quality of service.
	WanEgressQos int `pulumi:"wanEgressQos"`
	// The IPv4 gateway of the WAN.
	WanGateway string `pulumi:"wanGateway"`
	// The IPv6 gateway of the WAN.
	WanGatewayV6 string `pulumi:"wanGatewayV6"`
	// The IPv4 address of the WAN.
	WanIp string `pulumi:"wanIp"`
	// The IPv6 address of the WAN.
	WanIpv6 string `pulumi:"wanIpv6"`
	// The IPv4 netmask of the WAN.
	WanNetmask string `pulumi:"wanNetmask"`
	// Specifies the WAN network group. One of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
	WanNetworkgroup string `pulumi:"wanNetworkgroup"`
	// The IPv6 prefix length of the WAN. Must be between 1 and 128.
	WanPrefixlen int `pulumi:"wanPrefixlen"`
	// Specifies the IPV4 WAN connection type. One of either `disabled`, `static`, `dhcp`, or `pppoe`.
	WanType string `pulumi:"wanType"`
	// Specifies the IPV6 WAN connection type. Must be one of either `disabled`, `static`, or `dhcpv6`.
	WanTypeV6 string `pulumi:"wanTypeV6"`
	// Specifies the IPV4 WAN username.
	WanUsername string `pulumi:"wanUsername"`
	// Specifies the IPV4 WAN password.
	XWanPassword string `pulumi:"xWanPassword"`
}

func LookupNetworkOutput(ctx *pulumi.Context, args LookupNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkResult, error) {
			args := v.(LookupNetworkArgs)
			r, err := LookupNetwork(ctx, &args, opts...)
			var s LookupNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkResultOutput)
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkOutputArgs struct {
	// The ID of the network.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the network.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The name of the site to associate the network with.
	Site pulumi.StringPtrInput `pulumi:"site"`
}

func (LookupNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getNetwork.
type LookupNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkResult)(nil)).Elem()
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutput() LookupNetworkResultOutput {
	return o
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutputWithContext(ctx context.Context) LookupNetworkResultOutput {
	return o
}

// IPv4 addresses for the DNS server to be returned from the DHCP server.
func (o LookupNetworkResultOutput) DhcpDns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkResult) []string { return v.DhcpDns }).(pulumi.StringArrayOutput)
}

// whether DHCP is enabled or not on this network.
func (o LookupNetworkResultOutput) DhcpEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkResult) bool { return v.DhcpEnabled }).(pulumi.BoolOutput)
}

// lease time for DHCP addresses.
func (o LookupNetworkResultOutput) DhcpLease() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkResult) int { return v.DhcpLease }).(pulumi.IntOutput)
}

// The IPv4 address where the DHCP range of addresses starts.
func (o LookupNetworkResultOutput) DhcpStart() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.DhcpStart }).(pulumi.StringOutput)
}

// The IPv4 address where the DHCP range of addresses stops.
func (o LookupNetworkResultOutput) DhcpStop() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.DhcpStop }).(pulumi.StringOutput)
}

// Specifies the IPv6 addresses for the DNS server to be returned from the DHCP server. Used if `dhcpV6DnsAuto` is set to `false`.
func (o LookupNetworkResultOutput) DhcpV6Dns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkResult) []string { return v.DhcpV6Dns }).(pulumi.StringArrayOutput)
}

// Specifies DNS source to propagate. If set `false` the entries in `dhcpV6Dns` are used, the upstream entries otherwise
func (o LookupNetworkResultOutput) DhcpV6DnsAuto() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkResult) bool { return v.DhcpV6DnsAuto }).(pulumi.BoolOutput)
}

// Enable stateful DHCPv6 for static configuration.
func (o LookupNetworkResultOutput) DhcpV6Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkResult) bool { return v.DhcpV6Enabled }).(pulumi.BoolOutput)
}

// Specifies the lease time for DHCPv6 addresses.
func (o LookupNetworkResultOutput) DhcpV6Lease() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkResult) int { return v.DhcpV6Lease }).(pulumi.IntOutput)
}

// Start address of the DHCPv6 range. Used in static DHCPv6 configuration.
func (o LookupNetworkResultOutput) DhcpV6Start() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.DhcpV6Start }).(pulumi.StringOutput)
}

// End address of the DHCPv6 range. Used in static DHCPv6 configuration.
func (o LookupNetworkResultOutput) DhcpV6Stop() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.DhcpV6Stop }).(pulumi.StringOutput)
}

// Toggles on the DHCP boot options. will be set to true if you have dhcpd*boot*filename, and dhcpd*boot*server set.
func (o LookupNetworkResultOutput) DhcpdBootEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkResult) bool { return v.DhcpdBootEnabled }).(pulumi.BoolOutput)
}

// the file to PXE boot from on the dhcpd*boot*server.
func (o LookupNetworkResultOutput) DhcpdBootFilename() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.DhcpdBootFilename }).(pulumi.StringOutput)
}

// IPv4 address of a TFTP server to network boot from.
func (o LookupNetworkResultOutput) DhcpdBootServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.DhcpdBootServer }).(pulumi.StringOutput)
}

// The domain name of this network.
func (o LookupNetworkResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.DomainName }).(pulumi.StringOutput)
}

// The ID of the network.
func (o LookupNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies whether IGMP snooping is enabled or not.
func (o LookupNetworkResultOutput) IgmpSnooping() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkResult) bool { return v.IgmpSnooping }).(pulumi.BoolOutput)
}

// Specifies which type of IPv6 connection to use. Must be one of either `static`, `pd`, or `none`.
func (o LookupNetworkResultOutput) Ipv6InterfaceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Ipv6InterfaceType }).(pulumi.StringOutput)
}

// Specifies which WAN interface to use for IPv6 PD. Must be one of either `wan` or `wan2`.
func (o LookupNetworkResultOutput) Ipv6PdInterface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Ipv6PdInterface }).(pulumi.StringOutput)
}

// Specifies the IPv6 Prefix ID.
func (o LookupNetworkResultOutput) Ipv6PdPrefixid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Ipv6PdPrefixid }).(pulumi.StringOutput)
}

// Start address of the DHCPv6 range. Used if `ipv6InterfaceType` is set to `pd`.
func (o LookupNetworkResultOutput) Ipv6PdStart() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Ipv6PdStart }).(pulumi.StringOutput)
}

// End address of the DHCPv6 range. Used if `ipv6InterfaceType` is set to `pd`.
func (o LookupNetworkResultOutput) Ipv6PdStop() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Ipv6PdStop }).(pulumi.StringOutput)
}

// Specifies whether to enable router advertisements or not.
func (o LookupNetworkResultOutput) Ipv6RaEnable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkResult) bool { return v.Ipv6RaEnable }).(pulumi.BoolOutput)
}

// Lifetime in which the address can be used. Address becomes deprecated afterwards. Must be lower than or equal to `ipv6RaValidLifetime`
func (o LookupNetworkResultOutput) Ipv6RaPreferredLifetime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkResult) int { return v.Ipv6RaPreferredLifetime }).(pulumi.IntOutput)
}

// IPv6 router advertisement priority. Must be one of either `high`, `medium`, or `low`
func (o LookupNetworkResultOutput) Ipv6RaPriority() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Ipv6RaPriority }).(pulumi.StringOutput)
}

// Total lifetime in which the address can be used. Must be equal to or greater than `ipv6RaPreferredLifetime`.
func (o LookupNetworkResultOutput) Ipv6RaValidLifetime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkResult) int { return v.Ipv6RaValidLifetime }).(pulumi.IntOutput)
}

// Specifies the static IPv6 subnet (when ipv6*interface*type is 'static').
func (o LookupNetworkResultOutput) Ipv6StaticSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Ipv6StaticSubnet }).(pulumi.StringOutput)
}

// Specifies whether Multicast DNS (mDNS) is enabled or not on the network (Controller >=v7).
func (o LookupNetworkResultOutput) MulticastDns() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkResult) bool { return v.MulticastDns }).(pulumi.BoolOutput)
}

// The name of the network.
func (o LookupNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The group of the network.
func (o LookupNetworkResultOutput) NetworkGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.NetworkGroup }).(pulumi.StringOutput)
}

// The purpose of the network. One of `corporate`, `guest`, `wan`, or `vlan-only`.
func (o LookupNetworkResultOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Purpose }).(pulumi.StringOutput)
}

// The name of the site to associate the network with.
func (o LookupNetworkResultOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Site }).(pulumi.StringOutput)
}

// The subnet of the network (CIDR address).
func (o LookupNetworkResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Subnet }).(pulumi.StringOutput)
}

// The VLAN ID of the network.
func (o LookupNetworkResultOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkResult) int { return v.VlanId }).(pulumi.IntOutput)
}

// Specifies the IPv6 prefix size to request from ISP. Must be a number between 48 and 64.
func (o LookupNetworkResultOutput) WanDhcpV6PdSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkResult) int { return v.WanDhcpV6PdSize }).(pulumi.IntOutput)
}

// DNS servers IPs of the WAN.
func (o LookupNetworkResultOutput) WanDns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkResult) []string { return v.WanDns }).(pulumi.StringArrayOutput)
}

// Specifies the WAN egress quality of service.
func (o LookupNetworkResultOutput) WanEgressQos() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkResult) int { return v.WanEgressQos }).(pulumi.IntOutput)
}

// The IPv4 gateway of the WAN.
func (o LookupNetworkResultOutput) WanGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.WanGateway }).(pulumi.StringOutput)
}

// The IPv6 gateway of the WAN.
func (o LookupNetworkResultOutput) WanGatewayV6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.WanGatewayV6 }).(pulumi.StringOutput)
}

// The IPv4 address of the WAN.
func (o LookupNetworkResultOutput) WanIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.WanIp }).(pulumi.StringOutput)
}

// The IPv6 address of the WAN.
func (o LookupNetworkResultOutput) WanIpv6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.WanIpv6 }).(pulumi.StringOutput)
}

// The IPv4 netmask of the WAN.
func (o LookupNetworkResultOutput) WanNetmask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.WanNetmask }).(pulumi.StringOutput)
}

// Specifies the WAN network group. One of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
func (o LookupNetworkResultOutput) WanNetworkgroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.WanNetworkgroup }).(pulumi.StringOutput)
}

// The IPv6 prefix length of the WAN. Must be between 1 and 128.
func (o LookupNetworkResultOutput) WanPrefixlen() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkResult) int { return v.WanPrefixlen }).(pulumi.IntOutput)
}

// Specifies the IPV4 WAN connection type. One of either `disabled`, `static`, `dhcp`, or `pppoe`.
func (o LookupNetworkResultOutput) WanType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.WanType }).(pulumi.StringOutput)
}

// Specifies the IPV6 WAN connection type. Must be one of either `disabled`, `static`, or `dhcpv6`.
func (o LookupNetworkResultOutput) WanTypeV6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.WanTypeV6 }).(pulumi.StringOutput)
}

// Specifies the IPV4 WAN username.
func (o LookupNetworkResultOutput) WanUsername() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.WanUsername }).(pulumi.StringOutput)
}

// Specifies the IPV4 WAN password.
func (o LookupNetworkResultOutput) XWanPassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.XWanPassword }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkResultOutput{})
}
