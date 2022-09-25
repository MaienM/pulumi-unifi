// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `unifi.iam.User` manages a user (or "client" in the UI) of the network, these are identified by unique MAC addresses.
 *
 * Users are created in the controller when observed on the network, so the resource defaults to allowing itself to just take over management of a MAC address, but this can be turned off.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as unifi from "@pulumiverse/unifi";
 *
 * const test = new unifi.iam.User("test", {
 *     mac: "01:23:45:67:89:AB",
 *     note: "my note",
 *     fixedIp: "10.0.0.50",
 *     networkId: unifi_network.my_vlan.id,
 * });
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'unifi:iam/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * Specifies whether this resource should just take over control of an existing user. Defaults to `true`.
     */
    public readonly allowExisting!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether this user should be blocked from the network.
     */
    public readonly blocked!: pulumi.Output<boolean | undefined>;
    /**
     * Override the device fingerprint.
     */
    public readonly devIdOverride!: pulumi.Output<number | undefined>;
    /**
     * A fixed IPv4 address for this user.
     */
    public readonly fixedIp!: pulumi.Output<string | undefined>;
    /**
     * The hostname of the user.
     */
    public /*out*/ readonly hostname!: pulumi.Output<string>;
    /**
     * The IP address of the user.
     */
    public /*out*/ readonly ip!: pulumi.Output<string>;
    /**
     * The MAC address of the user.
     */
    public readonly mac!: pulumi.Output<string>;
    /**
     * The name of the user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network ID for this user.
     */
    public readonly networkId!: pulumi.Output<string | undefined>;
    /**
     * A note with additional information for the user.
     */
    public readonly note!: pulumi.Output<string | undefined>;
    /**
     * The name of the site to associate the user with.
     */
    public readonly site!: pulumi.Output<string>;
    /**
     * Specifies whether this resource should tell the controller to "forget" the user on destroy. Defaults to `false`.
     */
    public readonly skipForgetOnDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The user group ID for the user.
     */
    public readonly userGroupId!: pulumi.Output<string | undefined>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["allowExisting"] = state ? state.allowExisting : undefined;
            resourceInputs["blocked"] = state ? state.blocked : undefined;
            resourceInputs["devIdOverride"] = state ? state.devIdOverride : undefined;
            resourceInputs["fixedIp"] = state ? state.fixedIp : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["mac"] = state ? state.mac : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["note"] = state ? state.note : undefined;
            resourceInputs["site"] = state ? state.site : undefined;
            resourceInputs["skipForgetOnDestroy"] = state ? state.skipForgetOnDestroy : undefined;
            resourceInputs["userGroupId"] = state ? state.userGroupId : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.mac === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mac'");
            }
            resourceInputs["allowExisting"] = args ? args.allowExisting : undefined;
            resourceInputs["blocked"] = args ? args.blocked : undefined;
            resourceInputs["devIdOverride"] = args ? args.devIdOverride : undefined;
            resourceInputs["fixedIp"] = args ? args.fixedIp : undefined;
            resourceInputs["mac"] = args ? args.mac : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["note"] = args ? args.note : undefined;
            resourceInputs["site"] = args ? args.site : undefined;
            resourceInputs["skipForgetOnDestroy"] = args ? args.skipForgetOnDestroy : undefined;
            resourceInputs["userGroupId"] = args ? args.userGroupId : undefined;
            resourceInputs["hostname"] = undefined /*out*/;
            resourceInputs["ip"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * Specifies whether this resource should just take over control of an existing user. Defaults to `true`.
     */
    allowExisting?: pulumi.Input<boolean>;
    /**
     * Specifies whether this user should be blocked from the network.
     */
    blocked?: pulumi.Input<boolean>;
    /**
     * Override the device fingerprint.
     */
    devIdOverride?: pulumi.Input<number>;
    /**
     * A fixed IPv4 address for this user.
     */
    fixedIp?: pulumi.Input<string>;
    /**
     * The hostname of the user.
     */
    hostname?: pulumi.Input<string>;
    /**
     * The IP address of the user.
     */
    ip?: pulumi.Input<string>;
    /**
     * The MAC address of the user.
     */
    mac?: pulumi.Input<string>;
    /**
     * The name of the user.
     */
    name?: pulumi.Input<string>;
    /**
     * The network ID for this user.
     */
    networkId?: pulumi.Input<string>;
    /**
     * A note with additional information for the user.
     */
    note?: pulumi.Input<string>;
    /**
     * The name of the site to associate the user with.
     */
    site?: pulumi.Input<string>;
    /**
     * Specifies whether this resource should tell the controller to "forget" the user on destroy. Defaults to `false`.
     */
    skipForgetOnDestroy?: pulumi.Input<boolean>;
    /**
     * The user group ID for the user.
     */
    userGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Specifies whether this resource should just take over control of an existing user. Defaults to `true`.
     */
    allowExisting?: pulumi.Input<boolean>;
    /**
     * Specifies whether this user should be blocked from the network.
     */
    blocked?: pulumi.Input<boolean>;
    /**
     * Override the device fingerprint.
     */
    devIdOverride?: pulumi.Input<number>;
    /**
     * A fixed IPv4 address for this user.
     */
    fixedIp?: pulumi.Input<string>;
    /**
     * The MAC address of the user.
     */
    mac: pulumi.Input<string>;
    /**
     * The name of the user.
     */
    name?: pulumi.Input<string>;
    /**
     * The network ID for this user.
     */
    networkId?: pulumi.Input<string>;
    /**
     * A note with additional information for the user.
     */
    note?: pulumi.Input<string>;
    /**
     * The name of the site to associate the user with.
     */
    site?: pulumi.Input<string>;
    /**
     * Specifies whether this resource should tell the controller to "forget" the user on destroy. Defaults to `false`.
     */
    skipForgetOnDestroy?: pulumi.Input<boolean>;
    /**
     * The user group ID for the user.
     */
    userGroupId?: pulumi.Input<string>;
}
