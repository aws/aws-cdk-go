package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for Stack Synthesizers that can be used for more than one stack.
//
// Regular `IStackSynthesizer` instances can only be bound to a Stack once.
// `IReusableStackSynthesizer` instances.
//
// For backwards compatibility reasons, this class inherits from
// `IStackSynthesizer`, but if an object implements `IReusableStackSynthesizer`,
// no other methods than `reusableBind()` will be called.
type IReusableStackSynthesizer interface {
	IStackSynthesizer
	// Produce a bound Stack Synthesizer for the given stack.
	//
	// This method may be called more than once on the same object.
	ReusableBind(stack Stack) IBoundStackSynthesizer
}

// The jsii proxy for IReusableStackSynthesizer
type jsiiProxy_IReusableStackSynthesizer struct {
	jsiiProxy_IStackSynthesizer
}

func (i *jsiiProxy_IReusableStackSynthesizer) ReusableBind(stack Stack) IBoundStackSynthesizer {
	if err := i.validateReusableBindParameters(stack); err != nil {
		panic(err)
	}
	var returns IBoundStackSynthesizer

	_jsii_.Invoke(
		i,
		"reusableBind",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

