//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITaskDefinitionExtension) validateExtendParameters(taskDefinition TaskDefinition) error {
	return nil
}

