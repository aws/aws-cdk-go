//go:build no_runtime_type_checking

package awscloudwatchactions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsAction) validateBindParameters(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) error {
	return nil
}

func validateNewSnsActionParameters(topic awssns.ITopic) error {
	return nil
}

