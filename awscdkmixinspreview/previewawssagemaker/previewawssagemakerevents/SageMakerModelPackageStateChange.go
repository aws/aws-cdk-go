package previewawssagemakerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.sagemaker@SageMakerModelPackageStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerModelPackageStateChange := awscdkmixinspreview.Events.NewSageMakerModelPackageStateChange()
//
// Experimental.
type SageMakerModelPackageStateChange interface {
}

// The jsii proxy struct for SageMakerModelPackageStateChange
type jsiiProxy_SageMakerModelPackageStateChange struct {
	_ byte // padding
}

// Experimental.
func NewSageMakerModelPackageStateChange() SageMakerModelPackageStateChange {
	_init_.Initialize()

	j := jsiiProxy_SageMakerModelPackageStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerModelPackageStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerModelPackageStateChange_Override(s SageMakerModelPackageStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerModelPackageStateChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SageMaker Model Package State Change.
// Experimental.
func SageMakerModelPackageStateChange_EventPattern(options *SageMakerModelPackageStateChange_SageMakerModelPackageStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSageMakerModelPackageStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.SageMakerModelPackageStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

