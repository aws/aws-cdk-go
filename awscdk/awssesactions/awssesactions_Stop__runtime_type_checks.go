//go:build !no_runtime_type_checking

package awssesactions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsses"
)

func (s *jsiiProxy_Stop) validateBindParameters(_rule awsses.IReceiptRule) error {
	if _rule == nil {
		return fmt.Errorf("parameter _rule is required, but nil was provided")
	}

	return nil
}

func validateNewStopParameters(props *StopProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

