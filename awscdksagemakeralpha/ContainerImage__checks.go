//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_ContainerImage) validateBindParameters(scope constructs.Construct, model Model) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if model == nil {
		return fmt.Errorf("parameter model is required, but nil was provided")
	}

	return nil
}

func validateContainerImage_FromAssetParameters(directory *string, options *awsecrassets.DockerImageAssetOptions) error {
	if directory == nil {
		return fmt.Errorf("parameter directory is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateContainerImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	return nil
}

