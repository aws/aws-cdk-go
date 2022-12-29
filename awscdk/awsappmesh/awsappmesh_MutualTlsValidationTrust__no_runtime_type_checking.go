//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MutualTlsValidationTrust) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateMutualTlsValidationTrust_AcmParameters(certificateAuthorities *[]awsacmpca.ICertificateAuthority) error {
	return nil
}

func validateMutualTlsValidationTrust_FileParameters(certificateChain *string) error {
	return nil
}

func validateMutualTlsValidationTrust_SdsParameters(secretName *string) error {
	return nil
}

