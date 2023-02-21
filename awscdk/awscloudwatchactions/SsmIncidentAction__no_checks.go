//go:build no_runtime_type_checking

package awscloudwatchactions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmIncidentAction) validateBindParameters(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) error {
	return nil
}

func validateNewSsmIncidentActionParameters(responsePlanName *string) error {
	return nil
}

