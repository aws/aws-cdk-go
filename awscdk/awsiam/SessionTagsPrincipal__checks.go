//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (s *jsiiProxy_SessionTagsPrincipal) validateAddToAssumeRolePolicyParameters(doc PolicyDocument) error {
	if doc == nil {
		return fmt.Errorf("parameter doc is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SessionTagsPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SessionTagsPrincipal) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SessionTagsPrincipal) validateAppendDedupeParameters(append *string) error {
	if append == nil {
		return fmt.Errorf("parameter append is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SessionTagsPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validateNewSessionTagsPrincipalParameters(principal IPrincipal) error {
	if principal == nil {
		return fmt.Errorf("parameter principal is required, but nil was provided")
	}

	return nil
}

