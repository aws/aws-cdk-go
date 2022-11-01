//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

func (i *jsiiProxy_InlineApiDefinition) validateBindParameters(_scope awscdk.Construct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_InlineApiDefinition) validateBindAfterCreateParameters(_scope awscdk.Construct, _restApi IRestApi) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _restApi == nil {
		return fmt.Errorf("parameter _restApi is required, but nil was provided")
	}

	return nil
}

func validateInlineApiDefinition_FromAssetParameters(file *string, options *awss3assets.AssetOptions) error {
	if file == nil {
		return fmt.Errorf("parameter file is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInlineApiDefinition_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateInlineApiDefinition_FromInlineParameters(definition interface{}) error {
	if definition == nil {
		return fmt.Errorf("parameter definition is required, but nil was provided")
	}

	return nil
}

func validateNewInlineApiDefinitionParameters(definition interface{}) error {
	if definition == nil {
		return fmt.Errorf("parameter definition is required, but nil was provided")
	}

	return nil
}

