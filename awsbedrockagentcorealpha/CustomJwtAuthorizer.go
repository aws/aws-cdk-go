package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Custom JWT authorizer configuration implementation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var gatewayCustomClaim GatewayCustomClaim
//
//   customJwtAuthorizer := bedrock_agentcore_alpha.NewCustomJwtAuthorizer(&CustomJwtConfiguration{
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
// Experimental.
type CustomJwtAuthorizer interface {
	IGatewayAuthorizerConfig
	// The authorizer type.
	// Experimental.
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


// Experimental.
func NewCustomJwtAuthorizer(config *CustomJwtConfiguration) CustomJwtAuthorizer {
	_init_.Initialize()

	if err := validateNewCustomJwtAuthorizerParameters(config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomJwtAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CustomJwtAuthorizer",
		[]interface{}{config},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomJwtAuthorizer_Override(c CustomJwtAuthorizer, config *CustomJwtConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CustomJwtAuthorizer",
		[]interface{}{config},
		c,
	)
}

