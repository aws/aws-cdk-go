package awsbedrockagentcore


// Options for adding a policy via PolicyEngine.addPolicy().
//
// Example:
//   var policyEngine PolicyEngine
//   var gateway Gateway
//
//
//   // Allow specific tool actions on specific gateway
//   // Action names follow pattern: "ToolName__operation"
//   policyEngine.AddPolicy(jsii.String("SpecificToolPolicy"), &AddPolicyOptions{
//   	Statement: agentcore.NewPolicyStatement(&PolicyStatementProps{
//   		Effect: agentcore.PolicyEffect_PERMIT,
//   		Principal: agentcore.PolicyPrincipal_EntityType(jsii.String("AgentCore::OAuthUser")),
//   		Action: agentcore.PolicyAction_AnyOf([]*string{
//   			jsii.String("AgentCore::Action::WeatherTool__get_forecast"),
//   			jsii.String("AgentCore::Action::WeatherTool__get_current"),
//   		}),
//   		Resource: agentcore.PolicyResource_Instance(jsii.String("AgentCore::Gateway"), gateway.GatewayArn),
//   	}),
//   	Description: jsii.String("Allow specific weather tool operations"),
//   	ValidationMode: agentcore.PolicyValidationMode_FAIL_ON_ANY_FINDINGS(),
//   })
//
type AddPolicyOptions struct {
	// The Cedar policy statement for this policy.
	//
	// Build a type-safe statement with the `PolicyStatement` factories, or use
	// `PolicyStatement.fromCedar('...')` for raw Cedar. Raw Cedar is treated as trusted
	// input: the module does not escape, quote, or validate it.
	Statement PolicyStatement `field:"required" json:"statement" yaml:"statement"`
	// Optional description for the policy (max 4,096 characters).
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the policy.
	//
	// Valid characters: a-z, A-Z, 0-9, _ (underscore)
	// Must start with a letter, 1-48 characters.
	// Default: - Auto-generated unique name.
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Validation mode for the policy.
	// Default: PolicyValidationMode.FAIL_ON_ANY_FINDINGS
	//
	ValidationMode PolicyValidationMode `field:"optional" json:"validationMode" yaml:"validationMode"`
}

