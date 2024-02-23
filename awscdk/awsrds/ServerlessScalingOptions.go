package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for configuring scaling on an Aurora Serverless cluster.
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
type ServerlessScalingOptions struct {
	// The time before an Aurora Serverless database cluster is paused.
	//
	// A database cluster can be paused only when it is idle (it has no connections).
	// Auto pause time must be between 5 minutes and 1 day.
	//
	// If a DB cluster is paused for more than seven days, the DB cluster might be
	// backed up with a snapshot. In this case, the DB cluster is restored when there
	// is a request to connect to it.
	//
	// Set to 0 to disable.
	// Default: - automatic pause enabled after 5 minutes.
	//
	AutoPause awscdk.Duration `field:"optional" json:"autoPause" yaml:"autoPause"`
	// The maximum capacity for an Aurora Serverless database cluster.
	// Default: - determined by Aurora based on database engine.
	//
	MaxCapacity AuroraCapacityUnit `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum capacity for an Aurora Serverless database cluster.
	// Default: - determined by Aurora based on database engine.
	//
	MinCapacity AuroraCapacityUnit `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// The amount of time that Aurora Serverless v1 tries to find a scaling point to perform seamless scaling before enforcing the timeout action.
	// Default: - 5 minutes.
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The action to take when the timeout is reached.
	//
	// Selecting ForceApplyCapacityChange will force the capacity to the specified value as soon as possible, even without a scaling point.
	// Selecting RollbackCapacityChange will ignore the capacity change if a scaling point is not found. This is the default behavior.
	// Default: - TimeoutAction.ROLLBACK_CAPACITY_CHANGE
	//
	TimeoutAction TimeoutAction `field:"optional" json:"timeoutAction" yaml:"timeoutAction"`
}

