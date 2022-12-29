//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InlineCode) validateBindParameters(_scope constructs.Construct) error {
	return nil
}

func (i *jsiiProxy_InlineCode) validateBindToResourceParameters(_resource awscdk.CfnResource, _options *ResourceBindOptions) error {
	return nil
}

func validateInlineCode_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateInlineCode_FromAssetImageParameters(directory *string, props *AssetImageCodeProps) error {
	return nil
}

func validateInlineCode_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateInlineCode_FromCfnParametersParameters(props *CfnParametersCodeProps) error {
	return nil
}

func validateInlineCode_FromDockerBuildParameters(path *string, options *DockerBuildAssetOptions) error {
	return nil
}

func validateInlineCode_FromEcrImageParameters(repository awsecr.IRepository, props *EcrImageCodeProps) error {
	return nil
}

func validateInlineCode_FromInlineParameters(code *string) error {
	return nil
}

func validateNewInlineCodeParameters(code *string) error {
	return nil
}

