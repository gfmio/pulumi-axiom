// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PulumiAxiom.Axiom
{
    public static class GetNotifier
    {
        public static Task<GetNotifierResult> InvokeAsync(GetNotifierArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNotifierResult>("axiom:index/getNotifier:getNotifier", args ?? new GetNotifierArgs(), options.WithDefaults());

        public static Output<GetNotifierResult> Invoke(GetNotifierInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNotifierResult>("axiom:index/getNotifier:getNotifier", args ?? new GetNotifierInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNotifierArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Notifier identifier
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetNotifierArgs()
        {
        }
        public static new GetNotifierArgs Empty => new GetNotifierArgs();
    }

    public sealed class GetNotifierInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Notifier identifier
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetNotifierInvokeArgs()
        {
        }
        public static new GetNotifierInvokeArgs Empty => new GetNotifierInvokeArgs();
    }


    [OutputType]
    public sealed class GetNotifierResult
    {
        /// <summary>
        /// Notifier identifier
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Notifier name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties of the notifier
        /// </summary>
        public readonly Outputs.GetNotifierPropertiesResult Properties;

        [OutputConstructor]
        private GetNotifierResult(
            string id,

            string name,

            Outputs.GetNotifierPropertiesResult properties)
        {
            Id = id;
            Name = name;
            Properties = properties;
        }
    }
}
