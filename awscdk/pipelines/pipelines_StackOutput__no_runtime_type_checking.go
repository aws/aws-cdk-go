//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func validateNewStackOutputParameters(artifactFile awscodepipeline.ArtifactPath, outputName *string) error {
	return nil
}

