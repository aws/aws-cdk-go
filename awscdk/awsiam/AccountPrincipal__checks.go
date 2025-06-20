//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (a *jsiiProxy_AccountPrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	if document == nil {
		return fmt.Errorf("parameter document is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountPrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	if _statement == nil {
		return fmt.Errorf("parameter _statement is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountPrincipal) validateInOrganizationParameters(organizationId *string) error {
	if organizationId == nil {
		return fmt.Errorf("parameter organizationId is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validateNewAccountPrincipalParameters(accountId interface{}) error {
	if accountId == nil {
		return fmt.Errorf("parameter accountId is required, but nil was provided")
	}

	return nil
}

