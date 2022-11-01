//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsservicecatalog

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CloudFormationTemplate) validateBindParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateCloudFormationTemplate_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateCloudFormationTemplate_FromProductStackParameters(productStack ProductStack) error {
	if productStack == nil {
		return fmt.Errorf("parameter productStack is required, but nil was provided")
	}

	return nil
}

func validateCloudFormationTemplate_FromUrlParameters(url *string) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	return nil
}

