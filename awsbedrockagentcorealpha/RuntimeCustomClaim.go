package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a custom claim validation configuration for Runtime JWT authorizers.
//
// Custom claims allow you to validate additional fields in JWT tokens beyond
// the standard audience, client, and scope validations.
//
// Example:
//   var userPool UserPool
//   var userPoolClient UserPoolClient
//   var anotherUserPoolClient UserPoolClient
//
//
//   repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
//   	RepositoryName: jsii.String("test-agent-runtime"),
//   })
//   agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))
//
//   // Optional: Create custom claims for additional validation
//   customClaims := []RuntimeCustomClaim{
//   	agentcore.RuntimeCustomClaim_WithStringValue(jsii.String("department"), jsii.String("engineering")),
//   	agentcore.RuntimeCustomClaim_WithStringArrayValue(jsii.String("roles"), []*string{
//   		jsii.String("admin"),
//   	}, agentcore.CustomClaimOperator_CONTAINS),
//   	agentcore.RuntimeCustomClaim_WithStringArrayValue(jsii.String("permissions"), []*string{
//   		jsii.String("read"),
//   		jsii.String("write"),
//   	}, agentcore.CustomClaimOperator_CONTAINS_ANY),
//   }
//
//   runtime := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
//   	RuntimeName: jsii.String("myAgent"),
//   	AgentRuntimeArtifact: agentRuntimeArtifact,
//   	AuthorizerConfiguration: agentcore.RuntimeAuthorizerConfiguration_UsingCognito(userPool, []IUserPoolClient{
//   		userPoolClient,
//   		anotherUserPoolClient,
//   	}, []*string{
//   		jsii.String("audience1"),
//   	}, []*string{
//   		jsii.String("read"),
//   		jsii.String("write"),
//   	}, customClaims),
//   })
//
// Experimental.
type RuntimeCustomClaim interface {
}

// The jsii proxy struct for RuntimeCustomClaim
type jsiiProxy_RuntimeCustomClaim struct {
	_ byte // padding
}

// Create a custom claim with a string array value.
//
// String array claims can use CONTAINS (default) or CONTAINS_ANY operator.
//
// Returns: A RuntimeCustomClaim configured for string array validation.
// Experimental.
func RuntimeCustomClaim_WithStringArrayValue(name *string, values *[]*string, operator CustomClaimOperator) RuntimeCustomClaim {
	_init_.Initialize()

	if err := validateRuntimeCustomClaim_WithStringArrayValueParameters(name, values); err != nil {
		panic(err)
	}
	var returns RuntimeCustomClaim

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeCustomClaim",
		"withStringArrayValue",
		[]interface{}{name, values, operator},
		&returns,
	)

	return returns
}

// Create a custom claim with a string value.
//
// String claims must use the EQUALS operator.
//
// Returns: A RuntimeCustomClaim configured for string validation.
// Experimental.
func RuntimeCustomClaim_WithStringValue(name *string, value *string) RuntimeCustomClaim {
	_init_.Initialize()

	if err := validateRuntimeCustomClaim_WithStringValueParameters(name, value); err != nil {
		panic(err)
	}
	var returns RuntimeCustomClaim

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeCustomClaim",
		"withStringValue",
		[]interface{}{name, value},
		&returns,
	)

	return returns
}

