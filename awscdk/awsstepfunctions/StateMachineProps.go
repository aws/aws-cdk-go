package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for defining a State Machine.
//
// Example:
//   // Define a state machine with one Pass state
//   child := sfn.NewStateMachine(this, jsii.String("ChildStateMachine"), &StateMachineProps{
//   	Definition: sfn.Chain_Start(sfn.NewPass(this, jsii.String("PassState"))),
//   })
//
//   // Include the state machine in a Task state with callback pattern
//   task := tasks.NewStepFunctionsStartExecution(this, jsii.String("ChildTask"), &StepFunctionsStartExecutionProps{
//   	StateMachine: child,
//   	IntegrationPattern: sfn.IntegrationPattern_WAIT_FOR_TASK_TOKEN,
//   	Input: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"token": sfn.JsonPath_taskToken(),
//   		"foo": jsii.String("bar"),
//   	}),
//   	Name: jsii.String("MyExecutionName"),
//   })
//
//   // Define a second state machine with the Task state above
//   // Define a second state machine with the Task state above
//   sfn.NewStateMachine(this, jsii.String("ParentStateMachine"), &StateMachineProps{
//   	Definition: task,
//   })
//
type StateMachineProps struct {
	// Comment that describes this state machine.
	// Default: - No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Definition for this state machine.
	// Deprecated: use definitionBody: DefinitionBody.fromChainable()
	Definition IChainable `field:"optional" json:"definition" yaml:"definition"`
	// Definition for this state machine.
	DefinitionBody DefinitionBody `field:"optional" json:"definitionBody" yaml:"definitionBody"`
	// substitutions for the definition body as a key-value map.
	DefinitionSubstitutions *map[string]*string `field:"optional" json:"definitionSubstitutions" yaml:"definitionSubstitutions"`
	// Configures server-side encryption of the state machine definition and execution history.
	// Default: - data is transparently encrypted using an AWS owned key.
	//
	EncryptionConfiguration EncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Defines what execution history events are logged and where they are logged.
	// Default: No logging.
	//
	Logs *LogOptions `field:"optional" json:"logs" yaml:"logs"`
	// The name of the query language used by the state machine.
	//
	// If the state does not contain a `queryLanguage` field,
	// then it will use the query language specified in this `queryLanguage` field.
	// Default: - JSON_PATH.
	//
	QueryLanguage QueryLanguage `field:"optional" json:"queryLanguage" yaml:"queryLanguage"`
	// The removal policy to apply to state machine.
	// Default: RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The execution role for the state machine service.
	// Default: A role is automatically created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// A name for the state machine.
	// Default: A name is automatically generated.
	//
	StateMachineName *string `field:"optional" json:"stateMachineName" yaml:"stateMachineName"`
	// Type of the state machine.
	// Default: StateMachineType.STANDARD
	//
	StateMachineType StateMachineType `field:"optional" json:"stateMachineType" yaml:"stateMachineType"`
	// Maximum run time for this state machine.
	// Default: No timeout.
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Specifies whether Amazon X-Ray tracing is enabled for this state machine.
	// Default: false.
	//
	TracingEnabled *bool `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
}

