package assertions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/cxapi"
)

// Suite of assertions that can be run on a CDK Stack.
//
// Focused on asserting annotations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//
//   annotations := awscdk.Assertions.annotations.fromStack(stack)
//
// Experimental.
type Annotations interface {
	// Get the set of matching errors of a given construct path and message.
	// Experimental.
	FindError(constructPath *string, message interface{}) *[]*cxapi.SynthesisMessage
	// Get the set of matching infos of a given construct path and message.
	// Experimental.
	FindInfo(constructPath *string, message interface{}) *[]*cxapi.SynthesisMessage
	// Get the set of matching warning of a given construct path and message.
	// Experimental.
	FindWarning(constructPath *string, message interface{}) *[]*cxapi.SynthesisMessage
	// Assert that an error with the given message exists in the synthesized CDK `Stack`.
	// Experimental.
	HasError(constructPath *string, message interface{})
	// Assert that an info with the given message exists in the synthesized CDK `Stack`.
	// Experimental.
	HasInfo(constructPath *string, message interface{})
	// Assert that an error with the given message does not exist in the synthesized CDK `Stack`.
	// Experimental.
	HasNoError(constructPath *string, message interface{})
	// Assert that an info with the given message does not exist in the synthesized CDK `Stack`.
	// Experimental.
	HasNoInfo(constructPath *string, message interface{})
	// Assert that an warning with the given message does not exist in the synthesized CDK `Stack`.
	// Experimental.
	HasNoWarning(constructPath *string, message interface{})
	// Assert that an warning with the given message exists in the synthesized CDK `Stack`.
	// Experimental.
	HasWarning(constructPath *string, message interface{})
}

// The jsii proxy struct for Annotations
type jsiiProxy_Annotations struct {
	_ byte // padding
}

// Base your assertions on the messages returned by a synthesized CDK `Stack`.
// Experimental.
func Annotations_FromStack(stack awscdk.Stack) Annotations {
	_init_.Initialize()

	if err := validateAnnotations_FromStackParameters(stack); err != nil {
		panic(err)
	}
	var returns Annotations

	_jsii_.StaticInvoke(
		"monocdk.assertions.Annotations",
		"fromStack",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Annotations) FindError(constructPath *string, message interface{}) *[]*cxapi.SynthesisMessage {
	if err := a.validateFindErrorParameters(constructPath, message); err != nil {
		panic(err)
	}
	var returns *[]*cxapi.SynthesisMessage

	_jsii_.Invoke(
		a,
		"findError",
		[]interface{}{constructPath, message},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Annotations) FindInfo(constructPath *string, message interface{}) *[]*cxapi.SynthesisMessage {
	if err := a.validateFindInfoParameters(constructPath, message); err != nil {
		panic(err)
	}
	var returns *[]*cxapi.SynthesisMessage

	_jsii_.Invoke(
		a,
		"findInfo",
		[]interface{}{constructPath, message},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Annotations) FindWarning(constructPath *string, message interface{}) *[]*cxapi.SynthesisMessage {
	if err := a.validateFindWarningParameters(constructPath, message); err != nil {
		panic(err)
	}
	var returns *[]*cxapi.SynthesisMessage

	_jsii_.Invoke(
		a,
		"findWarning",
		[]interface{}{constructPath, message},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Annotations) HasError(constructPath *string, message interface{}) {
	if err := a.validateHasErrorParameters(constructPath, message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"hasError",
		[]interface{}{constructPath, message},
	)
}

func (a *jsiiProxy_Annotations) HasInfo(constructPath *string, message interface{}) {
	if err := a.validateHasInfoParameters(constructPath, message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"hasInfo",
		[]interface{}{constructPath, message},
	)
}

func (a *jsiiProxy_Annotations) HasNoError(constructPath *string, message interface{}) {
	if err := a.validateHasNoErrorParameters(constructPath, message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"hasNoError",
		[]interface{}{constructPath, message},
	)
}

func (a *jsiiProxy_Annotations) HasNoInfo(constructPath *string, message interface{}) {
	if err := a.validateHasNoInfoParameters(constructPath, message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"hasNoInfo",
		[]interface{}{constructPath, message},
	)
}

func (a *jsiiProxy_Annotations) HasNoWarning(constructPath *string, message interface{}) {
	if err := a.validateHasNoWarningParameters(constructPath, message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"hasNoWarning",
		[]interface{}{constructPath, message},
	)
}

func (a *jsiiProxy_Annotations) HasWarning(constructPath *string, message interface{}) {
	if err := a.validateHasWarningParameters(constructPath, message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"hasWarning",
		[]interface{}{constructPath, message},
	)
}

