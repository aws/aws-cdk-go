package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Abstract interface for gateway protocol configuration.
// Experimental.
type IGatewayProtocolConfig interface {
	// The protocol type.
	// Experimental.
	ProtocolType() *string
}

// The jsii proxy for IGatewayProtocolConfig
type jsiiProxy_IGatewayProtocolConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_IGatewayProtocolConfig) ProtocolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolType",
		&returns,
	)
	return returns
}

