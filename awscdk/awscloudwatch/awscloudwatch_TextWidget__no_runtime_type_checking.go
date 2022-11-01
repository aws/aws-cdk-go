//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TextWidget) validatePositionParameters(x *float64, y *float64) error {
	return nil
}

func validateNewTextWidgetParameters(props *TextWidgetProps) error {
	return nil
}

