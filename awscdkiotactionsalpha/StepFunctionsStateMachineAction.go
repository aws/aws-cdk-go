package awscdkiotactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2"
)

// The action to put the record from an MQTT message to the Step Functions State Machine.
//
// Example:
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("SM"), &StateMachineProps{
//   	DefinitionBody: stepfunctions.DefinitionBody_FromChainable(stepfunctions.NewWait(this, jsii.String("Hello"), &WaitProps{
//   		Time: stepfunctions.WaitTime_Duration(awscdk.Duration_Seconds(jsii.Number(10))),
//   	})),
//   })
//
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewStepFunctionsStateMachineAction(stateMachine),
//   	},
//   })
//
// Experimental.
type StepFunctionsStateMachineAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for StepFunctionsStateMachineAction
type jsiiProxy_StepFunctionsStateMachineAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewStepFunctionsStateMachineAction(stateMachine awsstepfunctions.IStateMachine, props *StepFunctionsStateMachineActionProps) StepFunctionsStateMachineAction {
	_init_.Initialize()

	if err := validateNewStepFunctionsStateMachineActionParameters(stateMachine, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepFunctionsStateMachineAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.StepFunctionsStateMachineAction",
		[]interface{}{stateMachine, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionsStateMachineAction_Override(s StepFunctionsStateMachineAction, stateMachine awsstepfunctions.IStateMachine, props *StepFunctionsStateMachineActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.StepFunctionsStateMachineAction",
		[]interface{}{stateMachine, props},
		s,
	)
}

