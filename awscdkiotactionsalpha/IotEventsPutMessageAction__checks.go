//go:build !no_runtime_type_checking

// Receipt rule actions for AWS IoT
package awscdkiotactionsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2"
)

func validateNewIotEventsPutMessageActionParameters(input awscdkioteventsalpha.IInput, props *IotEventsPutMessageActionProps) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

