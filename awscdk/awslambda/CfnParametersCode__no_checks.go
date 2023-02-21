//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnParametersCode) validateAssignParameters(location *awss3.Location) error {
	return nil
}

func (c *jsiiProxy_CfnParametersCode) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (c *jsiiProxy_CfnParametersCode) validateBindToResourceParameters(_resource awscdk.CfnResource, _options *ResourceBindOptions) error {
	return nil
}

func validateCfnParametersCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateCfnParametersCode_FromAssetImageParameters(directory *string, props *AssetImageCodeProps) error {
	return nil
}

func validateCfnParametersCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateCfnParametersCode_FromCfnParametersParameters(props *CfnParametersCodeProps) error {
	return nil
}

func validateCfnParametersCode_FromDockerBuildParameters(path *string, options *DockerBuildAssetOptions) error {
	return nil
}

func validateCfnParametersCode_FromEcrImageParameters(repository awsecr.IRepository, props *EcrImageCodeProps) error {
	return nil
}

func validateCfnParametersCode_FromInlineParameters(code *string) error {
	return nil
}

func validateNewCfnParametersCodeParameters(props *CfnParametersCodeProps) error {
	return nil
}

