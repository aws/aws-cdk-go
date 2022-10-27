//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsefs"
)

func validateFileSystem_FromEfsAccessPointParameters(ap awsefs.IAccessPoint, mountPath *string) error {
	if ap == nil {
		return fmt.Errorf("parameter ap is required, but nil was provided")
	}

	if mountPath == nil {
		return fmt.Errorf("parameter mountPath is required, but nil was provided")
	}

	return nil
}

func validateNewFileSystemParameters(config *FileSystemConfig) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

