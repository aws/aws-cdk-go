//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IMergeStrategy) validateApplyParameters(target *map[string]interface{}, source *map[string]interface{}, allowedKeys *[]*string) error {
	return nil
}

