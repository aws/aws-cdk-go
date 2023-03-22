//go:build !no_runtime_type_checking

package customresources

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

func validateAwsCustomResourcePolicy_FromSdkCallsParameters(options *SdkCallsPolicyOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAwsCustomResourcePolicy_FromStatementsParameters(statements *[]awsiam.PolicyStatement) error {
	if statements == nil {
		return fmt.Errorf("parameter statements is required, but nil was provided")
	}

	return nil
}

