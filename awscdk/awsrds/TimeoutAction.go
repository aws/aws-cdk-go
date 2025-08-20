package awsrds


// TimeoutAction defines the action to take when a timeout occurs if a scaling point is not found.
//
// Example:
//   var vpc vpc
//
//
//   cluster := rds.NewServerlessCluster(this, jsii.String("AnotherCluster"), &ServerlessClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA_POSTGRESQL(),
//   	CopyTagsToSnapshot: jsii.Boolean(true),
//   	 // whether to save the cluster tags when creating the snapshot. Default is 'true'
//   	ParameterGroup: rds.ParameterGroup_FromParameterGroupName(this, jsii.String("ParameterGroup"), jsii.String("default.aurora-postgresql11")),
//   	Vpc: Vpc,
//   	Scaling: &ServerlessScalingOptions{
//   		AutoPause: awscdk.Duration_Minutes(jsii.Number(10)),
//   		 // default is to pause after 5 minutes of idle time
//   		MinCapacity: rds.AuroraCapacityUnit_ACU_8,
//   		 // default is 2 Aurora capacity units (ACUs)
//   		MaxCapacity: rds.AuroraCapacityUnit_ACU_32,
//   		 // default is 16 Aurora capacity units (ACUs)
//   		Timeout: awscdk.Duration_Seconds(jsii.Number(100)),
//   		 // default is 5 minutes
//   		TimeoutAction: rds.TimeoutAction_FORCE_APPLY_CAPACITY_CHANGE,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v1.how-it-works.html#aurora-serverless.how-it-works.timeout-action
//
type TimeoutAction string

const (
	// FORCE_APPLY_CAPACITY_CHANGE sets the capacity to the specified value as soon as possible.
	//
	// Transactions may be interrupted, and connections to temporary tables and locks may be dropped.
	// Only select this option if your application can recover from dropped connections or incomplete transactions.
	TimeoutAction_FORCE_APPLY_CAPACITY_CHANGE TimeoutAction = "FORCE_APPLY_CAPACITY_CHANGE"
	// ROLLBACK_CAPACITY_CHANGE ignores the capacity change if a scaling point is not found.
	//
	// This is the default behavior.
	TimeoutAction_ROLLBACK_CAPACITY_CHANGE TimeoutAction = "ROLLBACK_CAPACITY_CHANGE"
)

