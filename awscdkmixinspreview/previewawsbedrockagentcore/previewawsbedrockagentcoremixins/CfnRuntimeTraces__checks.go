//go:build !no_runtime_type_checking

package previewawsbedrockagentcoremixins

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

func (c *jsiiProxy_CfnRuntimeTraces) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	if destination == nil {
		return fmt.Errorf("parameter destination is required, but nil was provided")
	}

	return nil
}

