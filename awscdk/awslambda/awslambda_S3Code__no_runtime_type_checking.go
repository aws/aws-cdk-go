//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Code) validateBindParameters(_scope constructs.Construct) error {
	return nil
}

func (s *jsiiProxy_S3Code) validateBindToResourceParameters(_resource awscdk.CfnResource, _options *ResourceBindOptions) error {
	return nil
}

func validateS3Code_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateS3Code_FromAssetImageParameters(directory *string, props *AssetImageCodeProps) error {
	return nil
}

func validateS3Code_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateS3Code_FromCfnParametersParameters(props *CfnParametersCodeProps) error {
	return nil
}

func validateS3Code_FromDockerBuildParameters(path *string, options *DockerBuildAssetOptions) error {
	return nil
}

func validateS3Code_FromEcrImageParameters(repository awsecr.IRepository, props *EcrImageCodeProps) error {
	return nil
}

func validateS3Code_FromInlineParameters(code *string) error {
	return nil
}

func validateNewS3CodeParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

