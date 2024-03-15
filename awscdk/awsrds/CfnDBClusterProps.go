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
//   	EnableGlobalWriteForwarding: jsii.Boolean(false),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html
//
type CfnDBClusterProps struct {
	// The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster.
	//
	// Valid for Cluster Type: Multi-AZ DB clusters only
	//
	// This setting is required to create a Multi-AZ DB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-allocatedstorage
	//
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster.
	//
	// IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other Amazon Web Services on your behalf.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-associatedroles
	//
	AssociatedRoles interface{} `field:"optional" json:"associatedRoles" yaml:"associatedRoles"`
	// Specifies whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window.
	//
	// By default, minor engine upgrades are applied automatically.
	//
	// Valid for Cluster Type: Multi-AZ DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-autominorversionupgrade
	//
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// A list of Availability Zones (AZs) where instances in the DB cluster can be created.
	//
	// For information on AWS Regions and Availability Zones, see [Choosing the Regions and Availability Zones](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html) in the *Amazon Aurora User Guide* .
	//
	// Valid for: Aurora DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-availabilityzones
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-backtrackwindow
	//
	// Default: - 0.
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-backupretentionperiod
	//
	// Default: - 1.
	//
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster.
	//
	// The default is not to copy them.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-copytagstosnapshot
	//
	CopyTagsToSnapshot interface{} `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// The name of your database.
	//
	// If you don't provide a name, then Amazon RDS won't create a database in this DB cluster. For naming constraints, see [Naming Constraints](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon Aurora User Guide* .
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-databasename
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-dbclusteridentifier
	//
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example `db.m6gd.xlarge` . Not all DB instance classes are available in all AWS Regions , or for all database engines.
	//
	// For the full list of DB instance classes and availability for your engine, see [DB instance class](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the *Amazon RDS User Guide* .
	//
	// This setting is required to create a Multi-AZ DB cluster.
	//
	// Valid for Cluster Type: Multi-AZ DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-dbclusterinstanceclass
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-dbclusterparametergroupname
	//
	// Default: - "default.aurora5.6"
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-dbinstanceparametergroupname
	//
	DbInstanceParameterGroupName *string `field:"optional" json:"dbInstanceParameterGroupName" yaml:"dbInstanceParameterGroupName"`
	// A DB subnet group that you want to associate with this DB cluster.
	//
	// If you are restoring a DB cluster to a point in time with `RestoreType` set to `copy-on-write` , and don't specify a DB subnet group name, then the DB cluster is restored with a default DB subnet group.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-dbsubnetgroupname
	//
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Reserved for future use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-dbsystemid
	//
	DbSystemId *string `field:"optional" json:"dbSystemId" yaml:"dbSystemId"`
	// A value that indicates whether the DB cluster has deletion protection enabled.
	//
	// The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-deletionprotection
	//
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Indicates the directory ID of the Active Directory to create the DB cluster.
	//
	// For Amazon Aurora DB clusters, Amazon RDS can use Kerberos authentication to authenticate users that connect to the DB cluster.
	//
	// For more information, see [Kerberos authentication](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html) in the *Amazon Aurora User Guide* .
	//
	// Valid for: Aurora DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Specifies the name of the IAM role to use when making API calls to the Directory Service.
	//
	// Valid for: Aurora DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-domainiamrolename
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-enablecloudwatchlogsexports
	//
	EnableCloudwatchLogsExports *[]*string `field:"optional" json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// Specifies whether to enable this DB cluster to forward write operations to the primary cluster of a global cluster (Aurora global database).
	//
	// By default, write operations are not allowed on Aurora DB clusters that are secondary clusters in an Aurora global database.
	//
	// You can set this value only on Aurora DB clusters that are members of an Aurora global database. With this parameter enabled, a secondary cluster can forward writes to the current primary cluster, and the resulting changes are replicated back to this cluster. For the primary DB cluster of an Aurora global database, this value is used immediately if the primary is demoted by a global cluster API operation, but it does nothing until then.
	//
	// Valid for Cluster Type: Aurora DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-enableglobalwriteforwarding
	//
	EnableGlobalWriteForwarding interface{} `field:"optional" json:"enableGlobalWriteForwarding" yaml:"enableGlobalWriteForwarding"`
	// Specifies whether to enable the HTTP endpoint for the DB cluster. By default, the HTTP endpoint isn't enabled.
	//
	// When enabled, the HTTP endpoint provides a connectionless web service API (RDS Data API) for running SQL queries on the DB cluster. You can also query your database from inside the RDS console with the RDS query editor.
	//
	// RDS Data API is supported with the following DB clusters:
	//
	// - Aurora PostgreSQL Serverless v2 and provisioned
	// - Aurora PostgreSQL and Aurora MySQL Serverless v1
	//
	// For more information, see [Using RDS Data API](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html) in the *Amazon Aurora User Guide* .
	//
	// Valid for Cluster Type: Aurora DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-enablehttpendpoint
	//
	EnableHttpEndpoint interface{} `field:"optional" json:"enableHttpEndpoint" yaml:"enableHttpEndpoint"`
	// A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	//
	// By default, mapping is disabled.
	//
	// For more information, see [IAM Database Authentication](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html) in the *Amazon Aurora User Guide.*
	//
	// Valid for: Aurora DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-enableiamdatabaseauthentication
	//
	EnableIamDatabaseAuthentication interface{} `field:"optional" json:"enableIamDatabaseAuthentication" yaml:"enableIamDatabaseAuthentication"`
	// The name of the database engine to be used for this DB cluster.
	//
	// Valid Values:
	//
	// - `aurora-mysql`
	// - `aurora-postgresql`
	// - `mysql`
	// - `postgres`
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-engine
	//
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The DB engine mode of the DB cluster, either `provisioned` or `serverless` .
	//
	// The `serverless` engine mode only supports Aurora Serverless v1.
	//
	// Limitations and requirements apply to some DB engine modes. For more information, see the following sections in the *Amazon Aurora User Guide* :
	//
	// - [Limitations of Aurora Serverless v1](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html#aurora-serverless.limitations)
	// - [Requirements for Aurora Serverless v2](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.requirements.html)
	// - [Limitations of parallel query](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-mysql-parallel-query.html#aurora-mysql-parallel-query-limitations)
	// - [Limitations of Aurora global databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database.limitations)
	//
	// Valid for: Aurora DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-enginemode
	//
	EngineMode *string `field:"optional" json:"engineMode" yaml:"engineMode"`
	// The version number of the database engine to use.
	//
	// To list all of the available engine versions for Aurora MySQL version 2 (5.7-compatible) and version 3 (8.0-compatible), use the following command:
	//
	// `aws rds describe-db-engine-versions --engine aurora-mysql --query "DBEngineVersions[].EngineVersion"`
	//
	// You can supply either `5.7` or `8.0` to use the default engine version for Aurora MySQL version 2 or version 3, respectively.
	//
	// To list all of the available engine versions for Aurora PostgreSQL, use the following command:
	//
	// `aws rds describe-db-engine-versions --engine aurora-postgresql --query "DBEngineVersions[].EngineVersion"`
	//
	// To list all of the available engine versions for RDS for MySQL, use the following command:
	//
	// `aws rds describe-db-engine-versions --engine mysql --query "DBEngineVersions[].EngineVersion"`
	//
	// To list all of the available engine versions for RDS for PostgreSQL, use the following command:
	//
	// `aws rds describe-db-engine-versions --engine postgres --query "DBEngineVersions[].EngineVersion"`
	//
	// *Aurora MySQL*
	//
	// For information, see [Database engine updates for Amazon Aurora MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Updates.html) in the *Amazon Aurora User Guide* .
	//
	// *Aurora PostgreSQL*
	//
	// For information, see [Amazon Aurora PostgreSQL releases and engine versions](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Updates.20180305.html) in the *Amazon Aurora User Guide* .
	//
	// *MySQL*
	//
	// For information, see [Amazon RDS for MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the *Amazon RDS User Guide* .
	//
	// *PostgreSQL*
	//
	// For information, see [Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts) in the *Amazon RDS User Guide* .
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-engineversion
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-globalclusteridentifier
	//
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster.
	//
	// For information about valid IOPS values, see [Provisioned IOPS storage](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS) in the *Amazon RDS User Guide* .
	//
	// This setting is required to create a Multi-AZ DB cluster.
	//
	// Valid for Cluster Type: Multi-AZ DB clusters only
	//
	// Constraints:
	//
	// - Must be a multiple between .5 and 50 of the storage amount for the DB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The Amazon Resource Name (ARN) of the AWS KMS key that is used to encrypt the database instances in the DB cluster, such as `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef` .
	//
	// If you enable the `StorageEncrypted` property but don't specify this property, the default KMS key is used. If you specify this property, you must set the `StorageEncrypted` property to `true` .
	//
	// If you specify the `SnapshotIdentifier` property, the `StorageEncrypted` property value is inherited from the snapshot, and if the DB cluster is encrypted, the specified `KmsKeyId` property is used.
	//
	// If you create a read replica of an encrypted DB cluster in another AWS Region, make sure to set `KmsKeyId` to a KMS key identifier that is valid in the destination AWS Region. This KMS key is used to encrypt the read replica in that AWS Region.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies whether to manage the master user password with AWS Secrets Manager.
	//
	// For more information, see [Password management with AWS Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide* and [Password management with AWS Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html) in the *Amazon Aurora User Guide.*
	//
	// Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
	//
	// Constraints:
	//
	// - Can't manage the master user password with AWS Secrets Manager if `MasterUserPassword` is specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-managemasteruserpassword
	//
	ManageMasterUserPassword interface{} `field:"optional" json:"manageMasterUserPassword" yaml:"manageMasterUserPassword"`
	// The name of the master user for the DB cluster.
	//
	// > If you specify the `SourceDBClusterIdentifier` , `SnapshotIdentifier` , or `GlobalClusterIdentifier` property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB cluster for the global database cluster, respectively.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-masterusername
	//
	MasterUsername *string `field:"optional" json:"masterUsername" yaml:"masterUsername"`
	// The master password for the DB instance.
	//
	// > If you specify the `SourceDBClusterIdentifier` , `SnapshotIdentifier` , or `GlobalClusterIdentifier` property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB cluster for the global database cluster, respectively.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-masteruserpassword
	//
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
	// The secret managed by RDS in AWS Secrets Manager for the master user password.
	//
	// For more information, see [Password management with AWS Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide* and [Password management with AWS Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html) in the *Amazon Aurora User Guide.*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-masterusersecret
	//
	MasterUserSecret interface{} `field:"optional" json:"masterUserSecret" yaml:"masterUserSecret"`
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster.
	//
	// To turn off collecting Enhanced Monitoring metrics, specify `0` .
	//
	// If `MonitoringRoleArn` is specified, also set `MonitoringInterval` to a value other than `0` .
	//
	// Valid for Cluster Type: Multi-AZ DB clusters only
	//
	// Valid Values: `0 | 1 | 5 | 10 | 15 | 30 | 60`
	//
	// Default: `0`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-monitoringinterval
	//
	// Default: - 0.
	//
	MonitoringInterval *float64 `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.
	//
	// An example is `arn:aws:iam:123456789012:role/emaccess` . For information on creating a monitoring role, see [Setting up and enabling Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the *Amazon RDS User Guide* .
	//
	// If `MonitoringInterval` is set to a value other than `0` , supply a `MonitoringRoleArn` value.
	//
	// Valid for Cluster Type: Multi-AZ DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-monitoringrolearn
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-networktype
	//
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Specifies whether to turn on Performance Insights for the DB cluster.
	//
	// For more information, see [Using Amazon Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the *Amazon RDS User Guide* .
	//
	// Valid for Cluster Type: Multi-AZ DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-performanceinsightsenabled
	//
	PerformanceInsightsEnabled interface{} `field:"optional" json:"performanceInsightsEnabled" yaml:"performanceInsightsEnabled"`
	// The AWS KMS key identifier for encryption of Performance Insights data.
	//
	// The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
	//
	// If you don't specify a value for `PerformanceInsightsKMSKeyId` , then Amazon RDS uses your default KMS key. There is a default KMS key for your AWS account . Your AWS account has a different default KMS key for each AWS Region .
	//
	// Valid for Cluster Type: Multi-AZ DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-performanceinsightskmskeyid
	//
	PerformanceInsightsKmsKeyId *string `field:"optional" json:"performanceInsightsKmsKeyId" yaml:"performanceInsightsKmsKeyId"`
	// The number of days to retain Performance Insights data.
	//
	// Valid for Cluster Type: Multi-AZ DB clusters only
	//
	// Valid Values:
	//
	// - `7`
	// - *month* * 31, where *month* is a number of months from 1-23. Examples: `93` (3 months * 31), `341` (11 months * 31), `589` (19 months * 31)
	// - `731`
	//
	// Default: `7` days
	//
	// If you specify a retention period that isn't valid, such as `94` , Amazon RDS issues an error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-performanceinsightsretentionperiod
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-port
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-preferredbackupwindow
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Specifies whether the DB cluster is publicly accessible.
	//
	// When the DB cluster is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB cluster's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB cluster's VPC. Access to the DB cluster is ultimately controlled by the security group it uses. That public access isn't permitted if the security group assigned to the DB cluster doesn't permit it.
	//
	// When the DB cluster isn't publicly accessible, it is an internal DB cluster with a DNS name that resolves to a private IP address.
	//
	// Valid for Cluster Type: Multi-AZ DB clusters only
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-publiclyaccessible
	//
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a read replica.
	//
	// Valid for: Aurora DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-replicationsourceidentifier
	//
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
	// This property must be used with `SourceDBClusterIdentifier` property. The resulting cluster will have the identifier that matches the value of the `DBclusterIdentifier` property.
	//
	// Example: `2015-03-07T23:45:00Z`
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-restoretotime
	//
	RestoreToTime *string `field:"optional" json:"restoreToTime" yaml:"restoreToTime"`
	// The type of restore to be performed. You can specify one of the following values:.
	//
	// - `full-copy` - The new DB cluster is restored as a full copy of the source DB cluster.
	// - `copy-on-write` - The new DB cluster is restored as a clone of the source DB cluster.
	//
	// If you don't specify a `RestoreType` value, then the new DB cluster is restored as a full copy of the source DB cluster.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-restoretype
	//
	// Default: - "full-copy".
	//
	RestoreType *string `field:"optional" json:"restoreType" yaml:"restoreType"`
	// The `ScalingConfiguration` property type specifies the scaling configuration of an Aurora Serverless DB cluster.
	//
	// This property is only supported for Aurora Serverless v1. For Aurora Serverless v2, Use the `ServerlessV2ScalingConfiguration` property.
	//
	// Valid for: Aurora DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-scalingconfiguration
	//
	ScalingConfiguration interface{} `field:"optional" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// The `ServerlessV2ScalingConfiguration` property type specifies the scaling configuration of an Aurora Serverless V2 DB cluster.
	//
	// This property is only supported for Aurora Serverless v2. For Aurora Serverless v1, Use the `ScalingConfiguration` property.
	//
	// Valid for: Aurora DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-serverlessv2scalingconfiguration
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-snapshotidentifier
	//
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// When restoring a DB cluster to a point in time, the identifier of the source DB cluster from which to restore.
	//
	// Constraints:
	//
	// - Must match the identifier of an existing DBCluster.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-sourcedbclusteridentifier
	//
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// The AWS Region which contains the source DB cluster when replicating a DB cluster. For example, `us-east-1` .
	//
	// Valid for: Aurora DB clusters only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-sourceregion
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-storageencrypted
	//
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// The storage type to associate with the DB cluster.
	//
	// For information on storage types for Aurora DB clusters, see [Storage configurations for Amazon Aurora DB clusters](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.StorageReliability.html#aurora-storage-type) . For information on storage types for Multi-AZ DB clusters, see [Settings for creating Multi-AZ DB clusters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/create-multi-az-db-cluster.html#create-multi-az-db-cluster-settings) .
	//
	// This setting is required to create a Multi-AZ DB cluster.
	//
	// When specified for a Multi-AZ DB cluster, a value for the `Iops` parameter is required.
	//
	// Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
	//
	// Valid Values:
	//
	// - Aurora DB clusters - `aurora | aurora-iopt1`
	// - Multi-AZ DB clusters - `io1 | io2 | gp3`
	//
	// Default:
	//
	// - Aurora DB clusters - `aurora`
	// - Multi-AZ DB clusters - `io1`
	//
	// > When you create an Aurora DB cluster with the storage type set to `aurora-iopt1` , the storage type is returned in the response. The storage type isn't returned when you set it to `aurora` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-storagetype
	//
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// An optional array of key-value pairs to apply to this DB cluster.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A value that indicates whether to restore the DB cluster to the latest restorable backup time.
	//
	// By default, the DB cluster is not restored to the latest restorable backup time.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-uselatestrestorabletime
	//
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
	// A list of EC2 VPC security groups to associate with this DB cluster.
	//
	// If you plan to update the resource, don't specify VPC security groups in a shared VPC.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-vpcsecuritygroupids
	//
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

