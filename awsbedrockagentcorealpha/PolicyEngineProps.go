package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a PolicyEngine resource.
//
// Example:
//   policyEngine := agentcore.NewPolicyEngine(this, jsii.String("MyEngine"), &PolicyEngineProps{
//   	PolicyEngineName: jsii.String("my_engine"),
//   })
//
//   lambdaRole := iam.NewRole(this, jsii.String("LambdaRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   })
//
//   // Grant read permissions
//   policyEngine.GrantRead(lambdaRole)
//
//   // Grant evaluation permissions
//   policyEngine.GrantEvaluate(lambdaRole)
//
// Experimental.
type PolicyEngineProps struct {
	// Optional description for the policy engine.
	//
	// Maximum 4,096 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Custom KMS key for encryption.
	//
	// [disable-awslint:prefer-ref-interface].
	// Default: - AWS owned key.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The name of the policy engine.
	//
	// Valid characters: a-z, A-Z, 0-9, _ (underscore)
	// Must start with a letter, 1-48 characters
	// Pattern: ^[A-Za-z][A-Za-z0-9_]*$.
	// Default: - Auto-generated unique name.
	//
	// Experimental.
	PolicyEngineName *string `field:"optional" json:"policyEngineName" yaml:"policyEngineName"`
	// Tags for the policy engine.
	//
	// Maximum 50 tags.
	// Default: - No tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

