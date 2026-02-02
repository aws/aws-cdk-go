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
//   repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
//   	RepositoryName: jsii.String("test-agent-runtime"),
//   })
//   agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))
//
//   // String claim - validates that the claim exactly equals the specified value
//   // Uses EQUALS operator automatically
//   departmentClaim := agentcore.RuntimeCustomClaim_WithStringValue(jsii.String("department"), jsii.String("engineering"))
//
//   // String array claim with CONTAINS operator (default)
//   // Validates that the claim array contains a specific string value
//   // IMPORTANT: CONTAINS requires exactly one value in the array parameter
//   rolesClaim := agentcore.RuntimeCustomClaim_WithStringArrayValue(jsii.String("roles"), []*string{
//   	jsii.String("admin"),
//   })
//
//   // String array claim with CONTAINS_ANY operator
//   // Validates that the claim array contains at least one of the specified values
//   // Use this when you want to check for multiple possible values
//   permissionsClaim := agentcore.RuntimeCustomClaim_WithStringArrayValue(jsii.String("permissions"), []*string{
//   	jsii.String("read"),
//   	jsii.String("write"),
//   }, agentcore.CustomClaimOperator_CONTAINS_ANY)
//
//   // Use custom claims in authorizer configuration
//   runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
//   	RuntimeName: jsii.String("myAgent"),
//   	AgentRuntimeArtifact: agentRuntimeArtifact,
//   	AuthorizerConfiguration: agentcore.RuntimeAuthorizerConfiguration_UsingJWT(jsii.String("https://example.com/.well-known/openid-configuration"), []*string{
//   		jsii.String("client1"),
//   		jsii.String("client2"),
//   	}, []*string{
//   		jsii.String("audience1"),
//   	}, []*string{
//   		jsii.String("read"),
//   		jsii.String("write"),
//   	}, []RuntimeCustomClaim{
//   		departmentClaim,
//   		rolesClaim,
//   		permissionsClaim,
//   	}),
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
// Experimental.
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

