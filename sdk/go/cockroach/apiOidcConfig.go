// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cockroach

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configuration to allow external OIDC providers to issue tokens for use with CC API.
type ApiOidcConfig struct {
	pulumi.CustomResourceState

	// The audience that CC API should accept for this API OIDC Configuration.
	Audience pulumi.StringOutput `pulumi:"audience"`
	// The JWT claim that should be used as the user identifier. Defaults to the subject.
	Claim pulumi.StringOutput `pulumi:"claim"`
	// The mapping rules to convert token user identifiers into a new form.
	IdentityMaps ApiOidcConfigIdentityMapArrayOutput `pulumi:"identityMaps"`
	// The issuer of tokens for the API OIDC Configuration. Usually this is a url.
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// The JSON Web Key Set used to check the signature of the JWTs.
	Jwks pulumi.StringOutput `pulumi:"jwks"`
}

// NewApiOidcConfig registers a new resource with the given unique name, arguments, and options.
func NewApiOidcConfig(ctx *pulumi.Context,
	name string, args *ApiOidcConfigArgs, opts ...pulumi.ResourceOption) (*ApiOidcConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Audience == nil {
		return nil, errors.New("invalid value for required argument 'Audience'")
	}
	if args.Issuer == nil {
		return nil, errors.New("invalid value for required argument 'Issuer'")
	}
	if args.Jwks == nil {
		return nil, errors.New("invalid value for required argument 'Jwks'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ApiOidcConfig
	err := ctx.RegisterResource("cockroach:index/apiOidcConfig:ApiOidcConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiOidcConfig gets an existing ApiOidcConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiOidcConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiOidcConfigState, opts ...pulumi.ResourceOption) (*ApiOidcConfig, error) {
	var resource ApiOidcConfig
	err := ctx.ReadResource("cockroach:index/apiOidcConfig:ApiOidcConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiOidcConfig resources.
type apiOidcConfigState struct {
	// The audience that CC API should accept for this API OIDC Configuration.
	Audience *string `pulumi:"audience"`
	// The JWT claim that should be used as the user identifier. Defaults to the subject.
	Claim *string `pulumi:"claim"`
	// The mapping rules to convert token user identifiers into a new form.
	IdentityMaps []ApiOidcConfigIdentityMap `pulumi:"identityMaps"`
	// The issuer of tokens for the API OIDC Configuration. Usually this is a url.
	Issuer *string `pulumi:"issuer"`
	// The JSON Web Key Set used to check the signature of the JWTs.
	Jwks *string `pulumi:"jwks"`
}

type ApiOidcConfigState struct {
	// The audience that CC API should accept for this API OIDC Configuration.
	Audience pulumi.StringPtrInput
	// The JWT claim that should be used as the user identifier. Defaults to the subject.
	Claim pulumi.StringPtrInput
	// The mapping rules to convert token user identifiers into a new form.
	IdentityMaps ApiOidcConfigIdentityMapArrayInput
	// The issuer of tokens for the API OIDC Configuration. Usually this is a url.
	Issuer pulumi.StringPtrInput
	// The JSON Web Key Set used to check the signature of the JWTs.
	Jwks pulumi.StringPtrInput
}

func (ApiOidcConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOidcConfigState)(nil)).Elem()
}

type apiOidcConfigArgs struct {
	// The audience that CC API should accept for this API OIDC Configuration.
	Audience string `pulumi:"audience"`
	// The JWT claim that should be used as the user identifier. Defaults to the subject.
	Claim *string `pulumi:"claim"`
	// The mapping rules to convert token user identifiers into a new form.
	IdentityMaps []ApiOidcConfigIdentityMap `pulumi:"identityMaps"`
	// The issuer of tokens for the API OIDC Configuration. Usually this is a url.
	Issuer string `pulumi:"issuer"`
	// The JSON Web Key Set used to check the signature of the JWTs.
	Jwks string `pulumi:"jwks"`
}

// The set of arguments for constructing a ApiOidcConfig resource.
type ApiOidcConfigArgs struct {
	// The audience that CC API should accept for this API OIDC Configuration.
	Audience pulumi.StringInput
	// The JWT claim that should be used as the user identifier. Defaults to the subject.
	Claim pulumi.StringPtrInput
	// The mapping rules to convert token user identifiers into a new form.
	IdentityMaps ApiOidcConfigIdentityMapArrayInput
	// The issuer of tokens for the API OIDC Configuration. Usually this is a url.
	Issuer pulumi.StringInput
	// The JSON Web Key Set used to check the signature of the JWTs.
	Jwks pulumi.StringInput
}

func (ApiOidcConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOidcConfigArgs)(nil)).Elem()
}

type ApiOidcConfigInput interface {
	pulumi.Input

	ToApiOidcConfigOutput() ApiOidcConfigOutput
	ToApiOidcConfigOutputWithContext(ctx context.Context) ApiOidcConfigOutput
}

func (*ApiOidcConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOidcConfig)(nil)).Elem()
}

func (i *ApiOidcConfig) ToApiOidcConfigOutput() ApiOidcConfigOutput {
	return i.ToApiOidcConfigOutputWithContext(context.Background())
}

func (i *ApiOidcConfig) ToApiOidcConfigOutputWithContext(ctx context.Context) ApiOidcConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOidcConfigOutput)
}

