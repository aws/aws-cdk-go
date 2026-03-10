package previewawsimagebuilderevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderCVEDetected",
		reflect.TypeOf((*EC2ImageBuilderCVEDetected)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EC2ImageBuilderCVEDetected{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderCVEDetected.EC2ImageBuilderCVEDetectedProps",
		reflect.TypeOf((*EC2ImageBuilderCVEDetected_EC2ImageBuilderCVEDetectedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderCVEDetected.FindingSeverityCounts",
		reflect.TypeOf((*EC2ImageBuilderCVEDetected_FindingSeverityCounts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderImageStateChange",
		reflect.TypeOf((*EC2ImageBuilderImageStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EC2ImageBuilderImageStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderImageStateChange.EC2ImageBuilderImageStateChangeProps",
		reflect.TypeOf((*EC2ImageBuilderImageStateChange_EC2ImageBuilderImageStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderImageStateChange.ImageState",
		reflect.TypeOf((*EC2ImageBuilderImageStateChange_ImageState)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderWorkflowStepWaiting",
		reflect.TypeOf((*EC2ImageBuilderWorkflowStepWaiting)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EC2ImageBuilderWorkflowStepWaiting{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderWorkflowStepWaiting.EC2ImageBuilderWorkflowStepWaitingProps",
		reflect.TypeOf((*EC2ImageBuilderWorkflowStepWaiting_EC2ImageBuilderWorkflowStepWaitingProps)(nil)).Elem(),
	)
}
