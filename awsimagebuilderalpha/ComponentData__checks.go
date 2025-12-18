//go:build !no_runtime_type_checking

package awsimagebuilderalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

func validateComponentData_FromAssetParameters(scope constructs.Construct, id *string, path *string, options *awss3assets.AssetOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateComponentData_FromComponentDocumentJsonObjectParameters(data *ComponentDocument) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func validateComponentData_FromInlineParameters(data *string) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	return nil
}

func validateComponentData_FromJsonObjectParameters(data *map[string]interface{}) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	return nil
}

func validateComponentData_FromS3Parameters(bucket awss3.IBucket, key *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

