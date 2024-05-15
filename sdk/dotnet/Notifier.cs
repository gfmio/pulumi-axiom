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
    [AxiomResourceType("axiom:index/notifier:Notifier")]
    public partial class Notifier : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Notifier name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of the notifier
        /// </summary>
        [Output("properties")]
        public Output<Outputs.NotifierProperties> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a Notifier resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Notifier(string name, NotifierArgs args, CustomResourceOptions? options = null)
            : base("axiom:index/notifier:Notifier", name, args ?? new NotifierArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Notifier(string name, Input<string> id, NotifierState? state = null, CustomResourceOptions? options = null)
            : base("axiom:index/notifier:Notifier", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/gfmio/pulumi-axiom",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Notifier resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Notifier Get(string name, Input<string> id, NotifierState? state = null, CustomResourceOptions? options = null)
        {
            return new Notifier(name, id, state, options);
        }
    }

    public sealed class NotifierArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Notifier name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The properties of the notifier
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.NotifierPropertiesArgs> Properties { get; set; } = null!;

        public NotifierArgs()
        {
        }
        public static new NotifierArgs Empty => new NotifierArgs();
    }

    public sealed class NotifierState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Notifier name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The properties of the notifier
        /// </summary>
        [Input("properties")]
        public Input<Inputs.NotifierPropertiesGetArgs>? Properties { get; set; }

        public NotifierState()
        {
        }
        public static new NotifierState Empty => new NotifierState();
    }
}