// ApiOidcConfigArrayInput is an input type that accepts ApiOidcConfigArray and ApiOidcConfigArrayOutput values.
// You can construct a concrete instance of `ApiOidcConfigArrayInput` via:
//
//	ApiOidcConfigArray{ ApiOidcConfigArgs{...} }
type ApiOidcConfigArrayInput interface {
	pulumi.Input

	ToApiOidcConfigArrayOutput() ApiOidcConfigArrayOutput
	ToApiOidcConfigArrayOutputWithContext(context.Context) ApiOidcConfigArrayOutput
}

type ApiOidcConfigArray []ApiOidcConfigInput

func (ApiOidcConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiOidcConfig)(nil)).Elem()
}

func (i ApiOidcConfigArray) ToApiOidcConfigArrayOutput() ApiOidcConfigArrayOutput {
	return i.ToApiOidcConfigArrayOutputWithContext(context.Background())
}

func (i ApiOidcConfigArray) ToApiOidcConfigArrayOutputWithContext(ctx context.Context) ApiOidcConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOidcConfigArrayOutput)
}

// ApiOidcConfigMapInput is an input type that accepts ApiOidcConfigMap and ApiOidcConfigMapOutput values.
// You can construct a concrete instance of `ApiOidcConfigMapInput` via:
//
//	ApiOidcConfigMap{ "key": ApiOidcConfigArgs{...} }
type ApiOidcConfigMapInput interface {
	pulumi.Input

	ToApiOidcConfigMapOutput() ApiOidcConfigMapOutput
	ToApiOidcConfigMapOutputWithContext(context.Context) ApiOidcConfigMapOutput
}

type ApiOidcConfigMap map[string]ApiOidcConfigInput

func (ApiOidcConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiOidcConfig)(nil)).Elem()
}

func (i ApiOidcConfigMap) ToApiOidcConfigMapOutput() ApiOidcConfigMapOutput {
	return i.ToApiOidcConfigMapOutputWithContext(context.Background())
}

func (i ApiOidcConfigMap) ToApiOidcConfigMapOutputWithContext(ctx context.Context) ApiOidcConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOidcConfigMapOutput)
}

type ApiOidcConfigOutput struct{ *pulumi.OutputState }

func (ApiOidcConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOidcConfig)(nil)).Elem()
}

func (o ApiOidcConfigOutput) ToApiOidcConfigOutput() ApiOidcConfigOutput {
	return o
}

func (o ApiOidcConfigOutput) ToApiOidcConfigOutputWithContext(ctx context.Context) ApiOidcConfigOutput {
	return o
}

// The audience that CC API should accept for this API OIDC Configuration.
func (o ApiOidcConfigOutput) Audience() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiOidcConfig) pulumi.StringOutput { return v.Audience }).(pulumi.StringOutput)
}

// The JWT claim that should be used as the user identifier. Defaults to the subject.
func (o ApiOidcConfigOutput) Claim() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiOidcConfig) pulumi.StringOutput { return v.Claim }).(pulumi.StringOutput)
}

// The mapping rules to convert token user identifiers into a new form.
func (o ApiOidcConfigOutput) IdentityMaps() ApiOidcConfigIdentityMapArrayOutput {
	return o.ApplyT(func(v *ApiOidcConfig) ApiOidcConfigIdentityMapArrayOutput { return v.IdentityMaps }).(ApiOidcConfigIdentityMapArrayOutput)
}

// The issuer of tokens for the API OIDC Configuration. Usually this is a url.
func (o ApiOidcConfigOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiOidcConfig) pulumi.StringOutput { return v.Issuer }).(pulumi.StringOutput)
}

// The JSON Web Key Set used to check the signature of the JWTs.
func (o ApiOidcConfigOutput) Jwks() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiOidcConfig) pulumi.StringOutput { return v.Jwks }).(pulumi.StringOutput)
}

type ApiOidcConfigArrayOutput struct{ *pulumi.OutputState }

func (ApiOidcConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiOidcConfig)(nil)).Elem()
}

func (o ApiOidcConfigArrayOutput) ToApiOidcConfigArrayOutput() ApiOidcConfigArrayOutput {
	return o
}

func (o ApiOidcConfigArrayOutput) ToApiOidcConfigArrayOutputWithContext(ctx context.Context) ApiOidcConfigArrayOutput {
	return o
}

func (o ApiOidcConfigArrayOutput) Index(i pulumi.IntInput) ApiOidcConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiOidcConfig {
		return vs[0].([]*ApiOidcConfig)[vs[1].(int)]
	}).(ApiOidcConfigOutput)
}

type ApiOidcConfigMapOutput struct{ *pulumi.OutputState }

func (ApiOidcConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiOidcConfig)(nil)).Elem()
}

func (o ApiOidcConfigMapOutput) ToApiOidcConfigMapOutput() ApiOidcConfigMapOutput {
	return o
}

func (o ApiOidcConfigMapOutput) ToApiOidcConfigMapOutputWithContext(ctx context.Context) ApiOidcConfigMapOutput {
	return o
}

func (o ApiOidcConfigMapOutput) MapIndex(k pulumi.StringInput) ApiOidcConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiOidcConfig {
		return vs[0].(map[string]*ApiOidcConfig)[vs[1].(string)]
	}).(ApiOidcConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOidcConfigInput)(nil)).Elem(), &ApiOidcConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOidcConfigArrayInput)(nil)).Elem(), ApiOidcConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOidcConfigMapInput)(nil)).Elem(), ApiOidcConfigMap{})
	pulumi.RegisterOutputType(ApiOidcConfigOutput{})
	pulumi.RegisterOutputType(ApiOidcConfigArrayOutput{})
	pulumi.RegisterOutputType(ApiOidcConfigMapOutput{})
}