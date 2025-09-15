//go:build no_runtime_type_checking

package awss3deployment

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeployTimeSubstitutedFile) validateAddSourceParameters(source ISource) error {
	return nil
}

func validateDeployTimeSubstitutedFile_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDeployTimeSubstitutedFileParameters(scope constructs.Construct, id *string, props *DeployTimeSubstitutedFileProps) error {
	return nil
}

