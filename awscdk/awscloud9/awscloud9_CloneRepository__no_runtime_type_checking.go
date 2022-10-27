//go:build no_runtime_type_checking

package awscloud9

// Building without runtime type checking enabled, so all the below just return nil

func validateCloneRepository_FromCodeCommitParameters(repository awscodecommit.IRepository, path *string) error {
	return nil
}

