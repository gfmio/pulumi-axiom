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

    public sealed class NotifierPropertiesEmailGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("emails", required: true)]
        private InputList<string>? _emails;

        /// <summary>
        /// The emails to be notified
        /// </summary>
        public InputList<string> Emails
        {
            get => _emails ?? (_emails = new InputList<string>());
            set => _emails = value;
        }

        public NotifierPropertiesEmailGetArgs()
        {
        }
        public static new NotifierPropertiesEmailGetArgs Empty => new NotifierPropertiesEmailGetArgs();
    }
}
