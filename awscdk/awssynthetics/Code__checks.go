//go:build !no_runtime_type_checking

package awssynthetics

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_Code) validateBindParameters(scope constructs.Construct, handler *string, family RuntimeFamily) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if handler == nil {
		return fmt.Errorf("parameter handler is required, but nil was provided")
	}

	if family == "" {
		return fmt.Errorf("parameter family is required, but nil was provided")
	}

	return nil
}

func validateCode_FromAssetParameters(assetPath *string, options *awss3assets.AssetOptions) error {
	if assetPath == nil {
		return fmt.Errorf("parameter assetPath is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateCode_FromBucketParameters(bucket interfacesawss3.IBucketRef, key *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateCode_FromInlineParameters(code *string) error {
	if code == nil {
		return fmt.Errorf("parameter code is required, but nil was provided")
	}

	return nil
}

