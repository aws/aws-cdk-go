//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"
)

func (c *jsiiProxy_ConditionalPolicyStatement) validateContextAttributeParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConditionalPolicyStatement) validatePrincipalAttributeParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConditionalPolicyStatement) validateResourceAttributeParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func validateNewConditionalPolicyStatementParameters(policyStatement PolicyStatement, conditionBuilder ConditionBuilder) error {
	if policyStatement == nil {
		return fmt.Errorf("parameter policyStatement is required, but nil was provided")
	}

	if conditionBuilder == nil {
		return fmt.Errorf("parameter conditionBuilder is required, but nil was provided")
	}

	return nil
}

