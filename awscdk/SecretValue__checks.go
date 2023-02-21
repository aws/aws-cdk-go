//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_SecretValue) validateNewErrorParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecretValue) validateResolveParameters(context IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func validateSecretValue_CfnDynamicReferenceParameters(ref CfnDynamicReference) error {
	if ref == nil {
		return fmt.Errorf("parameter ref is required, but nil was provided")
	}

	return nil
}

func validateSecretValue_CfnParameterParameters(param CfnParameter) error {
	if param == nil {
		return fmt.Errorf("parameter param is required, but nil was provided")
	}

	return nil
}

func validateSecretValue_IsSecretValueParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateSecretValue_PlainTextParameters(secret *string) error {
	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	return nil
}

func validateSecretValue_ResourceAttributeParameters(attr *string) error {
	if attr == nil {
		return fmt.Errorf("parameter attr is required, but nil was provided")
	}

	return nil
}

func validateSecretValue_SecretsManagerParameters(secretId *string, options *SecretsManagerSecretOptions) error {
	if secretId == nil {
		return fmt.Errorf("parameter secretId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateSecretValue_SsmSecureParameters(parameterName *string) error {
	if parameterName == nil {
		return fmt.Errorf("parameter parameterName is required, but nil was provided")
	}

	return nil
}

func validateSecretValue_UnsafePlainTextParameters(secret *string) error {
	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	return nil
}

func validateNewSecretValueParameters(protectedValue interface{}, options *IntrinsicProps) error {
	if protectedValue == nil {
		return fmt.Errorf("parameter protectedValue is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

