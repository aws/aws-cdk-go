//go:build no_runtime_type_checking

package awscodepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Artifact) validateAtPathParameters(fileName *string) error {
	return nil
}

func (a *jsiiProxy_Artifact) validateGetMetadataParameters(key *string) error {
	return nil
}

func (a *jsiiProxy_Artifact) validateGetParamParameters(jsonFile *string, keyName *string) error {
	return nil
}

func (a *jsiiProxy_Artifact) validateSetMetadataParameters(key *string, value interface{}) error {
	return nil
}

func validateArtifact_ArtifactParameters(name *string) error {
	return nil
}

