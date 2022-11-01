//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedKafkaEventSource) validateBindParameters(target awslambda.IFunction) error {
	return nil
}

func (m *jsiiProxy_ManagedKafkaEventSource) validateEnrichMappingOptionsParameters(options *awslambda.EventSourceMappingOptions) error {
	return nil
}

func validateNewManagedKafkaEventSourceParameters(props *ManagedKafkaEventSourceProps) error {
	return nil
}

