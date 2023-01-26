//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitIgnoreStrategy) validateAddParameters(pattern *string) error {
	return nil
}

func (g *jsiiProxy_GitIgnoreStrategy) validateIgnoresParameters(absoluteFilePath *string) error {
	return nil
}

func validateGitIgnoreStrategy_DockerParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

func validateGitIgnoreStrategy_FromCopyOptionsParameters(options *CopyOptions, absoluteRootPath *string) error {
	return nil
}

func validateGitIgnoreStrategy_GitParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

func validateGitIgnoreStrategy_GlobParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

func validateNewGitIgnoreStrategyParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

