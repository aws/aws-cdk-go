//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SelfManagedKafkaEventSource) validateBindParameters(target awslambda.IFunction) error {
	return nil
}

func (s *jsiiProxy_SelfManagedKafkaEventSource) validateEnrichMappingOptionsParameters(options *awslambda.EventSourceMappingOptions) error {
	return nil
}

func validateNewSelfManagedKafkaEventSourceParameters(props *SelfManagedKafkaEventSourceProps) error {
	return nil
}

