// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PulumiAxiom.Axiom.Inputs
{

    public sealed class NotifierPropertiesDiscordArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The discord channel
        /// </summary>
        [Input("discordChannel", required: true)]
        public Input<string> DiscordChannel { get; set; } = null!;

        /// <summary>
        /// The discord token
        /// </summary>
        [Input("discordToken", required: true)]
        public Input<string> DiscordToken { get; set; } = null!;

        public NotifierPropertiesDiscordArgs()
        {
        }
        public static new NotifierPropertiesDiscordArgs Empty => new NotifierPropertiesDiscordArgs();
    }
}
