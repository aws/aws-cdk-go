//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MutualTlsCertificate) validateBindParameters(_scope constructs.Construct) error {
	return nil
}

func validateMutualTlsCertificate_AcmParameters(certificate awscertificatemanager.ICertificate) error {
	return nil
}

func validateMutualTlsCertificate_FileParameters(certificateChainPath *string, privateKeyPath *string) error {
	return nil
}

func validateMutualTlsCertificate_SdsParameters(secretName *string) error {
	return nil
}

