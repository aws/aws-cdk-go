//go:build no_runtime_type_checking

package awsamplify

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitLabSourceCodeProvider) validateBindParameters(_app App) error {
	return nil
}

func validateNewGitLabSourceCodeProviderParameters(props *GitLabSourceCodeProviderProps) error {
	return nil
}

