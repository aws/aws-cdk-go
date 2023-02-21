// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Current resolution context for tokens.
type IResolveContext interface {
	// Use this postprocessor after the entire token structure has been resolved.
	RegisterPostProcessor(postProcessor IPostProcessor)
	// Resolve an inner object.
	Resolve(x interface{}, options *ResolveChangeContextOptions) interface{}
	// Path in the JSON document that is being constructed.
	DocumentPath() *[]*string
	// True when we are still preparing, false if we're rendering the final output.
	Preparing() *bool
	// The scope from which resolution has been initiated.
	Scope() constructs.IConstruct
}

// The jsii proxy for IResolveContext
type jsiiProxy_IResolveContext struct {
	_ byte // padding
}

func (i *jsiiProxy_IResolveContext) RegisterPostProcessor(postProcessor IPostProcessor) {
	if err := i.validateRegisterPostProcessorParameters(postProcessor); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerPostProcessor",
		[]interface{}{postProcessor},
	)
}

func (i *jsiiProxy_IResolveContext) Resolve(x interface{}, options *ResolveChangeContextOptions) interface{} {
	if err := i.validateResolveParameters(x, options); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{x, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IResolveContext) DocumentPath() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"documentPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolveContext) Preparing() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"preparing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolveContext) Scope() constructs.IConstruct {
	var returns constructs.IConstruct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

