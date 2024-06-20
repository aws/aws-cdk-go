package awsstepfunctions


// Two types of state machines are available in AWS Step Functions: EXPRESS AND STANDARD.
//
// Example:
//   distributedMap := sfn.NewDistributedMap(this, jsii.String("DistributedMap"), &DistributedMapProps{
//   	MapExecutionType: sfn.StateMachineType_EXPRESS,
//   })
//
//   distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass")), &ProcessorConfig{
//   	Mode: sfn.ProcessorMode_DISTRIBUTED,
//   	ExecutionType: sfn.ProcessorType_STANDARD,
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-standard-vs-express.html
//
// Default: STANDARD.
//
type StateMachineType string

const (
	// Express Workflows are ideal for high-volume, event processing workloads.
	StateMachineType_EXPRESS StateMachineType = "EXPRESS"
	// Standard Workflows are ideal for long-running, durable, and auditable workflows.
	StateMachineType_STANDARD StateMachineType = "STANDARD"
)

