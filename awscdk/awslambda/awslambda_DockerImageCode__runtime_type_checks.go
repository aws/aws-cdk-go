//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsecr"
)

func validateDockerImageCode_FromEcrParameters(repository awsecr.IRepository, props *EcrImageCodeProps) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateDockerImageCode_FromImageAssetParameters(directory *string, props *AssetImageCodeProps) error {
	if directory == nil {
		return fmt.Errorf("parameter directory is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

