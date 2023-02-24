//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BottleRocketImage) validateGetImageParameters(scope constructs.Construct) error {
	return nil
}

func validateNewBottleRocketImageParameters(props *BottleRocketImageProps) error {
	return nil
}

