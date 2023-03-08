//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsecrassets"
	"github.com/aws/constructs-go/constructs/v3"
)

func (d *jsiiProxy_DockerImage) validateBindParameters(task ISageMakerTask) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}

	return nil
}

func validateDockerImage_FromAssetParameters(scope constructs.Construct, id *string, props *awsecrassets.DockerImageAssetProps) error {
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

func validateDockerImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	return nil
}

func validateDockerImage_FromJsonExpressionParameters(expression *string) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	return nil
}

func validateDockerImage_FromRegistryParameters(imageUri *string) error {
	if imageUri == nil {
		return fmt.Errorf("parameter imageUri is required, but nil was provided")
	}

	return nil
}

