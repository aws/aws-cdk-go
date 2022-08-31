//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsiotactions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3PutObjectAction) validateBindParameters(rule awsiot.ITopicRule) error {
	return nil
}

func validateNewS3PutObjectActionParameters(bucket awss3.IBucket, props *S3PutObjectActionProps) error {
	return nil
}

