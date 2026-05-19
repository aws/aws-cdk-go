package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AWS IAM authorizer configuration implementation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamAuthorizer := awscdk.Aws_bedrockagentcore.NewIamAuthorizer()
//
type IamAuthorizer interface {
	IGatewayAuthorizerConfig
	// The authorizer type.
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


func NewIamAuthorizer() IamAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_IamAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.IamAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewIamAuthorizer_Override(i IamAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.IamAuthorizer",
		nil, // no parameters
		i,
	)
}

