//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Spacer) validatePositionParameters(_x *float64, _y *float64) error {
	return nil
}

func validateNewSpacerParameters(props *SpacerProps) error {
	return nil
}

