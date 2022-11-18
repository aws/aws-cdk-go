//go:build !no_runtime_type_checking

// Receipt rule actions for AWS IoT
package awscdkiotactionsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNewIotRepublishMqttActionParameters(topic *string, props *IotRepublishMqttActionProps) error {
	if topic == nil {
		return fmt.Errorf("parameter topic is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

