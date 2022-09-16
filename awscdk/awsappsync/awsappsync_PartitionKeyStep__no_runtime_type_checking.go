//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PartitionKeyStep) validateIsParameters(val *string) error {
	return nil
}

func validateNewPartitionKeyStepParameters(key *string) error {
	return nil
}

