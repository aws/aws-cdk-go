//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (g *jsiiProxy_GenericLogDriver) validateBindParameters(_scope constructs.Construct, _containerDefinition ContainerDefinition) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _containerDefinition == nil {
		return fmt.Errorf("parameter _containerDefinition is required, but nil was provided")
	}

	return nil
}

func validateGenericLogDriver_AwsLogsParameters(props *AwsLogDriverProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewGenericLogDriverParameters(props *GenericLogDriverProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

