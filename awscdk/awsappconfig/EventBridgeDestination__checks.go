//go:build !no_runtime_type_checking

package awsappconfig

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
)

func validateNewEventBridgeDestinationParameters(bus interfacesawsevents.IEventBusRef) error {
	if bus == nil {
		return fmt.Errorf("parameter bus is required, but nil was provided")
	}

	return nil
}

