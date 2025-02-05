//go:build !no_runtime_type_checking

package awscdkgluealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_S3Code) validateBindParameters(_scope constructs.Construct, grantable awsiam.IGrantable) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if grantable == nil {
		return fmt.Errorf("parameter grantable is required, but nil was provided")
	}

	return nil
}

func validateS3Code_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
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

func validateNewS3CodeParameters(bucket awss3.IBucket, key *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

