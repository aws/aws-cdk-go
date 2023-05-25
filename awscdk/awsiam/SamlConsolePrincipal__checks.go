//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (s *jsiiProxy_SamlConsolePrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	if document == nil {
		return fmt.Errorf("parameter document is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SamlConsolePrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SamlConsolePrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	if _statement == nil {
		return fmt.Errorf("parameter _statement is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SamlConsolePrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validateNewSamlConsolePrincipalParameters(samlProvider ISamlProvider) error {
	if samlProvider == nil {
		return fmt.Errorf("parameter samlProvider is required, but nil was provided")
	}

	return nil
}

