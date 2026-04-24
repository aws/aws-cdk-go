package awsbedrockagentcorealpha


// Properties for creating a Policy resource.
//
// Example:
//   var policyEngine PolicyEngine
//
//
//   // Option 1: Using definition property
//   advancedPolicy := agentcore.NewPolicy(this, jsii.String("AdvancedPolicy"), &PolicyProps{
//   	PolicyEngine: policyEngine,
//   	Definition: jsii.String("permit(principal, action, resource) when { context.custom > 10 };"),
//   	Description: jsii.String("Advanced policy with custom Cedar logic"),
//   })
//
//   // Option 2: Using fromCedar() with statement property
//   policyEngine.AddPolicy(jsii.String("CustomPolicy"), &AddPolicyOptions{
//   	Statement: agentcore.PolicyStatement_FromCedar(jsii.String("forbid(principal, action, resource) when { resource.confidential == true };")),
//   	Description: jsii.String("Custom policy from Cedar string"),
//   })
//
// Experimental.
type PolicyProps struct {
	// The policy engine this policy belongs to.
	//
	// [disable-awslint:prefer-ref-interface].
	// Experimental.
	PolicyEngine IPolicyEngine `field:"required" json:"policyEngine" yaml:"policyEngine"`
	// Cedar policy statement. The authorization policy written in Cedar policy language.
	//
	// Cedar supports permit and forbid rules with conditions.
	// The statement will be wrapped in a PolicyDefinition structure internally.
	//
	// Pass the raw Cedar statement as a string. For example:
	// - "permit(principal, action, resource);"
	// - "permit(principal in Group::\"Admins\", action == Action::\"InvokeModel\", resource) when { context.environment == \"production\" };"
	//
	// You must specify either `definition` or `statement`, but not both.
	// Default: - Must provide either definition or statement.
	//
	// Experimental.
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// Optional description for the policy.
	//
	// Maximum length of 4096.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the policy.
	//
	// Valid characters: a-z, A-Z, 0-9, _ (underscore)
	// Must start with a letter, 1-48 characters
	// Pattern: ^[A-Za-z][A-Za-z0-9_]*$.
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
	//
	// Controls how Cedar analyzer validation findings are handled.
	// Default: PolicyValidationMode.FAIL_ON_ANY_FINDINGS
	//
	// Experimental.
	ValidationMode PolicyValidationMode `field:"optional" json:"validationMode" yaml:"validationMode"`
}

