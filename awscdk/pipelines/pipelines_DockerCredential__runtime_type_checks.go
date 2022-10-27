//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

func (d *jsiiProxy_DockerCredential) validateGrantReadParameters(grantee awsiam.IGrantable, usage DockerCredentialUsage) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	if usage == "" {
		return fmt.Errorf("parameter usage is required, but nil was provided")
	}

	return nil
}

func validateDockerCredential_CustomRegistryParameters(registryDomain *string, secret awssecretsmanager.ISecret, opts *ExternalDockerCredentialOptions) error {
	if registryDomain == nil {
		return fmt.Errorf("parameter registryDomain is required, but nil was provided")
	}

	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func validateDockerCredential_DockerHubParameters(secret awssecretsmanager.ISecret, opts *ExternalDockerCredentialOptions) error {
	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func validateDockerCredential_EcrParameters(repositories *[]awsecr.IRepository, opts *EcrDockerCredentialOptions) error {
	if repositories == nil {
		return fmt.Errorf("parameter repositories is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

