//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateExposedSecret_FromParametersStoreParameters(optionName *string, parameter awsssm.IParameter) error {
	return nil
}

func validateExposedSecret_FromSecretsManagerParameters(optionName *string, secret awssecretsmanager.ISecret) error {
	return nil
}

func (j *jsiiProxy_ExposedSecret) validateSetOptionNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExposedSecret) validateSetSecretArnParameters(val *string) error {
	return nil
}

func validateNewExposedSecretParameters(optionName *string, secretArn *string) error {
	return nil
}

