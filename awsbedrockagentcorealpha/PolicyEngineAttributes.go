package awsbedrockagentcorealpha


// Attributes for importing an existing PolicyEngine.
//
// Example:
//   importedEngine := agentcore.PolicyEngine_FromPolicyEngineAttributes(this, jsii.String("ImportedEngine"), &PolicyEngineAttributes{
//   	PolicyEngineArn: jsii.String("policy-engine-arn"),
//   	KmsKeyArn: jsii.String("kms-arn"),
//   })
//
//   // Use the imported engine
//   policy := agentcore.NewPolicy(this, jsii.String("PolicyForImportedEngine"), &PolicyProps{
//   	PolicyEngine: importedEngine,
//   	Definition: jsii.String("permit(principal, action, resource);"),
//   })
//
// Experimental.
type PolicyEngineAttributes struct {
	// The ARN of the policy engine.
	// Experimental.
	PolicyEngineArn *string `field:"required" json:"policyEngineArn" yaml:"policyEngineArn"`
	// The KMS key ARN used for encryption (optional).
	// Default: - No KMS key.
	//
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

