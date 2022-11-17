// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `unifi.setting.Radius` manages settings for the built-in RADIUS server.
 */
export class Radius extends pulumi.CustomResource {
    /**
     * Get an existing Radius resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RadiusState, opts?: pulumi.CustomResourceOptions): Radius {
        return new Radius(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'unifi:setting/radius:Radius';

    /**
     * Returns true if the given object is an instance of Radius.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Radius {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Radius.__pulumiType;
    }

    /**
     * Enable RADIUS accounting Defaults to `false`.
     */
    public readonly accountingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The port for accounting communications. Defaults to `1813`.
     */
    public readonly accountingPort!: pulumi.Output<number | undefined>;
    /**
     * The port for authentication communications. Defaults to `1812`.
     */
    public readonly authPort!: pulumi.Output<number | undefined>;
    /**
     * RAIDUS server enabled. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Statistics will be collected from connected clients at this interval. Defaults to `3600`.
     */
    public readonly interimUpdateInterval!: pulumi.Output<number | undefined>;
    /**
     * RAIDUS secret passphrase. Defaults to ``.
     */
    public readonly secret!: pulumi.Output<string | undefined>;
    /**
     * The name of the site to associate the settings with.
     */
    public readonly site!: pulumi.Output<string>;
    /**
     * Encrypt communication between the server and the client. Defaults to `true`.
     */
    public readonly tunneledReply!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Radius resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RadiusArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RadiusArgs | RadiusState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RadiusState | undefined;
            resourceInputs["accountingEnabled"] = state ? state.accountingEnabled : undefined;
            resourceInputs["accountingPort"] = state ? state.accountingPort : undefined;
            resourceInputs["authPort"] = state ? state.authPort : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["interimUpdateInterval"] = state ? state.interimUpdateInterval : undefined;
            resourceInputs["secret"] = state ? state.secret : undefined;
            resourceInputs["site"] = state ? state.site : undefined;
            resourceInputs["tunneledReply"] = state ? state.tunneledReply : undefined;
        } else {
            const args = argsOrState as RadiusArgs | undefined;
            resourceInputs["accountingEnabled"] = args ? args.accountingEnabled : undefined;
            resourceInputs["accountingPort"] = args ? args.accountingPort : undefined;
            resourceInputs["authPort"] = args ? args.authPort : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["interimUpdateInterval"] = args ? args.interimUpdateInterval : undefined;
            resourceInputs["secret"] = args?.secret ? pulumi.secret(args.secret) : undefined;
            resourceInputs["site"] = args ? args.site : undefined;
            resourceInputs["tunneledReply"] = args ? args.tunneledReply : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Radius.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Radius resources.
 */
export interface RadiusState {
    /**
     * Enable RADIUS accounting Defaults to `false`.
     */
    accountingEnabled?: pulumi.Input<boolean>;
    /**
     * The port for accounting communications. Defaults to `1813`.
     */
    accountingPort?: pulumi.Input<number>;
    /**
     * The port for authentication communications. Defaults to `1812`.
     */
    authPort?: pulumi.Input<number>;
    /**
     * RAIDUS server enabled. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Statistics will be collected from connected clients at this interval. Defaults to `3600`.
     */
    interimUpdateInterval?: pulumi.Input<number>;
    /**
     * RAIDUS secret passphrase. Defaults to ``.
     */
    secret?: pulumi.Input<string>;
    /**
     * The name of the site to associate the settings with.
     */
    site?: pulumi.Input<string>;
    /**
     * Encrypt communication between the server and the client. Defaults to `true`.
     */
    tunneledReply?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Radius resource.
 */
export interface RadiusArgs {
    /**
     * Enable RADIUS accounting Defaults to `false`.
     */
    accountingEnabled?: pulumi.Input<boolean>;
    /**
     * The port for accounting communications. Defaults to `1813`.
     */
    accountingPort?: pulumi.Input<number>;
    /**
     * The port for authentication communications. Defaults to `1812`.
     */
    authPort?: pulumi.Input<number>;
    /**
     * RAIDUS server enabled. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Statistics will be collected from connected clients at this interval. Defaults to `3600`.
     */
    interimUpdateInterval?: pulumi.Input<number>;
    /**
     * RAIDUS secret passphrase. Defaults to ``.
     */
    secret?: pulumi.Input<string>;
    /**
     * The name of the site to associate the settings with.
     */
    site?: pulumi.Input<string>;
    /**
     * Encrypt communication between the server and the client. Defaults to `true`.
     */
    tunneledReply?: pulumi.Input<boolean>;
}
