//go:build !no_runtime_type_checking

package awss3deployment

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

func validateSource_AssetParameters(path *string, options *awss3assets.AssetOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateSource_BucketParameters(bucket awss3.IBucket, zipObjectKey *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if zipObjectKey == nil {
		return fmt.Errorf("parameter zipObjectKey is required, but nil was provided")
	}

	return nil
}

func validateSource_DataParameters(objectKey *string, data *string) error {
	if objectKey == nil {
		return fmt.Errorf("parameter objectKey is required, but nil was provided")
	}

	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	return nil
}

func validateSource_JsonDataParameters(objectKey *string, obj interface{}) error {
	if objectKey == nil {
		return fmt.Errorf("parameter objectKey is required, but nil was provided")
	}

	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

