// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package setting

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `setting.Radius` manages settings for the built-in RADIUS server.
type Radius struct {
	pulumi.CustomResourceState

	// Enable RADIUS accounting Defaults to `false`.
	AccountingEnabled pulumi.BoolPtrOutput `pulumi:"accountingEnabled"`
	// The port for accounting communications. Defaults to `1813`.
	AccountingPort pulumi.IntPtrOutput `pulumi:"accountingPort"`
	// The port for authentication communications. Defaults to `1812`.
	AuthPort pulumi.IntPtrOutput `pulumi:"authPort"`
	// RAIDUS server enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Statistics will be collected from connected clients at this interval. Defaults to `3600`.
	InterimUpdateInterval pulumi.IntPtrOutput `pulumi:"interimUpdateInterval"`
	// RAIDUS secret passphrase. Defaults to ``.
	Secret pulumi.StringPtrOutput `pulumi:"secret"`
	// The name of the site to associate the settings with.
	Site pulumi.StringOutput `pulumi:"site"`
	// Encrypt communication between the server and the client. Defaults to `true`.
	TunneledReply pulumi.BoolPtrOutput `pulumi:"tunneledReply"`
}

// NewRadius registers a new resource with the given unique name, arguments, and options.
func NewRadius(ctx *pulumi.Context,
	name string, args *RadiusArgs, opts ...pulumi.ResourceOption) (*Radius, error) {
	if args == nil {
		args = &RadiusArgs{}
	}

	if args.Secret != nil {
		args.Secret = pulumi.ToSecret(args.Secret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Radius
	err := ctx.RegisterResource("unifi:setting/radius:Radius", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRadius gets an existing Radius resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRadius(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RadiusState, opts ...pulumi.ResourceOption) (*Radius, error) {
	var resource Radius
	err := ctx.ReadResource("unifi:setting/radius:Radius", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Radius resources.
type radiusState struct {
	// Enable RADIUS accounting Defaults to `false`.
	AccountingEnabled *bool `pulumi:"accountingEnabled"`
	// The port for accounting communications. Defaults to `1813`.
	AccountingPort *int `pulumi:"accountingPort"`
	// The port for authentication communications. Defaults to `1812`.
	AuthPort *int `pulumi:"authPort"`
	// RAIDUS server enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Statistics will be collected from connected clients at this interval. Defaults to `3600`.
	InterimUpdateInterval *int `pulumi:"interimUpdateInterval"`
	// RAIDUS secret passphrase. Defaults to ``.
	Secret *string `pulumi:"secret"`
	// The name of the site to associate the settings with.
	Site *string `pulumi:"site"`
	// Encrypt communication between the server and the client. Defaults to `true`.
	TunneledReply *bool `pulumi:"tunneledReply"`
}

type RadiusState struct {
	// Enable RADIUS accounting Defaults to `false`.
	AccountingEnabled pulumi.BoolPtrInput
	// The port for accounting communications. Defaults to `1813`.
	AccountingPort pulumi.IntPtrInput
	// The port for authentication communications. Defaults to `1812`.
	AuthPort pulumi.IntPtrInput
	// RAIDUS server enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Statistics will be collected from connected clients at this interval. Defaults to `3600`.
	InterimUpdateInterval pulumi.IntPtrInput
	// RAIDUS secret passphrase. Defaults to ``.
	Secret pulumi.StringPtrInput
	// The name of the site to associate the settings with.
	Site pulumi.StringPtrInput
	// Encrypt communication between the server and the client. Defaults to `true`.
	TunneledReply pulumi.BoolPtrInput
}

func (RadiusState) ElementType() reflect.Type {
	return reflect.TypeOf((*radiusState)(nil)).Elem()
}

type radiusArgs struct {
	// Enable RADIUS accounting Defaults to `false`.
	AccountingEnabled *bool `pulumi:"accountingEnabled"`
	// The port for accounting communications. Defaults to `1813`.
	AccountingPort *int `pulumi:"accountingPort"`
	// The port for authentication communications. Defaults to `1812`.
	AuthPort *int `pulumi:"authPort"`
	// RAIDUS server enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Statistics will be collected from connected clients at this interval. Defaults to `3600`.
	InterimUpdateInterval *int `pulumi:"interimUpdateInterval"`
	// RAIDUS secret passphrase. Defaults to ``.
	Secret *string `pulumi:"secret"`
	// The name of the site to associate the settings with.
	Site *string `pulumi:"site"`
	// Encrypt communication between the server and the client. Defaults to `true`.
	TunneledReply *bool `pulumi:"tunneledReply"`
}

// The set of arguments for constructing a Radius resource.
type RadiusArgs struct {
	// Enable RADIUS accounting Defaults to `false`.
	AccountingEnabled pulumi.BoolPtrInput
	// The port for accounting communications. Defaults to `1813`.
	AccountingPort pulumi.IntPtrInput
	// The port for authentication communications. Defaults to `1812`.
	AuthPort pulumi.IntPtrInput
	// RAIDUS server enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Statistics will be collected from connected clients at this interval. Defaults to `3600`.
	InterimUpdateInterval pulumi.IntPtrInput
	// RAIDUS secret passphrase. Defaults to ``.
	Secret pulumi.StringPtrInput
	// The name of the site to associate the settings with.
	Site pulumi.StringPtrInput
	// Encrypt communication between the server and the client. Defaults to `true`.
	TunneledReply pulumi.BoolPtrInput
}

func (RadiusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*radiusArgs)(nil)).Elem()
}

type RadiusInput interface {
	pulumi.Input

	ToRadiusOutput() RadiusOutput
	ToRadiusOutputWithContext(ctx context.Context) RadiusOutput
}

func (*Radius) ElementType() reflect.Type {
	return reflect.TypeOf((**Radius)(nil)).Elem()
}

func (i *Radius) ToRadiusOutput() RadiusOutput {
	return i.ToRadiusOutputWithContext(context.Background())
}

func (i *Radius) ToRadiusOutputWithContext(ctx context.Context) RadiusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RadiusOutput)
}

// RadiusArrayInput is an input type that accepts RadiusArray and RadiusArrayOutput values.
// You can construct a concrete instance of `RadiusArrayInput` via:
//
//	RadiusArray{ RadiusArgs{...} }
type RadiusArrayInput interface {
	pulumi.Input

	ToRadiusArrayOutput() RadiusArrayOutput
	ToRadiusArrayOutputWithContext(context.Context) RadiusArrayOutput
}

type RadiusArray []RadiusInput

func (RadiusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Radius)(nil)).Elem()
}

