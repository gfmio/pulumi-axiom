// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { DatasetArgs, DatasetState } from "./dataset";
export type Dataset = import("./dataset").Dataset;
export const Dataset: typeof import("./dataset").Dataset = null as any;
utilities.lazyLoad(exports, ["Dataset"], () => require("./dataset"));

export { GetDatasetArgs, GetDatasetResult, GetDatasetOutputArgs } from "./getDataset";
export const getDataset: typeof import("./getDataset").getDataset = null as any;
export const getDatasetOutput: typeof import("./getDataset").getDatasetOutput = null as any;
utilities.lazyLoad(exports, ["getDataset","getDatasetOutput"], () => require("./getDataset"));

export { GetMonitorArgs, GetMonitorResult, GetMonitorOutputArgs } from "./getMonitor";
export const getMonitor: typeof import("./getMonitor").getMonitor = null as any;
export const getMonitorOutput: typeof import("./getMonitor").getMonitorOutput = null as any;
utilities.lazyLoad(exports, ["getMonitor","getMonitorOutput"], () => require("./getMonitor"));

export { GetNotifierArgs, GetNotifierResult, GetNotifierOutputArgs } from "./getNotifier";
export const getNotifier: typeof import("./getNotifier").getNotifier = null as any;
export const getNotifierOutput: typeof import("./getNotifier").getNotifierOutput = null as any;
utilities.lazyLoad(exports, ["getNotifier","getNotifierOutput"], () => require("./getNotifier"));

export { GetUserArgs, GetUserResult, GetUserOutputArgs } from "./getUser";
export const getUser: typeof import("./getUser").getUser = null as any;
export const getUserOutput: typeof import("./getUser").getUserOutput = null as any;
utilities.lazyLoad(exports, ["getUser","getUserOutput"], () => require("./getUser"));

export { MonitorArgs, MonitorState } from "./monitor";
export type Monitor = import("./monitor").Monitor;
export const Monitor: typeof import("./monitor").Monitor = null as any;
utilities.lazyLoad(exports, ["Monitor"], () => require("./monitor"));

export { NotifierArgs, NotifierState } from "./notifier";
export type Notifier = import("./notifier").Notifier;
export const Notifier: typeof import("./notifier").Notifier = null as any;
utilities.lazyLoad(exports, ["Notifier"], () => require("./notifier"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { UserArgs, UserState } from "./user";
export type User = import("./user").User;
export const User: typeof import("./user").User = null as any;
utilities.lazyLoad(exports, ["User"], () => require("./user"));


// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "axiom:index/dataset:Dataset":
                return new Dataset(name, <any>undefined, { urn })
            case "axiom:index/monitor:Monitor":
                return new Monitor(name, <any>undefined, { urn })
            case "axiom:index/notifier:Notifier":
                return new Notifier(name, <any>undefined, { urn })
            case "axiom:index/user:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("axiom", "index/dataset", _module)
pulumi.runtime.registerResourceModule("axiom", "index/monitor", _module)
pulumi.runtime.registerResourceModule("axiom", "index/notifier", _module)
pulumi.runtime.registerResourceModule("axiom", "index/user", _module)
pulumi.runtime.registerResourcePackage("axiom", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:axiom") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
