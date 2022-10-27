//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

func validateInitSource_FromAssetParameters(targetDirectory *string, path *string, options *InitSourceAssetOptions) error {
	if targetDirectory == nil {
		return fmt.Errorf("parameter targetDirectory is required, but nil was provided")
	}

	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitSource_FromExistingAssetParameters(targetDirectory *string, asset awss3assets.Asset, options *InitSourceOptions) error {
	if targetDirectory == nil {
		return fmt.Errorf("parameter targetDirectory is required, but nil was provided")
	}

	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitSource_FromGitHubParameters(targetDirectory *string, owner *string, repo *string, options *InitSourceOptions) error {
	if targetDirectory == nil {
		return fmt.Errorf("parameter targetDirectory is required, but nil was provided")
	}

	if owner == nil {
		return fmt.Errorf("parameter owner is required, but nil was provided")
	}

	if repo == nil {
		return fmt.Errorf("parameter repo is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitSource_FromS3ObjectParameters(targetDirectory *string, bucket awss3.IBucket, key *string, options *InitSourceOptions) error {
	if targetDirectory == nil {
		return fmt.Errorf("parameter targetDirectory is required, but nil was provided")
	}

	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitSource_FromUrlParameters(targetDirectory *string, url *string, options *InitSourceOptions) error {
	if targetDirectory == nil {
		return fmt.Errorf("parameter targetDirectory is required, but nil was provided")
	}

	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewInitSourceParameters(targetDirectory *string) error {
	if targetDirectory == nil {
		return fmt.Errorf("parameter targetDirectory is required, but nil was provided")
	}

	return nil
}

