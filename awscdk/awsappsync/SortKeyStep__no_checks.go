//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SortKeyStep) validateIsParameters(val *string) error {
	return nil
}

func validateNewSortKeyStepParameters(pkey Assign, skey *string) error {
	return nil
}

