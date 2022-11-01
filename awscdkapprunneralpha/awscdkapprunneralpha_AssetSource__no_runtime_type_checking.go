//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetSource) validateBindParameters(_scope constructs.Construct) error {
	return nil
}

func validateAssetSource_FromAssetParameters(props *AssetProps) error {
	return nil
}

func validateAssetSource_FromEcrParameters(props *EcrProps) error {
	return nil
}

func validateAssetSource_FromEcrPublicParameters(props *EcrPublicProps) error {
	return nil
}

func validateAssetSource_FromGitHubParameters(props *GithubRepositoryProps) error {
	return nil
}

func validateNewAssetSourceParameters(props *AssetProps) error {
	return nil
}

