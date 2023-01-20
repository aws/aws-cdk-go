//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GenericLinuxImage) validateGetImageParameters(scope constructs.Construct) error {
	return nil
}

func validateNewGenericLinuxImageParameters(amiMap *map[string]*string, props *GenericLinuxImageProps) error {
	return nil
}

