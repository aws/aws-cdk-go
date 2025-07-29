//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_DataConverterProcessor) validateSetTypeParameters(val DataConverterType) error {
	return nil
}

func validateNewDataConverterProcessorParameters(props *DataConverterProps) error {
	return nil
}

