//go:build no_runtime_type_checking

package awsamplify

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ISourceCodeProvider) validateBindParameters(app App) error {
	return nil
}

