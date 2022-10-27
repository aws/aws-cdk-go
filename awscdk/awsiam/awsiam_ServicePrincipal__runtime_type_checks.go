//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_ServicePrincipal) validateAddToAssumeRolePolicyParameters(document PolicyDocument) error {
	if document == nil {
		return fmt.Errorf("parameter document is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ServicePrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ServicePrincipal) validateAddToPrincipalPolicyParameters(_statement PolicyStatement) error {
	if _statement == nil {
		return fmt.Errorf("parameter _statement is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ServicePrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validateServicePrincipal_ServicePrincipalNameParameters(service *string) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	return nil
}

func validateNewServicePrincipalParameters(service *string, opts *ServicePrincipalOpts) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

