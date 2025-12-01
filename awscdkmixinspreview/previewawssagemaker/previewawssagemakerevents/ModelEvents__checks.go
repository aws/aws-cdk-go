//go:build !no_runtime_type_checking

package previewawssagemakerevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker"
)

func (m *jsiiProxy_ModelEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *ModelEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_ModelEvents) validateSageMakerTransformJobStateChangePatternParameters(options *ModelEvents_SageMakerTransformJobStateChange_SageMakerTransformJobStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateModelEvents_FromModelParameters(modelRef interfacesawssagemaker.IModelRef) error {
	if modelRef == nil {
		return fmt.Errorf("parameter modelRef is required, but nil was provided")
	}

	return nil
}

