//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GraphWidget) validateAddLeftMetricParameters(metric IMetric) error {
	return nil
}

func (g *jsiiProxy_GraphWidget) validateAddRightMetricParameters(metric IMetric) error {
	return nil
}

func (g *jsiiProxy_GraphWidget) validatePositionParameters(x *float64, y *float64) error {
	return nil
}

func validateNewGraphWidgetParameters(props *GraphWidgetProps) error {
	return nil
}

