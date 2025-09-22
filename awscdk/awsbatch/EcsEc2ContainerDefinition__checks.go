//go:build !no_runtime_type_checking

package awsbatch

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_EcsEc2ContainerDefinition) validateAddUlimitParameters(ulimit *Ulimit) error {
	if ulimit == nil {
		return fmt.Errorf("parameter ulimit is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(ulimit, func() string { return "parameter ulimit" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_EcsEc2ContainerDefinition) validateAddVolumeParameters(volume EcsVolume) error {
	if volume == nil {
		return fmt.Errorf("parameter volume is required, but nil was provided")
	}

	return nil
}

func validateEcsEc2ContainerDefinition_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewEcsEc2ContainerDefinitionParameters(scope constructs.Construct, id *string, props *EcsEc2ContainerDefinitionProps) error {
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

