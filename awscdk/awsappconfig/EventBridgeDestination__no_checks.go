//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func validateNewEventBridgeDestinationParameters(bus awsevents.IEventBus) error {
	return nil
}

