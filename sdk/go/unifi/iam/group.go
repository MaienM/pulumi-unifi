// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `iam.Group` manages a user group (called "client group" in the UI), which can be used to limit bandwidth for groups of users.
//
// ## Import
//
// # import using the ID
//
// ```sh
//
//	$ pulumi import unifi:iam/group:Group wifi 5fe6261995fe130013456a36
//
// ```
type Group struct {
	pulumi.CustomResourceState

	// The name of the user group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The QOS maximum download rate. Defaults to `-1`.
	QosRateMaxDown pulumi.IntPtrOutput `pulumi:"qosRateMaxDown"`
	// The QOS maximum upload rate. Defaults to `-1`.
	QosRateMaxUp pulumi.IntPtrOutput `pulumi:"qosRateMaxUp"`
	// The name of the site to associate the user group with.
	Site pulumi.StringOutput `pulumi:"site"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Group
	err := ctx.RegisterResource("unifi:iam/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("unifi:iam/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The name of the user group.
	Name *string `pulumi:"name"`
	// The QOS maximum download rate. Defaults to `-1`.
	QosRateMaxDown *int `pulumi:"qosRateMaxDown"`
	// The QOS maximum upload rate. Defaults to `-1`.
	QosRateMaxUp *int `pulumi:"qosRateMaxUp"`
	// The name of the site to associate the user group with.
	Site *string `pulumi:"site"`
}

type GroupState struct {
	// The name of the user group.
	Name pulumi.StringPtrInput
	// The QOS maximum download rate. Defaults to `-1`.
	QosRateMaxDown pulumi.IntPtrInput
	// The QOS maximum upload rate. Defaults to `-1`.
	QosRateMaxUp pulumi.IntPtrInput
	// The name of the site to associate the user group with.
	Site pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The name of the user group.
	Name *string `pulumi:"name"`
	// The QOS maximum download rate. Defaults to `-1`.
	QosRateMaxDown *int `pulumi:"qosRateMaxDown"`
	// The QOS maximum upload rate. Defaults to `-1`.
	QosRateMaxUp *int `pulumi:"qosRateMaxUp"`
	// The name of the site to associate the user group with.
	Site *string `pulumi:"site"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The name of the user group.
	Name pulumi.StringPtrInput
	// The QOS maximum download rate. Defaults to `-1`.
	QosRateMaxDown pulumi.IntPtrInput
	// The QOS maximum upload rate. Defaults to `-1`.
	QosRateMaxUp pulumi.IntPtrInput
	// The name of the site to associate the user group with.
	Site pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//	GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//	GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

// The name of the user group.
func (o GroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The QOS maximum download rate. Defaults to `-1`.
func (o GroupOutput) QosRateMaxDown() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.IntPtrOutput { return v.QosRateMaxDown }).(pulumi.IntPtrOutput)
}

// The QOS maximum upload rate. Defaults to `-1`.
func (o GroupOutput) QosRateMaxUp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.IntPtrOutput { return v.QosRateMaxUp }).(pulumi.IntPtrOutput)
}

// The name of the site to associate the user group with.
func (o GroupOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}