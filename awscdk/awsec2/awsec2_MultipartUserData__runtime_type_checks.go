//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (m *jsiiProxy_MultipartUserData) validateAddExecuteFileCommandParameters(params *ExecuteFileOptions) error {
	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(params, func() string { return "parameter params" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MultipartUserData) validateAddPartParameters(part MultipartBody) error {
	if part == nil {
		return fmt.Errorf("parameter part is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MultipartUserData) validateAddS3DownloadCommandParameters(params *S3DownloadOptions) error {
	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(params, func() string { return "parameter params" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MultipartUserData) validateAddSignalOnExitCommandParameters(resource awscdk.Resource) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MultipartUserData) validateAddUserDataPartParameters(userData UserData) error {
	if userData == nil {
		return fmt.Errorf("parameter userData is required, but nil was provided")
	}

	return nil
}

func validateMultipartUserData_CustomParameters(content *string) error {
	if content == nil {
		return fmt.Errorf("parameter content is required, but nil was provided")
	}

	return nil
}

func validateMultipartUserData_ForLinuxParameters(options *LinuxUserDataOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateMultipartUserData_ForOperatingSystemParameters(os OperatingSystemType) error {
	if os == "" {
		return fmt.Errorf("parameter os is required, but nil was provided")
	}

	return nil
}

func validateMultipartUserData_ForWindowsParameters(options *WindowsUserDataOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewMultipartUserDataParameters(opts *MultipartUserDataOptions) error {
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

