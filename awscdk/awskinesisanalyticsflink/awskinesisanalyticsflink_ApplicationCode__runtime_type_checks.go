//go:build !no_runtime_type_checking

package awskinesisanalyticsflink

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

func (a *jsiiProxy_ApplicationCode) validateBindParameters(scope awscdk.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateApplicationCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateApplicationCode_FromBucketParameters(bucket awss3.IBucket, fileKey *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if fileKey == nil {
		return fmt.Errorf("parameter fileKey is required, but nil was provided")
	}

	return nil
}

