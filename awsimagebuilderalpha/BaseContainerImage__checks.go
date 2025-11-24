//go:build !no_runtime_type_checking

package awsimagebuilderalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
)

func validateBaseContainerImage_FromDockerHubParameters(repository *string, tag *string) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	if tag == nil {
		return fmt.Errorf("parameter tag is required, but nil was provided")
	}

	return nil
}

func validateBaseContainerImage_FromEcrParameters(repository awsecr.IRepository, tag *string) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	if tag == nil {
		return fmt.Errorf("parameter tag is required, but nil was provided")
	}

	return nil
}

func validateBaseContainerImage_FromEcrPublicParameters(registryAlias *string, repositoryName *string, tag *string) error {
	if registryAlias == nil {
		return fmt.Errorf("parameter registryAlias is required, but nil was provided")
	}

	if repositoryName == nil {
		return fmt.Errorf("parameter repositoryName is required, but nil was provided")
	}

	if tag == nil {
		return fmt.Errorf("parameter tag is required, but nil was provided")
	}

	return nil
}

func validateBaseContainerImage_FromStringParameters(baseContainerImageString *string) error {
	if baseContainerImageString == nil {
		return fmt.Errorf("parameter baseContainerImageString is required, but nil was provided")
	}

	return nil
}

func validateNewBaseContainerImageParameters(image *string) error {
	if image == nil {
		return fmt.Errorf("parameter image is required, but nil was provided")
	}

	return nil
}

