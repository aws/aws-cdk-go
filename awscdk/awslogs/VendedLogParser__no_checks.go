//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_VendedLogParser) validateSetLogTypeParameters(val VendedLogType) error {
	return nil
}

func validateNewVendedLogParserParameters(props *VendedLogParserProps) error {
	return nil
}

