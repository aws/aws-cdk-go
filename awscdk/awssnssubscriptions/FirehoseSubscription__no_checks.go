//go:build no_runtime_type_checking

package awssnssubscriptions

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirehoseSubscription) validateBindParameters(topic awssns.ITopic) error {
	return nil
}

func validateNewFirehoseSubscriptionParameters(deliveryStream awskinesisfirehose.IDeliveryStream, props *FirehoseSubscriptionProps) error {
	return nil
}

