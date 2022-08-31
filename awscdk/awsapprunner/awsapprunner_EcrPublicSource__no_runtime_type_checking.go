//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapprunner

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcrPublicSource) validateBindParameters(_scope constructs.Construct) error {
	return nil
}

func validateEcrPublicSource_FromAssetParameters(props *AssetProps) error {
	return nil
}

func validateEcrPublicSource_FromEcrParameters(props *EcrProps) error {
	return nil
}

func validateEcrPublicSource_FromEcrPublicParameters(props *EcrPublicProps) error {
	return nil
}

func validateEcrPublicSource_FromGitHubParameters(props *GithubRepositoryProps) error {
	return nil
}

func validateNewEcrPublicSourceParameters(props *EcrPublicProps) error {
	return nil
}

