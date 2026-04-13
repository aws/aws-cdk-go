package previewawsimagebuilderevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.imagebuilder@EC2ImageBuilderWorkflowStepWaiting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2ImageBuilderWorkflowStepWaiting := awscdkmixinspreview.Events.NewEC2ImageBuilderWorkflowStepWaiting()
//
// Experimental.
type EC2ImageBuilderWorkflowStepWaiting interface {
}

// The jsii proxy struct for EC2ImageBuilderWorkflowStepWaiting
type jsiiProxy_EC2ImageBuilderWorkflowStepWaiting struct {
	_ byte // padding
}

// Experimental.
func NewEC2ImageBuilderWorkflowStepWaiting() EC2ImageBuilderWorkflowStepWaiting {
	_init_.Initialize()

	j := jsiiProxy_EC2ImageBuilderWorkflowStepWaiting{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderWorkflowStepWaiting",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEC2ImageBuilderWorkflowStepWaiting_Override(e EC2ImageBuilderWorkflowStepWaiting) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderWorkflowStepWaiting",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EC2 Image Builder Workflow Step Waiting.
// Experimental.
func EC2ImageBuilderWorkflowStepWaiting_EventPattern(options *EC2ImageBuilderWorkflowStepWaiting_EC2ImageBuilderWorkflowStepWaitingProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEC2ImageBuilderWorkflowStepWaiting_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderWorkflowStepWaiting",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

