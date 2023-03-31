//go:build no_runtime_type_checking

package awssnssubscriptions

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UrlSubscription) validateBindParameters(_topic awssns.ITopic) error {
	return nil
}

func validateNewUrlSubscriptionParameters(url *string, props *UrlSubscriptionProps) error {
	return nil
}

