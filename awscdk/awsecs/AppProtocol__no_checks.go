//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func validateAppProtocol_SetGrpcParameters(val AppProtocol) error {
	return nil
}

func validateAppProtocol_SetHttpParameters(val AppProtocol) error {
	return nil
}

func validateAppProtocol_SetHttp2Parameters(val AppProtocol) error {
	return nil
}

func validateNewAppProtocolParameters(value *string) error {
	return nil
}

