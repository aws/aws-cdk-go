package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Options for configuring scaling on an Aurora Serverless cluster.
//
// Example:
//   var vpc vpc
//
//
//   cluster := rds.NewServerlessCluster(this, jsii.String("AnotherCluster"), &ServerlessClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA_POSTGRESQL(),
//   	ParameterGroup: rds.ParameterGroup_FromParameterGroupName(this, jsii.String("ParameterGroup"), jsii.String("default.aurora-postgresql10")),
//   	Vpc: Vpc,
//   	Scaling: &ServerlessScalingOptions{
//   		AutoPause: awscdk.Duration_Minutes(jsii.Number(10)),
//   		 // default is to pause after 5 minutes of idle time
//   		MinCapacity: rds.AuroraCapacityUnit_ACU_8,
//   		 // default is 2 Aurora capacity units (ACUs)
//   		MaxCapacity: rds.AuroraCapacityUnit_ACU_32,
//   	},
//   })
//
// Experimental.
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
	// Experimental.
	AutoPause awscdk.Duration `field:"optional" json:"autoPause" yaml:"autoPause"`
	// The maximum capacity for an Aurora Serverless database cluster.
	// Experimental.
	MaxCapacity AuroraCapacityUnit `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum capacity for an Aurora Serverless database cluster.
	// Experimental.
	MinCapacity AuroraCapacityUnit `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}

