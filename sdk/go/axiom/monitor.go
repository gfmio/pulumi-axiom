// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package axiom

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/gfmio/pulumi-axiom/sdk/go/axiom/internal"
)

type Monitor struct {
	pulumi.CustomResourceState

	// If the monitor should trigger an alert if there is no data
	AlertOnNoData pulumi.BoolOutput `pulumi:"alertOnNoData"`
	// The query used inside the monitor
	AplQuery pulumi.StringOutput `pulumi:"aplQuery"`
	// Monitor description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The time the monitor will be disabled until
	DisabledUntil pulumi.StringPtrOutput `pulumi:"disabledUntil"`
	// How often the monitor should run
	IntervalMinutes pulumi.IntOutput `pulumi:"intervalMinutes"`
	// Monitor name
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of notifier id's to be used when this monitor triggers
	NotifierIds pulumi.StringArrayOutput `pulumi:"notifierIds"`
	// If the monitor should track non-time groups separately
	NotifyByGroup pulumi.BoolOutput `pulumi:"notifyByGroup"`
	// Operator used in monitor trigger evaluation
	Operator pulumi.StringOutput `pulumi:"operator"`
	// Query time range from now
	RangeMinutes pulumi.IntOutput `pulumi:"rangeMinutes"`
	// Determines whether the events triggered by the monitor are individually resolvable. This has no effect on threshold
	// monitors
	Resolvable pulumi.BoolOutput `pulumi:"resolvable"`
	// The threshold where the monitor should trigger
	Threshold pulumi.Float64Output `pulumi:"threshold"`
}

