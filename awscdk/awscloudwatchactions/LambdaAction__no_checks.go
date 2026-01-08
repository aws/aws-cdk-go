//go:build no_runtime_type_checking

package awscloudwatchactions

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaAction) validateBindParameters(scope constructs.Construct, alarm interfacesawscloudwatch.IAlarmRef) error {
	return nil
}

func validateNewLambdaActionParameters(lambdaFunction interface{}, props *LambdaActionProps) error {
	return nil
}

