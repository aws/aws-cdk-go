//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StreamEventSource) validateBindParameters(_target awslambda.IFunction) error {
	return nil
}

func (s *jsiiProxy_StreamEventSource) validateEnrichMappingOptionsParameters(options *awslambda.EventSourceMappingOptions) error {
	return nil
}

func validateNewStreamEventSourceParameters(props *StreamEventSourceProps) error {
	return nil
}

