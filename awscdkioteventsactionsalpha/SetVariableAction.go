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
//   				actions.NewSetVariableAction(jsii.String("MyVariable"), iotevents.Expression_InputAttribute(input, jsii.String("payload.temperature"))),
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

