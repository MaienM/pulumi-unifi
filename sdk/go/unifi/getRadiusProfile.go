// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package unifi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `RadiusProfile` data source can be used to retrieve the ID for a RADIUS profile by name.
func LookupRadiusProfile(ctx *pulumi.Context, args *LookupRadiusProfileArgs, opts ...pulumi.InvokeOption) (*LookupRadiusProfileResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRadiusProfileResult
	err := ctx.Invoke("unifi:index/getRadiusProfile:getRadiusProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRadiusProfile.
type LookupRadiusProfileArgs struct {
	// The name of the RADIUS profile to look up. Defaults to `Default`.
	Name *string `pulumi:"name"`
	// The name of the site the RADIUS profile is associated with.
	Site *string `pulumi:"site"`
}

// A collection of values returned by getRadiusProfile.
type LookupRadiusProfileResult struct {
	// The ID of this AP group.
	Id string `pulumi:"id"`
	// The name of the RADIUS profile to look up. Defaults to `Default`.
	Name *string `pulumi:"name"`
	// The name of the site the RADIUS profile is associated with.
	Site string `pulumi:"site"`
}

func LookupRadiusProfileOutput(ctx *pulumi.Context, args LookupRadiusProfileOutputArgs, opts ...pulumi.InvokeOption) LookupRadiusProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRadiusProfileResult, error) {
			args := v.(LookupRadiusProfileArgs)
			r, err := LookupRadiusProfile(ctx, &args, opts...)
			var s LookupRadiusProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRadiusProfileResultOutput)
}

// A collection of arguments for invoking getRadiusProfile.
type LookupRadiusProfileOutputArgs struct {
	// The name of the RADIUS profile to look up. Defaults to `Default`.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The name of the site the RADIUS profile is associated with.
	Site pulumi.StringPtrInput `pulumi:"site"`
}

func (LookupRadiusProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRadiusProfileArgs)(nil)).Elem()
}

// A collection of values returned by getRadiusProfile.
type LookupRadiusProfileResultOutput struct{ *pulumi.OutputState }

func (LookupRadiusProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRadiusProfileResult)(nil)).Elem()
}

func (o LookupRadiusProfileResultOutput) ToLookupRadiusProfileResultOutput() LookupRadiusProfileResultOutput {
	return o
}

func (o LookupRadiusProfileResultOutput) ToLookupRadiusProfileResultOutputWithContext(ctx context.Context) LookupRadiusProfileResultOutput {
	return o
}

// The ID of this AP group.
func (o LookupRadiusProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRadiusProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the RADIUS profile to look up. Defaults to `Default`.
func (o LookupRadiusProfileResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRadiusProfileResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The name of the site the RADIUS profile is associated with.
func (o LookupRadiusProfileResultOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRadiusProfileResult) string { return v.Site }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRadiusProfileResultOutput{})
}
