package awsbatch


// The Action to take when all specified conditions in a RetryStrategy are met.
//
// Example:
//   jobDefn := batch.NewEcsJobDefinition(this, jsii.String("JobDefn"), &EcsJobDefinitionProps{
//   	Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("containerDefn"), &EcsEc2ContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("public.ecr.aws/amazonlinux/amazonlinux:latest")),
//   		Memory: cdk.Size_Mebibytes(jsii.Number(2048)),
//   		Cpu: jsii.Number(256),
//   	}),
//   	RetryAttempts: jsii.Number(5),
//   	RetryStrategies: []RetryStrategy{
//   		batch.RetryStrategy_Of(batch.Action_EXIT, batch.Reason_CANNOT_PULL_CONTAINER()),
//   	},
//   })
//   jobDefn.addRetryStrategy(batch.RetryStrategy_Of(batch.Action_EXIT, batch.Reason_SPOT_INSTANCE_RECLAIMED()))
//   jobDefn.addRetryStrategy(batch.RetryStrategy_Of(batch.Action_EXIT, batch.Reason_CANNOT_PULL_CONTAINER()))
//   jobDefn.addRetryStrategy(batch.RetryStrategy_Of(batch.Action_EXIT, batch.Reason_Custom(&CustomReason{
//   	OnExitCode: jsii.String("40*"),
//   	OnReason: jsii.String("some reason"),
//   })))
//
type Action string

const (
	// The job will not retry.
	Action_EXIT Action = "EXIT"
	// The job will retry.
	//
	// It can be retried up to the number of times specified in `retryAttempts`.
	Action_RETRY Action = "RETRY"
)

