// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package axiom

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/gfmio/pulumi-axiom/sdk/go/axiom/internal"
)

func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("axiom:index/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// Users identifier
	Id string `pulumi:"id"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	// Users email
	Email string `pulumi:"email"`
	// Users identifier
	Id string `pulumi:"id"`
	// Users name
	Name string `pulumi:"name"`
	// Users role
	Role string `pulumi:"role"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

// A collection of arguments for invoking getUser.
type LookupUserOutputArgs struct {
	// Users identifier
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

// Users email
func (o LookupUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Email }).(pulumi.StringOutput)
}

// Users identifier
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// Users name
func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// Users role
func (o LookupUserResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}