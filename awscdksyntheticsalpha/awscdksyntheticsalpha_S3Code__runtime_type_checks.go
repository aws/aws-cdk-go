//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::Synthetics
package awscdksyntheticsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_S3Code) validateBindParameters(_scope constructs.Construct, _handler *string, _family RuntimeFamily) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _handler == nil {
		return fmt.Errorf("parameter _handler is required, but nil was provided")
	}

	if _family == "" {
		return fmt.Errorf("parameter _family is required, but nil was provided")
	}

	return nil
}

func validateS3Code_FromAssetParameters(assetPath *string, options *awss3assets.AssetOptions) error {
	if assetPath == nil {
		return fmt.Errorf("parameter assetPath is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateS3Code_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateS3Code_FromInlineParameters(code *string) error {
	if code == nil {
		return fmt.Errorf("parameter code is required, but nil was provided")
	}

	return nil
}

func validateNewS3CodeParameters(bucket awss3.IBucket, key *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

