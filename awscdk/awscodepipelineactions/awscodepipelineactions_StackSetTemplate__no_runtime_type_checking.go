//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func validateStackSetTemplate_FromArtifactPathParameters(artifactPath awscodepipeline.ArtifactPath) error {
	return nil
}