// NewMonitor registers a new resource with the given unique name, arguments, and options.
func NewMonitor(ctx *pulumi.Context,
	name string, args *MonitorArgs, opts ...pulumi.ResourceOption) (*Monitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlertOnNoData == nil {
		return nil, errors.New("invalid value for required argument 'AlertOnNoData'")
	}
	if args.AplQuery == nil {
		return nil, errors.New("invalid value for required argument 'AplQuery'")
	}
	if args.IntervalMinutes == nil {
		return nil, errors.New("invalid value for required argument 'IntervalMinutes'")
	}
	if args.NotifyByGroup == nil {
		return nil, errors.New("invalid value for required argument 'NotifyByGroup'")
	}
	if args.Operator == nil {
		return nil, errors.New("invalid value for required argument 'Operator'")
	}
	if args.RangeMinutes == nil {
		return nil, errors.New("invalid value for required argument 'RangeMinutes'")
	}
	if args.Threshold == nil {
		return nil, errors.New("invalid value for required argument 'Threshold'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Monitor
	err := ctx.RegisterResource("axiom:index/monitor:Monitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitor gets an existing Monitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitorState, opts ...pulumi.ResourceOption) (*Monitor, error) {
	var resource Monitor
	err := ctx.ReadResource("axiom:index/monitor:Monitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Monitor resources.
type monitorState struct {
	// If the monitor should trigger an alert if there is no data
	AlertOnNoData *bool `pulumi:"alertOnNoData"`
	// The query used inside the monitor
	AplQuery *string `pulumi:"aplQuery"`
	// Monitor description
	Description *string `pulumi:"description"`
	// The time the monitor will be disabled until
	DisabledUntil *string `pulumi:"disabledUntil"`
	// How often the monitor should run
	IntervalMinutes *int `pulumi:"intervalMinutes"`
	// Monitor name
	Name *string `pulumi:"name"`
	// A list of notifier id's to be used when this monitor triggers
	NotifierIds []string `pulumi:"notifierIds"`
	// If the monitor should track non-time groups separately
	NotifyByGroup *bool `pulumi:"notifyByGroup"`
	// Operator used in monitor trigger evaluation
	Operator *string `pulumi:"operator"`
	// Query time range from now
	RangeMinutes *int `pulumi:"rangeMinutes"`
	// Determines whether the events triggered by the monitor are individually resolvable. This has no effect on threshold
	// monitors
	Resolvable *bool `pulumi:"resolvable"`
	// The threshold where the monitor should trigger
	Threshold *float64 `pulumi:"threshold"`
}

type MonitorState struct {
	// If the monitor should trigger an alert if there is no data
	AlertOnNoData pulumi.BoolPtrInput
	// The query used inside the monitor
	AplQuery pulumi.StringPtrInput
	// Monitor description
	Description pulumi.StringPtrInput
	// The time the monitor will be disabled until
	DisabledUntil pulumi.StringPtrInput
	// How often the monitor should run
	IntervalMinutes pulumi.IntPtrInput
	// Monitor name
	Name pulumi.StringPtrInput
	// A list of notifier id's to be used when this monitor triggers
	NotifierIds pulumi.StringArrayInput
	// If the monitor should track non-time groups separately
	NotifyByGroup pulumi.BoolPtrInput
	// Operator used in monitor trigger evaluation
	Operator pulumi.StringPtrInput
	// Query time range from now
	RangeMinutes pulumi.IntPtrInput
	// Determines whether the events triggered by the monitor are individually resolvable. This has no effect on threshold
	// monitors
	Resolvable pulumi.BoolPtrInput
	// The threshold where the monitor should trigger
	Threshold pulumi.Float64PtrInput
}

func (MonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorState)(nil)).Elem()
}

type monitorArgs struct {
	// If the monitor should trigger an alert if there is no data
	AlertOnNoData bool `pulumi:"alertOnNoData"`
	// The query used inside the monitor
	AplQuery string `pulumi:"aplQuery"`
	// Monitor description
	Description *string `pulumi:"description"`
	// The time the monitor will be disabled until
	DisabledUntil *string `pulumi:"disabledUntil"`
	// How often the monitor should run
	IntervalMinutes int `pulumi:"intervalMinutes"`
	// Monitor name
	Name *string `pulumi:"name"`
	// A list of notifier id's to be used when this monitor triggers
	NotifierIds []string `pulumi:"notifierIds"`
	// If the monitor should track non-time groups separately
	NotifyByGroup bool `pulumi:"notifyByGroup"`
	// Operator used in monitor trigger evaluation
	Operator string `pulumi:"operator"`
	// Query time range from now
	RangeMinutes int `pulumi:"rangeMinutes"`
	// Determines whether the events triggered by the monitor are individually resolvable. This has no effect on threshold
	// monitors
	Resolvable *bool `pulumi:"resolvable"`
	// The threshold where the monitor should trigger
	Threshold float64 `pulumi:"threshold"`
}

// The set of arguments for constructing a Monitor resource.
type MonitorArgs struct {
	// If the monitor should trigger an alert if there is no data
	AlertOnNoData pulumi.BoolInput
	// The query used inside the monitor
	AplQuery pulumi.StringInput
	// Monitor description
	Description pulumi.StringPtrInput
	// The time the monitor will be disabled until
	DisabledUntil pulumi.StringPtrInput
	// How often the monitor should run
	IntervalMinutes pulumi.IntInput
	// Monitor name
	Name pulumi.StringPtrInput
	// A list of notifier id's to be used when this monitor triggers
	NotifierIds pulumi.StringArrayInput
	// If the monitor should track non-time groups separately
	NotifyByGroup pulumi.BoolInput
	// Operator used in monitor trigger evaluation
	Operator pulumi.StringInput
	// Query time range from now
	RangeMinutes pulumi.IntInput
	// Determines whether the events triggered by the monitor are individually resolvable. This has no effect on threshold
	// monitors
	Resolvable pulumi.BoolPtrInput
	// The threshold where the monitor should trigger
	Threshold pulumi.Float64Input
}

func (MonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorArgs)(nil)).Elem()
}

type MonitorInput interface {
	pulumi.Input

	ToMonitorOutput() MonitorOutput
	ToMonitorOutputWithContext(ctx context.Context) MonitorOutput
}

func (*Monitor) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (i *Monitor) ToMonitorOutput() MonitorOutput {
	return i.ToMonitorOutputWithContext(context.Background())
}

func (i *Monitor) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorOutput)
}

// MonitorArrayInput is an input type that accepts MonitorArray and MonitorArrayOutput values.
// You can construct a concrete instance of `MonitorArrayInput` via:
//
//	MonitorArray{ MonitorArgs{...} }
type MonitorArrayInput interface {
	pulumi.Input

	ToMonitorArrayOutput() MonitorArrayOutput
	ToMonitorArrayOutputWithContext(context.Context) MonitorArrayOutput
}

type MonitorArray []MonitorInput

func (MonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Monitor)(nil)).Elem()
}

func (i MonitorArray) ToMonitorArrayOutput() MonitorArrayOutput {
	return i.ToMonitorArrayOutputWithContext(context.Background())
}

func (i MonitorArray) ToMonitorArrayOutputWithContext(ctx context.Context) MonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorArrayOutput)
}

