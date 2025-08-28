//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ParserProcessor) validateSetTypeParameters(val ParserProcessorType) error {
	return nil
}

func validateNewParserProcessorParameters(props *ParserProcessorProps) error {
	return nil
}

