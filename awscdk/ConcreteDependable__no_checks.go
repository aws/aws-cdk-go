//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConcreteDependable) validateAddParameters(construct IConstruct) error {
	return nil
}

