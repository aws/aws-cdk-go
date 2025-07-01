//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfluentSchemaRegistry) validateBindParameters(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) error {
	return nil
}

func validateNewConfluentSchemaRegistryParameters(props *ConfluentSchemaRegistryProps) error {
	return nil
}

