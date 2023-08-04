package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for ``ServerlessClusterFromSnapshot``.
//
// Example:
//   var vpc vpc
//
//   rds.NewServerlessClusterFromSnapshot(this, jsii.String("Cluster"), &ServerlessClusterFromSnapshotProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA_MYSQL(),
//   	Vpc: Vpc,
//   	SnapshotIdentifier: jsii.String("mySnapshot"),
//   })
//
type ServerlessClusterFromSnapshotProps struct {
	// What kind of database to start.
	Engine IClusterEngine `field:"required" json:"engine" yaml:"engine"`
	// The identifier for the DB instance snapshot or DB cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot.
	// However, you can use only the ARN to specify a DB instance snapshot.
	SnapshotIdentifier *string `field:"required" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Automatic backup retention cannot be disabled on serverless clusters.
	// Must be a value from 1 day to 35 days.
	// Default: Duration.days(1)
	//
	BackupRetention awscdk.Duration `field:"optional" json:"backupRetention" yaml:"backupRetention"`
	// An optional identifier for the cluster.
	// Default: - A name is automatically generated.
	//
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Whether to copy tags to the snapshot when a snapshot is created.
	// Default: - true.
	//
	CopyTagsToSnapshot *bool `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Master user credentials.
	//
	// Note - It is not possible to change the master username for a snapshot;
	// however, it is possible to provide (or generate) a new password.
	// Default: - The existing username and password from the snapshot will be used.
	//
	Credentials SnapshotCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Name of a database which is automatically created inside the cluster.
	// Default: - Database is not created in cluster.
	//
	DefaultDatabaseName *string `field:"optional" json:"defaultDatabaseName" yaml:"defaultDatabaseName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	// Default: - true if removalPolicy is RETAIN, false otherwise.
	//
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Whether to enable the Data API.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html
	//
	// Default: false.
	//
	EnableDataApi *bool `field:"optional" json:"enableDataApi" yaml:"enableDataApi"`
	// Additional parameters to pass to the database engine.
	// Default: - no parameter group.
	//
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	// Default: - RemovalPolicy.SNAPSHOT (remove the cluster and instances, but retain a snapshot of the data)
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Scaling configuration of an Aurora Serverless database cluster.
	// Default: - Serverless cluster is automatically paused after 5 minutes of being idle.
	// minimum capacity: 2 ACU
	// maximum capacity: 16 ACU.
	//
	Scaling *ServerlessScalingOptions `field:"optional" json:"scaling" yaml:"scaling"`
	// Security group.
	// Default: - a new security group is created if `vpc` was provided.
	// If the `vpc` property was not provided, no VPC security groups will be associated with the DB cluster.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Existing subnet group for the cluster.
	// Default: - a new subnet group is created if `vpc` was provided.
	// If the `vpc` property was not provided, no subnet group will be associated with the DB cluster.
	//
	SubnetGroup ISubnetGroup `field:"optional" json:"subnetGroup" yaml:"subnetGroup"`
	// The VPC that this Aurora Serverless cluster has been created in.
	// Default: - the default VPC in the account and region will be used.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place the instances within the VPC.
	//
	// If provided, the `vpc` property must also be specified.
	// Default: - the VPC default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

