//go:build no_runtime_type_checking

package awscodepipeline

// Building without runtime type checking enabled, so all the below just return nil

func validateArtifactPath_ArtifactPathParameters(artifactName *string, fileName *string) error {
	return nil
}

func validateNewArtifactPathParameters(artifact Artifact, fileName *string) error {
	return nil
}

