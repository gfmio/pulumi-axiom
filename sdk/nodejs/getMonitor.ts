// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getMonitor(args: GetMonitorArgs, opts?: pulumi.InvokeOptions): Promise<GetMonitorResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("axiom:index/getMonitor:getMonitor", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getMonitor.
 */
export interface GetMonitorArgs {
    /**
     * Monitor identifier
     */
    id: string;
}

/**
 * A collection of values returned by getMonitor.
 */
export interface GetMonitorResult {
    /**
     * If the monitor should trigger an alert if there is no data
     */
    readonly alertOnNoData: boolean;
    /**
     * The query used inside the monitor
     */
    readonly aplQuery: string;
    /**
     * Monitor description
     */
    readonly description: string;
    /**
     * The time the monitor will be disabled until
     */
    readonly disabledUntil: string;
    /**
     * Monitor identifier
     */
    readonly id: string;
    /**
     * How often the monitor should run
     */
    readonly intervalMinutes: number;
    /**
     * Monitor name
     */
    readonly name: string;
    /**
     * A list of notifier id's to be used when this monitor triggers
     */
    readonly notifierIds: string[];
    readonly notifyByGroup: boolean;
    /**
     * Operator used in monitor trigger evaluation
     */
    readonly operator: string;
    /**
     * Query time range from now
     */
    readonly rangeMinutes: number;
    readonly resolvable: boolean;
    /**
     * The threshold where the monitor should trigger
     */
    readonly threshold: number;
}
export function getMonitorOutput(args: GetMonitorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMonitorResult> {
    return pulumi.output(args).apply((a: any) => getMonitor(a, opts))
}

/**
 * A collection of arguments for invoking getMonitor.
 */
export interface GetMonitorOutputArgs {
    /**
     * Monitor identifier
     */
    id: pulumi.Input<string>;
}