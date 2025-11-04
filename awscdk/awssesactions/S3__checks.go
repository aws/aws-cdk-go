//go:build !no_runtime_type_checking

package awssesactions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses"
)

func (s *jsiiProxy_S3) validateBindParameters(receiptRule awsses.IReceiptRule) error {
	if receiptRule == nil {
		return fmt.Errorf("parameter receiptRule is required, but nil was provided")
	}

	return nil
}

func validateNewS3Parameters(props *S3Props) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

