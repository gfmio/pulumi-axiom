// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/gfmio/pulumi-axiom/sdk/go/axiom/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// The Axiom API token.
func GetApiToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "axiom:apiToken")
}

// The base url of the axiom api.
func GetBaseUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "axiom:baseUrl")
}
