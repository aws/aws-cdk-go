//go:build no_runtime_type_checking

package awssnssubscriptions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqsSubscription) validateBindParameters(topic awssns.ITopic) error {
	return nil
}

func validateNewSqsSubscriptionParameters(queue awssqs.IQueue, props *SqsSubscriptionProps) error {
	return nil
}

