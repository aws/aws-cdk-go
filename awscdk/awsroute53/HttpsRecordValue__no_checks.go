//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func validateHttpsRecordValue_AliasParameters(targetName *string) error {
	return nil
}

func validateHttpsRecordValue_ServiceParameters(props *HttpsRecordServiceModeProps) error {
	return nil
}

