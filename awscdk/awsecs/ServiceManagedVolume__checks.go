//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_ServiceManagedVolume) validateMountInParameters(container ContainerDefinition, mountPoint *ContainerMountPoint) error {
	if container == nil {
		return fmt.Errorf("parameter container is required, but nil was provided")
	}

	if mountPoint == nil {
		return fmt.Errorf("parameter mountPoint is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(mountPoint, func() string { return "parameter mountPoint" }); err != nil {
		return err
	}

	return nil
}

func validateServiceManagedVolume_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewServiceManagedVolumeParameters(scope constructs.Construct, id *string, props *ServiceManagedVolumeProps) error {
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

