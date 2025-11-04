//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (p *jsiiProxy_PrincipalWithConditions) validateAddConditionParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrincipalWithConditions) validateAddConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrincipalWithConditions) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	if document == nil {
		return fmt.Errorf("parameter document is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrincipalWithConditions) validateAddToPolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrincipalWithConditions) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrincipalWithConditions) validateAppendDedupeParameters(append *string) error {
	if append == nil {
		return fmt.Errorf("parameter append is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrincipalWithConditions) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validateNewPrincipalWithConditionsParameters(principal IPrincipal, conditions *map[string]interface{}) error {
	if principal == nil {
		return fmt.Errorf("parameter principal is required, but nil was provided")
	}

	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

