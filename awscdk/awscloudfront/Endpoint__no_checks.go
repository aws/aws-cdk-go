//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func validateEndpoint_FromKinesisStreamParameters(stream awskinesis.IStream) error {
	return nil
}

