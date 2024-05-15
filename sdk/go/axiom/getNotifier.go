// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package axiom

import (
	"context"
	"reflect"

	"github.com/gfmio/pulumi-axiom/sdk/go/axiom/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNotifier(ctx *pulumi.Context, args *LookupNotifierArgs, opts ...pulumi.InvokeOption) (*LookupNotifierResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNotifierResult
	err := ctx.Invoke("axiom:index/getNotifier:getNotifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNotifier.
type LookupNotifierArgs struct {
	// Notifier identifier
	Id string `pulumi:"id"`
}

// A collection of values returned by getNotifier.
type LookupNotifierResult struct {
	// Notifier identifier
	Id string `pulumi:"id"`
	// Notifier name
	Name string `pulumi:"name"`
	// The properties of the notifier
	Properties GetNotifierProperties `pulumi:"properties"`
}

func LookupNotifierOutput(ctx *pulumi.Context, args LookupNotifierOutputArgs, opts ...pulumi.InvokeOption) LookupNotifierResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotifierResult, error) {
			args := v.(LookupNotifierArgs)
			r, err := LookupNotifier(ctx, &args, opts...)
			var s LookupNotifierResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotifierResultOutput)
}

// A collection of arguments for invoking getNotifier.
type LookupNotifierOutputArgs struct {
	// Notifier identifier
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupNotifierOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotifierArgs)(nil)).Elem()
}

// A collection of values returned by getNotifier.
type LookupNotifierResultOutput struct{ *pulumi.OutputState }

func (LookupNotifierResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotifierResult)(nil)).Elem()
}

func (o LookupNotifierResultOutput) ToLookupNotifierResultOutput() LookupNotifierResultOutput {
	return o
}

func (o LookupNotifierResultOutput) ToLookupNotifierResultOutputWithContext(ctx context.Context) LookupNotifierResultOutput {
	return o
}

// Notifier identifier
func (o LookupNotifierResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotifierResult) string { return v.Id }).(pulumi.StringOutput)
}

// Notifier name
func (o LookupNotifierResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotifierResult) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of the notifier
func (o LookupNotifierResultOutput) Properties() GetNotifierPropertiesOutput {
	return o.ApplyT(func(v LookupNotifierResult) GetNotifierProperties { return v.Properties }).(GetNotifierPropertiesOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotifierResultOutput{})
}
