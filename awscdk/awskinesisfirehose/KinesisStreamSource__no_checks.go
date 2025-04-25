//go:build no_runtime_type_checking

package awskinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisStreamSource) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateNewKinesisStreamSourceParameters(stream awskinesis.IStream) error {
	return nil
}

