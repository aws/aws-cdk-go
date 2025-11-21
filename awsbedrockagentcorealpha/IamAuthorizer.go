package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AWS IAM authorizer configuration implementation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   iamAuthorizer := bedrock_agentcore_alpha.NewIamAuthorizer()
//
// Experimental.
type IamAuthorizer interface {
	IGatewayAuthorizerConfig
	// The authorizer type.
	// Experimental.
	AuthorizerType() GatewayAuthorizerType
}

// The jsii proxy struct for IamAuthorizer
type jsiiProxy_IamAuthorizer struct {
	jsiiProxy_IGatewayAuthorizerConfig
}

func (j *jsiiProxy_IamAuthorizer) AuthorizerType() GatewayAuthorizerType {
	var returns GatewayAuthorizerType
	_jsii_.Get(
		j,
		"authorizerType",
		&returns,
	)
	return returns
}


// Experimental.
func NewIamAuthorizer() IamAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_IamAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IamAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewIamAuthorizer_Override(i IamAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.IamAuthorizer",
		nil, // no parameters
		i,
	)
}

