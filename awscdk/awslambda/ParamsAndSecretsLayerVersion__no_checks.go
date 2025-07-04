//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func validateParamsAndSecretsLayerVersion_FromVersionParameters(version ParamsAndSecretsVersions, options *ParamsAndSecretsOptions) error {
	return nil
}

func validateParamsAndSecretsLayerVersion_FromVersionArnParameters(arn *string, options *ParamsAndSecretsOptions) error {
	return nil
}

