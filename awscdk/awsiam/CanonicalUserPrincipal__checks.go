//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (c *jsiiProxy_CanonicalUserPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	if document == nil {
		return fmt.Errorf("parameter document is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CanonicalUserPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CanonicalUserPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	if _statement == nil {
		return fmt.Errorf("parameter _statement is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CanonicalUserPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validateNewCanonicalUserPrincipalParameters(canonicalUserId *string) error {
	if canonicalUserId == nil {
		return fmt.Errorf("parameter canonicalUserId is required, but nil was provided")
	}

	return nil
}

