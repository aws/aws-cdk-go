package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// No authorization configuration implementation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   noAuthAuthorizer := awscdk.Aws_bedrockagentcore.NewNoAuthAuthorizer()
//
type NoAuthAuthorizer interface {
	IGatewayAuthorizerConfig
	// The authorizer type.
	AuthorizerType() GatewayAuthorizerType
}

// The jsii proxy struct for NoAuthAuthorizer
type jsiiProxy_NoAuthAuthorizer struct {
	jsiiProxy_IGatewayAuthorizerConfig
}

func (j *jsiiProxy_NoAuthAuthorizer) AuthorizerType() GatewayAuthorizerType {
	var returns GatewayAuthorizerType
	_jsii_.Get(
		j,
		"authorizerType",
		&returns,
	)
	return returns
}


func NewNoAuthAuthorizer() NoAuthAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_NoAuthAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.NoAuthAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewNoAuthAuthorizer_Override(n NoAuthAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.NoAuthAuthorizer",
		nil, // no parameters
		n,
	)
}

