//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3CodeV2) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func (s *jsiiProxy_S3CodeV2) validateBindToResourceParameters(_resource awscdk.CfnResource, _options *ResourceBindOptions) error {
	return nil
}

func validateS3CodeV2_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateS3CodeV2_FromAssetImageParameters(directory *string, props *AssetImageCodeProps) error {
	return nil
}

func validateS3CodeV2_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateS3CodeV2_FromBucketV2Parameters(bucket awss3.IBucket, key *string, options *BucketOptions) error {
	return nil
}

func validateS3CodeV2_FromCfnParametersParameters(props *CfnParametersCodeProps) error {
	return nil
}

func validateS3CodeV2_FromCustomCommandParameters(output *string, command *[]*string, options *CustomCommandOptions) error {
	return nil
}

func validateS3CodeV2_FromDockerBuildParameters(path *string, options *DockerBuildAssetOptions) error {
	return nil
}

func validateS3CodeV2_FromEcrImageParameters(repository awsecr.IRepository, props *EcrImageCodeProps) error {
	return nil
}

func validateS3CodeV2_FromInlineParameters(code *string) error {
	return nil
}

func validateNewS3CodeV2Parameters(bucket awss3.IBucket, key *string, options *BucketOptions) error {
	return nil
}

