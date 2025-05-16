//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (o *jsiiProxy_OpenIdConnectPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	if document == nil {
		return fmt.Errorf("parameter document is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OpenIdConnectPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OpenIdConnectPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	if _statement == nil {
		return fmt.Errorf("parameter _statement is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OpenIdConnectPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validateNewOpenIdConnectPrincipalParameters(openIdConnectProvider IOpenIdConnectProvider) error {
	if openIdConnectProvider == nil {
		return fmt.Errorf("parameter openIdConnectProvider is required, but nil was provided")
	}

	return nil
}

