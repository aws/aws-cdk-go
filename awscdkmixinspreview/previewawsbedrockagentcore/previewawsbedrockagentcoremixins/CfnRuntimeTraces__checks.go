//go:build !no_runtime_type_checking

package previewawsbedrockagentcoremixins

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

func (c *jsiiProxy_CfnRuntimeTraces) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnRuntimeTracesDestProps) error {
	if destination == nil {
		return fmt.Errorf("parameter destination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnRuntimeTraces) validateToXRayParameters(props *CfnRuntimeTracesXRayProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

