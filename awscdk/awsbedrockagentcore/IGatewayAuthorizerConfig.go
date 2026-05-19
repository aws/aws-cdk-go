package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Abstract interface for gateway authorizer configuration.
type IGatewayAuthorizerConfig interface {
	// The authorizer type.
	AuthorizerType() GatewayAuthorizerType
}

// The jsii proxy for IGatewayAuthorizerConfig
type jsiiProxy_IGatewayAuthorizerConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_IGatewayAuthorizerConfig) AuthorizerType() GatewayAuthorizerType {
	var returns GatewayAuthorizerType
	_jsii_.Get(
		j,
		"authorizerType",
		&returns,
	)
	return returns
}

