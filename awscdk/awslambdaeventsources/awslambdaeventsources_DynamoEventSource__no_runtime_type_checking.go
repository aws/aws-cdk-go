//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamoEventSource) validateBindParameters(target awslambda.IFunction) error {
	return nil
}

func (d *jsiiProxy_DynamoEventSource) validateEnrichMappingOptionsParameters(options *awslambda.EventSourceMappingOptions) error {
	return nil
}

func validateNewDynamoEventSourceParameters(table awsdynamodb.ITable, props *DynamoEventSourceProps) error {
	return nil
}

