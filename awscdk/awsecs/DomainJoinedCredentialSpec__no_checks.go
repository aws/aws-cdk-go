//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func validateDomainJoinedCredentialSpec_ArnForS3ObjectParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateDomainJoinedCredentialSpec_ArnForSsmParameterParameters(parameter awsssm.IParameter) error {
	return nil
}

func validateDomainJoinedCredentialSpec_FromS3BucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateDomainJoinedCredentialSpec_FromSsmParameterParameters(parameter awsssm.IParameter) error {
	return nil
}

func validateNewDomainJoinedCredentialSpecParameters(fileLocation *string) error {
	return nil
}

