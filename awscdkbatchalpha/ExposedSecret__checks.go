//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

func validateExposedSecret_FromParametersStoreParameters(optionName *string, parameter awsssm.IParameter) error {
	if optionName == nil {
		return fmt.Errorf("parameter optionName is required, but nil was provided")
	}

	if parameter == nil {
		return fmt.Errorf("parameter parameter is required, but nil was provided")
	}

	return nil
}

func validateExposedSecret_FromSecretsManagerParameters(optionName *string, secret awssecretsmanager.ISecret) error {
	if optionName == nil {
		return fmt.Errorf("parameter optionName is required, but nil was provided")
	}

	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ExposedSecret) validateSetOptionNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ExposedSecret) validateSetSecretArnParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewExposedSecretParameters(optionName *string, secretArn *string) error {
	if optionName == nil {
		return fmt.Errorf("parameter optionName is required, but nil was provided")
	}

	if secretArn == nil {
		return fmt.Errorf("parameter secretArn is required, but nil was provided")
	}

	return nil
}

