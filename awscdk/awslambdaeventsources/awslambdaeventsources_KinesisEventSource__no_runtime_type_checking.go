//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisEventSource) validateBindParameters(target awslambda.IFunction) error {
	return nil
}

func (k *jsiiProxy_KinesisEventSource) validateEnrichMappingOptionsParameters(options *awslambda.EventSourceMappingOptions) error {
	return nil
}

func validateNewKinesisEventSourceParameters(stream awskinesis.IStream, props *KinesisEventSourceProps) error {
	return nil
}

