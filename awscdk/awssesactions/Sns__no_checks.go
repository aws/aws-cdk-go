//go:build no_runtime_type_checking

package awssesactions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Sns) validateBindParameters(_rule awsses.IReceiptRule) error {
	return nil
}

func validateNewSnsParameters(props *SnsProps) error {
	return nil
}

