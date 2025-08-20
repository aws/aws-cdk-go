//go:build !no_runtime_type_checking

package awsbedrockalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

func validateS3ApiSchema_FromInlineParameters(schema *string) error {
	if schema == nil {
		return fmt.Errorf("parameter schema is required, but nil was provided")
	}

	return nil
}

func validateS3ApiSchema_FromLocalAssetParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateS3ApiSchema_FromS3FileParameters(bucket awss3.IBucket, objectKey *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if objectKey == nil {
		return fmt.Errorf("parameter objectKey is required, but nil was provided")
	}

	return nil
}

func validateNewS3ApiSchemaParameters(location *awss3.Location) error {
	if location == nil {
		return fmt.Errorf("parameter location is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(location, func() string { return "parameter location" }); err != nil {
		return err
	}

	return nil
}

