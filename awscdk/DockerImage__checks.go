//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (d *jsiiProxy_DockerImage) validateCpParameters(imagePath *string) error {
	if imagePath == nil {
		return fmt.Errorf("parameter imagePath is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DockerImage) validateRunParameters(options *DockerRunOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDockerImage_FromBuildParameters(path *string, options *DockerBuildOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDockerImage_FromRegistryParameters(image *string) error {
	if image == nil {
		return fmt.Errorf("parameter image is required, but nil was provided")
	}

	return nil
}

func validateNewDockerImageParameters(image *string) error {
	if image == nil {
		return fmt.Errorf("parameter image is required, but nil was provided")
	}

	return nil
}

