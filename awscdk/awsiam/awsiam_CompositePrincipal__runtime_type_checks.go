//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (c *jsiiProxy_CompositePrincipal) validateAddToAssumeRolePolicyParameters(doc PolicyDocument) error {
	if doc == nil {
		return fmt.Errorf("parameter doc is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CompositePrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CompositePrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	if _statement == nil {
		return fmt.Errorf("parameter _statement is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CompositePrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

