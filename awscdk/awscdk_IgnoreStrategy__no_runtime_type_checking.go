//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IgnoreStrategy) validateAddParameters(pattern *string) error {
	return nil
}

func (i *jsiiProxy_IgnoreStrategy) validateIgnoresParameters(absoluteFilePath *string) error {
	return nil
}

func validateIgnoreStrategy_DockerParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

func validateIgnoreStrategy_FromCopyOptionsParameters(options *CopyOptions, absoluteRootPath *string) error {
	return nil
}

func validateIgnoreStrategy_GitParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

func validateIgnoreStrategy_GlobParameters(absoluteRootPath *string, patterns *[]*string) error {
	return nil
}

