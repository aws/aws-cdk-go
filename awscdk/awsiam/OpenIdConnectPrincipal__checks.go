//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
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

func (o *jsiiProxy_OpenIdConnectPrincipal) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OpenIdConnectPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validateNewOpenIdConnectPrincipalParameters(openIdConnectProvider interfacesawsiam.IOIDCProviderRef) error {
	if openIdConnectProvider == nil {
		return fmt.Errorf("parameter openIdConnectProvider is required, but nil was provided")
	}

	return nil
}

