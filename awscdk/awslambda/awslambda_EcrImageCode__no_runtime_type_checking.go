//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcrImageCode) validateBindParameters(_arg constructs.Construct) error {
	return nil
}

func (e *jsiiProxy_EcrImageCode) validateBindToResourceParameters(_resource awscdk.CfnResource, _options *ResourceBindOptions) error {
	return nil
}

func validateEcrImageCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateEcrImageCode_FromAssetImageParameters(directory *string, props *AssetImageCodeProps) error {
	return nil
}

func validateEcrImageCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateEcrImageCode_FromCfnParametersParameters(props *CfnParametersCodeProps) error {
	return nil
}

func validateEcrImageCode_FromDockerBuildParameters(path *string, options *DockerBuildAssetOptions) error {
	return nil
}

func validateEcrImageCode_FromEcrImageParameters(repository awsecr.IRepository, props *EcrImageCodeProps) error {
	return nil
}

func validateEcrImageCode_FromInlineParameters(code *string) error {
	return nil
}

func validateNewEcrImageCodeParameters(repository awsecr.IRepository, props *EcrImageCodeProps) error {
	return nil
}

