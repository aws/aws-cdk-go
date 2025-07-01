package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Access log destination for a RestApi Stage.
type IAccessLogDestination interface {
	// Binds this destination to the RestApi Stage.
	Bind(stage IStage) *AccessLogDestinationConfig
}

// The jsii proxy for IAccessLogDestination
type jsiiProxy_IAccessLogDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_IAccessLogDestination) Bind(stage IStage) *AccessLogDestinationConfig {
	if err := i.validateBindParameters(stage); err != nil {
		panic(err)
	}
	var returns *AccessLogDestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{stage},
		&returns,
	)

	return returns
}