func (i RadiusArray) ToRadiusArrayOutput() RadiusArrayOutput {
	return i.ToRadiusArrayOutputWithContext(context.Background())
}

func (i RadiusArray) ToRadiusArrayOutputWithContext(ctx context.Context) RadiusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RadiusArrayOutput)
}

// RadiusMapInput is an input type that accepts RadiusMap and RadiusMapOutput values.
// You can construct a concrete instance of `RadiusMapInput` via:
//
//	RadiusMap{ "key": RadiusArgs{...} }
type RadiusMapInput interface {
	pulumi.Input

	ToRadiusMapOutput() RadiusMapOutput
	ToRadiusMapOutputWithContext(context.Context) RadiusMapOutput
}

type RadiusMap map[string]RadiusInput

func (RadiusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Radius)(nil)).Elem()
}

func (i RadiusMap) ToRadiusMapOutput() RadiusMapOutput {
	return i.ToRadiusMapOutputWithContext(context.Background())
}

func (i RadiusMap) ToRadiusMapOutputWithContext(ctx context.Context) RadiusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RadiusMapOutput)
}

type RadiusOutput struct{ *pulumi.OutputState }

func (RadiusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Radius)(nil)).Elem()
}

func (o RadiusOutput) ToRadiusOutput() RadiusOutput {
	return o
}

func (o RadiusOutput) ToRadiusOutputWithContext(ctx context.Context) RadiusOutput {
	return o
}

// Enable RADIUS accounting Defaults to `false`.
func (o RadiusOutput) AccountingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Radius) pulumi.BoolPtrOutput { return v.AccountingEnabled }).(pulumi.BoolPtrOutput)
}

// The port for accounting communications. Defaults to `1813`.
func (o RadiusOutput) AccountingPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Radius) pulumi.IntPtrOutput { return v.AccountingPort }).(pulumi.IntPtrOutput)
}

// The port for authentication communications. Defaults to `1812`.
func (o RadiusOutput) AuthPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Radius) pulumi.IntPtrOutput { return v.AuthPort }).(pulumi.IntPtrOutput)
}

// RAIDUS server enabled. Defaults to `true`.
func (o RadiusOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Radius) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Statistics will be collected from connected clients at this interval. Defaults to `3600`.
func (o RadiusOutput) InterimUpdateInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Radius) pulumi.IntPtrOutput { return v.InterimUpdateInterval }).(pulumi.IntPtrOutput)
}

// RAIDUS secret passphrase. Defaults to “.
func (o RadiusOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Radius) pulumi.StringPtrOutput { return v.Secret }).(pulumi.StringPtrOutput)
}

// The name of the site to associate the settings with.
func (o RadiusOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *Radius) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

// Encrypt communication between the server and the client. Defaults to `true`.
func (o RadiusOutput) TunneledReply() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Radius) pulumi.BoolPtrOutput { return v.TunneledReply }).(pulumi.BoolPtrOutput)
}

type RadiusArrayOutput struct{ *pulumi.OutputState }

func (RadiusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Radius)(nil)).Elem()
}

func (o RadiusArrayOutput) ToRadiusArrayOutput() RadiusArrayOutput {
	return o
}

func (o RadiusArrayOutput) ToRadiusArrayOutputWithContext(ctx context.Context) RadiusArrayOutput {
	return o
}

func (o RadiusArrayOutput) Index(i pulumi.IntInput) RadiusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Radius {
		return vs[0].([]*Radius)[vs[1].(int)]
	}).(RadiusOutput)
}

type RadiusMapOutput struct{ *pulumi.OutputState }

func (RadiusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Radius)(nil)).Elem()
}

func (o RadiusMapOutput) ToRadiusMapOutput() RadiusMapOutput {
	return o
}

func (o RadiusMapOutput) ToRadiusMapOutputWithContext(ctx context.Context) RadiusMapOutput {
	return o
}

func (o RadiusMapOutput) MapIndex(k pulumi.StringInput) RadiusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Radius {
		return vs[0].(map[string]*Radius)[vs[1].(string)]
	}).(RadiusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RadiusInput)(nil)).Elem(), &Radius{})
	pulumi.RegisterInputType(reflect.TypeOf((*RadiusArrayInput)(nil)).Elem(), RadiusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RadiusMapInput)(nil)).Elem(), RadiusMap{})
	pulumi.RegisterOutputType(RadiusOutput{})
	pulumi.RegisterOutputType(RadiusArrayOutput{})
	pulumi.RegisterOutputType(RadiusMapOutput{})
}
