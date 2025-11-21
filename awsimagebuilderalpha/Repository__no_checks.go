//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateRepository_FromEcrParameters(repository awsecr.IRepository) error {
	return nil
}

