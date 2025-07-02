package awscdkpipestargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkpipestargetsalpha/v2/internal"
)

// An EventBridge Pipes target that sends messages to an EventBridge event bus.
//
// Example:
//   var sourceQueue queue
//   var targetEventBus eventBus
//
//
//   eventBusTarget := targets.NewEventBridgeTarget(targetEventBus)
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: eventBusTarget,
//   })
//
// Experimental.
type EventBridgeTarget interface {
	awscdkpipesalpha.ITarget
	// The ARN of the target resource.
	// Experimental.
	TargetArn() *string
	// Bind this target to a pipe.
	// Experimental.
	Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig
	// Grant the pipe role to push to the target.
	// Experimental.
	GrantPush(grantee awsiam.IRole)
}

// The jsii proxy struct for EventBridgeTarget
type jsiiProxy_EventBridgeTarget struct {
	internal.Type__awscdkpipesalphaITarget
}

func (j *jsiiProxy_EventBridgeTarget) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewEventBridgeTarget(eventBus awsevents.IEventBus, parameters *EventBridgeTargetParameters) EventBridgeTarget {
	_init_.Initialize()

	if err := validateNewEventBridgeTargetParameters(eventBus, parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventBridgeTarget{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.EventBridgeTarget",
		[]interface{}{eventBus, parameters},
		&j,
	)

	return &j
}

// Experimental.
func NewEventBridgeTarget_Override(e EventBridgeTarget, eventBus awsevents.IEventBus, parameters *EventBridgeTargetParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-targets-alpha.EventBridgeTarget",
		[]interface{}{eventBus, parameters},
		e,
	)
}

func (e *jsiiProxy_EventBridgeTarget) Bind(pipe awscdkpipesalpha.IPipe) *awscdkpipesalpha.TargetConfig {
	if err := e.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *awscdkpipesalpha.TargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventBridgeTarget) GrantPush(grantee awsiam.IRole) {
	if err := e.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"grantPush",
		[]interface{}{grantee},
	)
}

