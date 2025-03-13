//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerDefinition) validateBindParameters(task ISageMakerTask) error {
	return nil
}

func validateNewContainerDefinitionParameters(options *ContainerDefinitionOptions) error {
	return nil
}

