package previewawsecsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ecs@ECSTaskStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCSTaskStateChange := awscdkmixinspreview.Events.NewECSTaskStateChange()
//
// Experimental.
type ECSTaskStateChange interface {
}

// The jsii proxy struct for ECSTaskStateChange
type jsiiProxy_ECSTaskStateChange struct {
	_ byte // padding
}

// Experimental.
func NewECSTaskStateChange() ECSTaskStateChange {
	_init_.Initialize()

	j := jsiiProxy_ECSTaskStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSTaskStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewECSTaskStateChange_Override(e ECSTaskStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSTaskStateChange",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for ECS Task State Change.
// Experimental.
func ECSTaskStateChange_EventPattern(options *ECSTaskStateChange_ECSTaskStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateECSTaskStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecs.events.ECSTaskStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

