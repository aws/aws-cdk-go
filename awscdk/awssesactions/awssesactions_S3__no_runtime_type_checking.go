//go:build no_runtime_type_checking

package awssesactions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3) validateBindParameters(rule awsses.IReceiptRule) error {
	return nil
}

func validateNewS3Parameters(props *S3Props) error {
	return nil
}

