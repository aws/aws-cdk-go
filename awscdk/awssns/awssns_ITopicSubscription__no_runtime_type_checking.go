//go:build no_runtime_type_checking

package awssns

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITopicSubscription) validateBindParameters(topic ITopic) error {
	return nil
}

