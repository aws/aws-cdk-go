//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (f *jsiiProxy_FederatedPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	if document == nil {
		return fmt.Errorf("parameter document is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FederatedPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FederatedPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	if _statement == nil {
		return fmt.Errorf("parameter _statement is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FederatedPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validateNewFederatedPrincipalParameters(federated *string) error {
	if federated == nil {
		return fmt.Errorf("parameter federated is required, but nil was provided")
	}

	return nil
}

