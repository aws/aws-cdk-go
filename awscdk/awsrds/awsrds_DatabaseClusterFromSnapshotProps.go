package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Properties for ``DatabaseClusterFromSnapshot``.
//
// Example:
//   var vpc vpc
//
//   rds.NewDatabaseClusterFromSnapshot(this, jsii.String("Database"), &databaseClusterFromSnapshotProps{
//   	engine: rds.databaseClusterEngine.aurora(&auroraClusterEngineProps{
//   		version: rds.auroraEngineVersion_VER_1_22_2(),
//   	}),
//   	instanceProps: &instanceProps{
//   		vpc: vpc,
//   	},
//   	snapshotIdentifier: jsii.String("mySnapshot"),
//   })
//
// Experimental.
type DatabaseClusterFromSnapshotProps struct {
	// What kind of database to start.
	// Experimental.
	Engine IClusterEngine `field:"required" json:"engine" yaml:"engine"`
	// Settings for the individual instances that are launched.
	// Experimental.
	InstanceProps *InstanceProps `field:"required" json:"instanceProps" yaml:"instanceProps"`
	// The identifier for the DB instance snapshot or DB cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot.
	// However, you can use only the ARN to specify a DB instance snapshot.
	// Experimental.
	SnapshotIdentifier *string `field:"required" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// The number of seconds to set a cluster's target backtrack window to.
	//
	// This feature is only supported by the Aurora MySQL database engine and
	// cannot be enabled on existing clusters.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Managing.Backtrack.html
	//
	// Experimental.
	BacktrackWindow awscdk.Duration `field:"optional" json:"backtrackWindow" yaml:"backtrackWindow"`
	// Backup settings.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow
	//
	// Experimental.
	Backup *BackupProps `field:"optional" json:"backup" yaml:"backup"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// Experimental.
	CloudwatchLogsExports *[]*string `field:"optional" json:"cloudwatchLogsExports" yaml:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	// Experimental.
	CloudwatchLogsRetention awslogs.RetentionDays `field:"optional" json:"cloudwatchLogsRetention" yaml:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	CloudwatchLogsRetentionRole awsiam.IRole `field:"optional" json:"cloudwatchLogsRetentionRole" yaml:"cloudwatchLogsRetentionRole"`
	// An optional identifier for the cluster.
	// Experimental.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Whether to copy tags to the snapshot when a snapshot is created.
	// Experimental.
	CopyTagsToSnapshot *bool `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Credentials for the administrative user.
	// Experimental.
	Credentials Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Name of a database which is automatically created inside the cluster.
	// Experimental.
	DefaultDatabaseName *string `field:"optional" json:"defaultDatabaseName" yaml:"defaultDatabaseName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	// Experimental.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	// Experimental.
	IamAuthentication *bool `field:"optional" json:"iamAuthentication" yaml:"iamAuthentication"`
	// Base identifier for instances.
	//
	// Every replica is named by appending the replica number to this string, 1-based.
	// Experimental.
	InstanceIdentifierBase *string `field:"optional" json:"instanceIdentifierBase" yaml:"instanceIdentifierBase"`
	// How many replicas/instances to create.
	//
	// Has to be at least 1.
	// Experimental.
	Instances *float64 `field:"optional" json:"instances" yaml:"instances"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instances.
	// Experimental.
	MonitoringInterval awscdk.Duration `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// Role that will be used to manage DB instances monitoring.
	// Experimental.
	MonitoringRole awsiam.IRole `field:"optional" json:"monitoringRole" yaml:"monitoringRole"`
	// Additional parameters to pass to the database engine.
	// Experimental.
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The parameters in the DBClusterParameterGroup to create automatically.
	//
	// You can only specify parameterGroup or parameters but not both.
	// You need to use a versioned engine to auto-generate a DBClusterParameterGroup.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// What port to listen on.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A preferred maintenance window day/time range. Should be specified as a range ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC).
	//
	// Example: 'Sun:23:45-Mon:00:15'.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance
	//
	// Experimental.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// S3 buckets that you want to load data into. This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For MySQL:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-s3-export.html
	//
	// Experimental.
	S3ExportBuckets *[]awss3.IBucket `field:"optional" json:"s3ExportBuckets" yaml:"s3ExportBuckets"`
	// Role that will be associated with this DB cluster to enable S3 export.
	//
	// This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For MySQL:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-s3-export.html
	//
	// Experimental.
	S3ExportRole awsiam.IRole `field:"optional" json:"s3ExportRole" yaml:"s3ExportRole"`
	// S3 buckets that you want to load data from. This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For MySQL:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html
	//
	// Experimental.
	S3ImportBuckets *[]awss3.IBucket `field:"optional" json:"s3ImportBuckets" yaml:"s3ImportBuckets"`
	// Role that will be associated with this DB cluster to enable S3 import.
	//
	// This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For MySQL:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html
	//
	// Experimental.
	S3ImportRole awsiam.IRole `field:"optional" json:"s3ImportRole" yaml:"s3ImportRole"`
	// Whether to enable storage encryption.
	// Experimental.
	StorageEncrypted *bool `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// The KMS key for storage encryption.
	//
	// If specified, {@link storageEncrypted} will be set to `true`.
	// Experimental.
	StorageEncryptionKey awskms.IKey `field:"optional" json:"storageEncryptionKey" yaml:"storageEncryptionKey"`
	// Existing subnet group for the cluster.
	// Experimental.
	SubnetGroup ISubnetGroup `field:"optional" json:"subnetGroup" yaml:"subnetGroup"`
}

