// Receipt Detector Model actions for AWS IoT Events
package awscdkioteventsactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2"
)

// The action to delete an existing timer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iotevents_actions_alpha "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"
//
//   clearTimerAction := iotevents_actions_alpha.NewClearTimerAction(jsii.String("timerName"))
//
// Experimental.
type ClearTimerAction interface {
	awscdkioteventsalpha.IAction
}

// The jsii proxy struct for ClearTimerAction
type jsiiProxy_ClearTimerAction struct {
	internal.Type__awscdkioteventsalphaIAction
}

// Experimental.
func NewClearTimerAction(timerName *string) ClearTimerAction {
	_init_.Initialize()

	j := jsiiProxy_ClearTimerAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.ClearTimerAction",
		[]interface{}{timerName},
		&j,
	)

	return &j
}

// Experimental.
func NewClearTimerAction_Override(c ClearTimerAction, timerName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.ClearTimerAction",
		[]interface{}{timerName},
		c,
	)
}

