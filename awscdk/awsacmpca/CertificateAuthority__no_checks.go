//go:build no_runtime_type_checking

package awsacmpca

// Building without runtime type checking enabled, so all the below just return nil

func validateCertificateAuthority_FromCertificateAuthorityArnParameters(scope constructs.Construct, id *string, certificateAuthorityArn *string) error {
	return nil
}

