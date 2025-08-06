package awsstepfunctions


// Two types of state machines are available in AWS Step Functions: EXPRESS AND STANDARD.
//
// Example:
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   kmsKey := kms.NewKey(this, jsii.String("Key"))
//   stateMachine := sfn.NewStateMachine(this, jsii.String("StateMachineWithCMKEncryptionConfiguration"), &StateMachineProps{
//   	StateMachineName: jsii.String("StateMachineWithCMKEncryptionConfiguration"),
//   	DefinitionBody: sfn.DefinitionBody_FromChainable(sfn.Chain_Start(sfn.NewPass(this, jsii.String("Pass")))),
//   	StateMachineType: sfn.StateMachineType_STANDARD,
//   	EncryptionConfiguration: sfn.NewCustomerManagedEncryptionConfiguration(kmsKey, cdk.Duration_Seconds(jsii.Number(60))),
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

