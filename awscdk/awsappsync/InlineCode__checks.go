//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_InlineCode) validateBindParameters(_scope constructs.Construct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateInlineCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInlineCode_FromInlineParameters(code *string) error {
	if code == nil {
		return fmt.Errorf("parameter code is required, but nil was provided")
	}

	return nil
}

func validateNewInlineCodeParameters(code *string) error {
	if code == nil {
		return fmt.Errorf("parameter code is required, but nil was provided")
	}

	return nil
}

