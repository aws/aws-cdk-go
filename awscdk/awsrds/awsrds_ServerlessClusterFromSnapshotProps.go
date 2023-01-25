package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Properties for ``ServerlessClusterFromSnapshot``.
//
// Example:
//   var vpc vpc
//
//   rds.NewServerlessClusterFromSnapshot(this, jsii.String("Cluster"), &serverlessClusterFromSnapshotProps{
//   	engine: rds.databaseClusterEngine_AURORA_MYSQL(),
//   	vpc: vpc,
//   	snapshotIdentifier: jsii.String("mySnapshot"),
//   })
//
// Experimental.
type ServerlessClusterFromSnapshotProps struct {
	// What kind of database to start.
	// Experimental.
	Engine IClusterEngine `field:"required" json:"engine" yaml:"engine"`
	// The identifier for the DB instance snapshot or DB cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot.
	// However, you can use only the ARN to specify a DB instance snapshot.
	// Experimental.
	SnapshotIdentifier *string `field:"required" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Automatic backup retention cannot be disabled on serverless clusters.
	// Must be a value from 1 day to 35 days.
	// Experimental.
	BackupRetention awscdk.Duration `field:"optional" json:"backupRetention" yaml:"backupRetention"`
	// An optional identifier for the cluster.
	// Experimental.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Master user credentials.
	//
	// Note - It is not possible to change the master username for a snapshot;
	// however, it is possible to provide (or generate) a new password.
	// Experimental.
	Credentials SnapshotCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Name of a database which is automatically created inside the cluster.
	// Experimental.
	DefaultDatabaseName *string `field:"optional" json:"defaultDatabaseName" yaml:"defaultDatabaseName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	// Experimental.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Whether to enable the Data API.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html
	//
	// Experimental.
	EnableDataApi *bool `field:"optional" json:"enableDataApi" yaml:"enableDataApi"`
	// Additional parameters to pass to the database engine.
	// Experimental.
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Scaling configuration of an Aurora Serverless database cluster.
	// Experimental.
	Scaling *ServerlessScalingOptions `field:"optional" json:"scaling" yaml:"scaling"`
	// Security group.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Existing subnet group for the cluster.
	// Experimental.
	SubnetGroup ISubnetGroup `field:"optional" json:"subnetGroup" yaml:"subnetGroup"`
	// The VPC that this Aurora Serverless cluster has been created in.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place the instances within the VPC.
	//
	// If provided, the `vpc` property must also be specified.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

