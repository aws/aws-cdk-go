//go:build !no_runtime_type_checking

package awsmediapackagev2alpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/constructs-go/constructs/v10"
)

func (m *jsiiProxy_MediaPackageV2Origin) validateBindParameters(scope constructs.Construct, options *awscloudfront.OriginBindOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewMediaPackageV2OriginParameters(endpoint IOriginEndpoint, props *MediaPackageV2OriginProps) error {
	if endpoint == nil {
		return fmt.Errorf("parameter endpoint is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

