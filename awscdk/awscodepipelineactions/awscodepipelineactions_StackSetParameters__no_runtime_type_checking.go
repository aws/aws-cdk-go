//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func validateStackSetParameters_FromArtifactPathParameters(artifactPath awscodepipeline.ArtifactPath) error {
	return nil
}

func validateStackSetParameters_FromLiteralParameters(parameters *map[string]*string) error {
	return nil
}

