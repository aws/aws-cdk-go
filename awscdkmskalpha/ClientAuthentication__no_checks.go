//go:build no_runtime_type_checking

package awscdkmskalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateClientAuthentication_SaslParameters(props *SaslAuthProps) error {
	return nil
}

func validateClientAuthentication_SaslTlsParameters(saslTlsProps *SaslTlsAuthProps) error {
	return nil
}

func validateClientAuthentication_TlsParameters(props *TlsAuthProps) error {
	return nil
}

