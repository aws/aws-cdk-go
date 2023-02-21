//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/constructs-go/constructs/v10"
)

func (t *jsiiProxy_TagParameterContainerImage) validateBindParameters(scope constructs.Construct, containerDefinition ContainerDefinition) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if containerDefinition == nil {
		return fmt.Errorf("parameter containerDefinition is required, but nil was provided")
	}

	return nil
}

func validateTagParameterContainerImage_FromAssetParameters(directory *string, props *AssetImageProps) error {
	if directory == nil {
		return fmt.Errorf("parameter directory is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateTagParameterContainerImage_FromDockerImageAssetParameters(asset awsecrassets.DockerImageAsset) error {
	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}

	return nil
}

func validateTagParameterContainerImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	return nil
}

func validateTagParameterContainerImage_FromRegistryParameters(name *string, props *RepositoryImageProps) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateTagParameterContainerImage_FromTarballParameters(tarballFile *string) error {
	if tarballFile == nil {
		return fmt.Errorf("parameter tarballFile is required, but nil was provided")
	}

	return nil
}

func validateNewTagParameterContainerImageParameters(repository awsecr.IRepository) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	return nil
}

