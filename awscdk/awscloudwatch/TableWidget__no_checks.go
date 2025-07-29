//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TableWidget) validateAddMetricParameters(metric IMetric) error {
	return nil
}

func (t *jsiiProxy_TableWidget) validatePositionParameters(x *float64, y *float64) error {
	return nil
}

func validateNewTableWidgetParameters(props *TableWidgetProps) error {
	return nil
}

