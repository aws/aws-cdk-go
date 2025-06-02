//go:build no_runtime_type_checking

package awscdkpipestargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudWatchLogsTarget) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (c *jsiiProxy_CloudWatchLogsTarget) validateGrantPushParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewCloudWatchLogsTargetParameters(logGroup awslogs.ILogGroup, parameters *CloudWatchLogsTargetParameters) error {
	return nil
}

