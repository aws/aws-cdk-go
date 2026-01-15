//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudWatchLogGroup) validateBindParameters(rule interfacesawsevents.IRuleRef) error {
	return nil
}

func validateNewCloudWatchLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *LogGroupProps) error {
	return nil
}

