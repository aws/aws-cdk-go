//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetImageCode) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (a *jsiiProxy_AssetImageCode) validateBindToResourceParameters(resource awscdk.CfnResource, options *ResourceBindOptions) error {
	return nil
}

func validateAssetImageCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateAssetImageCode_FromAssetImageParameters(directory *string, props *AssetImageCodeProps) error {
	return nil
}

func validateAssetImageCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateAssetImageCode_FromCfnParametersParameters(props *CfnParametersCodeProps) error {
	return nil
}

func validateAssetImageCode_FromDockerBuildParameters(path *string, options *DockerBuildAssetOptions) error {
	return nil
}

func validateAssetImageCode_FromEcrImageParameters(repository awsecr.IRepository, props *EcrImageCodeProps) error {
	return nil
}

func validateAssetImageCode_FromInlineParameters(code *string) error {
	return nil
}

func validateNewAssetImageCodeParameters(directory *string, props *AssetImageCodeProps) error {
	return nil
}

