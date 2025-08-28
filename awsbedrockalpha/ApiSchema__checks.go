//go:build !no_runtime_type_checking

package awsbedrockalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

func validateApiSchema_FromInlineParameters(schema *string) error {
	if schema == nil {
		return fmt.Errorf("parameter schema is required, but nil was provided")
	}

	return nil
}

func validateApiSchema_FromLocalAssetParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateApiSchema_FromS3FileParameters(bucket awss3.IBucket, objectKey *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if objectKey == nil {
		return fmt.Errorf("parameter objectKey is required, but nil was provided")
	}

	return nil
}

func validateNewApiSchemaParameters(s3File *awss3.Location) error {
	if err := _jsii_.ValidateStruct(s3File, func() string { return "parameter s3File" }); err != nil {
		return err
	}

	return nil
}

