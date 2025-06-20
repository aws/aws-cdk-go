//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func validateAccessLogFormat_CustomParameters(format *string) error {
	return nil
}

func validateAccessLogFormat_JsonWithStandardFieldsParameters(fields *JsonWithStandardFieldProps) error {
	return nil
}

