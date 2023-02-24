//go:build no_runtime_type_checking

package awssesactions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Stop) validateBindParameters(_rule awsses.IReceiptRule) error {
	return nil
}

func validateNewStopParameters(props *StopProps) error {
	return nil
}

