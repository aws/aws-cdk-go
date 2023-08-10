//go:build !no_runtime_type_checking

package awscdkappconfigalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

func validateConfigurationSource_FromBucketParameters(bucket awss3.IBucket, objectKey *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if objectKey == nil {
		return fmt.Errorf("parameter objectKey is required, but nil was provided")
	}

	return nil
}

func validateConfigurationSource_FromCfnDocumentParameters(document awsssm.CfnDocument) error {
	if document == nil {
		return fmt.Errorf("parameter document is required, but nil was provided")
	}

	return nil
}

func validateConfigurationSource_FromParameterParameters(parameter awsssm.IParameter) error {
	if parameter == nil {
		return fmt.Errorf("parameter parameter is required, but nil was provided")
	}

	return nil
}

func validateConfigurationSource_FromPipelineParameters(pipeline awscodepipeline.IPipeline) error {
	if pipeline == nil {
		return fmt.Errorf("parameter pipeline is required, but nil was provided")
	}

	return nil
}

func validateConfigurationSource_FromSecretParameters(secret awssecretsmanager.ISecret) error {
	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	return nil
}

