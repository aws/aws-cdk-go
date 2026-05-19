package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Custom JWT authorizer configuration implementation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var gatewayCustomClaim GatewayCustomClaim
//
//   customJwtAuthorizer := awscdk.Aws_bedrockagentcore.NewCustomJwtAuthorizer(&CustomJwtConfiguration{
//   	DiscoveryUrl: jsii.String("discoveryUrl"),
//
//   	// the properties below are optional
//   	AllowedAudience: []*string{
//   		jsii.String("allowedAudience"),
//   	},
//   	AllowedClients: []*string{
//   		jsii.String("allowedClients"),
//   	},
//   	AllowedScopes: []*string{
//   		jsii.String("allowedScopes"),
//   	},
//   	CustomClaims: []GatewayCustomClaim{
//   		gatewayCustomClaim,
//   	},
//   })
//
type CustomJwtAuthorizer interface {
	IGatewayAuthorizerConfig
	// The authorizer type.
	AuthorizerType() GatewayAuthorizerType
}

// The jsii proxy struct for CustomJwtAuthorizer
type jsiiProxy_CustomJwtAuthorizer struct {
	jsiiProxy_IGatewayAuthorizerConfig
}

func (j *jsiiProxy_CustomJwtAuthorizer) AuthorizerType() GatewayAuthorizerType {
	var returns GatewayAuthorizerType
	_jsii_.Get(
		j,
		"authorizerType",
		&returns,
	)
	return returns
}


func NewCustomJwtAuthorizer(config *CustomJwtConfiguration) CustomJwtAuthorizer {
	_init_.Initialize()

	if err := validateNewCustomJwtAuthorizerParameters(config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomJwtAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.CustomJwtAuthorizer",
		[]interface{}{config},
		&j,
	)

	return &j
}

func NewCustomJwtAuthorizer_Override(c CustomJwtAuthorizer, config *CustomJwtConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.CustomJwtAuthorizer",
		[]interface{}{config},
		c,
	)
}

