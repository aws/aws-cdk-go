//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAlarmAction) validateBindParameters(scope awscdk.Construct, alarm IAlarm) error {
	return nil
}

