//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretValue) validateNewErrorParameters(message *string) error {
	return nil
}

func (s *jsiiProxy_SecretValue) validateResolveParameters(context IResolveContext) error {
	return nil
}

func validateSecretValue_CfnDynamicReferenceParameters(ref CfnDynamicReference) error {
	return nil
}

func validateSecretValue_CfnParameterParameters(param CfnParameter) error {
	return nil
}

func validateSecretValue_IsSecretValueParameters(x interface{}) error {
	return nil
}

func validateSecretValue_PlainTextParameters(secret *string) error {
	return nil
}

func validateSecretValue_ResourceAttributeParameters(attr *string) error {
	return nil
}

func validateSecretValue_SecretsManagerParameters(secretId *string, options *SecretsManagerSecretOptions) error {
	return nil
}

func validateSecretValue_SsmSecureParameters(parameterName *string) error {
	return nil
}

func validateSecretValue_UnsafePlainTextParameters(secret *string) error {
	return nil
}

func validateNewSecretValueParameters(protectedValue interface{}, options *IntrinsicProps) error {
	return nil
}

