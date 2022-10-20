//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::Glue
package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Code) validateBindParameters(_scope constructs.Construct, grantable awsiam.IGrantable) error {
	return nil
}

func validateS3Code_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateS3Code_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateNewS3CodeParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

