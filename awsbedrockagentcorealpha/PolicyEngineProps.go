package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a PolicyEngine resource.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   policyEngine := agentcore.NewPolicyEngine(this, jsii.String("MyPolicyEngine"), &PolicyEngineProps{
//   	PolicyEngineName: jsii.String("my_policy_engine"),
//   })
//
//   allowAllPolicy := agentcore.NewPolicy(this, jsii.String("AllowAllPolicy"), &PolicyProps{
//   	PolicyEngine: policyEngine,
//   	PolicyName: jsii.String("allow_all"),
//   	Statement: agentcore.PolicyStatement_Permit().ForAllPrincipals().OnAllActions().OnResource(jsii.String("AgentCore::Gateway"), gateway.GatewayArn),
//   	Description: jsii.String("Allow all actions on specific gateway (development only)"),
//   	ValidationMode: agentcore.PolicyValidationMode_IGNORE_ALL_FINDINGS(),
//   })
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

