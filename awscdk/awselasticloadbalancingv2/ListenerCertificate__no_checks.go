//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func validateListenerCertificate_FromArnParameters(certificateArn *string) error {
	return nil
}

func validateListenerCertificate_FromCertificateManagerParameters(acmCertificate awscertificatemanager.ICertificate) error {
	return nil
}

func validateNewListenerCertificateParameters(certificateArn *string) error {
	return nil
}

