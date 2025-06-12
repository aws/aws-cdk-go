//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func validateDomainlessCredentialSpec_ArnForS3ObjectParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateDomainlessCredentialSpec_ArnForSsmParameterParameters(parameter awsssm.IParameter) error {
	return nil
}

func validateDomainlessCredentialSpec_FromS3BucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateDomainlessCredentialSpec_FromSsmParameterParameters(parameter awsssm.IParameter) error {
	return nil
}

func validateNewDomainlessCredentialSpecParameters(fileLocation *string) error {
	return nil
}

