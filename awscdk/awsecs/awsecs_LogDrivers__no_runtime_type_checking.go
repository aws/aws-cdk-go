//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func validateLogDrivers_AwsLogsParameters(props *AwsLogDriverProps) error {
	return nil
}

func validateLogDrivers_FirelensParameters(props *FireLensLogDriverProps) error {
	return nil
}

func validateLogDrivers_FluentdParameters(props *FluentdLogDriverProps) error {
	return nil
}

func validateLogDrivers_GelfParameters(props *GelfLogDriverProps) error {
	return nil
}

func validateLogDrivers_JournaldParameters(props *JournaldLogDriverProps) error {
	return nil
}

func validateLogDrivers_JsonFileParameters(props *JsonFileLogDriverProps) error {
	return nil
}

func validateLogDrivers_SplunkParameters(props *SplunkLogDriverProps) error {
	return nil
}

func validateLogDrivers_SyslogParameters(props *SyslogLogDriverProps) error {
	return nil
}

