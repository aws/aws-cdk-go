//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GenericSSMParameterImage) validateGetImageParameters(scope constructs.Construct) error {
	return nil
}

func validateNewGenericSSMParameterImageParameters(parameterName *string, os OperatingSystemType) error {
	return nil
}

