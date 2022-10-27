//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func validateViewerCertificate_FromAcmCertificateParameters(certificate awscertificatemanager.ICertificate, options *ViewerCertificateOptions) error {
	return nil
}

func validateViewerCertificate_FromIamCertificateParameters(iamCertificateId *string, options *ViewerCertificateOptions) error {
	return nil
}

