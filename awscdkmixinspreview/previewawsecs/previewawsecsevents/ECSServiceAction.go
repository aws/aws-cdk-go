package previewawsecsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ecs@ECSServiceAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCSServiceAction := awscdkmixinspreview.Events.NewECSServiceAction()
//
// Experimental.
type ECSServiceAction interface {
}

// The jsii proxy struct for ECSServiceAction
type jsiiProxy_ECSServiceAction struct {
	_ byte // padding
}

// Experimental.
func NewECSServiceAction() ECSServiceAction {
	_init_.Initialize()

	j := jsiiProxy_ECSServiceAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSServiceAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewECSServiceAction_Override(e ECSServiceAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSServiceAction",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for ECS Service Action.
// Experimental.
func ECSServiceAction_EcsServiceActionPattern(options *ECSServiceAction_ECSServiceActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateECSServiceAction_EcsServiceActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSServiceAction",
		"ecsServiceActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

