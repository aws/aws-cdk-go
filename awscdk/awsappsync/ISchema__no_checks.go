//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ISchema) validateBindParameters(api interfacesawsappsync.IGraphQLApiRef, options *SchemaBindOptions) error {
	return nil
}

