//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KafkaDlq) validateBindParameters(target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) error {
	return nil
}

func validateNewKafkaDlqParameters(topicName *string) error {
	return nil
}

