// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class VirtualSwitch extends pulumi.CustomResource {
    /**
     * Get an existing VirtualSwitch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualSwitch {
        return new VirtualSwitch(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'esxi-native:index:VirtualSwitch';

    /**
     * Returns true if the given object is an instance of VirtualSwitch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualSwitch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualSwitch.__pulumiType;
    }

    /**
     * Forged transmits (true=Accept/false=Reject).
     */
    public readonly forgedTransmits!: pulumi.Output<boolean | undefined>;
    /**
     * Virtual Switch Link Discovery Mode.
     */
    public readonly linkDiscoveryMode!: pulumi.Output<string | undefined>;
    /**
     * MAC address changes (true=Accept/false=Reject).
     */
    public readonly macChanges!: pulumi.Output<boolean | undefined>;
    /**
     * Virtual Switch mtu. (1280-9000)
     */
    public readonly mtu!: pulumi.Output<number | undefined>;
    /**
     * Virtual Switch name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Virtual Switch number of ports. (1-4096)
     */
    public readonly ports!: pulumi.Output<number | undefined>;
    /**
     * Promiscuous mode (true=Accept/false=Reject).
     */
    public readonly promiscuousMode!: pulumi.Output<boolean | undefined>;
    /**
     * Uplink configuration.
     */
    public readonly uplinks!: pulumi.Output<outputs.Uplink[] | undefined>;

    /**
     * Create a VirtualSwitch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VirtualSwitchArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["forgedTransmits"] = args ? args.forgedTransmits : undefined;
            resourceInputs["linkDiscoveryMode"] = args ? args.linkDiscoveryMode : undefined;
            resourceInputs["macChanges"] = args ? args.macChanges : undefined;
            resourceInputs["mtu"] = args ? args.mtu : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ports"] = args ? args.ports : undefined;
            resourceInputs["promiscuousMode"] = args ? args.promiscuousMode : undefined;
            resourceInputs["uplinks"] = args ? args.uplinks : undefined;
        } else {
            resourceInputs["forgedTransmits"] = undefined /*out*/;
            resourceInputs["linkDiscoveryMode"] = undefined /*out*/;
            resourceInputs["macChanges"] = undefined /*out*/;
            resourceInputs["mtu"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["ports"] = undefined /*out*/;
            resourceInputs["promiscuousMode"] = undefined /*out*/;
            resourceInputs["uplinks"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VirtualSwitch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualSwitch resource.
 */
export interface VirtualSwitchArgs {
    /**
     * Forged transmits (true=Accept/false=Reject).
     */
    forgedTransmits?: pulumi.Input<boolean>;
    /**
     * Virtual Switch Link Discovery Mode.
     */
    linkDiscoveryMode?: pulumi.Input<string>;
    /**
     * MAC address changes (true=Accept/false=Reject).
     */
    macChanges?: pulumi.Input<boolean>;
    /**
     * Virtual Switch mtu. (1280-9000)
     */
    mtu?: pulumi.Input<number>;
    /**
     * Virtual Switch name.
     */
    name?: pulumi.Input<string>;
    /**
     * Virtual Switch number of ports. (1-4096)
     */
    ports?: pulumi.Input<number>;
    /**
     * Promiscuous mode (true=Accept/false=Reject).
     */
    promiscuousMode?: pulumi.Input<boolean>;
    /**
     * Uplink configuration.
     */
    uplinks?: pulumi.Input<pulumi.Input<inputs.UplinkArgs>[]>;
}
