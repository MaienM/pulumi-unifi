// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * `unifi.setting.Mgmt` manages settings for a unifi site.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as unifi from "@maienm/pulumi-unifi";
 *
 * const exampleSite = new unifi.Site("exampleSite", {description: "example"});
 * const exampleMgmt = new unifi.setting.Mgmt("exampleMgmt", {
 *     site: exampleSite.name,
 *     autoUpgrade: true,
 * });
 * ```
 */
export class Mgmt extends pulumi.CustomResource {
    /**
     * Get an existing Mgmt resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MgmtState, opts?: pulumi.CustomResourceOptions): Mgmt {
        return new Mgmt(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'unifi:setting/mgmt:Mgmt';

    /**
     * Returns true if the given object is an instance of Mgmt.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Mgmt {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Mgmt.__pulumiType;
    }

    /**
     * Automatically upgrade device firmware.
     */
    public readonly autoUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the site to associate the settings with.
     */
    public readonly site!: pulumi.Output<string>;
    /**
     * Enable SSH authentication.
     */
    public readonly sshEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * SSH key.
     */
    public readonly sshKeys!: pulumi.Output<outputs.setting.MgmtSshKey[] | undefined>;

    /**
     * Create a Mgmt resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MgmtArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MgmtArgs | MgmtState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MgmtState | undefined;
            resourceInputs["autoUpgrade"] = state ? state.autoUpgrade : undefined;
            resourceInputs["site"] = state ? state.site : undefined;
            resourceInputs["sshEnabled"] = state ? state.sshEnabled : undefined;
            resourceInputs["sshKeys"] = state ? state.sshKeys : undefined;
        } else {
            const args = argsOrState as MgmtArgs | undefined;
            resourceInputs["autoUpgrade"] = args ? args.autoUpgrade : undefined;
            resourceInputs["site"] = args ? args.site : undefined;
            resourceInputs["sshEnabled"] = args ? args.sshEnabled : undefined;
            resourceInputs["sshKeys"] = args ? args.sshKeys : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Mgmt.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Mgmt resources.
 */
export interface MgmtState {
    /**
     * Automatically upgrade device firmware.
     */
    autoUpgrade?: pulumi.Input<boolean>;
    /**
     * The name of the site to associate the settings with.
     */
    site?: pulumi.Input<string>;
    /**
     * Enable SSH authentication.
     */
    sshEnabled?: pulumi.Input<boolean>;
    /**
     * SSH key.
     */
    sshKeys?: pulumi.Input<pulumi.Input<inputs.setting.MgmtSshKey>[]>;
}

/**
 * The set of arguments for constructing a Mgmt resource.
 */
export interface MgmtArgs {
    /**
     * Automatically upgrade device firmware.
     */
    autoUpgrade?: pulumi.Input<boolean>;
    /**
     * The name of the site to associate the settings with.
     */
    site?: pulumi.Input<string>;
    /**
     * Enable SSH authentication.
     */
    sshEnabled?: pulumi.Input<boolean>;
    /**
     * SSH key.
     */
    sshKeys?: pulumi.Input<pulumi.Input<inputs.setting.MgmtSshKey>[]>;
}
