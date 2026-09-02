//go:build !no_runtime_type_checking

package previewawssecurityhubmixins

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

func (c *jsiiProxy_CfnHubSecurityFindingLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnHubSecurityFindingLogsDestProps) error {
	if destination == nil {
		return fmt.Errorf("parameter destination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnHubSecurityFindingLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnHubSecurityFindingLogsLogGroupProps) error {
	if logGroup == nil {
		return fmt.Errorf("parameter logGroup is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

