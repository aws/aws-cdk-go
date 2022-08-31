package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Access log destination for a RestApi Stage.
// Experimental.
type IAccessLogDestination interface {
	// Binds this destination to the RestApi Stage.
	// Experimental.
	Bind(stage IStage) *AccessLogDestinationConfig
}

// The jsii proxy for IAccessLogDestination
type jsiiProxy_IAccessLogDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_IAccessLogDestination) Bind(stage IStage) *AccessLogDestinationConfig {
	var returns *AccessLogDestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{stage},
		&returns,
	)

	return returns
}

