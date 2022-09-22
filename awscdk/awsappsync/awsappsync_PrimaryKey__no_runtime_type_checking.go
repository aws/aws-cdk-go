//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func validatePrimaryKey_PartitionParameters(key *string) error {
	return nil
}

func validateNewPrimaryKeyParameters(pkey Assign) error {
	return nil
}

