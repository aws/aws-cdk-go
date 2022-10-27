//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func validateStateMachineInput_FilePathParameters(inputFile awscodepipeline.ArtifactPath) error {
	return nil
}

func validateStateMachineInput_LiteralParameters(object *map[string]interface{}) error {
	return nil
}

