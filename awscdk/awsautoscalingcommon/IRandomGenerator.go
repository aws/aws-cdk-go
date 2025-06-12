package awsautoscalingcommon

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IRandomGenerator interface {
	NextBoolean() *bool
	NextInt(min *float64, max *float64) *float64
}

// The jsii proxy for IRandomGenerator
type jsiiProxy_IRandomGenerator struct {
	_ byte // padding
}

func (i *jsiiProxy_IRandomGenerator) NextBoolean() *bool {
	var returns *bool

	_jsii_.Invoke(
		i,
		"nextBoolean",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRandomGenerator) NextInt(min *float64, max *float64) *float64 {
	if err := i.validateNextIntParameters(min, max); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"nextInt",
		[]interface{}{min, max},
		&returns,
	)

	return returns
}

