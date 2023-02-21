//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TlsValidationTrust) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateTlsValidationTrust_AcmParameters(certificateAuthorities *[]awsacmpca.ICertificateAuthority) error {
	return nil
}

func validateTlsValidationTrust_FileParameters(certificateChain *string) error {
	return nil
}

func validateTlsValidationTrust_SdsParameters(secretName *string) error {
	return nil
}

