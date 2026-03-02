//go:build !no_runtime_type_checking

package previewawsstepfunctionsmixins

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

func (c *jsiiProxy_CfnStateMachineStandardLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnStateMachineStandardLogsDestProps) error {
	if destination == nil {
		return fmt.Errorf("parameter destination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnStateMachineStandardLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnStateMachineStandardLogsLogGroupProps) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

