// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.unifi.firewall;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumiverse.unifi.Utilities;
import com.pulumiverse.unifi.firewall.RuleArgs;
import com.pulumiverse.unifi.firewall.inputs.RuleState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `unifi.firewall.Rule` manages an individual firewall rule on the gateway.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.unifi.firewall.Rule;
 * import com.pulumi.unifi.firewall.RuleArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var ipAddress = config.get(&#34;ipAddress&#34;);
 *         var dropAll = new Rule(&#34;dropAll&#34;, RuleArgs.builder()        
 *             .action(&#34;drop&#34;)
 *             .ruleset(&#34;LAN_IN&#34;)
 *             .ruleIndex(2011)
 *             .protocol(&#34;all&#34;)
 *             .dstAddress(ipAddress)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * import using the ID from the controller API/UI
 * 
 * ```sh
 *  $ pulumi import unifi:firewall/rule:Rule my_rule 5f7080eb6b8969064f80494f
 * ```
 * 
 */
@ResourceType(type="unifi:firewall/rule:Rule")
public class Rule extends com.pulumi.resources.CustomResource {
    /**
     * The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
     * 
     */
    @Export(name="action", refs={String.class}, tree="[0]")
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
    @Export(name="dstAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dstAddress;

    /**
     * @return The destination address of the firewall rule.
     * 
     */
    public Output<Optional<String>> dstAddress() {
        return Codegen.optional(this.dstAddress);
    }
    /**
     * The IPv6 destination address of the firewall rule.
     * 
     */
    @Export(name="dstAddressIpv6", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dstAddressIpv6;

    /**
     * @return The IPv6 destination address of the firewall rule.
     * 
     */
    public Output<Optional<String>> dstAddressIpv6() {
        return Codegen.optional(this.dstAddressIpv6);
    }
    /**
     * The destination firewall group IDs of the firewall rule.
     * 
     */
    @Export(name="dstFirewallGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> dstFirewallGroupIds;

    /**
     * @return The destination firewall group IDs of the firewall rule.
     * 
     */
    public Output<Optional<List<String>>> dstFirewallGroupIds() {
        return Codegen.optional(this.dstFirewallGroupIds);
    }
    /**
     * The destination network ID of the firewall rule.
     * 
     */
    @Export(name="dstNetworkId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dstNetworkId;

    /**
     * @return The destination network ID of the firewall rule.
     * 
     */
    public Output<Optional<String>> dstNetworkId() {
        return Codegen.optional(this.dstNetworkId);
    }
    /**
     * The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     * 
     */
    @Export(name="dstNetworkType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dstNetworkType;

    /**
     * @return The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     * 
     */
    public Output<Optional<String>> dstNetworkType() {
        return Codegen.optional(this.dstNetworkType);
    }
    /**
     * The destination port of the firewall rule.
     * 
     */
    @Export(name="dstPort", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dstPort;

    /**
     * @return The destination port of the firewall rule.
     * 
     */
    public Output<Optional<String>> dstPort() {
        return Codegen.optional(this.dstPort);
    }
    /**
     * ICMP type name.
     * 
     */
    @Export(name="icmpTypename", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> icmpTypename;

    /**
     * @return ICMP type name.
     * 
     */
    public Output<Optional<String>> icmpTypename() {
        return Codegen.optional(this.icmpTypename);
    }
    /**
     * ICMPv6 type name.
     * 
     */
    @Export(name="icmpV6Typename", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> icmpV6Typename;

    /**
     * @return ICMPv6 type name.
     * 
     */
    public Output<Optional<String>> icmpV6Typename() {
        return Codegen.optional(this.icmpV6Typename);
    }
    /**
     * Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
     * 
     */
    @Export(name="ipSec", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipSec;

    /**
     * @return Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
     * 
     */
    public Output<Optional<String>> ipSec() {
        return Codegen.optional(this.ipSec);
    }
    /**
     * Enable logging for the firewall rule.
     * 
     */
    @Export(name="logging", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> logging;

    /**
     * @return Enable logging for the firewall rule.
     * 
     */
    public Output<Optional<Boolean>> logging() {
        return Codegen.optional(this.logging);
    }
    /**
     * The name of the firewall rule.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the firewall rule.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The protocol of the rule.
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> protocol;

    /**
     * @return The protocol of the rule.
     * 
     */
    public Output<Optional<String>> protocol() {
        return Codegen.optional(this.protocol);
    }
    /**
     * The IPv6 protocol of the rule.
     * 
     */
    @Export(name="protocolV6", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> protocolV6;

    /**
     * @return The IPv6 protocol of the rule.
     * 
     */
    public Output<Optional<String>> protocolV6() {
        return Codegen.optional(this.protocolV6);
    }
    /**
     * The index of the rule. Must be &gt;= 2000 &lt; 3000 or &gt;= 4000 &lt; 5000.
     * 
     */
    @Export(name="ruleIndex", refs={Integer.class}, tree="[0]")
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
    @Export(name="ruleset", refs={String.class}, tree="[0]")
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
    @Export(name="site", refs={String.class}, tree="[0]")
    private Output<String> site;

    /**
     * @return The name of the site to associate the firewall rule with.
     * 
     */
    public Output<String> site() {
        return this.site;
    }
    /**
     * The source address for the firewall rule.
     * 
     */
    @Export(name="srcAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> srcAddress;

    /**
     * @return The source address for the firewall rule.
     * 
     */
    public Output<Optional<String>> srcAddress() {
        return Codegen.optional(this.srcAddress);
    }
    /**
     * The IPv6 source address for the firewall rule.
     * 
     */
    @Export(name="srcAddressIpv6", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> srcAddressIpv6;

    /**
     * @return The IPv6 source address for the firewall rule.
     * 
     */
    public Output<Optional<String>> srcAddressIpv6() {
        return Codegen.optional(this.srcAddressIpv6);
    }
    /**
     * The source firewall group IDs for the firewall rule.
     * 
     */
    @Export(name="srcFirewallGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> srcFirewallGroupIds;

    /**
     * @return The source firewall group IDs for the firewall rule.
     * 
     */
    public Output<Optional<List<String>>> srcFirewallGroupIds() {
        return Codegen.optional(this.srcFirewallGroupIds);
    }
    /**
     * The source MAC address of the firewall rule.
     * 
     */
    @Export(name="srcMac", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> srcMac;

    /**
     * @return The source MAC address of the firewall rule.
     * 
     */
    public Output<Optional<String>> srcMac() {
        return Codegen.optional(this.srcMac);
    }
    /**
     * The source network ID for the firewall rule.
     * 
     */
    @Export(name="srcNetworkId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> srcNetworkId;

    /**
     * @return The source network ID for the firewall rule.
     * 
     */
    public Output<Optional<String>> srcNetworkId() {
        return Codegen.optional(this.srcNetworkId);
    }
    /**
     * The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     * 
     */
    @Export(name="srcNetworkType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> srcNetworkType;

    /**
     * @return The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     * 
     */
    public Output<Optional<String>> srcNetworkType() {
        return Codegen.optional(this.srcNetworkType);
    }
    /**
     * The source port of the firewall rule.
     * 
     */
    @Export(name="srcPort", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> srcPort;

    /**
     * @return The source port of the firewall rule.
     * 
     */
    public Output<Optional<String>> srcPort() {
        return Codegen.optional(this.srcPort);
    }
    /**
     * Match where the state is established.
     * 
     */
    @Export(name="stateEstablished", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stateEstablished;

    /**
     * @return Match where the state is established.
     * 
     */
    public Output<Optional<Boolean>> stateEstablished() {
        return Codegen.optional(this.stateEstablished);
    }
    /**
     * Match where the state is invalid.
     * 
     */
    @Export(name="stateInvalid", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stateInvalid;

    /**
     * @return Match where the state is invalid.
     * 
     */
    public Output<Optional<Boolean>> stateInvalid() {
        return Codegen.optional(this.stateInvalid);
    }
    /**
     * Match where the state is new.
     * 
     */
    @Export(name="stateNew", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stateNew;

    /**
     * @return Match where the state is new.
     * 
     */
    public Output<Optional<Boolean>> stateNew() {
        return Codegen.optional(this.stateNew);
    }
    /**
     * Match where the state is related.
     * 
     */
    @Export(name="stateRelated", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stateRelated;

    /**
     * @return Match where the state is related.
     * 
     */
    public Output<Optional<Boolean>> stateRelated() {
        return Codegen.optional(this.stateRelated);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Rule(String name) {
        this(name, RuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Rule(String name, RuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Rule(String name, RuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("unifi:firewall/rule:Rule", name, args == null ? RuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Rule(String name, Output<String> id, @Nullable RuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("unifi:firewall/rule:Rule", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Rule get(String name, Output<String> id, @Nullable RuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Rule(name, id, state, options);
    }
}
