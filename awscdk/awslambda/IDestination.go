package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Lambda destination.
type IDestination interface {
	// Binds this destination to the Lambda function.
	Bind(scope constructs.Construct, fn IFunction, options *DestinationOptions) *DestinationConfig
}

// The jsii proxy for IDestination
type jsiiProxy_IDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_IDestination) Bind(scope constructs.Construct, fn IFunction, options *DestinationOptions) *DestinationConfig {
	if err := i.validateBindParameters(scope, fn, options); err != nil {
		panic(err)
	}
	var returns *DestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, fn, options},
		&returns,
	)

	return returns
}

