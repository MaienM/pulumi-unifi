// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `unifi.StaticRoute` manages a static route.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as unifi from "@maienm/pulumi-unifi";
 *
 * const nexthop = new unifi.StaticRoute("nexthop", {
 *     type: "nexthop-route",
 *     network: "172.17.0.0/16",
 *     distance: 1,
 *     nextHop: "172.16.0.1",
 * });
 * const blackhole = new unifi.StaticRoute("blackhole", {
 *     type: "blackhole",
 *     network: _var.blackhole_cidr,
 *     distance: 1,
 * });
 * const _interface = new unifi.StaticRoute("interface", {
 *     type: "interface-route",
 *     network: _var.wan2_cidr,
 *     distance: 1,
 *     "interface": "WAN2",
 * });
 * ```
 */
export class StaticRoute extends pulumi.CustomResource {
    /**
     * Get an existing StaticRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StaticRouteState, opts?: pulumi.CustomResourceOptions): StaticRoute {
        return new StaticRoute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'unifi:index/staticRoute:StaticRoute';

    /**
     * Returns true if the given object is an instance of StaticRoute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StaticRoute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StaticRoute.__pulumiType;
    }

    /**
     * The distance of the static route.
     */
    public readonly distance!: pulumi.Output<number>;
    /**
     * The interface of the static route (only valid for `interface-route` type). This can be `WAN1`, `WAN2`, or a network ID.
     */
    public readonly interface!: pulumi.Output<string | undefined>;
    /**
     * The name of the static route.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network subnet address.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The next hop of the static route (only valid for `nexthop-route` type).
     */
    public readonly nextHop!: pulumi.Output<string | undefined>;
    /**
     * The name of the site to associate the static route with.
     */
    public readonly site!: pulumi.Output<string>;
    /**
     * The type of static route. Can be `interface-route`, `nexthop-route`, or `blackhole`.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a StaticRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StaticRouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StaticRouteArgs | StaticRouteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StaticRouteState | undefined;
            resourceInputs["distance"] = state ? state.distance : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["network"] = state ? state.network : undefined;
            resourceInputs["nextHop"] = state ? state.nextHop : undefined;
            resourceInputs["site"] = state ? state.site : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as StaticRouteArgs | undefined;
            if ((!args || args.distance === undefined) && !opts.urn) {
                throw new Error("Missing required property 'distance'");
            }
            if ((!args || args.network === undefined) && !opts.urn) {
                throw new Error("Missing required property 'network'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["distance"] = args ? args.distance : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["nextHop"] = args ? args.nextHop : undefined;
            resourceInputs["site"] = args ? args.site : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StaticRoute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StaticRoute resources.
 */
export interface StaticRouteState {
    /**
     * The distance of the static route.
     */
    distance?: pulumi.Input<number>;
    /**
     * The interface of the static route (only valid for `interface-route` type). This can be `WAN1`, `WAN2`, or a network ID.
     */
    interface?: pulumi.Input<string>;
    /**
     * The name of the static route.
     */
    name?: pulumi.Input<string>;
    /**
     * The network subnet address.
     */
    network?: pulumi.Input<string>;
    /**
     * The next hop of the static route (only valid for `nexthop-route` type).
     */
    nextHop?: pulumi.Input<string>;
    /**
     * The name of the site to associate the static route with.
     */
    site?: pulumi.Input<string>;
    /**
     * The type of static route. Can be `interface-route`, `nexthop-route`, or `blackhole`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StaticRoute resource.
 */
export interface StaticRouteArgs {
    /**
     * The distance of the static route.
     */
    distance: pulumi.Input<number>;
    /**
     * The interface of the static route (only valid for `interface-route` type). This can be `WAN1`, `WAN2`, or a network ID.
     */
    interface?: pulumi.Input<string>;
    /**
     * The name of the static route.
     */
    name?: pulumi.Input<string>;
    /**
     * The network subnet address.
     */
    network: pulumi.Input<string>;
    /**
     * The next hop of the static route (only valid for `nexthop-route` type).
     */
    nextHop?: pulumi.Input<string>;
    /**
     * The name of the site to associate the static route with.
     */
    site?: pulumi.Input<string>;
    /**
     * The type of static route. Can be `interface-route`, `nexthop-route`, or `blackhole`.
     */
    type: pulumi.Input<string>;
}
