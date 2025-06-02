//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (w *jsiiProxy_WebIdentityPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	if document == nil {
		return fmt.Errorf("parameter document is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WebIdentityPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WebIdentityPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	if _statement == nil {
		return fmt.Errorf("parameter _statement is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WebIdentityPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validateNewWebIdentityPrincipalParameters(identityProvider *string) error {
	if identityProvider == nil {
		return fmt.Errorf("parameter identityProvider is required, but nil was provided")
	}

	return nil
}

