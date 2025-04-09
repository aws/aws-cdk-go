//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Column) validateAddWidgetParameters(w IWidget) error {
	return nil
}

func (c *jsiiProxy_Column) validatePositionParameters(x *float64, y *float64) error {
	return nil
}

