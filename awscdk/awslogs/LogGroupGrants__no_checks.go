//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogGroupGrants) validateActionsParameters(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) error {
	return nil
}

func (l *jsiiProxy_LogGroupGrants) validateReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (l *jsiiProxy_LogGroupGrants) validateWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateLogGroupGrants_FromLogGroupParameters(resource interfacesawslogs.ILogGroupRef) error {
	return nil
}

