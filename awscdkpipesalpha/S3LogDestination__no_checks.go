//go:build no_runtime_type_checking

package awscdkpipesalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3LogDestination) validateBindParameters(_pipe IPipe) error {
	return nil
}

func (s *jsiiProxy_S3LogDestination) validateGrantPushParameters(pipeRole awsiam.IRole) error {
	return nil
}

func validateNewS3LogDestinationParameters(parameters *S3LogDestinationProps) error {
	return nil
}

