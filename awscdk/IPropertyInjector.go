package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// This interface define an inject function that operates on a Construct's Property.
//
// The Construct must have a constructUniqueId to uniquely identify itself.
type IPropertyInjector interface {
	// The injector to be applied to the constructor properties of the Construct.
	Inject(originalProps interface{}, context *InjectionContext) interface{}
	// The unique Id of the Construct class.
	ConstructUniqueId() *string
}

// The jsii proxy for IPropertyInjector
type jsiiProxy_IPropertyInjector struct {
	_ byte // padding
}

func (i *jsiiProxy_IPropertyInjector) Inject(originalProps interface{}, context *InjectionContext) interface{} {
	if err := i.validateInjectParameters(originalProps, context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"inject",
		[]interface{}{originalProps, context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPropertyInjector) ConstructUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructUniqueId",
		&returns,
	)
	return returns
}

