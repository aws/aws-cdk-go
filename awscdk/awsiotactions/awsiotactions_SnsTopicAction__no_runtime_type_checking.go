//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsiotactions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsTopicAction) validateBindParameters(rule awsiot.ITopicRule) error {
	return nil
}

func validateNewSnsTopicActionParameters(topic awssns.ITopic, props *SnsTopicActionProps) error {
	return nil
}

