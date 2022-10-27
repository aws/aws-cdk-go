//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

func validateInitFile_FromAssetParameters(targetFileName *string, path *string, options *InitFileAssetOptions) error {
	if targetFileName == nil {
		return fmt.Errorf("parameter targetFileName is required, but nil was provided")
	}

	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitFile_FromExistingAssetParameters(targetFileName *string, asset awss3assets.Asset, options *InitFileOptions) error {
	if targetFileName == nil {
		return fmt.Errorf("parameter targetFileName is required, but nil was provided")
	}

	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitFile_FromFileInlineParameters(targetFileName *string, sourceFileName *string, options *InitFileOptions) error {
	if targetFileName == nil {
		return fmt.Errorf("parameter targetFileName is required, but nil was provided")
	}

	if sourceFileName == nil {
		return fmt.Errorf("parameter sourceFileName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitFile_FromObjectParameters(fileName *string, obj *map[string]interface{}, options *InitFileOptions) error {
	if fileName == nil {
		return fmt.Errorf("parameter fileName is required, but nil was provided")
	}

	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitFile_FromS3ObjectParameters(fileName *string, bucket awss3.IBucket, key *string, options *InitFileOptions) error {
	if fileName == nil {
		return fmt.Errorf("parameter fileName is required, but nil was provided")
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

func validateInitFile_FromStringParameters(fileName *string, content *string, options *InitFileOptions) error {
	if fileName == nil {
		return fmt.Errorf("parameter fileName is required, but nil was provided")
	}

	if content == nil {
		return fmt.Errorf("parameter content is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitFile_FromUrlParameters(fileName *string, url *string, options *InitFileOptions) error {
	if fileName == nil {
		return fmt.Errorf("parameter fileName is required, but nil was provided")
	}

	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInitFile_SymlinkParameters(fileName *string, target *string, options *InitFileOptions) error {
	if fileName == nil {
		return fmt.Errorf("parameter fileName is required, but nil was provided")
	}

	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewInitFileParameters(fileName *string, options *InitFileOptions) error {
	if fileName == nil {
		return fmt.Errorf("parameter fileName is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

