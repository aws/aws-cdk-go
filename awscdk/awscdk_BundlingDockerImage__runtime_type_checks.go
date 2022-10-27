//go:build !no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (b *jsiiProxy_BundlingDockerImage) validateCpParameters(imagePath *string) error {
	if imagePath == nil {
		return fmt.Errorf("parameter imagePath is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BundlingDockerImage) validateRunParameters(options *DockerRunOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateBundlingDockerImage_FromAssetParameters(path *string, options *DockerBuildOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateBundlingDockerImage_FromRegistryParameters(image *string) error {
	if image == nil {
		return fmt.Errorf("parameter image is required, but nil was provided")
	}

	return nil
}

func validateNewBundlingDockerImageParameters(image *string) error {
	if image == nil {
		return fmt.Errorf("parameter image is required, but nil was provided")
	}

	return nil
}

