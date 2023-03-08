//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DockerIgnoreStrategy) validateAddParameters(pattern *string) error {
	return nil
}

func (d *jsiiProxy_DockerIgnoreStrategy) validateIgnoresParameters(absoluteFilePath *string) error {
	return nil
}

func validateDockerIgnoreStrategy_DockerParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

func validateDockerIgnoreStrategy_FromCopyOptionsParameters(options *CopyOptions, absoluteRootPath *string) error {
	return nil
}

func validateDockerIgnoreStrategy_GitParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

func validateDockerIgnoreStrategy_GlobParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

func validateNewDockerIgnoreStrategyParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

