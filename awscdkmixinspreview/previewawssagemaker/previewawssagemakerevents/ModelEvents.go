package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker"
)

// EventBridge event patterns for Model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var modelRef IModelRef
//
//   modelEvents := awscdkmixinspreview.Events.ModelEvents_FromModel(modelRef)
//
// Experimental.
type ModelEvents interface {
	// EventBridge event pattern for Model AWS API Call via CloudTrail.
	// Experimental.
	AwsAPICallViaCloudTrailPattern(options *ModelEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern
	// EventBridge event pattern for Model SageMaker Transform Job State Change.
	// Experimental.
	SageMakerTransformJobStateChangePattern(options *ModelEvents_SageMakerTransformJobStateChange_SageMakerTransformJobStateChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for ModelEvents
type jsiiProxy_ModelEvents struct {
	_ byte // padding
}

// Create ModelEvents from a Model reference.
// Experimental.
func ModelEvents_FromModel(modelRef interfacesawssagemaker.IModelRef) ModelEvents {
	_init_.Initialize()

	if err := validateModelEvents_FromModelParameters(modelRef); err != nil {
		panic(err)
	}
	var returns ModelEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents",
		"fromModel",
		[]interface{}{modelRef},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelEvents) AwsAPICallViaCloudTrailPattern(options *ModelEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern {
	if err := m.validateAwsAPICallViaCloudTrailPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		m,
		"awsAPICallViaCloudTrailPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelEvents) SageMakerTransformJobStateChangePattern(options *ModelEvents_SageMakerTransformJobStateChange_SageMakerTransformJobStateChangeProps) *awsevents.EventPattern {
	if err := m.validateSageMakerTransformJobStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		m,
		"sageMakerTransformJobStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

