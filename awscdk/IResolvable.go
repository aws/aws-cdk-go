// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for values that can be resolvable later.
//
// Tokens are special objects that participate in synthesis.
type IResolvable interface {
	// Produce the Token's value at resolution time.
	Resolve(context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	ToString() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// This may return an array with a single informational element indicating how
	// to get this property populated, if it was skipped for performance reasons.
	CreationStack() *[]*string
	// The type that this token will likely resolve to.
	TypeHint() ResolutionTypeHint
}

// The jsii proxy for IResolvable
type jsiiProxy_IResolvable struct {
	_ byte // padding
}

func (i *jsiiProxy_IResolvable) Resolve(context IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResolvable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IResolvable) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolvable) TypeHint() ResolutionTypeHint {
	var returns ResolutionTypeHint
	_jsii_.Get(
		j,
		"typeHint",
		&returns,
	)
	return returns
}

