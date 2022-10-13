package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDBCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBClusterProps := &cfnDBClusterProps{
//   	engine: jsii.String("engine"),
//
//   	// the properties below are optional
//   	allocatedStorage: jsii.Number(123),
//   	associatedRoles: []interface{}{
//   		&dBClusterRoleProperty{
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			featureName: jsii.String("featureName"),
//   		},
//   	},
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	backtrackWindow: jsii.Number(123),
//   	backupRetentionPeriod: jsii.Number(123),
//   	copyTagsToSnapshot: jsii.Boolean(false),
//   	databaseName: jsii.String("databaseName"),
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	dbClusterInstanceClass: jsii.String("dbClusterInstanceClass"),
//   	dbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	dbInstanceParameterGroupName: jsii.String("dbInstanceParameterGroupName"),
//   	dbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	deletionProtection: jsii.Boolean(false),
//   	enableCloudwatchLogsExports: []*string{
//   		jsii.String("enableCloudwatchLogsExports"),
//   	},
//   	enableHttpEndpoint: jsii.Boolean(false),
//   	enableIamDatabaseAuthentication: jsii.Boolean(false),
//   	engineMode: jsii.String("engineMode"),
//   	engineVersion: jsii.String("engineVersion"),
//   	globalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   	iops: jsii.Number(123),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	masterUsername: jsii.String("masterUsername"),
//   	masterUserPassword: jsii.String("masterUserPassword"),
//   	monitoringInterval: jsii.Number(123),
//   	monitoringRoleArn: jsii.String("monitoringRoleArn"),
//   	performanceInsightsEnabled: jsii.Boolean(false),
//   	performanceInsightsKmsKeyId: jsii.String("performanceInsightsKmsKeyId"),
//   	performanceInsightsRetentionPeriod: jsii.Number(123),
//   	port: jsii.Number(123),
//   	preferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	replicationSourceIdentifier: jsii.String("replicationSourceIdentifier"),
//   	restoreType: jsii.String("restoreType"),
//   	scalingConfiguration: &scalingConfigurationProperty{
//   		autoPause: jsii.Boolean(false),
//   		maxCapacity: jsii.Number(123),
//   		minCapacity: jsii.Number(123),
//   		secondsUntilAutoPause: jsii.Number(123),
//   	},
//   	serverlessV2ScalingConfiguration: &serverlessV2ScalingConfigurationProperty{
//   		maxCapacity: jsii.Number(123),
//   		minCapacity: jsii.Number(123),
//   	},
//   	snapshotIdentifier: jsii.String("snapshotIdentifier"),
//   	sourceDbClusterIdentifier: jsii.String("sourceDbClusterIdentifier"),
//   	sourceRegion: jsii.String("sourceRegion"),
//   	storageEncrypted: jsii.Boolean(false),
//   	storageType: jsii.String("storageType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	useLatestRestorableTime: jsii.Boolean(false),
//   	vpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
type CfnDBClusterProps struct {
	// The name of the database engine to be used for this DB cluster.
	//
	// Valid Values: `aurora` (for MySQL 5.6-compatible Aurora), `aurora-mysql` (for MySQL 5.7-compatible Aurora), and `aurora-postgresql`
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// `AWS::RDS::DBCluster.AllocatedStorage`.
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster.
	//
	// IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other Amazon Web Services on your behalf.
	AssociatedRoles interface{} `field:"optional" json:"associatedRoles" yaml:"associatedRoles"`
	// `AWS::RDS::DBCluster.AutoMinorVersionUpgrade`.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// A list of Availability Zones (AZs) where instances in the DB cluster can be created.
	//
	// For information on AWS Regions and Availability Zones, see [Choosing the Regions and Availability Zones](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html) in the *Amazon Aurora User Guide* .
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// The target backtrack window, in seconds. To disable backtracking, set this value to 0.
	//
	// > Currently, Backtrack is only supported for Aurora MySQL DB clusters.
	//
	// Default: 0
	//
	// Constraints:
	//
	// - If specified, this value must be set to a number from 0 to 259,200 (72 hours).
	BacktrackWindow *float64 `field:"optional" json:"backtrackWindow" yaml:"backtrackWindow"`
	// The number of days for which automated backups are retained.
	//
	// Default: 1
	//
	// Constraints:
	//
	// - Must be a value from 1 to 35.
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster.
	//
	// The default is not to copy them.
	CopyTagsToSnapshot interface{} `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// The name of your database.
	//
	// If you don't provide a name, then Amazon RDS won't create a database in this DB cluster. For naming constraints, see [Naming Constraints](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon Aurora User Guide* .
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The DB cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain from 1 to 63 letters, numbers, or hyphens.
	// - First character must be a letter.
	// - Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `my-cluster1`.
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// `AWS::RDS::DBCluster.DBClusterInstanceClass`.
	DbClusterInstanceClass *string `field:"optional" json:"dbClusterInstanceClass" yaml:"dbClusterInstanceClass"`
	// The name of the DB cluster parameter group to associate with this DB cluster.
	//
	// > If you apply a parameter group to an existing DB cluster, then its DB instances might need to reboot. This can result in an outage while the DB instances are rebooting.
	// >
	// > If you apply a change to parameter group associated with a stopped DB cluster, then the update stack waits until the DB cluster is started.
	//
	// To list all of the available DB cluster parameter group names, use the following command:
	//
	// `aws rds describe-db-cluster-parameter-groups --query "DBClusterParameterGroups[].DBClusterParameterGroupName" --output text`
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// `AWS::RDS::DBCluster.DBInstanceParameterGroupName`.
	DbInstanceParameterGroupName *string `field:"optional" json:"dbInstanceParameterGroupName" yaml:"dbInstanceParameterGroupName"`
	// A DB subnet group that you want to associate with this DB cluster.
	//
	// If you are restoring a DB cluster to a point in time with `RestoreType` set to `copy-on-write` , and don't specify a DB subnet group name, then the DB cluster is restored with a default DB subnet group.
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// A value that indicates whether the DB cluster has deletion protection enabled.
	//
	// The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	//
	// The values in the list depend on the DB engine being used. For more information, see [Publishing Database Logs to Amazon CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the *Amazon Aurora User Guide* .
	//
	// *Aurora MySQL*
	//
	// Valid values: `audit` , `error` , `general` , `slowquery`
	//
	// *Aurora PostgreSQL*
	//
	// Valid values: `postgresql`.
	EnableCloudwatchLogsExports *[]*string `field:"optional" json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// A value that indicates whether to enable the HTTP endpoint for an Aurora Serverless DB cluster.
	//
	// By default, the HTTP endpoint is disabled.
	//
	// When enabled, the HTTP endpoint provides a connectionless web service API for running SQL queries on the Aurora Serverless DB cluster. You can also query your database from inside the RDS console with the query editor.
	//
	// For more information, see [Using the Data API for Aurora Serverless](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html) in the *Amazon Aurora User Guide* .
	EnableHttpEndpoint interface{} `field:"optional" json:"enableHttpEndpoint" yaml:"enableHttpEndpoint"`
	// A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	//
	// By default, mapping is disabled.
	//
	// For more information, see [IAM Database Authentication](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html) in the *Amazon Aurora User Guide.*
	EnableIamDatabaseAuthentication interface{} `field:"optional" json:"enableIamDatabaseAuthentication" yaml:"enableIamDatabaseAuthentication"`
	// The DB engine mode of the DB cluster, either `provisioned` , `serverless` , `parallelquery` , `global` , or `multimaster` .
	//
	// The `serverless` engine mode only supports Aurora Serverless v1. Currently, AWS CloudFormation doesn't support Aurora Serverless v2.
	//
	// The `parallelquery` engine mode isn't required for Aurora MySQL version 1.23 and higher 1.x versions, and version 2.09 and higher 2.x versions.
	//
	// The `global` engine mode isn't required for Aurora MySQL version 1.22 and higher 1.x versions, and `global` engine mode isn't required for any 2.x versions.
	//
	// The `multimaster` engine mode only applies for DB clusters created with Aurora MySQL version 5.6.10a.
	//
	// For Aurora PostgreSQL, the `global` engine mode isn't required, and both the `parallelquery` and the `multimaster` engine modes currently aren't supported.
	//
	// Limitations and requirements apply to some DB engine modes. For more information, see the following sections in the *Amazon Aurora User Guide* :
	//
	// - [Limitations of Aurora Serverless](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html#aurora-serverless.limitations)
	// - [Limitations of Parallel Query](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-mysql-parallel-query.html#aurora-mysql-parallel-query-limitations)
	// - [Limitations of Aurora Global Databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database.limitations)
	// - [Limitations of Multi-Master Clusters](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-multi-master.html#aurora-multi-master-limitations)
	EngineMode *string `field:"optional" json:"engineMode" yaml:"engineMode"`
	// The version number of the database engine to use.
	//
	// To list all of the available engine versions for `aurora` (for MySQL 5.6-compatible Aurora), use the following command:
	//
	// `aws rds describe-db-engine-versions --engine aurora --query "DBEngineVersions[].EngineVersion"`
	//
	// To list all of the available engine versions for `aurora-mysql` (for MySQL 5.7-compatible Aurora), use the following command:
	//
	// `aws rds describe-db-engine-versions --engine aurora-mysql --query "DBEngineVersions[].EngineVersion"`
	//
	// To list all of the available engine versions for `aurora-postgresql` , use the following command:
	//
	// `aws rds describe-db-engine-versions --engine aurora-postgresql --query "DBEngineVersions[].EngineVersion"`
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// If you are configuring an Aurora global database cluster and want your Aurora DB cluster to be a secondary member in the global database cluster, specify the global cluster ID of the global database cluster.
	//
	// To define the primary database cluster of the global cluster, use the [AWS::RDS::GlobalCluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html) resource.
	//
	// If you aren't configuring a global database cluster, don't specify this property.
	//
	// > To remove the DB cluster from a global database cluster, specify an empty value for the `GlobalClusterIdentifier` property.
	//
	// For information about Aurora global databases, see [Working with Amazon Aurora Global Databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html) in the *Amazon Aurora User Guide* .
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// `AWS::RDS::DBCluster.Iops`.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The Amazon Resource Name (ARN) of the AWS KMS key that is used to encrypt the database instances in the DB cluster, such as `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef` .
	//
	// If you enable the `StorageEncrypted` property but don't specify this property, the default KMS key is used. If you specify this property, you must set the `StorageEncrypted` property to `true` .
	//
	// If you specify the `SnapshotIdentifier` property, the `StorageEncrypted` property value is inherited from the snapshot, and if the DB cluster is encrypted, the specified `KmsKeyId` property is used.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the master user for the DB cluster.
	//
	// > If you specify the `SourceDBClusterIdentifier` , `SnapshotIdentifier` , or `GlobalClusterIdentifier` property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB cluster for the global database cluster, respectively.
	MasterUsername *string `field:"optional" json:"masterUsername" yaml:"masterUsername"`
	// The master password for the DB instance.
	//
	// > If you specify the `SourceDBClusterIdentifier` , `SnapshotIdentifier` , or `GlobalClusterIdentifier` property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB cluster for the global database cluster, respectively.
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
	// `AWS::RDS::DBCluster.MonitoringInterval`.
	MonitoringInterval *float64 `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// `AWS::RDS::DBCluster.MonitoringRoleArn`.
	MonitoringRoleArn *string `field:"optional" json:"monitoringRoleArn" yaml:"monitoringRoleArn"`
	// `AWS::RDS::DBCluster.PerformanceInsightsEnabled`.
	PerformanceInsightsEnabled interface{} `field:"optional" json:"performanceInsightsEnabled" yaml:"performanceInsightsEnabled"`
	// `AWS::RDS::DBCluster.PerformanceInsightsKmsKeyId`.
	PerformanceInsightsKmsKeyId *string `field:"optional" json:"performanceInsightsKmsKeyId" yaml:"performanceInsightsKmsKeyId"`
	// `AWS::RDS::DBCluster.PerformanceInsightsRetentionPeriod`.
	PerformanceInsightsRetentionPeriod *float64 `field:"optional" json:"performanceInsightsRetentionPeriod" yaml:"performanceInsightsRetentionPeriod"`
	// The port number on which the DB instances in the DB cluster accept connections.
	//
	// Default:
	//
	// - When `EngineMode` is `provisioned` , `3306` (for both Aurora MySQL and Aurora PostgreSQL)
	// - When `EngineMode` is `serverless` :
	//
	// - `3306` when `Engine` is `aurora` or `aurora-mysql`
	// - `5432` when `Engine` is `aurora-postgresql`
	//
	// > The `No interruption` on update behavior only applies to DB clusters. If you are updating a DB instance, see [Port](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-port) for the AWS::RDS::DBInstance resource.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The daily time range during which automated backups are created.
	//
	// For more information, see [Backup Window](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.Backups.BackupWindow) in the *Amazon Aurora User Guide.*
	//
	// Constraints:
	//
	// - Must be in the format `hh24:mi-hh24:mi` .
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week. To see the time blocks available, see [Adjusting the Preferred DB Cluster Maintenance Window](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora) in the *Amazon Aurora User Guide.*
	//
	// Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// `AWS::RDS::DBCluster.PubliclyAccessible`.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a read replica.
	ReplicationSourceIdentifier *string `field:"optional" json:"replicationSourceIdentifier" yaml:"replicationSourceIdentifier"`
	// The type of restore to be performed. You can specify one of the following values:.
	//
	// - `full-copy` - The new DB cluster is restored as a full copy of the source DB cluster.
	// - `copy-on-write` - The new DB cluster is restored as a clone of the source DB cluster.
	//
	// Constraints: You can't specify `copy-on-write` if the engine version of the source DB cluster is earlier than 1.11.
	//
	// If you don't specify a `RestoreType` value, then the new DB cluster is restored as a full copy of the source DB cluster.
	RestoreType *string `field:"optional" json:"restoreType" yaml:"restoreType"`
	// The `ScalingConfiguration` property type specifies the scaling configuration of an Aurora Serverless DB cluster.
	//
	// Currently, AWS CloudFormation only supports Aurora Serverless v1. AWS CloudFormation doesn't support Aurora Serverless v2.
	ScalingConfiguration interface{} `field:"optional" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// `AWS::RDS::DBCluster.ServerlessV2ScalingConfiguration`.
	ServerlessV2ScalingConfiguration interface{} `field:"optional" json:"serverlessV2ScalingConfiguration" yaml:"serverlessV2ScalingConfiguration"`
	// The identifier for the DB snapshot or DB cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot. However, you can use only the ARN to specify a DB snapshot.
	//
	// After you restore a DB cluster with a `SnapshotIdentifier` property, you must specify the same `SnapshotIdentifier` property for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed. However, if you don't specify the `SnapshotIdentifier` property, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB cluster is restored from the specified `SnapshotIdentifier` property, and the original DB cluster is deleted.
	//
	// If you specify the `SnapshotIdentifier` property to restore a DB cluster (as opposed to specifying it for DB cluster updates), then don't specify the following properties:
	//
	// - `GlobalClusterIdentifier`
	// - `MasterUsername`
	// - `MasterUserPassword`
	// - `ReplicationSourceIdentifier`
	// - `RestoreType`
	// - `SourceDBClusterIdentifier`
	// - `SourceRegion`
	// - `StorageEncrypted` (for an encrypted snapshot)
	// - `UseLatestRestorableTime`
	//
	// Constraints:
	//
	// - Must match the identifier of an existing Snapshot.
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// When restoring a DB cluster to a point in time, the identifier of the source DB cluster from which to restore.
	//
	// Constraints:
	//
	// - Must match the identifier of an existing DBCluster.
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// The AWS Region which contains the source DB cluster when replicating a DB cluster.
	//
	// For example, `us-east-1` .
	SourceRegion *string `field:"optional" json:"sourceRegion" yaml:"sourceRegion"`
	// Indicates whether the DB cluster is encrypted.
	//
	// If you specify the `KmsKeyId` property, then you must enable encryption.
	//
	// If you specify the `SourceDBClusterIdentifier` property, don't specify this property. The value is inherited from the source DB cluster, and if the DB cluster is encrypted, the specified `KmsKeyId` property is used.
	//
	// If you specify the `SnapshotIdentifier` and the specified snapshot is encrypted, don't specify this property. The value is inherited from the snapshot, and the specified `KmsKeyId` property is used.
	//
	// If you specify the `SnapshotIdentifier` and the specified snapshot isn't encrypted, you can use this property to specify that the restored DB cluster is encrypted. Specify the `KmsKeyId` property for the KMS key to use for encryption. If you don't want the restored DB cluster to be encrypted, then don't set this property or set it to `false` .
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// `AWS::RDS::DBCluster.StorageType`.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Tags to assign to the DB cluster.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A value that indicates whether to restore the DB cluster to the latest restorable backup time.
	//
	// By default, the DB cluster is not restored to the latest restorable backup time.
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
	// A list of EC2 VPC security groups to associate with this DB cluster.
	//
	// If you plan to update the resource, don't specify VPC security groups in a shared VPC.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

