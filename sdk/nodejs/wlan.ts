// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `unifi.Wlan` manages a WiFi network / SSID.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as unifi from "@pulumi/unifi";
 * import * as unifi from "@pulumiverse/unifi";
 *
 * const config = new pulumi.Config();
 * const vlanId = config.getNumber("vlanId") || 10;
 * const defaultApGroup = unifi.getApGroup({});
 * const defaultGroup = unifi.iam.getGroup({});
 * const vlan = new unifi.Network("vlan", {
 *     purpose: "corporate",
 *     subnet: "10.0.0.1/24",
 *     vlanId: vlanId,
 *     dhcpStart: "10.0.0.6",
 *     dhcpStop: "10.0.0.254",
 *     dhcpEnabled: true,
 * });
 * const wifi = new unifi.Wlan("wifi", {
 *     passphrase: "12345678",
 *     security: "wpapsk",
 *     wpa3Support: true,
 *     wpa3Transition: true,
 *     pmfMode: "optional",
 *     networkId: vlan.id,
 *     apGroupIds: [defaultApGroup.then(defaultApGroup => defaultApGroup.id)],
 *     userGroupId: defaultGroup.then(defaultGroup => defaultGroup.id),
 * });
 * ```
 *
 * ## Import
 *
 * # import from provider configured site
 *
 * ```sh
 *  $ pulumi import unifi:index/wlan:Wlan mywlan 5dc28e5e9106d105bdc87217
 * ```
 *
 * # import from another site
 *
 * ```sh
 *  $ pulumi import unifi:index/wlan:Wlan mywlan bfa2l6i7:5dc28e5e9106d105bdc87217
 * ```
 */
export class Wlan extends pulumi.CustomResource {
    /**
     * Get an existing Wlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WlanState, opts?: pulumi.CustomResourceOptions): Wlan {
        return new Wlan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'unifi:index/wlan:Wlan';

    /**
     * Returns true if the given object is an instance of Wlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Wlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Wlan.__pulumiType;
    }

    /**
     * IDs of the AP groups to use for this network.
     */
    public readonly apGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Indicates whether or not to hide the SSID from broadcast.
     */
    public readonly hideSsid!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates that this is a guest WLAN and should use guest behaviors.
     */
    public readonly isGuest!: pulumi.Output<boolean | undefined>;
    /**
     * Isolates stations on layer 2 (ethernet) level Defaults to `false`.
     */
    public readonly l2Isolation!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether or not the MAC filter is turned of for the network.
     */
    public readonly macFilterEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of MAC addresses to filter (only valid if `macFilterEnabled` is `true`).
     */
    public readonly macFilterLists!: pulumi.Output<string[] | undefined>;
    /**
     * MAC address filter policy (only valid if `macFilterEnabled` is `true`). Defaults to `deny`.
     */
    public readonly macFilterPolicy!: pulumi.Output<string | undefined>;
    /**
     * Set minimum data rate control for 2G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `1000`, `2000`, `5500`, `6000`, `9000`, `11000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
     */
    public readonly minimumDataRate2gKbps!: pulumi.Output<number | undefined>;
    /**
     * Set minimum data rate control for 5G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `6000`, `9000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
     */
    public readonly minimumDataRate5gKbps!: pulumi.Output<number | undefined>;
    /**
     * Indicates whether or not Multicast Enhance is turned of for the network.
     */
    public readonly multicastEnhance!: pulumi.Output<boolean | undefined>;
    /**
     * The SSID of the network.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the network for this SSID
     */
    public readonly networkId!: pulumi.Output<string | undefined>;
    /**
     * Connect high performance clients to 5 GHz only Defaults to `true`.
     */
    public readonly no2ghzOui!: pulumi.Output<boolean | undefined>;
    /**
     * The passphrase for the network, this is only required if `security` is not set to `open`.
     */
    public readonly passphrase!: pulumi.Output<string | undefined>;
    /**
     * Enable Protected Management Frames. This cannot be disabled if using WPA 3. Valid values are `required`, `optional` and `disabled`. Defaults to `disabled`.
     */
    public readonly pmfMode!: pulumi.Output<string | undefined>;
    /**
     * ID of the RADIUS profile to use when security `wpaeap`. You can query this via the `unifi.getRadiusProfile` data source.
     */
    public readonly radiusProfileId!: pulumi.Output<string | undefined>;
    /**
     * Start and stop schedules for the WLAN
     */
    public readonly schedules!: pulumi.Output<outputs.WlanSchedule[] | undefined>;
    /**
     * The type of WiFi security for this network. Valid values are: `wpapsk`, `wpaeap`, and `open`.
     */
    public readonly security!: pulumi.Output<string>;
    /**
     * The name of the site to associate the wlan with.
     */
    public readonly site!: pulumi.Output<string>;
    /**
     * Enable Unscheduled Automatic Power Save Delivery Defaults to `false`.
     */
    public readonly uapsd!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the user group to use for this network.
     */
    public readonly userGroupId!: pulumi.Output<string>;
    /**
     * Radio band your WiFi network will use.
     */
    public readonly wlanBand!: pulumi.Output<string | undefined>;
    /**
     * Enable WPA 3 support (security must be `wpapsk` and PMF must be turned on).
     */
    public readonly wpa3Support!: pulumi.Output<boolean | undefined>;
    /**
     * Enable WPA 3 and WPA 2 support (security must be `wpapsk` and `wpa3Support` must be true).
     */
    public readonly wpa3Transition!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Wlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WlanArgs | WlanState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WlanState | undefined;
            resourceInputs["apGroupIds"] = state ? state.apGroupIds : undefined;
            resourceInputs["hideSsid"] = state ? state.hideSsid : undefined;
            resourceInputs["isGuest"] = state ? state.isGuest : undefined;
            resourceInputs["l2Isolation"] = state ? state.l2Isolation : undefined;
            resourceInputs["macFilterEnabled"] = state ? state.macFilterEnabled : undefined;
            resourceInputs["macFilterLists"] = state ? state.macFilterLists : undefined;
            resourceInputs["macFilterPolicy"] = state ? state.macFilterPolicy : undefined;
            resourceInputs["minimumDataRate2gKbps"] = state ? state.minimumDataRate2gKbps : undefined;
            resourceInputs["minimumDataRate5gKbps"] = state ? state.minimumDataRate5gKbps : undefined;
            resourceInputs["multicastEnhance"] = state ? state.multicastEnhance : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["no2ghzOui"] = state ? state.no2ghzOui : undefined;
            resourceInputs["passphrase"] = state ? state.passphrase : undefined;
            resourceInputs["pmfMode"] = state ? state.pmfMode : undefined;
            resourceInputs["radiusProfileId"] = state ? state.radiusProfileId : undefined;
            resourceInputs["schedules"] = state ? state.schedules : undefined;
            resourceInputs["security"] = state ? state.security : undefined;
            resourceInputs["site"] = state ? state.site : undefined;
            resourceInputs["uapsd"] = state ? state.uapsd : undefined;
            resourceInputs["userGroupId"] = state ? state.userGroupId : undefined;
            resourceInputs["wlanBand"] = state ? state.wlanBand : undefined;
            resourceInputs["wpa3Support"] = state ? state.wpa3Support : undefined;
            resourceInputs["wpa3Transition"] = state ? state.wpa3Transition : undefined;
        } else {
            const args = argsOrState as WlanArgs | undefined;
            if ((!args || args.security === undefined) && !opts.urn) {
                throw new Error("Missing required property 'security'");
            }
            if ((!args || args.userGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userGroupId'");
            }
            resourceInputs["apGroupIds"] = args ? args.apGroupIds : undefined;
            resourceInputs["hideSsid"] = args ? args.hideSsid : undefined;
            resourceInputs["isGuest"] = args ? args.isGuest : undefined;
            resourceInputs["l2Isolation"] = args ? args.l2Isolation : undefined;
            resourceInputs["macFilterEnabled"] = args ? args.macFilterEnabled : undefined;
            resourceInputs["macFilterLists"] = args ? args.macFilterLists : undefined;
            resourceInputs["macFilterPolicy"] = args ? args.macFilterPolicy : undefined;
            resourceInputs["minimumDataRate2gKbps"] = args ? args.minimumDataRate2gKbps : undefined;
            resourceInputs["minimumDataRate5gKbps"] = args ? args.minimumDataRate5gKbps : undefined;
            resourceInputs["multicastEnhance"] = args ? args.multicastEnhance : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["no2ghzOui"] = args ? args.no2ghzOui : undefined;
            resourceInputs["passphrase"] = args ? args.passphrase : undefined;
            resourceInputs["pmfMode"] = args ? args.pmfMode : undefined;
            resourceInputs["radiusProfileId"] = args ? args.radiusProfileId : undefined;
            resourceInputs["schedules"] = args ? args.schedules : undefined;
            resourceInputs["security"] = args ? args.security : undefined;
            resourceInputs["site"] = args ? args.site : undefined;
            resourceInputs["uapsd"] = args ? args.uapsd : undefined;
            resourceInputs["userGroupId"] = args ? args.userGroupId : undefined;
            resourceInputs["wlanBand"] = args ? args.wlanBand : undefined;
            resourceInputs["wpa3Support"] = args ? args.wpa3Support : undefined;
            resourceInputs["wpa3Transition"] = args ? args.wpa3Transition : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Wlan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Wlan resources.
 */
export interface WlanState {
    /**
     * IDs of the AP groups to use for this network.
     */
    apGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates whether or not to hide the SSID from broadcast.
     */
    hideSsid?: pulumi.Input<boolean>;
    /**
     * Indicates that this is a guest WLAN and should use guest behaviors.
     */
    isGuest?: pulumi.Input<boolean>;
    /**
     * Isolates stations on layer 2 (ethernet) level Defaults to `false`.
     */
    l2Isolation?: pulumi.Input<boolean>;
    /**
     * Indicates whether or not the MAC filter is turned of for the network.
     */
    macFilterEnabled?: pulumi.Input<boolean>;
    /**
     * List of MAC addresses to filter (only valid if `macFilterEnabled` is `true`).
     */
    macFilterLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * MAC address filter policy (only valid if `macFilterEnabled` is `true`). Defaults to `deny`.
     */
    macFilterPolicy?: pulumi.Input<string>;
    /**
     * Set minimum data rate control for 2G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `1000`, `2000`, `5500`, `6000`, `9000`, `11000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
     */
    minimumDataRate2gKbps?: pulumi.Input<number>;
    /**
     * Set minimum data rate control for 5G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `6000`, `9000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
     */
    minimumDataRate5gKbps?: pulumi.Input<number>;
    /**
     * Indicates whether or not Multicast Enhance is turned of for the network.
     */
    multicastEnhance?: pulumi.Input<boolean>;
    /**
     * The SSID of the network.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the network for this SSID
     */
    networkId?: pulumi.Input<string>;
    /**
     * Connect high performance clients to 5 GHz only Defaults to `true`.
     */
    no2ghzOui?: pulumi.Input<boolean>;
    /**
     * The passphrase for the network, this is only required if `security` is not set to `open`.
     */
    passphrase?: pulumi.Input<string>;
    /**
     * Enable Protected Management Frames. This cannot be disabled if using WPA 3. Valid values are `required`, `optional` and `disabled`. Defaults to `disabled`.
     */
    pmfMode?: pulumi.Input<string>;
    /**
     * ID of the RADIUS profile to use when security `wpaeap`. You can query this via the `unifi.getRadiusProfile` data source.
     */
    radiusProfileId?: pulumi.Input<string>;
    /**
     * Start and stop schedules for the WLAN
     */
    schedules?: pulumi.Input<pulumi.Input<inputs.WlanSchedule>[]>;
    /**
     * The type of WiFi security for this network. Valid values are: `wpapsk`, `wpaeap`, and `open`.
     */
    security?: pulumi.Input<string>;
    /**
     * The name of the site to associate the wlan with.
     */
    site?: pulumi.Input<string>;
    /**
     * Enable Unscheduled Automatic Power Save Delivery Defaults to `false`.
     */
    uapsd?: pulumi.Input<boolean>;
    /**
     * ID of the user group to use for this network.
     */
    userGroupId?: pulumi.Input<string>;
    /**
     * Radio band your WiFi network will use.
     */
    wlanBand?: pulumi.Input<string>;
    /**
     * Enable WPA 3 support (security must be `wpapsk` and PMF must be turned on).
     */
    wpa3Support?: pulumi.Input<boolean>;
    /**
     * Enable WPA 3 and WPA 2 support (security must be `wpapsk` and `wpa3Support` must be true).
     */
    wpa3Transition?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Wlan resource.
 */
export interface WlanArgs {
    /**
     * IDs of the AP groups to use for this network.
     */
    apGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates whether or not to hide the SSID from broadcast.
     */
    hideSsid?: pulumi.Input<boolean>;
    /**
     * Indicates that this is a guest WLAN and should use guest behaviors.
     */
    isGuest?: pulumi.Input<boolean>;
    /**
     * Isolates stations on layer 2 (ethernet) level Defaults to `false`.
     */
    l2Isolation?: pulumi.Input<boolean>;
    /**
     * Indicates whether or not the MAC filter is turned of for the network.
     */
    macFilterEnabled?: pulumi.Input<boolean>;
    /**
     * List of MAC addresses to filter (only valid if `macFilterEnabled` is `true`).
     */
    macFilterLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * MAC address filter policy (only valid if `macFilterEnabled` is `true`). Defaults to `deny`.
     */
    macFilterPolicy?: pulumi.Input<string>;
    /**
     * Set minimum data rate control for 2G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `1000`, `2000`, `5500`, `6000`, `9000`, `11000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
     */
    minimumDataRate2gKbps?: pulumi.Input<number>;
    /**
     * Set minimum data rate control for 5G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `6000`, `9000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
     */
    minimumDataRate5gKbps?: pulumi.Input<number>;
    /**
     * Indicates whether or not Multicast Enhance is turned of for the network.
     */
    multicastEnhance?: pulumi.Input<boolean>;
    /**
     * The SSID of the network.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the network for this SSID
     */
    networkId?: pulumi.Input<string>;
    /**
     * Connect high performance clients to 5 GHz only Defaults to `true`.
     */
    no2ghzOui?: pulumi.Input<boolean>;
    /**
     * The passphrase for the network, this is only required if `security` is not set to `open`.
     */
    passphrase?: pulumi.Input<string>;
    /**
     * Enable Protected Management Frames. This cannot be disabled if using WPA 3. Valid values are `required`, `optional` and `disabled`. Defaults to `disabled`.
     */
    pmfMode?: pulumi.Input<string>;
    /**
     * ID of the RADIUS profile to use when security `wpaeap`. You can query this via the `unifi.getRadiusProfile` data source.
     */
    radiusProfileId?: pulumi.Input<string>;
    /**
     * Start and stop schedules for the WLAN
     */
    schedules?: pulumi.Input<pulumi.Input<inputs.WlanSchedule>[]>;
    /**
     * The type of WiFi security for this network. Valid values are: `wpapsk`, `wpaeap`, and `open`.
     */
    security: pulumi.Input<string>;
    /**
     * The name of the site to associate the wlan with.
     */
    site?: pulumi.Input<string>;
    /**
     * Enable Unscheduled Automatic Power Save Delivery Defaults to `false`.
     */
    uapsd?: pulumi.Input<boolean>;
    /**
     * ID of the user group to use for this network.
     */
    userGroupId: pulumi.Input<string>;
    /**
     * Radio band your WiFi network will use.
     */
    wlanBand?: pulumi.Input<string>;
    /**
     * Enable WPA 3 support (security must be `wpapsk` and PMF must be turned on).
     */
    wpa3Support?: pulumi.Input<boolean>;
    /**
     * Enable WPA 3 and WPA 2 support (security must be `wpapsk` and `wpa3Support` must be true).
     */
    wpa3Transition?: pulumi.Input<boolean>;
}
