// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PulumiAxiom.Axiom.Outputs
{

    [OutputType]
    public sealed class NotifierPropertiesDiscord
    {
        /// <summary>
        /// The discord channel
        /// </summary>
        public readonly string DiscordChannel;
        /// <summary>
        /// The discord token
        /// </summary>
        public readonly string DiscordToken;

        [OutputConstructor]
        private NotifierPropertiesDiscord(
            string discordChannel,

            string discordToken)
        {
            DiscordChannel = discordChannel;
            DiscordToken = discordToken;
        }
    }
}