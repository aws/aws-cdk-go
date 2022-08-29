// Receipt Detector Model actions for AWS IoT Events
package awscdkioteventsactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2"
)

// The action to reset an existing timer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iotevents_actions_alpha "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"
//
//   resetTimerAction := iotevents_actions_alpha.NewResetTimerAction(jsii.String("timerName"))
//
// Experimental.
type ResetTimerAction interface {
	awscdkioteventsalpha.IAction
}

// The jsii proxy struct for ResetTimerAction
type jsiiProxy_ResetTimerAction struct {
	internal.Type__awscdkioteventsalphaIAction
}

// Experimental.
func NewResetTimerAction(timerName *string) ResetTimerAction {
	_init_.Initialize()

	j := jsiiProxy_ResetTimerAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.ResetTimerAction",
		[]interface{}{timerName},
		&j,
	)

	return &j
}

// Experimental.
func NewResetTimerAction_Override(r ResetTimerAction, timerName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.ResetTimerAction",
		[]interface{}{timerName},
		r,
	)
}

