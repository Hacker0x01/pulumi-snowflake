// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as snowflake from "@pulumi/snowflake";
 *
 * const account = new snowflake.ManagedAccount("account", {
 *     adminName: "admin",
 *     adminPassword: "secret",
 *     cloud: "aws",
 *     comment: "A managed account.",
 *     locator: "managed-account",
 *     region: "us-west-2",
 *     type: "READER",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import snowflake:index/managedAccount:ManagedAccount example name
 * ```
 */
export class ManagedAccount extends pulumi.CustomResource {
    /**
     * Get an existing ManagedAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagedAccountState, opts?: pulumi.CustomResourceOptions): ManagedAccount {
        return new ManagedAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'snowflake:index/managedAccount:ManagedAccount';

    /**
     * Returns true if the given object is an instance of ManagedAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedAccount.__pulumiType;
    }

    /**
     * Identifier, as well as login name, for the initial user in the managed account. This user serves as the account administrator for the account.
     */
    public readonly adminName!: pulumi.Output<string>;
    /**
     * Password for the initial user in the managed account.
     */
    public readonly adminPassword!: pulumi.Output<string>;
    /**
     * Cloud in which the managed account is located.
     */
    public /*out*/ readonly cloud!: pulumi.Output<string>;
    /**
     * Specifies a comment for the managed account.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Date and time when the managed account was created.
     */
    public /*out*/ readonly createdOn!: pulumi.Output<string>;
    /**
     * Display name of the managed account.
     */
    public /*out*/ readonly locator!: pulumi.Output<string>;
    /**
     * Identifier for the managed account; must be unique for your account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Snowflake Region in which the managed account is located.
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * Specifies the type of managed account.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * URL for accessing the managed account, particularly through the web interface.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a ManagedAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedAccountArgs | ManagedAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ManagedAccountState | undefined;
            resourceInputs["adminName"] = state ? state.adminName : undefined;
            resourceInputs["adminPassword"] = state ? state.adminPassword : undefined;
            resourceInputs["cloud"] = state ? state.cloud : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["createdOn"] = state ? state.createdOn : undefined;
            resourceInputs["locator"] = state ? state.locator : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ManagedAccountArgs | undefined;
            if ((!args || args.adminName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminName'");
            }
            if ((!args || args.adminPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminPassword'");
            }
            resourceInputs["adminName"] = args ? args.adminName : undefined;
            resourceInputs["adminPassword"] = args ? args.adminPassword : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["cloud"] = undefined /*out*/;
            resourceInputs["createdOn"] = undefined /*out*/;
            resourceInputs["locator"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ManagedAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagedAccount resources.
 */
export interface ManagedAccountState {
    /**
     * Identifier, as well as login name, for the initial user in the managed account. This user serves as the account administrator for the account.
     */
    adminName?: pulumi.Input<string>;
    /**
     * Password for the initial user in the managed account.
     */
    adminPassword?: pulumi.Input<string>;
    /**
     * Cloud in which the managed account is located.
     */
    cloud?: pulumi.Input<string>;
    /**
     * Specifies a comment for the managed account.
     */
    comment?: pulumi.Input<string>;
    /**
     * Date and time when the managed account was created.
     */
    createdOn?: pulumi.Input<string>;
    /**
     * Display name of the managed account.
     */
    locator?: pulumi.Input<string>;
    /**
     * Identifier for the managed account; must be unique for your account.
     */
    name?: pulumi.Input<string>;
    /**
     * Snowflake Region in which the managed account is located.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the type of managed account.
     */
    type?: pulumi.Input<string>;
    /**
     * URL for accessing the managed account, particularly through the web interface.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ManagedAccount resource.
 */
export interface ManagedAccountArgs {
    /**
     * Identifier, as well as login name, for the initial user in the managed account. This user serves as the account administrator for the account.
     */
    adminName: pulumi.Input<string>;
    /**
     * Password for the initial user in the managed account.
     */
    adminPassword: pulumi.Input<string>;
    /**
     * Specifies a comment for the managed account.
     */
    comment?: pulumi.Input<string>;
    /**
     * Identifier for the managed account; must be unique for your account.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the type of managed account.
     */
    type?: pulumi.Input<string>;
}
