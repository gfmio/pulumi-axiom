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

    public sealed class NotifierPropertiesDiscordWebhookGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The discord webhook URL
        /// </summary>
        [Input("discordWebhookUrl", required: true)]
        public Input<string> DiscordWebhookUrl { get; set; } = null!;

        public NotifierPropertiesDiscordWebhookGetArgs()
        {
        }
        public static new NotifierPropertiesDiscordWebhookGetArgs Empty => new NotifierPropertiesDiscordWebhookGetArgs();
    }
}
