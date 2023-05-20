//go:build !no_runtime_type_checking

package awscdkservicecatalogappregistryalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (a *jsiiProxy_ApplicationAssociator) validateAssociateStageParameters(stage awscdk.Stage) error {
	if stage == nil {
		return fmt.Errorf("parameter stage is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationAssociator) validateIsStageAssociatedParameters(stage awscdk.Stage) error {
	if stage == nil {
		return fmt.Errorf("parameter stage is required, but nil was provided")
	}

	return nil
}

func validateApplicationAssociator_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewApplicationAssociatorParameters(scope awscdk.App, id *string, props *ApplicationAssociatorProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

