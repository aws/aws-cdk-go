//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func validateDockerImageCode_FromEcrParameters(repository awsecr.IRepository, props *EcrImageCodeProps) error {
	return nil
}

func validateDockerImageCode_FromImageAssetParameters(directory *string, props *AssetImageCodeProps) error {
	return nil
}

