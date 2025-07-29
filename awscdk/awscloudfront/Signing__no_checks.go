//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func validateNewSigningParameters(protocol SigningProtocol, behavior SigningBehavior) error {
	return nil
}

