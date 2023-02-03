//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetCode) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (a *jsiiProxy_AssetCode) validateBindToResourceParameters(resource awscdk.CfnResource, options *ResourceBindOptions) error {
	return nil
}

func validateAssetCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateAssetCode_FromAssetImageParameters(directory *string, props *AssetImageCodeProps) error {
	return nil
}

func validateAssetCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateAssetCode_FromCfnParametersParameters(props *CfnParametersCodeProps) error {
	return nil
}

func validateAssetCode_FromDockerBuildParameters(path *string, options *DockerBuildAssetOptions) error {
	return nil
}

func validateAssetCode_FromEcrImageParameters(repository awsecr.IRepository, props *EcrImageCodeProps) error {
	return nil
}

func validateAssetCode_FromInlineParameters(code *string) error {
	return nil
}

func validateNewAssetCodeParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

