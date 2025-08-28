//go:build no_runtime_type_checking

package awscdkiotactionsalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateNewCloudWatchLogsActionParameters(logGroup awslogs.ILogGroup, props *CloudWatchLogsActionProps) error {
	return nil
}

