//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Artifacts) validateBindParameters(_scope constructs.Construct, _project IProject) error {
	return nil
}

func validateArtifacts_S3Parameters(props *S3ArtifactsProps) error {
	return nil
}

func validateNewArtifactsParameters(props *ArtifactsProps) error {
	return nil
}

