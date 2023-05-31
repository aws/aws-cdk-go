package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDBCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBClusterProps := &CfnDBClusterProps{
//   	AllocatedStorage: jsii.Number(123),
//   	AssociatedRoles: []interface{}{
//   		&DBClusterRoleProperty{
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			FeatureName: jsii.String("featureName"),
//   		},
//   	},
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	BacktrackWindow: jsii.Number(123),
//   	BackupRetentionPeriod: jsii.Number(123),
//   	CopyTagsToSnapshot: jsii.Boolean(false),
//   	DatabaseName: jsii.String("databaseName"),
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	DbClusterInstanceClass: jsii.String("dbClusterInstanceClass"),
//   	DbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	DbInstanceParameterGroupName: jsii.String("dbInstanceParameterGroupName"),
//   	DbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	DbSystemId: jsii.String("dbSystemId"),
//   	DeletionProtection: jsii.Boolean(false),
//   	Domain: jsii.String("domain"),
//   	DomainIamRoleName: jsii.String("domainIamRoleName"),
//   	EnableCloudwatchLogsExports: []*string{
//   		jsii.String("enableCloudwatchLogsExports"),
//   	},
//   	EnableHttpEndpoint: jsii.Boolean(false),
//   	EnableIamDatabaseAuthentication: jsii.Boolean(false),
//   	Engine: jsii.String("engine"),
//   	EngineMode: jsii.String("engineMode"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   	Iops: jsii.Number(123),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ManageMasterUserPassword: jsii.Boolean(false),
//   	MasterUsername: jsii.String("masterUsername"),
//   	MasterUserPassword: jsii.String("masterUserPassword"),
//   	MasterUserSecret: &MasterUserSecretProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	MonitoringInterval: jsii.Number(123),
//   	MonitoringRoleArn: jsii.String("monitoringRoleArn"),
//   	NetworkType: jsii.String("networkType"),
//   	PerformanceInsightsEnabled: jsii.Boolean(false),
//   	PerformanceInsightsKmsKeyId: jsii.String("performanceInsightsKmsKeyId"),
//   	PerformanceInsightsRetentionPeriod: jsii.Number(123),
//   	Port: jsii.Number(123),
//   	PreferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	PubliclyAccessible: jsii.Boolean(false),
//   	ReplicationSourceIdentifier: jsii.String("replicationSourceIdentifier"),
//   	RestoreToTime: jsii.String("restoreToTime"),
//   	RestoreType: jsii.String("restoreType"),
//   	ScalingConfiguration: &ScalingConfigurationProperty{
//   		AutoPause: jsii.Boolean(false),
//   		MaxCapacity: jsii.Number(123),
//   		MinCapacity: jsii.Number(123),
//   		SecondsBeforeTimeout: jsii.Number(123),
//   		SecondsUntilAutoPause: jsii.Number(123),
//   		TimeoutAction: jsii.String("timeoutAction"),
//   	},
//   	ServerlessV2ScalingConfiguration: &ServerlessV2ScalingConfigurationProperty{
//   		MaxCapacity: jsii.Number(123),
//   		MinCapacity: jsii.Number(123),
//   	},
//   	SnapshotIdentifier: jsii.String("snapshotIdentifier"),
//   	SourceDbClusterIdentifier: jsii.String("sourceDbClusterIdentifier"),
//   	SourceRegion: jsii.String("sourceRegion"),
//   	StorageEncrypted: jsii.Boolean(false),
//   	StorageType: jsii.String("storageType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseLatestRestorableTime: jsii.Boolean(false),
//   	VpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
type CfnDBClusterProps struct {
	// The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster.
	//
	// This setting is required to create a Multi-AZ DB cluster.
	//
	// Valid for: Multi-AZ DB clusters only.
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster.
	//
	// IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other Amazon Web Services on your behalf.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	AssociatedRoles interface{} `field:"optional" json:"associatedRoles" yaml:"associatedRoles"`
	// A value that indicates whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window.
	//
	// By default, minor engine upgrades are applied automatically.
	//
	// Valid for: Multi-AZ DB clusters only.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// A list of Availability Zones (AZs) where instances in the DB cluster can be created.
	//
	// For information on AWS Regions and Availability Zones, see [Choosing the Regions and Availability Zones](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html) in the *Amazon Aurora User Guide* .
	//
	// Valid for: Aurora DB clusters only.
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
	//
	// Valid for: Aurora MySQL DB clusters only.
	BacktrackWindow *float64 `field:"optional" json:"backtrackWindow" yaml:"backtrackWindow"`
	// The number of days for which automated backups are retained.
	//
	// Default: 1
	//
	// Constraints:
	//
	// - Must be a value from 1 to 35
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster.
	//
	// The default is not to copy them.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	CopyTagsToSnapshot interface{} `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// The name of your database.
	//
	// If you don't provide a name, then Amazon RDS won't create a database in this DB cluster. For naming constraints, see [Naming Constraints](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon Aurora User Guide* .
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The DB cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain from 1 to 63 letters, numbers, or hyphens.
	// - First character must be a letter.
	// - Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `my-cluster1`
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example db.m6gd.xlarge. Not all DB instance classes are available in all AWS Regions , or for all database engines.
	//
	// For the full list of DB instance classes and availability for your engine, see [DB instance class](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the *Amazon RDS User Guide* .
	//
	// This setting is required to create a Multi-AZ DB cluster.
	//
	// Valid for: Multi-AZ DB clusters only.
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
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// The name of the DB parameter group to apply to all instances of the DB cluster.
	//
	// > When you apply a parameter group using the `DBInstanceParameterGroupName` parameter, the DB cluster isn't rebooted automatically. Also, parameter changes are applied immediately rather than during the next maintenance window.
	//
	// Default: The existing name setting
	//
	// Constraints:
	//
	// - The DB parameter group must be in the same DB parameter group family as this DB cluster.
	DbInstanceParameterGroupName *string `field:"optional" json:"dbInstanceParameterGroupName" yaml:"dbInstanceParameterGroupName"`
	// A DB subnet group that you want to associate with this DB cluster.
	//
	// If you are restoring a DB cluster to a point in time with `RestoreType` set to `copy-on-write` , and don't specify a DB subnet group name, then the DB cluster is restored with a default DB subnet group.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Reserved for future use.
	DbSystemId *string `field:"optional" json:"dbSystemId" yaml:"dbSystemId"`
	// A value that indicates whether the DB cluster has deletion protection enabled.
	//
	// The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Indicates the directory ID of the Active Directory to create the DB cluster.
	//
	// For Amazon Aurora DB clusters, Amazon RDS can use Kerberos authentication to authenticate users that connect to the DB cluster.
	//
	// For more information, see [Kerberos authentication](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html) in the *Amazon Aurora User Guide* .
	//
	// Valid for: Aurora DB clusters only.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Specifies the name of the IAM role to use when making API calls to the Directory Service.
	//
	// Valid for: Aurora DB clusters only.
	DomainIamRoleName *string `field:"optional" json:"domainIamRoleName" yaml:"domainIamRoleName"`
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
	// Valid values: `postgresql`
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	EnableCloudwatchLogsExports *[]*string `field:"optional" json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// A value that indicates whether to enable the HTTP endpoint for an Aurora Serverless DB cluster.
	//
	// By default, the HTTP endpoint is disabled.
	//
	// When enabled, the HTTP endpoint provides a connectionless web service API for running SQL queries on the Aurora Serverless DB cluster. You can also query your database from inside the RDS console with the query editor.
	//
	// For more information, see [Using the Data API for Aurora Serverless](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html) in the *Amazon Aurora User Guide* .
	//
	// Valid for: Aurora DB clusters only.
	EnableHttpEndpoint interface{} `field:"optional" json:"enableHttpEndpoint" yaml:"enableHttpEndpoint"`
	// A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	//
	// By default, mapping is disabled.
	//
	// For more information, see [IAM Database Authentication](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html) in the *Amazon Aurora User Guide.*
	//
	// Valid for: Aurora DB clusters only.
	EnableIamDatabaseAuthentication interface{} `field:"optional" json:"enableIamDatabaseAuthentication" yaml:"enableIamDatabaseAuthentication"`
	// The name of the database engine to be used for this DB cluster.
	//
	// Valid Values:
	//
	// - `aurora` (for MySQL 5.6-compatible Aurora)
	// - `aurora-mysql` (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora)
	// - `aurora-postgresql`
	// - `mysql`
	// - `postgres`
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The DB engine mode of the DB cluster, either `provisioned` , `serverless` , `parallelquery` , `global` , or `multimaster` .
	//
	// The `serverless` engine mode only applies for Aurora Serverless v1 DB clusters.
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
	//
	// Valid for: Aurora DB clusters only.
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
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
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
	//
	// Valid for: Aurora DB clusters only.
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster.
	//
	// For information about valid IOPS values, see [Provisioned IOPS storage](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS) in the *Amazon RDS User Guide* .
	//
	// This setting is required to create a Multi-AZ DB cluster.
	//
	// Constraints: Must be a multiple between .5 and 50 of the storage amount for the DB cluster.
	//
	// Valid for: Multi-AZ DB clusters only.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The Amazon Resource Name (ARN) of the AWS KMS key that is used to encrypt the database instances in the DB cluster, such as `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef` .
	//
	// If you enable the `StorageEncrypted` property but don't specify this property, the default KMS key is used. If you specify this property, you must set the `StorageEncrypted` property to `true` .
	//
	// If you specify the `SnapshotIdentifier` property, the `StorageEncrypted` property value is inherited from the snapshot, and if the DB cluster is encrypted, the specified `KmsKeyId` property is used.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A value that indicates whether to manage the master user password with AWS Secrets Manager.
	//
	// For more information, see [Password management with AWS Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide* and [Password management with AWS Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html) in the *Amazon Aurora User Guide.*
	//
	// Constraints:
	//
	// - Can't manage the master user password with AWS Secrets Manager if `MasterUserPassword` is specified.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	ManageMasterUserPassword interface{} `field:"optional" json:"manageMasterUserPassword" yaml:"manageMasterUserPassword"`
	// The name of the master user for the DB cluster.
	//
	// > If you specify the `SourceDBClusterIdentifier` , `SnapshotIdentifier` , or `GlobalClusterIdentifier` property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB cluster for the global database cluster, respectively.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	MasterUsername *string `field:"optional" json:"masterUsername" yaml:"masterUsername"`
	// The master password for the DB instance.
	//
	// > If you specify the `SourceDBClusterIdentifier` , `SnapshotIdentifier` , or `GlobalClusterIdentifier` property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB cluster for the global database cluster, respectively.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
	// Contains the secret managed by RDS in AWS Secrets Manager for the master user password.
	//
	// For more information, see [Password management with AWS Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide* and [Password management with AWS Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html) in the *Amazon Aurora User Guide.*
	MasterUserSecret interface{} `field:"optional" json:"masterUserSecret" yaml:"masterUserSecret"`
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster.
	//
	// To turn off collecting Enhanced Monitoring metrics, specify 0. The default is 0.
	//
	// If `MonitoringRoleArn` is specified, also set `MonitoringInterval` to a value other than 0.
	//
	// Valid Values: `0, 1, 5, 10, 15, 30, 60`
	//
	// Valid for: Multi-AZ DB clusters only.
	MonitoringInterval *float64 `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.
	//
	// An example is `arn:aws:iam:123456789012:role/emaccess` . For information on creating a monitoring role, see [Setting up and enabling Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the *Amazon RDS User Guide* .
	//
	// If `MonitoringInterval` is set to a value other than 0, supply a `MonitoringRoleArn` value.
	//
	// Valid for: Multi-AZ DB clusters only.
	MonitoringRoleArn *string `field:"optional" json:"monitoringRoleArn" yaml:"monitoringRoleArn"`
	// The network type of the DB cluster.
	//
	// Valid values:
	//
	// - `IPV4`
	// - `DUAL`
	//
	// The network type is determined by the `DBSubnetGroup` specified for the DB cluster. A `DBSubnetGroup` can support only the IPv4 protocol or the IPv4 and IPv6 protocols ( `DUAL` ).
	//
	// For more information, see [Working with a DB instance in a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the *Amazon Aurora User Guide.*
	//
	// Valid for: Aurora DB clusters only.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// A value that indicates whether to turn on Performance Insights for the DB cluster.
	//
	// For more information, see [Using Amazon Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the *Amazon RDS User Guide* .
	//
	// Valid for: Multi-AZ DB clusters only.
	PerformanceInsightsEnabled interface{} `field:"optional" json:"performanceInsightsEnabled" yaml:"performanceInsightsEnabled"`
	// The AWS KMS key identifier for encryption of Performance Insights data.
	//
	// The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
	//
	// If you don't specify a value for `PerformanceInsightsKMSKeyId` , then Amazon RDS uses your default KMS key. There is a default KMS key for your AWS account . Your AWS account has a different default KMS key for each AWS Region .
	//
	// Valid for: Multi-AZ DB clusters only.
	PerformanceInsightsKmsKeyId *string `field:"optional" json:"performanceInsightsKmsKeyId" yaml:"performanceInsightsKmsKeyId"`
	// The number of days to retain Performance Insights data. The default is 7 days. The following values are valid:.
	//
	// - 7
	// - *month* * 31, where *month* is a number of months from 1-23
	// - 731
	//
	// For example, the following values are valid:
	//
	// - 93 (3 months * 31)
	// - 341 (11 months * 31)
	// - 589 (19 months * 31)
	// - 731
	//
	// If you specify a retention period such as 94, which isn't a valid value, RDS issues an error.
	//
	// Valid for: Multi-AZ DB clusters only.
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
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
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
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
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
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// A value that indicates whether the DB cluster is publicly accessible.
	//
	// When the DB cluster is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB cluster's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB cluster's VPC. Access to the DB cluster is ultimately controlled by the security group it uses. That public access isn't permitted if the security group assigned to the DB cluster doesn't permit it.
	//
	// When the DB cluster isn't publicly accessible, it is an internal DB cluster with a DNS name that resolves to a private IP address.
	//
	// Default: The default behavior varies depending on whether `DBSubnetGroupName` is specified.
	//
	// If `DBSubnetGroupName` isn't specified, and `PubliclyAccessible` isn't specified, the following applies:
	//
	// - If the default VPC in the target Region doesn’t have an internet gateway attached to it, the DB cluster is private.
	// - If the default VPC in the target Region has an internet gateway attached to it, the DB cluster is public.
	//
	// If `DBSubnetGroupName` is specified, and `PubliclyAccessible` isn't specified, the following applies:
	//
	// - If the subnets are part of a VPC that doesn’t have an internet gateway attached to it, the DB cluster is private.
	// - If the subnets are part of a VPC that has an internet gateway attached to it, the DB cluster is public.
	//
	// Valid for: Multi-AZ DB clusters only.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a read replica.
	//
	// Valid for: Aurora DB clusters only.
	ReplicationSourceIdentifier *string `field:"optional" json:"replicationSourceIdentifier" yaml:"replicationSourceIdentifier"`
	// The date and time to restore the DB cluster to.
	//
	// Valid Values: Value must be a time in Universal Coordinated Time (UTC) format
	//
	// Constraints:
	//
	// - Must be before the latest restorable time for the DB instance
	// - Must be specified if `UseLatestRestorableTime` parameter isn't provided
	// - Can't be specified if the `UseLatestRestorableTime` parameter is enabled
	// - Can't be specified if the `RestoreType` parameter is `copy-on-write`
	//
	// Example: `2015-03-07T23:45:00Z`
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	RestoreToTime *string `field:"optional" json:"restoreToTime" yaml:"restoreToTime"`
	// The type of restore to be performed. You can specify one of the following values:.
	//
	// - `full-copy` - The new DB cluster is restored as a full copy of the source DB cluster.
	// - `copy-on-write` - The new DB cluster is restored as a clone of the source DB cluster.
	//
	// Constraints: You can't specify `copy-on-write` if the engine version of the source DB cluster is earlier than 1.11.
	//
	// If you don't specify a `RestoreType` value, then the new DB cluster is restored as a full copy of the source DB cluster.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	RestoreType *string `field:"optional" json:"restoreType" yaml:"restoreType"`
	// The `ScalingConfiguration` property type specifies the scaling configuration of an Aurora Serverless DB cluster.
	//
	// This property is only supported for Aurora Serverless v1. For Aurora Serverless v2, use `ServerlessV2ScalingConfiguration` property.
	//
	// Valid for: Aurora DB clusters only.
	ScalingConfiguration interface{} `field:"optional" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// The `ServerlessV2ScalingConfiguration` property type specifies the scaling configuration of an Aurora Serverless V2 DB cluster.
	//
	// This property is only supported for Aurora Serverless v2. For Aurora Serverless v1, use `ScalingConfiguration` property.
	//
	// Valid for: Aurora DB clusters only.
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
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// When restoring a DB cluster to a point in time, the identifier of the source DB cluster from which to restore.
	//
	// Constraints:
	//
	// - Must match the identifier of an existing DBCluster.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// The AWS Region which contains the source DB cluster when replicating a DB cluster. For example, `us-east-1` .
	//
	// Valid for: Aurora DB clusters only.
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
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// Specifies the storage type to be associated with the DB cluster.
	//
	// This setting is required to create a Multi-AZ DB cluster.
	//
	// When specified for a Multi-AZ DB cluster, a value for the `Iops` parameter is required.
	//
	// Valid values: `aurora` , `aurora-iopt1` (Aurora DB clusters); `io1` (Multi-AZ DB clusters)
	//
	// Default: `aurora` (Aurora DB clusters); `io1` (Multi-AZ DB clusters)
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters
	//
	// For more information on storage types for Aurora DB clusters, see [Storage configurations for Amazon Aurora DB clusters](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.StorageReliability.html#aurora-storage-type) . For more information on storage types for Multi-AZ DB clusters, see [Settings for creating Multi-AZ DB clusters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/create-multi-az-db-cluster.html#create-multi-az-db-cluster-settings) .
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// An optional array of key-value pairs to apply to this DB cluster.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A value that indicates whether to restore the DB cluster to the latest restorable backup time.
	//
	// By default, the DB cluster is not restored to the latest restorable backup time.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
	// A list of EC2 VPC security groups to associate with this DB cluster.
	//
	// If you plan to update the resource, don't specify VPC security groups in a shared VPC.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

