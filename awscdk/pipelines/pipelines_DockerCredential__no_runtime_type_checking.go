//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DockerCredential) validateGrantReadParameters(grantee awsiam.IGrantable, usage DockerCredentialUsage) error {
	return nil
}

func validateDockerCredential_CustomRegistryParameters(registryDomain *string, secret awssecretsmanager.ISecret, opts *ExternalDockerCredentialOptions) error {
	return nil
}

func validateDockerCredential_DockerHubParameters(secret awssecretsmanager.ISecret, opts *ExternalDockerCredentialOptions) error {
	return nil
}

func validateDockerCredential_EcrParameters(repositories *[]awsecr.IRepository, opts *EcrDockerCredentialOptions) error {
	return nil
}

