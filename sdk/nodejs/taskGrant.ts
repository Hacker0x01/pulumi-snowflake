// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class TaskGrant extends pulumi.CustomResource {
    /**
     * Get an existing TaskGrant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TaskGrantState, opts?: pulumi.CustomResourceOptions): TaskGrant {
        return new TaskGrant(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'snowflake:index/taskGrant:TaskGrant';

    /**
     * Returns true if the given object is an instance of TaskGrant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TaskGrant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TaskGrant.__pulumiType;
    }

    /**
     * The name of the database containing the current or future tasks on which to grant privileges.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * When this is set to true and a schema_name is provided, apply this grant on all future tasks in the given schema. When
     * this is true and no schema_name is provided apply this grant on all future tasks in the given database. The task_name
     * field must be unset in order to use on_future.
     */
    public readonly onFuture!: pulumi.Output<boolean | undefined>;
    /**
     * The privilege to grant on the current or future task.
     */
    public readonly privilege!: pulumi.Output<string | undefined>;
    /**
     * Grants privilege to these roles.
     */
    public readonly roles!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the schema containing the current or future tasks on which to grant privileges.
     */
    public readonly schemaName!: pulumi.Output<string>;
    /**
     * The name of the task on which to grant privileges immediately (only valid if on_future is false).
     */
    public readonly taskName!: pulumi.Output<string | undefined>;
    /**
     * When this is set to true, allows the recipient role to grant the privileges to other roles.
     */
    public readonly withGrantOption!: pulumi.Output<boolean | undefined>;

    /**
     * Create a TaskGrant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TaskGrantArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TaskGrantArgs | TaskGrantState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TaskGrantState | undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["onFuture"] = state ? state.onFuture : undefined;
            inputs["privilege"] = state ? state.privilege : undefined;
            inputs["roles"] = state ? state.roles : undefined;
            inputs["schemaName"] = state ? state.schemaName : undefined;
            inputs["taskName"] = state ? state.taskName : undefined;
            inputs["withGrantOption"] = state ? state.withGrantOption : undefined;
        } else {
            const args = argsOrState as TaskGrantArgs | undefined;
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.schemaName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schemaName'");
            }
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["onFuture"] = args ? args.onFuture : undefined;
            inputs["privilege"] = args ? args.privilege : undefined;
            inputs["roles"] = args ? args.roles : undefined;
            inputs["schemaName"] = args ? args.schemaName : undefined;
            inputs["taskName"] = args ? args.taskName : undefined;
            inputs["withGrantOption"] = args ? args.withGrantOption : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TaskGrant.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TaskGrant resources.
 */
export interface TaskGrantState {
    /**
     * The name of the database containing the current or future tasks on which to grant privileges.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * When this is set to true and a schema_name is provided, apply this grant on all future tasks in the given schema. When
     * this is true and no schema_name is provided apply this grant on all future tasks in the given database. The task_name
     * field must be unset in order to use on_future.
     */
    onFuture?: pulumi.Input<boolean>;
    /**
     * The privilege to grant on the current or future task.
     */
    privilege?: pulumi.Input<string>;
    /**
     * Grants privilege to these roles.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the schema containing the current or future tasks on which to grant privileges.
     */
    schemaName?: pulumi.Input<string>;
    /**
     * The name of the task on which to grant privileges immediately (only valid if on_future is false).
     */
    taskName?: pulumi.Input<string>;
    /**
     * When this is set to true, allows the recipient role to grant the privileges to other roles.
     */
    withGrantOption?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a TaskGrant resource.
 */
export interface TaskGrantArgs {
    /**
     * The name of the database containing the current or future tasks on which to grant privileges.
     */
    databaseName: pulumi.Input<string>;
    /**
     * When this is set to true and a schema_name is provided, apply this grant on all future tasks in the given schema. When
     * this is true and no schema_name is provided apply this grant on all future tasks in the given database. The task_name
     * field must be unset in order to use on_future.
     */
    onFuture?: pulumi.Input<boolean>;
    /**
     * The privilege to grant on the current or future task.
     */
    privilege?: pulumi.Input<string>;
    /**
     * Grants privilege to these roles.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the schema containing the current or future tasks on which to grant privileges.
     */
    schemaName: pulumi.Input<string>;
    /**
     * The name of the task on which to grant privileges immediately (only valid if on_future is false).
     */
    taskName?: pulumi.Input<string>;
    /**
     * When this is set to true, allows the recipient role to grant the privileges to other roles.
     */
    withGrantOption?: pulumi.Input<boolean>;
}
