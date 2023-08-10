//go:build no_runtime_type_checking

package awscdkappconfigalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateConfigurationSource_FromBucketParameters(bucket awss3.IBucket, objectKey *string) error {
	return nil
}

func validateConfigurationSource_FromCfnDocumentParameters(document awsssm.CfnDocument) error {
	return nil
}

func validateConfigurationSource_FromParameterParameters(parameter awsssm.IParameter) error {
	return nil
}

func validateConfigurationSource_FromPipelineParameters(pipeline awscodepipeline.IPipeline) error {
	return nil
}

func validateConfigurationSource_FromSecretParameters(secret awssecretsmanager.ISecret) error {
	return nil
}

