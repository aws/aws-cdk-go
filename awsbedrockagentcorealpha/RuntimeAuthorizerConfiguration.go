package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Abstract base class for runtime authorizer configurations.
//
// Provides static factory methods to create different authentication types.
//
// Example:
//   repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
//   	RepositoryName: jsii.String("test-agent-runtime"),
//   })
//   agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))
//
//   runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
//   	RuntimeName: jsii.String("myAgent"),
//   	AgentRuntimeArtifact: agentRuntimeArtifact,
//   	AuthorizerConfiguration: agentcore.RuntimeAuthorizerConfiguration_UsingCognito(jsii.String("us-west-2_ABC123"), jsii.String("client123"), jsii.String("us-west-2")),
//   })
//
// Experimental.
type RuntimeAuthorizerConfiguration interface {
}

// The jsii proxy struct for RuntimeAuthorizerConfiguration
type jsiiProxy_RuntimeAuthorizerConfiguration struct {
	_ byte // padding
}

// Experimental.
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
// Experimental.
func RuntimeAuthorizerConfiguration_UsingCognito(userPoolId *string, clientId *string, region *string, allowedAudience *[]*string) RuntimeAuthorizerConfiguration {
	_init_.Initialize()

	if err := validateRuntimeAuthorizerConfiguration_UsingCognitoParameters(userPoolId, clientId); err != nil {
		panic(err)
	}
	var returns RuntimeAuthorizerConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeAuthorizerConfiguration",
		"usingCognito",
		[]interface{}{userPoolId, clientId, region, allowedAudience},
		&returns,
	)

	return returns
}

// Use IAM authentication (default).
//
// Requires AWS credentials to sign requests using SigV4.
//
// Returns: RuntimeAuthorizerConfiguration for IAM authentication.
// Experimental.
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
// Experimental.
func RuntimeAuthorizerConfiguration_UsingJWT(discoveryUrl *string, allowedClients *[]*string, allowedAudience *[]*string) RuntimeAuthorizerConfiguration {
	_init_.Initialize()

	if err := validateRuntimeAuthorizerConfiguration_UsingJWTParameters(discoveryUrl); err != nil {
		panic(err)
	}
	var returns RuntimeAuthorizerConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeAuthorizerConfiguration",
		"usingJWT",
		[]interface{}{discoveryUrl, allowedClients, allowedAudience},
		&returns,
	)

	return returns
}

// Use OAuth 2.0 authentication. Supports various OAuth providers.
//
// Returns: RuntimeAuthorizerConfiguration for OAuth authentication.
// Experimental.
func RuntimeAuthorizerConfiguration_UsingOAuth(discoveryUrl *string, clientId *string, allowedAudience *[]*string) RuntimeAuthorizerConfiguration {
	_init_.Initialize()

	if err := validateRuntimeAuthorizerConfiguration_UsingOAuthParameters(discoveryUrl, clientId); err != nil {
		panic(err)
	}
	var returns RuntimeAuthorizerConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeAuthorizerConfiguration",
		"usingOAuth",
		[]interface{}{discoveryUrl, clientId, allowedAudience},
		&returns,
	)

	return returns
}

