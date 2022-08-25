package awsstepfunctions


// Two types of state machines are available in AWS Step Functions: EXPRESS AND STANDARD.
//
// Example:
//   stateMachineDefinition := stepfunctions.NewPass(this, jsii.String("PassState"))
//
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("StateMachine"), &stateMachineProps{
//   	definition: stateMachineDefinition,
//   	stateMachineType: stepfunctions.stateMachineType_EXPRESS,
//   })
//
//   apigateway.NewStepFunctionsRestApi(this, jsii.String("StepFunctionsRestApi"), &stepFunctionsRestApiProps{
//   	deploy: jsii.Boolean(true),
//   	stateMachine: stateMachine,
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-standard-vs-express.html
//
type StateMachineType string

const (
	// Express Workflows are ideal for high-volume, event processing workloads.
	StateMachineType_EXPRESS StateMachineType = "EXPRESS"
	// Standard Workflows are ideal for long-running, durable, and auditable workflows.
	StateMachineType_STANDARD StateMachineType = "STANDARD"
)

