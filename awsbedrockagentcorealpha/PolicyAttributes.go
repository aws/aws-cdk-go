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
// Experimental.
type PolicyAttributes struct {
	// The ARN of the policy.
	// Experimental.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
	// The policy engine this policy belongs to [disable-awslint:prefer-ref-interface].
	// Experimental.
	PolicyEngine IPolicyEngine `field:"required" json:"policyEngine" yaml:"policyEngine"`
}

