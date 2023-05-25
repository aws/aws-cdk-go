//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClientAttributes) validateWithStandardAttributesParameters(attributes *StandardAttributesMask) error {
	return nil
}

