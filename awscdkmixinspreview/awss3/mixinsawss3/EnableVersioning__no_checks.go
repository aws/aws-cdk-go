//go:build no_runtime_type_checking

package mixinsawss3

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EnableVersioning) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (e *jsiiProxy_EnableVersioning) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

