package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.sagemaker@SageMakerAlgorithmStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerAlgorithmStateChange := awscdkmixinspreview.Events.NewSageMakerAlgorithmStateChange()
//
// Experimental.
type SageMakerAlgorithmStateChange interface {
}

// The jsii proxy struct for SageMakerAlgorithmStateChange
type jsiiProxy_SageMakerAlgorithmStateChange struct {
	_ byte // padding
}

// Experimental.
func NewSageMakerAlgorithmStateChange() SageMakerAlgorithmStateChange {
	_init_.Initialize()

	j := jsiiProxy_SageMakerAlgorithmStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerAlgorithmStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerAlgorithmStateChange_Override(s SageMakerAlgorithmStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerAlgorithmStateChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SageMaker Algorithm State Change.
// Experimental.
func SageMakerAlgorithmStateChange_SageMakerAlgorithmStateChangePattern(options *SageMakerAlgorithmStateChange_SageMakerAlgorithmStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSageMakerAlgorithmStateChange_SageMakerAlgorithmStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerAlgorithmStateChange",
		"sageMakerAlgorithmStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

