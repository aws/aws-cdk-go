//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func validateStackInstances_FromArtifactPathParameters(artifactPath awscodepipeline.ArtifactPath, regions *[]*string) error {
	return nil
}

func validateStackInstances_InAccountsParameters(accounts *[]*string, regions *[]*string) error {
	return nil
}

func validateStackInstances_InOrganizationalUnitsParameters(ous *[]*string, regions *[]*string) error {
	return nil
}

