//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlobIgnoreStrategy) validateAddParameters(pattern *string) error {
	return nil
}

func (g *jsiiProxy_GlobIgnoreStrategy) validateIgnoresParameters(absoluteFilePath *string) error {
	return nil
}

func validateGlobIgnoreStrategy_DockerParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

func validateGlobIgnoreStrategy_FromCopyOptionsParameters(options *CopyOptions, absoluteRootPath *string) error {
	return nil
}

func validateGlobIgnoreStrategy_GitParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

func validateGlobIgnoreStrategy_GlobParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

func validateNewGlobIgnoreStrategyParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

