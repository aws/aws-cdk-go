//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func validateCredentialSpec_ArnForS3ObjectParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateCredentialSpec_ArnForSsmParameterParameters(parameter awsssm.IParameter) error {
	return nil
}

func validateNewCredentialSpecParameters(prefixId *string, fileLocation *string) error {
	return nil
}

