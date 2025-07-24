//go:build !no_runtime_type_checking

package awscloudwatch

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_SearchExpression) validateWithParameters(props *SearchExpressionOptions) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewSearchExpressionParameters(props *SearchExpressionProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

