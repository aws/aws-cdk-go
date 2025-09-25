//go:build no_runtime_type_checking

package awscdkapprunneralpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Source) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateSource_FromAssetParameters(props *AssetProps) error {
	return nil
}

func validateSource_FromEcrParameters(props *EcrProps) error {
	return nil
}

func validateSource_FromEcrPublicParameters(props *EcrPublicProps) error {
	return nil
}

func validateSource_FromGitHubParameters(props *GithubRepositoryProps) error {
	return nil
}

