//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (u *jsiiProxy_UserData) validateAddExecuteFileCommandParameters(params *ExecuteFileOptions) error {
	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(params, func() string { return "parameter params" }); err != nil {
		return err
	}

	return nil
}

func (u *jsiiProxy_UserData) validateAddS3DownloadCommandParameters(params *S3DownloadOptions) error {
	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(params, func() string { return "parameter params" }); err != nil {
		return err
	}

	return nil
}

func (u *jsiiProxy_UserData) validateAddSignalOnExitCommandParameters(resource awscdk.Resource) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

func validateUserData_CustomParameters(content *string) error {
	if content == nil {
		return fmt.Errorf("parameter content is required, but nil was provided")
	}

	return nil
}

func validateUserData_ForLinuxParameters(options *LinuxUserDataOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateUserData_ForOperatingSystemParameters(os OperatingSystemType) error {
	if os == "" {
		return fmt.Errorf("parameter os is required, but nil was provided")
	}

	return nil
}

func validateUserData_ForWindowsParameters(options *WindowsUserDataOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

