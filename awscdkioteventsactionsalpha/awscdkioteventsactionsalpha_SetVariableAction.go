// Receipt Detector Model actions for AWS IoT Events
package awscdkioteventsactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2"
)

// The action to create a variable with a specified value.
//
// Example:
//   import iotevents "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
//   import actions "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"
//
//   var input iInput
//
//   state := iotevents.NewState(&stateProps{
//   	stateName: jsii.String("MyState"),
//   	onEnter: []event{
//   		&event{
//   			eventName: jsii.String("test-event"),
//   			condition: iotevents.expression.currentInput(input),
//   			actions: []iAction{
//   				actions.NewSetVariableAction(jsii.String("MyVariable"), iotevents.*expression.inputAttribute(input, jsii.String("payload.temperature"))),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type SetVariableAction interface {
	awscdkioteventsalpha.IAction
}

// The jsii proxy struct for SetVariableAction
type jsiiProxy_SetVariableAction struct {
	internal.Type__awscdkioteventsalphaIAction
}

// Experimental.
func NewSetVariableAction(variableName *string, value awscdkioteventsalpha.Expression) SetVariableAction {
	_init_.Initialize()

	if err := validateNewSetVariableActionParameters(variableName, value); err != nil {
		panic(err)
	}
	j := jsiiProxy_SetVariableAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.SetVariableAction",
		[]interface{}{variableName, value},
		&j,
	)

	return &j
}

// Experimental.
func NewSetVariableAction_Override(s SetVariableAction, variableName *string, value awscdkioteventsalpha.Expression) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.SetVariableAction",
		[]interface{}{variableName, value},
		s,
	)
}

