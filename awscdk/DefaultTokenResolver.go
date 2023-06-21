package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Default resolver implementation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var fragmentConcatenator iFragmentConcatenator
//
//   defaultTokenResolver := cdk.NewDefaultTokenResolver(fragmentConcatenator)
//
type DefaultTokenResolver interface {
	ITokenResolver
	// Resolve a tokenized list.
	ResolveList(xs *[]*string, context IResolveContext) interface{}
	// Resolve string fragments to Tokens.
	ResolveString(fragments TokenizedStringFragments, context IResolveContext) interface{}
	// Default Token resolution.
	//
	// Resolve the Token, recurse into whatever it returns,
	// then finally post-process it.
	ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{}
}

// The jsii proxy struct for DefaultTokenResolver
type jsiiProxy_DefaultTokenResolver struct {
	jsiiProxy_ITokenResolver
}

func NewDefaultTokenResolver(concat IFragmentConcatenator) DefaultTokenResolver {
	_init_.Initialize()

	if err := validateNewDefaultTokenResolverParameters(concat); err != nil {
		panic(err)
	}
	j := jsiiProxy_DefaultTokenResolver{}

	_jsii_.Create(
		"aws-cdk-lib.DefaultTokenResolver",
		[]interface{}{concat},
		&j,
	)

	return &j
}

func NewDefaultTokenResolver_Override(d DefaultTokenResolver, concat IFragmentConcatenator) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.DefaultTokenResolver",
		[]interface{}{concat},
		d,
	)
}

func (d *jsiiProxy_DefaultTokenResolver) ResolveList(xs *[]*string, context IResolveContext) interface{} {
	if err := d.validateResolveListParameters(xs, context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveList",
		[]interface{}{xs, context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultTokenResolver) ResolveString(fragments TokenizedStringFragments, context IResolveContext) interface{} {
	if err := d.validateResolveStringParameters(fragments, context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveString",
		[]interface{}{fragments, context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultTokenResolver) ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{} {
	if err := d.validateResolveTokenParameters(t, context, postProcessor); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveToken",
		[]interface{}{t, context, postProcessor},
		&returns,
	)

	return returns
}

