//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GithubSource) validateBindParameters(_scope constructs.Construct) error {
	return nil
}

func validateGithubSource_FromAssetParameters(props *AssetProps) error {
	return nil
}

func validateGithubSource_FromEcrParameters(props *EcrProps) error {
	return nil
}

func validateGithubSource_FromEcrPublicParameters(props *EcrPublicProps) error {
	return nil
}

func validateGithubSource_FromGitHubParameters(props *GithubRepositoryProps) error {
	return nil
}

func validateNewGithubSourceParameters(props *GithubRepositoryProps) error {
	return nil
}

