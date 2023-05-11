//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DockerImage) validateCpParameters(imagePath *string) error {
	return nil
}

func (d *jsiiProxy_DockerImage) validateRunParameters(options *DockerRunOptions) error {
	return nil
}

func validateDockerImage_FromBuildParameters(path *string, options *DockerBuildOptions) error {
	return nil
}

func validateDockerImage_FromRegistryParameters(image *string) error {
	return nil
}

func validateNewDockerImageParameters(image *string) error {
	return nil
}

