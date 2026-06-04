//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CrossStackReferences) validateConsumeParameters(strength ReferenceStrength) error {
	return nil
}

func (c *jsiiProxy_CrossStackReferences) validateProduceParameters(strength ReferenceStrength) error {
	return nil
}

func validateCrossStackReferences_OfParameters(scope constructs.IConstruct) error {
	return nil
}

