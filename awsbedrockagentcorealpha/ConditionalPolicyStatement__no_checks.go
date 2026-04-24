//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConditionalPolicyStatement) validateContextAttributeParameters(attribute *string) error {
	return nil
}

func (c *jsiiProxy_ConditionalPolicyStatement) validatePrincipalAttributeParameters(attribute *string) error {
	return nil
}

func (c *jsiiProxy_ConditionalPolicyStatement) validateResourceAttributeParameters(attribute *string) error {
	return nil
}

func validateNewConditionalPolicyStatementParameters(policyStatement PolicyStatement, conditionBuilder ConditionBuilder) error {
	return nil
}

