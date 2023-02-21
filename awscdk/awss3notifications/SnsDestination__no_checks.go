//go:build no_runtime_type_checking

package awss3notifications

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsDestination) validateBindParameters(_scope constructs.Construct, bucket awss3.IBucket) error {
	return nil
}

func validateNewSnsDestinationParameters(topic awssns.ITopic) error {
	return nil
}

