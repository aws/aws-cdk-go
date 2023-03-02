//go:build !no_runtime_type_checking

package awscloudwatch

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_Spacer) validatePositionParameters(_x *float64, _y *float64) error {
	if _x == nil {
		return fmt.Errorf("parameter _x is required, but nil was provided")
	}

	if _y == nil {
		return fmt.Errorf("parameter _y is required, but nil was provided")
	}

	return nil
}

func validateNewSpacerParameters(props *SpacerProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

