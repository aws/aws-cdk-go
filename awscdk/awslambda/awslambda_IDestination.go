package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// A Lambda destination.
// Experimental.
type IDestination interface {
	// Binds this destination to the Lambda function.
	// Experimental.
	Bind(scope awscdk.Construct, fn IFunction, options *DestinationOptions) *DestinationConfig
}

// The jsii proxy for IDestination
type jsiiProxy_IDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_IDestination) Bind(scope awscdk.Construct, fn IFunction, options *DestinationOptions) *DestinationConfig {
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

