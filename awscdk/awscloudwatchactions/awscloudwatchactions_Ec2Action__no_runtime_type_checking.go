//go:build no_runtime_type_checking

package awscloudwatchactions

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2Action) validateBindParameters(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) error {
	return nil
}

func validateNewEc2ActionParameters(instanceAction Ec2InstanceAction) error {
	return nil
}

