//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/awsssm"
)

func (s *jsiiProxy_Secret) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateSecret_FromSecretsManagerParameters(secret awssecretsmanager.ISecret) error {
	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	return nil
}

func validateSecret_FromSecretsManagerVersionParameters(secret awssecretsmanager.ISecret, versionInfo *SecretVersionInfo) error {
	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	if versionInfo == nil {
		return fmt.Errorf("parameter versionInfo is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(versionInfo, func() string { return "parameter versionInfo" }); err != nil {
		return err
	}

	return nil
}

func validateSecret_FromSsmParameterParameters(parameter awsssm.IParameter) error {
	if parameter == nil {
		return fmt.Errorf("parameter parameter is required, but nil was provided")
	}

	return nil
}

