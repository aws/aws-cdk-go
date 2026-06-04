package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
)

// Abstract base class for runtime authorizer configurations.
//
// Provides static factory methods to create different authentication types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var runtimeCustomClaim RuntimeCustomClaim
//   var userPool UserPool
//   var userPoolClient UserPoolClient
//
//   runtimeAuthorizerConfiguration := bedrock_agentcore_alpha.RuntimeAuthorizerConfiguration_UsingCognito(userPool, []IUserPoolClient{
//   	userPoolClient,
//   }, []*string{
//   	jsii.String("allowedAudience"),
//   }, []*string{
//   	jsii.String("allowedScopes"),
//   }, []RuntimeCustomClaim{
//   	runtimeCustomClaim,
//   })
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type RuntimeAuthorizerConfiguration interface {
}

// The jsii proxy struct for RuntimeAuthorizerConfiguration
type jsiiProxy_RuntimeAuthorizerConfiguration struct {
	_ byte // padding
}

// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func NewRuntimeAuthorizerConfiguration_Override(r RuntimeAuthorizerConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeAuthorizerConfiguration",
		nil, // no parameters
		r,
	)
}

// Use AWS Cognito User Pool authentication.
//
// Validates Cognito-issued JWT tokens.
//
// Returns: RuntimeAuthorizerConfiguration for Cognito authentication.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func RuntimeAuthorizerConfiguration_UsingCognito(userPool awscognito.IUserPool, userPoolClients *[]awscognito.IUserPoolClient, allowedAudience *[]*string, allowedScopes *[]*string, customClaims *[]RuntimeCustomClaim) RuntimeAuthorizerConfiguration {
	_init_.Initialize()

	if err := validateRuntimeAuthorizerConfiguration_UsingCognitoParameters(userPool, userPoolClients); err != nil {
		panic(err)
	}
	var returns RuntimeAuthorizerConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeAuthorizerConfiguration",
		"usingCognito",
		[]interface{}{userPool, userPoolClients, allowedAudience, allowedScopes, customClaims},
		&returns,
	)

	return returns
}

// Use IAM authentication (default).
//
// Requires AWS credentials to sign requests using SigV4.
//
// Returns: RuntimeAuthorizerConfiguration for IAM authentication.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func RuntimeAuthorizerConfiguration_UsingIAM() RuntimeAuthorizerConfiguration {
	_init_.Initialize()

	var returns RuntimeAuthorizerConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeAuthorizerConfiguration",
		"usingIAM",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Use custom JWT authentication.
//
// Validates JWT tokens against the specified OIDC provider.
//
// Returns: RuntimeAuthorizerConfiguration for JWT authentication.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func RuntimeAuthorizerConfiguration_UsingJWT(discoveryUrl *string, allowedClients *[]*string, allowedAudience *[]*string, allowedScopes *[]*string, customClaims *[]RuntimeCustomClaim) RuntimeAuthorizerConfiguration {
	_init_.Initialize()

	if err := validateRuntimeAuthorizerConfiguration_UsingJWTParameters(discoveryUrl); err != nil {
		panic(err)
	}
	var returns RuntimeAuthorizerConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeAuthorizerConfiguration",
		"usingJWT",
		[]interface{}{discoveryUrl, allowedClients, allowedAudience, allowedScopes, customClaims},
		&returns,
	)

	return returns
}

// Use OAuth 2.0 authentication. Supports various OAuth providers.
//
// Returns: RuntimeAuthorizerConfiguration for OAuth authentication.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func RuntimeAuthorizerConfiguration_UsingOAuth(discoveryUrl *string, clientId *string, allowedAudience *[]*string, allowedScopes *[]*string, customClaims *[]RuntimeCustomClaim) RuntimeAuthorizerConfiguration {
	_init_.Initialize()

	if err := validateRuntimeAuthorizerConfiguration_UsingOAuthParameters(discoveryUrl, clientId); err != nil {
		panic(err)
	}
	var returns RuntimeAuthorizerConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeAuthorizerConfiguration",
		"usingOAuth",
		[]interface{}{discoveryUrl, clientId, allowedAudience, allowedScopes, customClaims},
		&returns,
	)

	return returns
}

