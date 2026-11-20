package awsbedrockagentcorealpha


// Attributes for importing an existing Policy.
//
// Example:
//   importedEngine := agentcore.PolicyEngine_FromPolicyEngineAttributes(this, jsii.String("ImportedEngine"), &PolicyEngineAttributes{
//   	PolicyEngineArn: jsii.String("policy-engine/my-engine-id"),
//   })
//
//   importedPolicy := agentcore.Policy_FromPolicyAttributes(this, jsii.String("ImportedPolicy"), &PolicyAttributes{
//   	PolicyArn: jsii.String("my-policy-arn"),
//   	PolicyEngine: importedEngine,
//   })
//
//   // Grant permissions to the imported policy
//   role := iam.NewRole(this, jsii.String("PolicyRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   })
//
//   importedPolicy.GrantRead(role)
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type PolicyAttributes struct {
	// The ARN of the policy.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
	// The policy engine this policy belongs to [disable-awslint:prefer-ref-interface].
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	PolicyEngine IPolicyEngine `field:"required" json:"policyEngine" yaml:"policyEngine"`
}

