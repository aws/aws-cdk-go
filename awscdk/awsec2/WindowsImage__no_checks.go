//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WindowsImage) validateGetImageParameters(scope constructs.Construct) error {
	return nil
}

func validateNewWindowsImageParameters(version WindowsVersion, props *WindowsImageProps) error {
	return nil
}

