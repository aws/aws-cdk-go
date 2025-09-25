//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func validateSvcbRecordValue_AliasParameters(targetName *string) error {
	return nil
}

func validateSvcbRecordValue_ServiceParameters(props *SvcbRecordServiceModeProps) error {
	return nil
}

