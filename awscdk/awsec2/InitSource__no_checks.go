//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateInitSource_FromAssetParameters(targetDirectory *string, path *string, options *InitSourceAssetOptions) error {
	return nil
}

func validateInitSource_FromExistingAssetParameters(targetDirectory *string, asset awss3assets.Asset, options *InitSourceOptions) error {
	return nil
}

func validateInitSource_FromGitHubParameters(targetDirectory *string, owner *string, repo *string, options *InitSourceOptions) error {
	return nil
}

func validateInitSource_FromS3ObjectParameters(targetDirectory *string, bucket awss3.IBucket, key *string, options *InitSourceOptions) error {
	return nil
}

func validateInitSource_FromUrlParameters(targetDirectory *string, url *string, options *InitSourceOptions) error {
	return nil
}

func validateNewInitSourceParameters(targetDirectory *string) error {
	return nil
}

