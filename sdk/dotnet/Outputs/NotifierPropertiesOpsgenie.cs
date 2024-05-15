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
    public sealed class NotifierPropertiesOpsgenie
    {
        /// <summary>
        /// The opsgenie API key
        /// </summary>
        public readonly string ApiKey;
        /// <summary>
        /// The opsgenie is EU
        /// </summary>
        public readonly bool IsEu;

        [OutputConstructor]
        private NotifierPropertiesOpsgenie(
            string apiKey,

            bool isEu)
        {
            ApiKey = apiKey;
            IsEu = isEu;
        }
    }
}