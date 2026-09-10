package awsbedrockagentcore


// Properties for creating a Policy resource.
//
// Example:
//   var policyEngine PolicyEngine
//
//
//   advancedPolicy := agentcore.NewPolicy(this, jsii.String("AdvancedPolicy"), &PolicyProps{
//   	PolicyEngine: policyEngine,
//   	Statement: agentcore.PolicyStatement_FromCedar(jsii.String("permit(principal, action, resource) when { context.custom > 10 };")),
//   	Description: jsii.String("Advanced policy with custom Cedar logic"),
//   })
//
//   policyEngine.AddPolicy(jsii.String("CustomPolicy"), &AddPolicyOptions{
//   	Statement: agentcore.PolicyStatement_*FromCedar(jsii.String("forbid(principal, action, resource) when { resource.confidential == true };")),
//   	Description: jsii.String("Custom policy from Cedar string"),
//   })
//
type PolicyProps struct {
	// The policy engine this policy belongs to.
	//
	// [disable-awslint:prefer-ref-interface].
	PolicyEngine IPolicyEngine `field:"required" json:"policyEngine" yaml:"policyEngine"`
	// The Cedar policy statement for this policy.
	//
	// Build a type-safe statement with the `PolicyStatement` factories, which validate
	// at synthesis time and reject values that cannot be represented safely in Cedar.
	//
	// For raw Cedar (features this API does not model, or migrating an existing policy),
	// use `PolicyStatement.fromCedar('...')`. That string is used exactly as given: the
	// module does not escape, quote, or validate it, so it is treated as trusted input
	// and you own its correctness and safety. Do not assemble it from values that come
	// from outside your application, such as a request body or a database record.
	Statement PolicyStatement `field:"required" json:"statement" yaml:"statement"`
	// Optional description for the policy.
	//
	// Maximum length of 4096.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the policy.
	//
	// Valid characters: a-z, A-Z, 0-9, _ (underscore)
	// Must start with a letter, 1-48 characters
	// Pattern: ^[A-Za-z][A-Za-z0-9_]*$.
	// Default: - Auto-generated unique name.
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Validation mode for the policy.
	//
	// Controls how Cedar analyzer validation findings are handled.
	// Default: PolicyValidationMode.FAIL_ON_ANY_FINDINGS
	//
	ValidationMode PolicyValidationMode `field:"optional" json:"validationMode" yaml:"validationMode"`
}