// MonitorMapInput is an input type that accepts MonitorMap and MonitorMapOutput values.
// You can construct a concrete instance of `MonitorMapInput` via:
//
//	MonitorMap{ "key": MonitorArgs{...} }
type MonitorMapInput interface {
	pulumi.Input

	ToMonitorMapOutput() MonitorMapOutput
	ToMonitorMapOutputWithContext(context.Context) MonitorMapOutput
}

type MonitorMap map[string]MonitorInput

func (MonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Monitor)(nil)).Elem()
}

func (i MonitorMap) ToMonitorMapOutput() MonitorMapOutput {
	return i.ToMonitorMapOutputWithContext(context.Background())
}

func (i MonitorMap) ToMonitorMapOutputWithContext(ctx context.Context) MonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorMapOutput)
}

type MonitorOutput struct{ *pulumi.OutputState }

func (MonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (o MonitorOutput) ToMonitorOutput() MonitorOutput {
	return o
}

func (o MonitorOutput) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return o
}

// If the monitor should trigger an alert if there is no data
func (o MonitorOutput) AlertOnNoData() pulumi.BoolOutput {
	return o.ApplyT(func(v *Monitor) pulumi.BoolOutput { return v.AlertOnNoData }).(pulumi.BoolOutput)
}

// The query used inside the monitor
func (o MonitorOutput) AplQuery() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.AplQuery }).(pulumi.StringOutput)
}

// Monitor description
func (o MonitorOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The time the monitor will be disabled until
func (o MonitorOutput) DisabledUntil() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringPtrOutput { return v.DisabledUntil }).(pulumi.StringPtrOutput)
}

// How often the monitor should run
func (o MonitorOutput) IntervalMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v *Monitor) pulumi.IntOutput { return v.IntervalMinutes }).(pulumi.IntOutput)
}

// Monitor name
func (o MonitorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of notifier id's to be used when this monitor triggers
func (o MonitorOutput) NotifierIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringArrayOutput { return v.NotifierIds }).(pulumi.StringArrayOutput)
}

// If the monitor should track non-time groups separately
func (o MonitorOutput) NotifyByGroup() pulumi.BoolOutput {
	return o.ApplyT(func(v *Monitor) pulumi.BoolOutput { return v.NotifyByGroup }).(pulumi.BoolOutput)
}

// Operator used in monitor trigger evaluation
func (o MonitorOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Operator }).(pulumi.StringOutput)
}

// Query time range from now
func (o MonitorOutput) RangeMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v *Monitor) pulumi.IntOutput { return v.RangeMinutes }).(pulumi.IntOutput)
}

// Determines whether the events triggered by the monitor are individually resolvable. This has no effect on threshold
// monitors
func (o MonitorOutput) Resolvable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Monitor) pulumi.BoolOutput { return v.Resolvable }).(pulumi.BoolOutput)
}

// The threshold where the monitor should trigger
func (o MonitorOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v *Monitor) pulumi.Float64Output { return v.Threshold }).(pulumi.Float64Output)
}

type MonitorArrayOutput struct{ *pulumi.OutputState }

func (MonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Monitor)(nil)).Elem()
}

func (o MonitorArrayOutput) ToMonitorArrayOutput() MonitorArrayOutput {
	return o
}

func (o MonitorArrayOutput) ToMonitorArrayOutputWithContext(ctx context.Context) MonitorArrayOutput {
	return o
}

func (o MonitorArrayOutput) Index(i pulumi.IntInput) MonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Monitor {
		return vs[0].([]*Monitor)[vs[1].(int)]
	}).(MonitorOutput)
}

type MonitorMapOutput struct{ *pulumi.OutputState }

func (MonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Monitor)(nil)).Elem()
}

func (o MonitorMapOutput) ToMonitorMapOutput() MonitorMapOutput {
	return o
}

func (o MonitorMapOutput) ToMonitorMapOutputWithContext(ctx context.Context) MonitorMapOutput {
	return o
}

func (o MonitorMapOutput) MapIndex(k pulumi.StringInput) MonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Monitor {
		return vs[0].(map[string]*Monitor)[vs[1].(string)]
	}).(MonitorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorInput)(nil)).Elem(), &Monitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorArrayInput)(nil)).Elem(), MonitorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorMapInput)(nil)).Elem(), MonitorMap{})
	pulumi.RegisterOutputType(MonitorOutput{})
	pulumi.RegisterOutputType(MonitorArrayOutput{})
	pulumi.RegisterOutputType(MonitorMapOutput{})
}