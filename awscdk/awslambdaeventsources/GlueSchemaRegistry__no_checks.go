//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueSchemaRegistry) validateBindParameters(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) error {
	return nil
}

func validateNewGlueSchemaRegistryParameters(props *GlueSchemaRegistryProps) error {
	return nil
}

