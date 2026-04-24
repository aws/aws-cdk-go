package awsbedrockagentcorealpha


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
//   	Statement: agentcore.PolicyStatement_Permit().ForPrincipal(jsii.String("AgentCore::OAuthUser::your-client-id")).OnActions([]*string{
//   		jsii.String("AgentCore::Action::WeatherTool__get_forecast"),
//   		jsii.String("AgentCore::Action::WeatherTool__get_current"),
//   	}).OnResource(jsii.String("AgentCore::Gateway"), gateway.GatewayArn),
//   	Description: jsii.String("Allow specific weather tool operations"),
//   	ValidationMode: agentcore.PolicyValidationMode_FAIL_ON_ANY_FINDINGS(),
//   })
//
// Experimental.
type AddPolicyOptions struct {
	// Cedar policy statement (35-153,600 characters).
	//
	// You must specify either `definition` or `statement`, but not both.
	// Default: - Must provide either definition or statement.
	//
	// Experimental.
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// Optional description for the policy (max 4,096 characters).
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the policy.
	//
	// Valid characters: a-z, A-Z, 0-9, _ (underscore)
	// Must start with a letter, 1-48 characters.
	// Default: - Auto-generated unique name.
	//
	// Experimental.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Type-safe Cedar policy statement built using PolicyStatement builder.
	//
	// Use this for a type-safe, form-like API to build Cedar policies without
	// writing raw Cedar syntax. The builder validates at synthesis time.
	//
	// You must specify either `definition` or `statement`, but not both.
	// Default: - Must provide either definition or statement.
	//
	// Experimental.
	Statement PolicyStatement `field:"optional" json:"statement" yaml:"statement"`
	// Validation mode for the policy.
	// Default: PolicyValidationMode.FAIL_ON_ANY_FINDINGS
	//
	// Experimental.
	ValidationMode PolicyValidationMode `field:"optional" json:"validationMode" yaml:"validationMode"`
}

