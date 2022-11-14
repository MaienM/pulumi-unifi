// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.unifi.firewall;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final RuleArgs Empty = new RuleArgs();

    /**
     * The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
     * 
     */
    @Import(name="action", required=true)
    private Output<String> action;

    /**
     * @return The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
     * 
     */
    public Output<String> action() {
        return this.action;
    }

    /**
     * The destination address of the firewall rule.
     * 
     */
    @Import(name="dstAddress")
    private @Nullable Output<String> dstAddress;

    /**
     * @return The destination address of the firewall rule.
     * 
     */
    public Optional<Output<String>> dstAddress() {
        return Optional.ofNullable(this.dstAddress);
    }

    /**
     * The IPv6 destination address of the firewall rule.
     * 
     */
    @Import(name="dstAddressIpv6")
    private @Nullable Output<String> dstAddressIpv6;

    /**
     * @return The IPv6 destination address of the firewall rule.
     * 
     */
    public Optional<Output<String>> dstAddressIpv6() {
        return Optional.ofNullable(this.dstAddressIpv6);
    }

    /**
     * The destination firewall group IDs of the firewall rule.
     * 
     */
    @Import(name="dstFirewallGroupIds")
    private @Nullable Output<List<String>> dstFirewallGroupIds;

    /**
     * @return The destination firewall group IDs of the firewall rule.
     * 
     */
    public Optional<Output<List<String>>> dstFirewallGroupIds() {
        return Optional.ofNullable(this.dstFirewallGroupIds);
    }

    /**
     * The destination network ID of the firewall rule.
     * 
     */
    @Import(name="dstNetworkId")
    private @Nullable Output<String> dstNetworkId;

    /**
     * @return The destination network ID of the firewall rule.
     * 
     */
    public Optional<Output<String>> dstNetworkId() {
        return Optional.ofNullable(this.dstNetworkId);
    }

    /**
     * The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     * 
     */
    @Import(name="dstNetworkType")
    private @Nullable Output<String> dstNetworkType;

    /**
     * @return The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     * 
     */
    public Optional<Output<String>> dstNetworkType() {
        return Optional.ofNullable(this.dstNetworkType);
    }

    /**
     * The destination port of the firewall rule.
     * 
     */
    @Import(name="dstPort")
    private @Nullable Output<String> dstPort;

    /**
     * @return The destination port of the firewall rule.
     * 
     */
    public Optional<Output<String>> dstPort() {
        return Optional.ofNullable(this.dstPort);
    }

    /**
     * ICMP type name.
     * 
     */
    @Import(name="icmpTypename")
    private @Nullable Output<String> icmpTypename;

    /**
     * @return ICMP type name.
     * 
     */
    public Optional<Output<String>> icmpTypename() {
        return Optional.ofNullable(this.icmpTypename);
    }

    /**
     * ICMPv6 type name.
     * 
     */
    @Import(name="icmpV6Typename")
    private @Nullable Output<String> icmpV6Typename;

    /**
     * @return ICMPv6 type name.
     * 
     */
    public Optional<Output<String>> icmpV6Typename() {
        return Optional.ofNullable(this.icmpV6Typename);
    }

    /**
     * Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
     * 
     */
    @Import(name="ipSec")
    private @Nullable Output<String> ipSec;

    /**
     * @return Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
     * 
     */
    public Optional<Output<String>> ipSec() {
        return Optional.ofNullable(this.ipSec);
    }

    /**
     * Enable logging for the firewall rule.
     * 
     */
    @Import(name="logging")
    private @Nullable Output<Boolean> logging;

    /**
     * @return Enable logging for the firewall rule.
     * 
     */
    public Optional<Output<Boolean>> logging() {
        return Optional.ofNullable(this.logging);
    }

    /**
     * The name of the firewall rule.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the firewall rule.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The protocol of the rule.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return The protocol of the rule.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * The IPv6 protocol of the rule.
     * 
     */
    @Import(name="protocolV6")
    private @Nullable Output<String> protocolV6;

    /**
     * @return The IPv6 protocol of the rule.
     * 
     */
    public Optional<Output<String>> protocolV6() {
        return Optional.ofNullable(this.protocolV6);
    }

    /**
     * The index of the rule. Must be &gt;= 2000 &lt; 3000 or &gt;= 4000 &lt; 5000.
     * 
     */
    @Import(name="ruleIndex", required=true)
    private Output<Integer> ruleIndex;

    /**
     * @return The index of the rule. Must be &gt;= 2000 &lt; 3000 or &gt;= 4000 &lt; 5000.
     * 
     */
    public Output<Integer> ruleIndex() {
        return this.ruleIndex;
    }

    /**
     * The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
     * 
     */
    @Import(name="ruleset", required=true)
    private Output<String> ruleset;

    /**
     * @return The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
     * 
     */
    public Output<String> ruleset() {
        return this.ruleset;
    }

    /**
     * The name of the site to associate the firewall rule with.
     * 
     */
    @Import(name="site")
    private @Nullable Output<String> site;

    /**
     * @return The name of the site to associate the firewall rule with.
     * 
     */
    public Optional<Output<String>> site() {
        return Optional.ofNullable(this.site);
    }

    /**
     * The source address for the firewall rule.
     * 
     */
    @Import(name="srcAddress")
    private @Nullable Output<String> srcAddress;

    /**
     * @return The source address for the firewall rule.
     * 
     */
    public Optional<Output<String>> srcAddress() {
        return Optional.ofNullable(this.srcAddress);
    }

    /**
     * The IPv6 source address for the firewall rule.
     * 
     */
    @Import(name="srcAddressIpv6")
    private @Nullable Output<String> srcAddressIpv6;

    /**
     * @return The IPv6 source address for the firewall rule.
     * 
     */
    public Optional<Output<String>> srcAddressIpv6() {
        return Optional.ofNullable(this.srcAddressIpv6);
    }

    /**
     * The source firewall group IDs for the firewall rule.
     * 
     */
    @Import(name="srcFirewallGroupIds")
    private @Nullable Output<List<String>> srcFirewallGroupIds;

    /**
     * @return The source firewall group IDs for the firewall rule.
     * 
     */
    public Optional<Output<List<String>>> srcFirewallGroupIds() {
        return Optional.ofNullable(this.srcFirewallGroupIds);
    }

    /**
     * The source MAC address of the firewall rule.
     * 
     */
    @Import(name="srcMac")
    private @Nullable Output<String> srcMac;

    /**
     * @return The source MAC address of the firewall rule.
     * 
     */
    public Optional<Output<String>> srcMac() {
        return Optional.ofNullable(this.srcMac);
    }

    /**
     * The source network ID for the firewall rule.
     * 
     */
    @Import(name="srcNetworkId")
    private @Nullable Output<String> srcNetworkId;

    /**
     * @return The source network ID for the firewall rule.
     * 
     */
    public Optional<Output<String>> srcNetworkId() {
        return Optional.ofNullable(this.srcNetworkId);
    }

    /**
     * The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     * 
     */
    @Import(name="srcNetworkType")
    private @Nullable Output<String> srcNetworkType;

    /**
     * @return The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     * 
     */
    public Optional<Output<String>> srcNetworkType() {
        return Optional.ofNullable(this.srcNetworkType);
    }

    /**
     * Match where the state is established.
     * 
     */
    @Import(name="stateEstablished")
    private @Nullable Output<Boolean> stateEstablished;

    /**
     * @return Match where the state is established.
     * 
     */
    public Optional<Output<Boolean>> stateEstablished() {
        return Optional.ofNullable(this.stateEstablished);
    }

    /**
     * Match where the state is invalid.
     * 
     */
    @Import(name="stateInvalid")
    private @Nullable Output<Boolean> stateInvalid;

    /**
     * @return Match where the state is invalid.
     * 
     */
    public Optional<Output<Boolean>> stateInvalid() {
        return Optional.ofNullable(this.stateInvalid);
    }

    /**
     * Match where the state is new.
     * 
     */
    @Import(name="stateNew")
    private @Nullable Output<Boolean> stateNew;

    /**
     * @return Match where the state is new.
     * 
     */
    public Optional<Output<Boolean>> stateNew() {
        return Optional.ofNullable(this.stateNew);
    }

    /**
     * Match where the state is related.
     * 
     */
    @Import(name="stateRelated")
    private @Nullable Output<Boolean> stateRelated;

    /**
     * @return Match where the state is related.
     * 
     */
    public Optional<Output<Boolean>> stateRelated() {
        return Optional.ofNullable(this.stateRelated);
    }

    private RuleArgs() {}

    private RuleArgs(RuleArgs $) {
        this.action = $.action;
        this.dstAddress = $.dstAddress;
        this.dstAddressIpv6 = $.dstAddressIpv6;
        this.dstFirewallGroupIds = $.dstFirewallGroupIds;
        this.dstNetworkId = $.dstNetworkId;
        this.dstNetworkType = $.dstNetworkType;
        this.dstPort = $.dstPort;
        this.icmpTypename = $.icmpTypename;
        this.icmpV6Typename = $.icmpV6Typename;
        this.ipSec = $.ipSec;
        this.logging = $.logging;
        this.name = $.name;
        this.protocol = $.protocol;
        this.protocolV6 = $.protocolV6;
        this.ruleIndex = $.ruleIndex;
        this.ruleset = $.ruleset;
        this.site = $.site;
        this.srcAddress = $.srcAddress;
        this.srcAddressIpv6 = $.srcAddressIpv6;
        this.srcFirewallGroupIds = $.srcFirewallGroupIds;
        this.srcMac = $.srcMac;
        this.srcNetworkId = $.srcNetworkId;
        this.srcNetworkType = $.srcNetworkType;
        this.stateEstablished = $.stateEstablished;
        this.stateInvalid = $.stateInvalid;
        this.stateNew = $.stateNew;
        this.stateRelated = $.stateRelated;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleArgs $;

        public Builder() {
            $ = new RuleArgs();
        }

        public Builder(RuleArgs defaults) {
            $ = new RuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param action The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
         * 
         * @return builder
         * 
         */
        public Builder action(Output<String> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
         * 
         * @return builder
         * 
         */
        public Builder action(String action) {
            return action(Output.of(action));
        }

        /**
         * @param dstAddress The destination address of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder dstAddress(@Nullable Output<String> dstAddress) {
            $.dstAddress = dstAddress;
            return this;
        }

        /**
         * @param dstAddress The destination address of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder dstAddress(String dstAddress) {
            return dstAddress(Output.of(dstAddress));
        }

        /**
         * @param dstAddressIpv6 The IPv6 destination address of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder dstAddressIpv6(@Nullable Output<String> dstAddressIpv6) {
            $.dstAddressIpv6 = dstAddressIpv6;
            return this;
        }

        /**
         * @param dstAddressIpv6 The IPv6 destination address of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder dstAddressIpv6(String dstAddressIpv6) {
            return dstAddressIpv6(Output.of(dstAddressIpv6));
        }

        /**
         * @param dstFirewallGroupIds The destination firewall group IDs of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder dstFirewallGroupIds(@Nullable Output<List<String>> dstFirewallGroupIds) {
            $.dstFirewallGroupIds = dstFirewallGroupIds;
            return this;
        }

        /**
         * @param dstFirewallGroupIds The destination firewall group IDs of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder dstFirewallGroupIds(List<String> dstFirewallGroupIds) {
            return dstFirewallGroupIds(Output.of(dstFirewallGroupIds));
        }

        /**
         * @param dstFirewallGroupIds The destination firewall group IDs of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder dstFirewallGroupIds(String... dstFirewallGroupIds) {
            return dstFirewallGroupIds(List.of(dstFirewallGroupIds));
        }

        /**
         * @param dstNetworkId The destination network ID of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder dstNetworkId(@Nullable Output<String> dstNetworkId) {
            $.dstNetworkId = dstNetworkId;
            return this;
        }

        /**
         * @param dstNetworkId The destination network ID of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder dstNetworkId(String dstNetworkId) {
            return dstNetworkId(Output.of(dstNetworkId));
        }

        /**
         * @param dstNetworkType The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
         * 
         * @return builder
         * 
         */
        public Builder dstNetworkType(@Nullable Output<String> dstNetworkType) {
            $.dstNetworkType = dstNetworkType;
            return this;
        }

        /**
         * @param dstNetworkType The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
         * 
         * @return builder
         * 
         */
        public Builder dstNetworkType(String dstNetworkType) {
            return dstNetworkType(Output.of(dstNetworkType));
        }

        /**
         * @param dstPort The destination port of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder dstPort(@Nullable Output<String> dstPort) {
            $.dstPort = dstPort;
            return this;
        }

        /**
         * @param dstPort The destination port of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder dstPort(String dstPort) {
            return dstPort(Output.of(dstPort));
        }

        /**
         * @param icmpTypename ICMP type name.
         * 
         * @return builder
         * 
         */
        public Builder icmpTypename(@Nullable Output<String> icmpTypename) {
            $.icmpTypename = icmpTypename;
            return this;
        }

        /**
         * @param icmpTypename ICMP type name.
         * 
         * @return builder
         * 
         */
        public Builder icmpTypename(String icmpTypename) {
            return icmpTypename(Output.of(icmpTypename));
        }

        /**
         * @param icmpV6Typename ICMPv6 type name.
         * 
         * @return builder
         * 
         */
        public Builder icmpV6Typename(@Nullable Output<String> icmpV6Typename) {
            $.icmpV6Typename = icmpV6Typename;
            return this;
        }

        /**
         * @param icmpV6Typename ICMPv6 type name.
         * 
         * @return builder
         * 
         */
        public Builder icmpV6Typename(String icmpV6Typename) {
            return icmpV6Typename(Output.of(icmpV6Typename));
        }

        /**
         * @param ipSec Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
         * 
         * @return builder
         * 
         */
        public Builder ipSec(@Nullable Output<String> ipSec) {
            $.ipSec = ipSec;
            return this;
        }

        /**
         * @param ipSec Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
         * 
         * @return builder
         * 
         */
        public Builder ipSec(String ipSec) {
            return ipSec(Output.of(ipSec));
        }

        /**
         * @param logging Enable logging for the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder logging(@Nullable Output<Boolean> logging) {
            $.logging = logging;
            return this;
        }

        /**
         * @param logging Enable logging for the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder logging(Boolean logging) {
            return logging(Output.of(logging));
        }

        /**
         * @param name The name of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param protocol The protocol of the rule.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The protocol of the rule.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param protocolV6 The IPv6 protocol of the rule.
         * 
         * @return builder
         * 
         */
        public Builder protocolV6(@Nullable Output<String> protocolV6) {
            $.protocolV6 = protocolV6;
            return this;
        }

        /**
         * @param protocolV6 The IPv6 protocol of the rule.
         * 
         * @return builder
         * 
         */
        public Builder protocolV6(String protocolV6) {
            return protocolV6(Output.of(protocolV6));
        }

        /**
         * @param ruleIndex The index of the rule. Must be &gt;= 2000 &lt; 3000 or &gt;= 4000 &lt; 5000.
         * 
         * @return builder
         * 
         */
        public Builder ruleIndex(Output<Integer> ruleIndex) {
            $.ruleIndex = ruleIndex;
            return this;
        }

        /**
         * @param ruleIndex The index of the rule. Must be &gt;= 2000 &lt; 3000 or &gt;= 4000 &lt; 5000.
         * 
         * @return builder
         * 
         */
        public Builder ruleIndex(Integer ruleIndex) {
            return ruleIndex(Output.of(ruleIndex));
        }

        /**
         * @param ruleset The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
         * 
         * @return builder
         * 
         */
        public Builder ruleset(Output<String> ruleset) {
            $.ruleset = ruleset;
            return this;
        }

        /**
         * @param ruleset The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
         * 
         * @return builder
         * 
         */
        public Builder ruleset(String ruleset) {
            return ruleset(Output.of(ruleset));
        }

        /**
         * @param site The name of the site to associate the firewall rule with.
         * 
         * @return builder
         * 
         */
        public Builder site(@Nullable Output<String> site) {
            $.site = site;
            return this;
        }

        /**
         * @param site The name of the site to associate the firewall rule with.
         * 
         * @return builder
         * 
         */
        public Builder site(String site) {
            return site(Output.of(site));
        }

        /**
         * @param srcAddress The source address for the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder srcAddress(@Nullable Output<String> srcAddress) {
            $.srcAddress = srcAddress;
            return this;
        }

        /**
         * @param srcAddress The source address for the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder srcAddress(String srcAddress) {
            return srcAddress(Output.of(srcAddress));
        }

        /**
         * @param srcAddressIpv6 The IPv6 source address for the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder srcAddressIpv6(@Nullable Output<String> srcAddressIpv6) {
            $.srcAddressIpv6 = srcAddressIpv6;
            return this;
        }

        /**
         * @param srcAddressIpv6 The IPv6 source address for the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder srcAddressIpv6(String srcAddressIpv6) {
            return srcAddressIpv6(Output.of(srcAddressIpv6));
        }

        /**
         * @param srcFirewallGroupIds The source firewall group IDs for the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder srcFirewallGroupIds(@Nullable Output<List<String>> srcFirewallGroupIds) {
            $.srcFirewallGroupIds = srcFirewallGroupIds;
            return this;
        }

        /**
         * @param srcFirewallGroupIds The source firewall group IDs for the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder srcFirewallGroupIds(List<String> srcFirewallGroupIds) {
            return srcFirewallGroupIds(Output.of(srcFirewallGroupIds));
        }

        /**
         * @param srcFirewallGroupIds The source firewall group IDs for the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder srcFirewallGroupIds(String... srcFirewallGroupIds) {
            return srcFirewallGroupIds(List.of(srcFirewallGroupIds));
        }

        /**
         * @param srcMac The source MAC address of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder srcMac(@Nullable Output<String> srcMac) {
            $.srcMac = srcMac;
            return this;
        }

        /**
         * @param srcMac The source MAC address of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder srcMac(String srcMac) {
            return srcMac(Output.of(srcMac));
        }

        /**
         * @param srcNetworkId The source network ID for the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder srcNetworkId(@Nullable Output<String> srcNetworkId) {
            $.srcNetworkId = srcNetworkId;
            return this;
        }

        /**
         * @param srcNetworkId The source network ID for the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder srcNetworkId(String srcNetworkId) {
            return srcNetworkId(Output.of(srcNetworkId));
        }

        /**
         * @param srcNetworkType The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
         * 
         * @return builder
         * 
         */
        public Builder srcNetworkType(@Nullable Output<String> srcNetworkType) {
            $.srcNetworkType = srcNetworkType;
            return this;
        }

        /**
         * @param srcNetworkType The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
         * 
         * @return builder
         * 
         */
        public Builder srcNetworkType(String srcNetworkType) {
            return srcNetworkType(Output.of(srcNetworkType));
        }

        /**
         * @param stateEstablished Match where the state is established.
         * 
         * @return builder
         * 
         */
        public Builder stateEstablished(@Nullable Output<Boolean> stateEstablished) {
            $.stateEstablished = stateEstablished;
            return this;
        }

        /**
         * @param stateEstablished Match where the state is established.
         * 
         * @return builder
         * 
         */
        public Builder stateEstablished(Boolean stateEstablished) {
            return stateEstablished(Output.of(stateEstablished));
        }

        /**
         * @param stateInvalid Match where the state is invalid.
         * 
         * @return builder
         * 
         */
        public Builder stateInvalid(@Nullable Output<Boolean> stateInvalid) {
            $.stateInvalid = stateInvalid;
            return this;
        }

        /**
         * @param stateInvalid Match where the state is invalid.
         * 
         * @return builder
         * 
         */
        public Builder stateInvalid(Boolean stateInvalid) {
            return stateInvalid(Output.of(stateInvalid));
        }

        /**
         * @param stateNew Match where the state is new.
         * 
         * @return builder
         * 
         */
        public Builder stateNew(@Nullable Output<Boolean> stateNew) {
            $.stateNew = stateNew;
            return this;
        }

        /**
         * @param stateNew Match where the state is new.
         * 
         * @return builder
         * 
         */
        public Builder stateNew(Boolean stateNew) {
            return stateNew(Output.of(stateNew));
        }

        /**
         * @param stateRelated Match where the state is related.
         * 
         * @return builder
         * 
         */
        public Builder stateRelated(@Nullable Output<Boolean> stateRelated) {
            $.stateRelated = stateRelated;
            return this;
        }

        /**
         * @param stateRelated Match where the state is related.
         * 
         * @return builder
         * 
         */
        public Builder stateRelated(Boolean stateRelated) {
            return stateRelated(Output.of(stateRelated));
        }

        public RuleArgs build() {
            $.action = Objects.requireNonNull($.action, "expected parameter 'action' to be non-null");
            $.ruleIndex = Objects.requireNonNull($.ruleIndex, "expected parameter 'ruleIndex' to be non-null");
            $.ruleset = Objects.requireNonNull($.ruleset, "expected parameter 'ruleset' to be non-null");
            return $;
        }
    }

}
