//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogQueryWidget) validatePositionParameters(x *float64, y *float64) error {
	return nil
}

func validateNewLogQueryWidgetParameters(props *LogQueryWidgetProps) error {
	return nil
}

