// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package axiom

import (
	"context"
	"reflect"

	"github.com/gfmio/pulumi-axiom/sdk/go/axiom/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataset(ctx *pulumi.Context, args *LookupDatasetArgs, opts ...pulumi.InvokeOption) (*LookupDatasetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatasetResult
	err := ctx.Invoke("axiom:index/getDataset:getDataset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataset.
type LookupDatasetArgs struct {
	// Dataset identifier
	Id string `pulumi:"id"`
}

// A collection of values returned by getDataset.
type LookupDatasetResult struct {
	// Dataset description
	Description string `pulumi:"description"`
	// Dataset identifier
	Id string `pulumi:"id"`
	// Dataset name
	Name string `pulumi:"name"`
}

func LookupDatasetOutput(ctx *pulumi.Context, args LookupDatasetOutputArgs, opts ...pulumi.InvokeOption) LookupDatasetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatasetResult, error) {
			args := v.(LookupDatasetArgs)
			r, err := LookupDataset(ctx, &args, opts...)
			var s LookupDatasetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatasetResultOutput)
}

// A collection of arguments for invoking getDataset.
type LookupDatasetOutputArgs struct {
	// Dataset identifier
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDatasetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatasetArgs)(nil)).Elem()
}

// A collection of values returned by getDataset.
type LookupDatasetResultOutput struct{ *pulumi.OutputState }

func (LookupDatasetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatasetResult)(nil)).Elem()
}

func (o LookupDatasetResultOutput) ToLookupDatasetResultOutput() LookupDatasetResultOutput {
	return o
}

func (o LookupDatasetResultOutput) ToLookupDatasetResultOutputWithContext(ctx context.Context) LookupDatasetResultOutput {
	return o
}

// Dataset description
func (o LookupDatasetResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Description }).(pulumi.StringOutput)
}

// Dataset identifier
func (o LookupDatasetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Dataset name
func (o LookupDatasetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatasetResultOutput{})
}
