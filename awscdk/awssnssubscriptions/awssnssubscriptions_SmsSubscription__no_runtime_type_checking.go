//go:build no_runtime_type_checking

package awssnssubscriptions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SmsSubscription) validateBindParameters(_topic awssns.ITopic) error {
	return nil
}

func validateNewSmsSubscriptionParameters(phoneNumber *string, props *SmsSubscriptionProps) error {
	return nil
}

