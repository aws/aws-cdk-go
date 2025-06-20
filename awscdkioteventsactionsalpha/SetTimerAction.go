package awscdkioteventsactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2"
)

// The action to create a timer with duration in seconds.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
//   import actions "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"
//
//   var input iInput
//
//   state := iotevents.NewState(&StateProps{
//   	StateName: jsii.String("MyState"),
//   	OnEnter: []event{
//   		&event{
//   			EventName: jsii.String("test-event"),
//   			Condition: iotevents.Expression_CurrentInput(input),
//   			Actions: []iAction{
//   				actions.NewSetTimerAction(jsii.String("MyTimer"), map[string]interface{}{
//   					"duration": cdk.Duration_seconds(jsii.Number(60)),
//   				}),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type SetTimerAction interface {
	awscdkioteventsalpha.IAction
}

// The jsii proxy struct for SetTimerAction
type jsiiProxy_SetTimerAction struct {
	internal.Type__awscdkioteventsalphaIAction
}

// Experimental.
func NewSetTimerAction(timerName *string, timerDuration TimerDuration) SetTimerAction {
	_init_.Initialize()

	if err := validateNewSetTimerActionParameters(timerName, timerDuration); err != nil {
		panic(err)
	}
	j := jsiiProxy_SetTimerAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.SetTimerAction",
		[]interface{}{timerName, timerDuration},
		&j,
	)

	return &j
}

// Experimental.
func NewSetTimerAction_Override(s SetTimerAction, timerName *string, timerDuration TimerDuration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.SetTimerAction",
		[]interface{}{timerName, timerDuration},
		s,
	)
}

