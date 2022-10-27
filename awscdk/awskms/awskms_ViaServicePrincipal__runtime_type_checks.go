//go:build !no_runtime_type_checking

package awskms

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

func (v *jsiiProxy_ViaServicePrincipal) validateAddToAssumeRolePolicyParameters(document awsiam.PolicyDocument) error {
	if document == nil {
		return fmt.Errorf("parameter document is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_ViaServicePrincipal) validateAddToPolicyParameters(statement awsiam.PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_ViaServicePrincipal) validateAddToPrincipalPolicyParameters(_statement awsiam.PolicyStatement) error {
	if _statement == nil {
		return fmt.Errorf("parameter _statement is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_ViaServicePrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validateNewViaServicePrincipalParameters(serviceName *string) error {
	if serviceName == nil {
		return fmt.Errorf("parameter serviceName is required, but nil was provided")
	}

	return nil
}

