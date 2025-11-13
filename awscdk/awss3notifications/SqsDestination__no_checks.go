//go:build no_runtime_type_checking

package awss3notifications

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqsDestination) validateBindParameters(scope constructs.Construct, bucket interfacesawss3.IBucketRef) error {
	return nil
}

func validateNewSqsDestinationParameters(queue awssqs.IQueue) error {
	return nil
}

