//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GenericWindowsImage) validateGetImageParameters(scope constructs.Construct) error {
	return nil
}

func validateNewGenericWindowsImageParameters(amiMap *map[string]*string, props *GenericWindowsImageProps) error {
	return nil
}

