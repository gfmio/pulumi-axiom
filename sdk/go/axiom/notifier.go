// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package axiom

import (
	"context"
	"reflect"

	"errors"
	"github.com/gfmio/pulumi-axiom/sdk/go/axiom/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Notifier struct {
	pulumi.CustomResourceState

	// Notifier name
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the notifier
	Properties NotifierPropertiesOutput `pulumi:"properties"`
}

// NewNotifier registers a new resource with the given unique name, arguments, and options.
func NewNotifier(ctx *pulumi.Context,
	name string, args *NotifierArgs, opts ...pulumi.ResourceOption) (*Notifier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Notifier
	err := ctx.RegisterResource("axiom:index/notifier:Notifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotifier gets an existing Notifier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotifierState, opts ...pulumi.ResourceOption) (*Notifier, error) {
	var resource Notifier
	err := ctx.ReadResource("axiom:index/notifier:Notifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Notifier resources.
type notifierState struct {
	// Notifier name
	Name *string `pulumi:"name"`
	// The properties of the notifier
	Properties *NotifierProperties `pulumi:"properties"`
}

type NotifierState struct {
	// Notifier name
	Name pulumi.StringPtrInput
	// The properties of the notifier
	Properties NotifierPropertiesPtrInput
}

func (NotifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*notifierState)(nil)).Elem()
}

type notifierArgs struct {
	// Notifier name
	Name *string `pulumi:"name"`
	// The properties of the notifier
	Properties NotifierProperties `pulumi:"properties"`
}

// The set of arguments for constructing a Notifier resource.
type NotifierArgs struct {
	// Notifier name
	Name pulumi.StringPtrInput
	// The properties of the notifier
	Properties NotifierPropertiesInput
}

func (NotifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notifierArgs)(nil)).Elem()
}

type NotifierInput interface {
	pulumi.Input

	ToNotifierOutput() NotifierOutput
	ToNotifierOutputWithContext(ctx context.Context) NotifierOutput
}

func (*Notifier) ElementType() reflect.Type {
	return reflect.TypeOf((**Notifier)(nil)).Elem()
}

func (i *Notifier) ToNotifierOutput() NotifierOutput {
	return i.ToNotifierOutputWithContext(context.Background())
}

func (i *Notifier) ToNotifierOutputWithContext(ctx context.Context) NotifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotifierOutput)
}

// NotifierArrayInput is an input type that accepts NotifierArray and NotifierArrayOutput values.
// You can construct a concrete instance of `NotifierArrayInput` via:
//
//	NotifierArray{ NotifierArgs{...} }
type NotifierArrayInput interface {
	pulumi.Input

	ToNotifierArrayOutput() NotifierArrayOutput
	ToNotifierArrayOutputWithContext(context.Context) NotifierArrayOutput
}

type NotifierArray []NotifierInput

func (NotifierArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Notifier)(nil)).Elem()
}

func (i NotifierArray) ToNotifierArrayOutput() NotifierArrayOutput {
	return i.ToNotifierArrayOutputWithContext(context.Background())
}

func (i NotifierArray) ToNotifierArrayOutputWithContext(ctx context.Context) NotifierArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotifierArrayOutput)
}

// NotifierMapInput is an input type that accepts NotifierMap and NotifierMapOutput values.
// You can construct a concrete instance of `NotifierMapInput` via:
//
//	NotifierMap{ "key": NotifierArgs{...} }
type NotifierMapInput interface {
	pulumi.Input

	ToNotifierMapOutput() NotifierMapOutput
	ToNotifierMapOutputWithContext(context.Context) NotifierMapOutput
}

type NotifierMap map[string]NotifierInput

func (NotifierMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Notifier)(nil)).Elem()
}

func (i NotifierMap) ToNotifierMapOutput() NotifierMapOutput {
	return i.ToNotifierMapOutputWithContext(context.Background())
}

func (i NotifierMap) ToNotifierMapOutputWithContext(ctx context.Context) NotifierMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotifierMapOutput)
}

type NotifierOutput struct{ *pulumi.OutputState }

func (NotifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Notifier)(nil)).Elem()
}

func (o NotifierOutput) ToNotifierOutput() NotifierOutput {
	return o
}

func (o NotifierOutput) ToNotifierOutputWithContext(ctx context.Context) NotifierOutput {
	return o
}

// Notifier name
func (o NotifierOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Notifier) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The properties of the notifier
func (o NotifierOutput) Properties() NotifierPropertiesOutput {
	return o.ApplyT(func(v *Notifier) NotifierPropertiesOutput { return v.Properties }).(NotifierPropertiesOutput)
}

type NotifierArrayOutput struct{ *pulumi.OutputState }

func (NotifierArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Notifier)(nil)).Elem()
}

func (o NotifierArrayOutput) ToNotifierArrayOutput() NotifierArrayOutput {
	return o
}

func (o NotifierArrayOutput) ToNotifierArrayOutputWithContext(ctx context.Context) NotifierArrayOutput {
	return o
}

func (o NotifierArrayOutput) Index(i pulumi.IntInput) NotifierOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Notifier {
		return vs[0].([]*Notifier)[vs[1].(int)]
	}).(NotifierOutput)
}

type NotifierMapOutput struct{ *pulumi.OutputState }

func (NotifierMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Notifier)(nil)).Elem()
}

func (o NotifierMapOutput) ToNotifierMapOutput() NotifierMapOutput {
	return o
}

func (o NotifierMapOutput) ToNotifierMapOutputWithContext(ctx context.Context) NotifierMapOutput {
	return o
}

func (o NotifierMapOutput) MapIndex(k pulumi.StringInput) NotifierOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Notifier {
		return vs[0].(map[string]*Notifier)[vs[1].(string)]
	}).(NotifierOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotifierInput)(nil)).Elem(), &Notifier{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotifierArrayInput)(nil)).Elem(), NotifierArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotifierMapInput)(nil)).Elem(), NotifierMap{})
	pulumi.RegisterOutputType(NotifierOutput{})
	pulumi.RegisterOutputType(NotifierArrayOutput{})
	pulumi.RegisterOutputType(NotifierMapOutput{})
}
