//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisConsumerEventSource) validateBindParameters(target awslambda.IFunction) error {
	return nil
}

func (k *jsiiProxy_KinesisConsumerEventSource) validateEnrichMappingOptionsParameters(options *awslambda.EventSourceMappingOptions) error {
	return nil
}

func validateNewKinesisConsumerEventSourceParameters(streamConsumer awskinesis.IStreamConsumer, props *KinesisEventSourceProps) error {
	return nil
}

