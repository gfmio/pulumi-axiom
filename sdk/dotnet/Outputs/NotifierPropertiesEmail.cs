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
    public sealed class NotifierPropertiesEmail
    {
        /// <summary>
        /// The emails to be notified
        /// </summary>
        public readonly ImmutableArray<string> Emails;

        [OutputConstructor]
        private NotifierPropertiesEmail(ImmutableArray<string> emails)
        {
            Emails = emails;
        }
    }
}