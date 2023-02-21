//go:build !no_runtime_type_checking

package awscodecommit

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_Code) validateBindParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateCode_FromAssetParameters(asset awss3assets.Asset) error {
	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}

	return nil
}

func validateCode_FromDirectoryParameters(directoryPath *string) error {
	if directoryPath == nil {
		return fmt.Errorf("parameter directoryPath is required, but nil was provided")
	}

	return nil
}

func validateCode_FromZipFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

