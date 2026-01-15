//go:build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func validateEventDestination_CloudWatchDimensionsParameters(dimensions *[]*CloudWatchDimension) error {
	return nil
}

func validateEventDestination_EventBusParameters(eventBus interfacesawsevents.IEventBusRef) error {
	return nil
}

func validateEventDestination_FirehoseDeliveryStreamParameters(stream *FirehoseDeliveryStreamDestination) error {
	return nil
}

func validateEventDestination_SnsTopicParameters(topic awssns.ITopic) error {
	return nil
}

