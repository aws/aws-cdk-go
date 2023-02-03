//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LinuxParameters) validateAddDevicesParameters(device *[]*Device) error {
	for idx_263a4d, v := range *device {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter device[%#v]", idx_263a4d) }); err != nil {
			return err
		}
	}

	return nil
}

func (l *jsiiProxy_LinuxParameters) validateAddTmpfsParameters(tmpfs *[]*Tmpfs) error {
	for idx_02dae1, v := range *tmpfs {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter tmpfs[%#v]", idx_02dae1) }); err != nil {
			return err
		}
	}

	return nil
}

func validateLinuxParameters_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewLinuxParametersParameters(scope constructs.Construct, id *string, props *LinuxParametersProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

