//go:build no_runtime_type_checking

package awsamplify

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeCommitSourceCodeProvider) validateBindParameters(app App) error {
	return nil
}

func validateNewCodeCommitSourceCodeProviderParameters(props *CodeCommitSourceCodeProviderProps) error {
	return nil
}

