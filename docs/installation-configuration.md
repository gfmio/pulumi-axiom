---
title: Axiom Installation & Configuration
meta_desc: Information on how to install the Axiom provider.
layout: installation
---

## Installation

The Pulumi Axiom provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@gfmio/pulumi-axiom`](https://www.npmjs.com/package/@gfmio/pulumi-axiom)
* Python: [`pulumi_axiom`](https://pypi.org/project/pulumi_axiom/)
* Go: [`github.com/gfmio/pulumi-axiom/sdk/go/axiom`](https://pkg.go.dev/github.com/gfmio/pulumi-axiom/sdk/go/axiom)
* .NET: [`PulumiAxiom.Axiom`](https://www.nuget.org/packages/PulumiAxiom.Axiom)


## Configuration

> Note:  
> Replace the following **sample content**, with the configuration options
> of the wrapped Terraform provider and remove this note.

The following configuration points are available for the `axiom` provider:

- `axiom:apiKey` (environment: `axiom_API_KEY`) - the API key for `axiom`
- `axiom:region` (environment: `axiom_REGION`) - the region in which to deploy resources

### Provider Binary

The Axiom provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource axiom <version>
```

Replace the version string `<version>` with your desired version.
