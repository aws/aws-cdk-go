package previewawsecsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ecs@ECSContainerInstanceStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCSContainerInstanceStateChange := awscdkmixinspreview.Events.NewECSContainerInstanceStateChange()
//
// Experimental.
type ECSContainerInstanceStateChange interface {
}

// The jsii proxy struct for ECSContainerInstanceStateChange
type jsiiProxy_ECSContainerInstanceStateChange struct {
	_ byte // padding
}

// Experimental.
func NewECSContainerInstanceStateChange() ECSContainerInstanceStateChange {
	_init_.Initialize()

	j := jsiiProxy_ECSContainerInstanceStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSContainerInstanceStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewECSContainerInstanceStateChange_Override(e ECSContainerInstanceStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSContainerInstanceStateChange",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for ECS Container Instance State Change.
// Experimental.
func ECSContainerInstanceStateChange_EventPattern(options *ECSContainerInstanceStateChange_ECSContainerInstanceStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateECSContainerInstanceStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSContainerInstanceStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

