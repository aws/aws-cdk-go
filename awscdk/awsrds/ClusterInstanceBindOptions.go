package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for binding the instance to the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//   var subnetGroup subnetGroup
//
//   clusterInstanceBindOptions := &ClusterInstanceBindOptions{
//   	MonitoringInterval: cdk.Duration_Minutes(jsii.Number(30)),
//   	MonitoringRole: role,
//   	PromotionTier: jsii.Number(123),
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	SubnetGroup: subnetGroup,
//   }
//
type ClusterInstanceBindOptions struct {
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instances.
	// Default: no enhanced monitoring.
	//
	MonitoringInterval awscdk.Duration `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// Role that will be used to manage DB instances monitoring.
	// Default: - A role is automatically created for you.
	//
	MonitoringRole awsiam.IRole `field:"optional" json:"monitoringRole" yaml:"monitoringRole"`
	// The promotion tier of the cluster instance.
	//
	// This matters more for serverlessV2 instances. If a serverless
	// instance is in tier 0-1 then it will scale with the writer.
	//
	// For provisioned instances this just determines the failover priority.
	// If multiple instances have the same priority then one will be picked at random.
	// Default: 2.
	//
	PromotionTier *float64 `field:"optional" json:"promotionTier" yaml:"promotionTier"`
	// The removal policy on the cluster.
	// Default: - RemovalPolicy.DESTROY (cluster snapshot can restore)
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Existing subnet group for the cluster.
	//
	// This is only needed when using the isFromLegacyInstanceProps.
	// Default: - cluster subnet group is used.
	//
	SubnetGroup ISubnetGroup `field:"optional" json:"subnetGroup" yaml:"subnetGroup"`
}

