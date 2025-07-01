//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsTopic) validateBindParameters(_rule awsevents.IRule) error {
	return nil
}

func validateNewSnsTopicParameters(topic awssns.ITopic, props *SnsTopicProps) error {
	return nil
}

