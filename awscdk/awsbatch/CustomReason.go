package awsbatch


// The corresponding Action will only be taken if *all* of the conditions specified here are met.
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
type CustomReason struct {
	// A glob string that will match on the job exit code.
	//
	// For example, `'40*'` will match 400, 404, 40123456789012.
	// Default: - will not match on the exit code.
	//
	OnExitCode *string `field:"optional" json:"onExitCode" yaml:"onExitCode"`
	// A glob string that will match on the reason returned by the exiting job For example, `'CannotPullContainerError*'` indicates that container needed to start the job could not be pulled.
	// Default: - will not match on the reason.
	//
	OnReason *string `field:"optional" json:"onReason" yaml:"onReason"`
	// A glob string that will match on the statusReason returned by the exiting job.
	//
	// For example, `'Host EC2*'` indicates that the spot instance has been reclaimed.
	// Default: - will not match on the status reason.
	//
	OnStatusReason *string `field:"optional" json:"onStatusReason" yaml:"onStatusReason"`
}

