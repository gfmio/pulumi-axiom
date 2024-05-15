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
    public sealed class NotifierProperties
    {
        public readonly Outputs.NotifierPropertiesDiscord? Discord;
        public readonly Outputs.NotifierPropertiesDiscordWebhook? DiscordWebhook;
        public readonly Outputs.NotifierPropertiesEmail? Email;
        public readonly Outputs.NotifierPropertiesOpsgenie? Opsgenie;
        public readonly Outputs.NotifierPropertiesPagerduty? Pagerduty;
        public readonly Outputs.NotifierPropertiesSlack? Slack;
        public readonly Outputs.NotifierPropertiesWebhook? Webhook;

        [OutputConstructor]
        private NotifierProperties(
            Outputs.NotifierPropertiesDiscord? discord,

            Outputs.NotifierPropertiesDiscordWebhook? discordWebhook,

            Outputs.NotifierPropertiesEmail? email,

            Outputs.NotifierPropertiesOpsgenie? opsgenie,

            Outputs.NotifierPropertiesPagerduty? pagerduty,

            Outputs.NotifierPropertiesSlack? slack,

            Outputs.NotifierPropertiesWebhook? webhook)
        {
            Discord = discord;
            DiscordWebhook = discordWebhook;
            Email = email;
            Opsgenie = opsgenie;
            Pagerduty = pagerduty;
            Slack = slack;
            Webhook = webhook;
        }
    }
}