//go:build no_runtime_type_checking

package awssnssubscriptions

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EmailSubscription) validateBindParameters(_topic awssns.ITopic) error {
	return nil
}

func validateNewEmailSubscriptionParameters(emailAddress *string, props *EmailSubscriptionProps) error {
	return nil
}

