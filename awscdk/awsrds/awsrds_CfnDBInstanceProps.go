package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDBInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBInstanceProps := &cfnDBInstanceProps{
//   	allocatedStorage: jsii.String("allocatedStorage"),
//   	allowMajorVersionUpgrade: jsii.Boolean(false),
//   	associatedRoles: []interface{}{
//   		&dBInstanceRoleProperty{
//   			featureName: jsii.String("featureName"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	backupRetentionPeriod: jsii.Number(123),
//   	caCertificateIdentifier: jsii.String("caCertificateIdentifier"),
//   	characterSetName: jsii.String("characterSetName"),
//   	copyTagsToSnapshot: jsii.Boolean(false),
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	dbInstanceClass: jsii.String("dbInstanceClass"),
//   	dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   	dbName: jsii.String("dbName"),
//   	dbParameterGroupName: jsii.String("dbParameterGroupName"),
//   	dbSecurityGroups: []*string{
//   		jsii.String("dbSecurityGroups"),
//   	},
//   	dbSnapshotIdentifier: jsii.String("dbSnapshotIdentifier"),
//   	dbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	deleteAutomatedBackups: jsii.Boolean(false),
//   	deletionProtection: jsii.Boolean(false),
//   	domain: jsii.String("domain"),
//   	domainIamRoleName: jsii.String("domainIamRoleName"),
//   	enableCloudwatchLogsExports: []*string{
//   		jsii.String("enableCloudwatchLogsExports"),
//   	},
//   	enableIamDatabaseAuthentication: jsii.Boolean(false),
//   	enablePerformanceInsights: jsii.Boolean(false),
//   	engine: jsii.String("engine"),
//   	engineVersion: jsii.String("engineVersion"),
//   	iops: jsii.Number(123),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	licenseModel: jsii.String("licenseModel"),
//   	masterUsername: jsii.String("masterUsername"),
//   	masterUserPassword: jsii.String("masterUserPassword"),
//   	maxAllocatedStorage: jsii.Number(123),
//   	monitoringInterval: jsii.Number(123),
//   	monitoringRoleArn: jsii.String("monitoringRoleArn"),
//   	multiAz: jsii.Boolean(false),
//   	optionGroupName: jsii.String("optionGroupName"),
//   	performanceInsightsKmsKeyId: jsii.String("performanceInsightsKmsKeyId"),
//   	performanceInsightsRetentionPeriod: jsii.Number(123),
//   	port: jsii.String("port"),
//   	preferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	processorFeatures: []interface{}{
//   		&processorFeatureProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	promotionTier: jsii.Number(123),
//   	publiclyAccessible: jsii.Boolean(false),
//   	sourceDbInstanceIdentifier: jsii.String("sourceDbInstanceIdentifier"),
//   	sourceRegion: jsii.String("sourceRegion"),
//   	storageEncrypted: jsii.Boolean(false),
//   	storageType: jsii.String("storageType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timezone: jsii.String("timezone"),
//   	useDefaultProcessorFeatures: jsii.Boolean(false),
//   	vpcSecurityGroups: []*string{
//   		jsii.String("vpcSecurityGroups"),
//   	},
//   }
//
type CfnDBInstanceProps struct {
	// The amount of storage (in gigabytes) to be initially allocated for the database instance.
	//
	// > If any value is set in the `Iops` parameter, `AllocatedStorage` must be at least 100 GiB, which corresponds to the minimum Iops value of 1,000. If you increase the `Iops` value (in 1,000 IOPS increments), then you must also increase the `AllocatedStorage` value (in 100-GiB increments).
	//
	// *Amazon Aurora*
	//
	// Not applicable. Aurora cluster volumes automatically grow as the amount of data in your database increases, though you are only charged for the space that you use in an Aurora cluster volume.
	//
	// *MySQL*
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	// - General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
	// - Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	// - Magnetic storage (standard): Must be an integer from 5 to 3072.
	//
	// *MariaDB*
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	// - General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
	// - Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	// - Magnetic storage (standard): Must be an integer from 5 to 3072.
	//
	// *PostgreSQL*
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	// - General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
	// - Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	// - Magnetic storage (standard): Must be an integer from 5 to 3072.
	//
	// *Oracle*
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	// - General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
	// - Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	// - Magnetic storage (standard): Must be an integer from 10 to 3072.
	//
	// *SQL Server*
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	// - General Purpose (SSD) storage (gp2):
	//
	// - Enterprise and Standard editions: Must be an integer from 20 to 16384.
	// - Web and Express editions: Must be an integer from 20 to 16384.
	// - Provisioned IOPS storage (io1):
	//
	// - Enterprise and Standard editions: Must be an integer from 20 to 16384.
	// - Web and Express editions: Must be an integer from 20 to 16384.
	// - Magnetic storage (standard):
	//
	// - Enterprise and Standard editions: Must be an integer from 20 to 1024.
	// - Web and Express editions: Must be an integer from 20 to 1024.
	AllocatedStorage *string `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// A value that indicates whether major version upgrades are allowed.
	//
	// Changing this parameter doesn't result in an outage and the change is asynchronously applied as soon as possible.
	//
	// Constraints: Major version upgrades must be allowed when specifying a value for the `EngineVersion` parameter that is a different major version than the DB instance's current version.
	AllowMajorVersionUpgrade interface{} `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// The AWS Identity and Access Management (IAM) roles associated with the DB instance.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The associated roles are managed by the DB cluster.
	AssociatedRoles interface{} `field:"optional" json:"associatedRoles" yaml:"associatedRoles"`
	// A value that indicates whether minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	//
	// By default, minor engine upgrades are applied automatically.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The Availability Zone (AZ) where the database will be created.
	//
	// For information on AWS Regions and Availability Zones, see [Regions and Availability Zones](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html) .
	//
	// *Amazon Aurora*
	//
	// Each Aurora DB cluster hosts copies of its storage in three separate Availability Zones. Specify one of these Availability Zones. Aurora automatically chooses an appropriate Availability Zone if you don't specify one.
	//
	// Default: A random, system-chosen Availability Zone in the endpoint's AWS Region .
	//
	// Example: `us-east-1d`
	//
	// Constraint: The `AvailabilityZone` parameter can't be specified if the DB instance is a Multi-AZ deployment. The specified Availability Zone must be in the same AWS Region as the current endpoint.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The number of days for which automated backups are retained.
	//
	// Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The retention period for automated backups is managed by the DB cluster.
	//
	// Default: 1
	//
	// Constraints:
	//
	// - Must be a value from 0 to 35
	// - Can't be set to 0 if the DB instance is a source to read replicas.
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// The identifier of the CA certificate for this DB instance.
	//
	// > Specifying or updating this property triggers a reboot.
	//
	// For more information about CA certificate identifiers for RDS DB engines, see [Rotating Your SSL/TLS Certificate](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL-certificate-rotation.html) in the *Amazon RDS User Guide* .
	//
	// For more information about CA certificate identifiers for Aurora DB engines, see [Rotating Your SSL/TLS Certificate](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL-certificate-rotation.html) in the *Amazon Aurora User Guide* .
	CaCertificateIdentifier *string `field:"optional" json:"caCertificateIdentifier" yaml:"caCertificateIdentifier"`
	// For supported engines, indicates that the DB instance should be associated with the specified character set.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The character set is managed by the DB cluster. For more information, see [AWS::RDS::DBCluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html) .
	CharacterSetName *string `field:"optional" json:"characterSetName" yaml:"characterSetName"`
	// A value that indicates whether to copy tags from the DB instance to snapshots of the DB instance.
	//
	// By default, tags are not copied.
	//
	// *Amazon Aurora*
	//
	// Not applicable. Copying tags to snapshots is managed by the DB cluster. Setting this value for an Aurora DB instance has no effect on the DB cluster setting.
	CopyTagsToSnapshot interface{} `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// The identifier of the DB cluster that the instance will belong to.
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The compute and memory capacity of the DB instance, for example, `db.m4.large` . Not all DB instance classes are available in all AWS Regions, or for all database engines.
	//
	// For the full list of DB instance classes, and availability for your engine, see [DB Instance Class](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the *Amazon RDS User Guide.* For more information about DB instance class pricing and AWS Region support for DB instance classes, see [Amazon RDS Pricing](https://docs.aws.amazon.com/rds/pricing/) .
	DbInstanceClass *string `field:"optional" json:"dbInstanceClass" yaml:"dbInstanceClass"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation converts it to lowercase. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the DB instance. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// For information about constraints that apply to DB instance identifiers, see [Naming constraints in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon RDS User Guide* .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DbInstanceIdentifier *string `field:"optional" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// The meaning of this parameter differs according to the database engine you use.
	//
	// > If you specify the `[DBSnapshotIdentifier](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsnapshotidentifier)` property, this property only applies to RDS for Oracle.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The database name is managed by the DB cluster.
	//
	// *MySQL*
	//
	// The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance.
	//
	// Constraints:
	//
	// - Must contain 1 to 64 letters or numbers.
	// - Can't be a word reserved by the specified database engine
	//
	// *MariaDB*
	//
	// The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance.
	//
	// Constraints:
	//
	// - Must contain 1 to 64 letters or numbers.
	// - Can't be a word reserved by the specified database engine
	//
	// *PostgreSQL*
	//
	// The name of the database to create when the DB instance is created. If this parameter is not specified, the default `postgres` database is created in the DB instance.
	//
	// Constraints:
	//
	// - Must contain 1 to 63 letters, numbers, or underscores.
	// - Must begin with a letter or an underscore. Subsequent characters can be letters, underscores, or digits (0-9).
	// - Can't be a word reserved by the specified database engine
	//
	// *Oracle*
	//
	// The Oracle System ID (SID) of the created DB instance. If you specify `null` , the default value `ORCL` is used. You can't specify the string NULL, or any other reserved word, for `DBName` .
	//
	// Default: `ORCL`
	//
	// Constraints:
	//
	// - Can't be longer than 8 characters
	//
	// *SQL Server*
	//
	// Not applicable. Must be null.
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// The name of an existing DB parameter group or a reference to an [AWS::RDS::DBParameterGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html) resource created in the template.
	//
	// To list all of the available DB parameter group names, use the following command:
	//
	// `aws rds describe-db-parameter-groups --query "DBParameterGroups[].DBParameterGroupName" --output text`
	//
	// > If any of the data members of the referenced parameter group are changed during an update, the DB instance might need to be restarted, which causes some interruption. If the parameter group contains static parameters, whether they were changed or not, an update triggers a reboot.
	//
	// If you don't specify a value for the `DBParameterGroupName` property, the default DB parameter group for the specified engine and engine version is used.
	DbParameterGroupName *string `field:"optional" json:"dbParameterGroupName" yaml:"dbParameterGroupName"`
	// A list of the DB security groups to assign to the DB instance.
	//
	// The list can include both the name of existing DB security groups or references to AWS::RDS::DBSecurityGroup resources created in the template.
	//
	// If you set DBSecurityGroups, you must not set VPCSecurityGroups, and vice versa. Also, note that the DBSecurityGroups property exists only for backwards compatibility with older regions and is no longer recommended for providing security information to an RDS DB instance. Instead, use VPCSecurityGroups.
	//
	// > If you specify this property, AWS CloudFormation sends only the following properties (if specified) to Amazon RDS during create operations:
	// >
	// > - `AllocatedStorage`
	// > - `AutoMinorVersionUpgrade`
	// > - `AvailabilityZone`
	// > - `BackupRetentionPeriod`
	// > - `CharacterSetName`
	// > - `DBInstanceClass`
	// > - `DBName`
	// > - `DBParameterGroupName`
	// > - `DBSecurityGroups`
	// > - `DBSubnetGroupName`
	// > - `Engine`
	// > - `EngineVersion`
	// > - `Iops`
	// > - `LicenseModel`
	// > - `MasterUsername`
	// > - `MasterUserPassword`
	// > - `MultiAZ`
	// > - `OptionGroupName`
	// > - `PreferredBackupWindow`
	// > - `PreferredMaintenanceWindow`
	// >
	// > All other properties are ignored. Specify a virtual private cloud (VPC) security group if you want to submit other properties, such as `StorageType` , `StorageEncrypted` , or `KmsKeyId` . If you're already using the `DBSecurityGroups` property, you can't use these other properties by updating your DB instance to use a VPC security group. You must recreate the DB instance.
	DbSecurityGroups *[]*string `field:"optional" json:"dbSecurityGroups" yaml:"dbSecurityGroups"`
	// The name or Amazon Resource Name (ARN) of the DB snapshot that's used to restore the DB instance.
	//
	// If you're restoring from a shared manual DB snapshot, you must specify the ARN of the snapshot.
	//
	// By specifying this property, you can create a DB instance from the specified DB snapshot. If the `DBSnapshotIdentifier` property is an empty string or the `AWS::RDS::DBInstance` declaration has no `DBSnapshotIdentifier` property, AWS CloudFormation creates a new database. If the property contains a value (other than an empty string), AWS CloudFormation creates a database from the specified snapshot. If a snapshot with the specified name doesn't exist, AWS CloudFormation can't create the database and it rolls back the stack.
	//
	// Some DB instance properties aren't valid when you restore from a snapshot, such as the `MasterUsername` and `MasterUserPassword` properties. For information about the properties that you can specify, see the `RestoreDBInstanceFromDBSnapshot` action in the *Amazon RDS API Reference* .
	//
	// After you restore a DB instance with a `DBSnapshotIdentifier` property, you must specify the same `DBSnapshotIdentifier` property for any future updates to the DB instance. When you specify this property for an update, the DB instance is not restored from the DB snapshot again, and the data in the database is not changed. However, if you don't specify the `DBSnapshotIdentifier` property, an empty DB instance is created, and the original DB instance is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB instance is restored from the specified `DBSnapshotIdentifier` property, and the original DB instance is deleted.
	//
	// If you specify the `DBSnapshotIdentifier` property to restore a DB instance (as opposed to specifying it for DB instance updates), then don't specify the following properties:
	//
	// - `CharacterSetName`
	// - `DBClusterIdentifier`
	// - `DBName`
	// - `DeleteAutomatedBackups`
	// - `EnablePerformanceInsights`
	// - `KmsKeyId`
	// - `MasterUsername`
	// - `MasterUserPassword`
	// - `PerformanceInsightsKMSKeyId`
	// - `PerformanceInsightsRetentionPeriod`
	// - `PromotionTier`
	// - `SourceDBInstanceIdentifier`
	// - `SourceRegion`
	// - `StorageEncrypted` (for an encrypted snapshot)
	// - `Timezone`
	//
	// *Amazon Aurora*
	//
	// Not applicable. Snapshot restore is managed by the DB cluster.
	DbSnapshotIdentifier *string `field:"optional" json:"dbSnapshotIdentifier" yaml:"dbSnapshotIdentifier"`
	// A DB subnet group to associate with the DB instance.
	//
	// If you update this value, the new subnet group must be a subnet group in a new VPC.
	//
	// If there's no DB subnet group, then the DB instance isn't a VPC DB instance.
	//
	// For more information about using Amazon RDS in a VPC, see [Using Amazon RDS with Amazon Virtual Private Cloud (VPC)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.html) in the *Amazon RDS User Guide* .
	//
	// *Amazon Aurora*
	//
	// Not applicable. The DB subnet group is managed by the DB cluster. If specified, the setting must match the DB cluster setting.
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// A value that indicates whether to remove automated backups immediately after the DB instance is deleted.
	//
	// This parameter isn't case-sensitive. The default is to remove automated backups immediately after the DB instance is deleted.
	DeleteAutomatedBackups interface{} `field:"optional" json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// A value that indicates whether the DB instance has deletion protection enabled.
	//
	// The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled. For more information, see [Deleting a DB Instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html) .
	//
	// *Amazon Aurora*
	//
	// Not applicable. You can enable or disable deletion protection for the DB cluster. For more information, see `CreateDBCluster` . DB instances in a DB cluster can be deleted even when deletion protection is enabled for the DB cluster.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	//
	// Currently, only Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active Directory Domain.
	//
	// For more information, see [Kerberos Authentication](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html) in the *Amazon RDS User Guide* .
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Specify the name of the IAM role to be used when making API calls to the Directory Service.
	//
	// This setting doesn't apply to RDS Custom.
	DomainIamRoleName *string `field:"optional" json:"domainIamRoleName" yaml:"domainIamRoleName"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	//
	// The values in the list depend on the DB engine being used. For more information, see [Publishing Database Logs to Amazon CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the *Amazon Relational Database Service User Guide* .
	//
	// *Amazon Aurora*
	//
	// Not applicable. CloudWatch Logs exports are managed by the DB cluster.
	//
	// *MariaDB*
	//
	// Valid values: `audit` , `error` , `general` , `slowquery`
	//
	// *Microsoft SQL Server*
	//
	// Valid values: `agent` , `error`
	//
	// *MySQL*
	//
	// Valid values: `audit` , `error` , `general` , `slowquery`
	//
	// *Oracle*
	//
	// Valid values: `alert` , `audit` , `listener` , `trace`
	//
	// *PostgreSQL*
	//
	// Valid values: `postgresql` , `upgrade`.
	EnableCloudwatchLogsExports *[]*string `field:"optional" json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	//
	// By default, mapping is disabled.
	//
	// For more information, see [IAM Database Authentication for MySQL and PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html) in the *Amazon RDS User Guide.*
	//
	// *Amazon Aurora*
	//
	// Not applicable. Mapping AWS IAM accounts to database accounts is managed by the DB cluster.
	EnableIamDatabaseAuthentication interface{} `field:"optional" json:"enableIamDatabaseAuthentication" yaml:"enableIamDatabaseAuthentication"`
	// A value that indicates whether to enable Performance Insights for the DB instance.
	//
	// For more information, see [Using Amazon Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the *Amazon RDS User Guide* .
	//
	// This setting doesn't apply to RDS Custom.
	EnablePerformanceInsights interface{} `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// The name of the database engine that you want to use for this DB instance.
	//
	// > When you are creating a DB instance, the `Engine` property is required.
	//
	// Valid Values:
	//
	// - `aurora` (for MySQL 5.6-compatible Aurora)
	// - `aurora-mysql` (for MySQL 5.7-compatible Aurora)
	// - `aurora-postgresql`
	// - `mariadb`
	// - `mysql`
	// - `oracle-ee`
	// - `oracle-se2`
	// - `oracle-se1`
	// - `oracle-se`
	// - `postgres`
	// - `sqlserver-ee`
	// - `sqlserver-se`
	// - `sqlserver-ex`
	// - `sqlserver-web`.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The version number of the database engine to use.
	//
	// For a list of valid engine versions, use the `DescribeDBEngineVersions` action.
	//
	// The following are the database engines and links to information about the major and minor versions that are available with Amazon RDS. Not every database engine is available for every AWS Region.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The version number of the database engine to be used by the DB instance is managed by the DB cluster.
	//
	// *MariaDB*
	//
	// See [MariaDB on Amazon RDS Versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt) in the *Amazon RDS User Guide.*
	//
	// *Microsoft SQL Server*
	//
	// See [Microsoft SQL Server Versions on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport) in the *Amazon RDS User Guide.*
	//
	// *MySQL*
	//
	// See [MySQL on Amazon RDS Versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the *Amazon RDS User Guide.*
	//
	// *Oracle*
	//
	// See [Oracle Database Engine Release Notes](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html) in the *Amazon RDS User Guide.*
	//
	// *PostgreSQL*
	//
	// See [Supported PostgreSQL Database Versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts.General.DBVersions) in the *Amazon RDS User Guide.*
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	//
	// If you specify this property, you must follow the range of allowed ratios of your requested IOPS rate to the amount of storage that you allocate (IOPS to allocated storage). For example, you can provision an Oracle database instance with 1000 IOPS and 200 GiB of storage (a ratio of 5:1), or specify 2000 IOPS with 200 GiB of storage (a ratio of 10:1). For more information, see [Amazon RDS Provisioned IOPS Storage to Improve Performance](https://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/CHAP_Storage.html#USER_PIOPS) in the *Amazon RDS User Guide* .
	//
	// > If you specify `io1` for the `StorageType` property, then you must also specify the `Iops` property.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The ARN of the AWS KMS key that's used to encrypt the DB instance, such as `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef` .
	//
	// If you enable the StorageEncrypted property but don't specify this property, AWS CloudFormation uses the default KMS key. If you specify this property, you must set the StorageEncrypted property to true.
	//
	// If you specify the `SourceDBInstanceIdentifier` property, the value is inherited from the source DB instance if the read replica is created in the same region.
	//
	// If you create an encrypted read replica in a different AWS Region, then you must specify a KMS key for the destination AWS Region. KMS encryption keys are specific to the region that they're created in, and you can't use encryption keys from one region in another region.
	//
	// If you specify the `SnapshotIdentifier` property, the `StorageEncrypted` property value is inherited from the snapshot, and if the DB instance is encrypted, the specified `KmsKeyId` property is used.
	//
	// If you specify `DBSecurityGroups` , AWS CloudFormation ignores this property. To specify both a security group and this property, you must use a VPC security group. For more information about Amazon RDS and VPC, see [Using Amazon RDS with Amazon VPC](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html) in the *Amazon RDS User Guide* .
	//
	// *Amazon Aurora*
	//
	// Not applicable. The KMS key identifier is managed by the DB cluster.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// License model information for this DB instance.
	//
	// Valid values:
	//
	// - Aurora MySQL - `general-public-license`
	// - Aurora PostgreSQL - `postgresql-license`
	// - MariaDB - `general-public-license`
	// - Microsoft SQL Server - `license-included`
	// - MySQL - `general-public-license`
	// - Oracle - `bring-your-own-license` or `license-included`
	// - PostgreSQL - `postgresql-license`
	//
	// > If you've specified `DBSecurityGroups` and then you update the license model, AWS CloudFormation replaces the underlying DB instance. This will incur some interruptions to database availability.
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// The master user name for the DB instance.
	//
	// > If you specify the `SourceDBInstanceIdentifier` or `DBSnapshotIdentifier` property, don't specify this property. The value is inherited from the source DB instance or snapshot.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The name for the master user is managed by the DB cluster.
	//
	// *MariaDB*
	//
	// Constraints:
	//
	// - Required for MariaDB.
	// - Must be 1 to 16 letters or numbers.
	// - Can't be a reserved word for the chosen database engine.
	//
	// *Microsoft SQL Server*
	//
	// Constraints:
	//
	// - Required for SQL Server.
	// - Must be 1 to 128 letters or numbers.
	// - The first character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// *MySQL*
	//
	// Constraints:
	//
	// - Required for MySQL.
	// - Must be 1 to 16 letters or numbers.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// *Oracle*
	//
	// Constraints:
	//
	// - Required for Oracle.
	// - Must be 1 to 30 letters or numbers.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// *PostgreSQL*
	//
	// Constraints:
	//
	// - Required for PostgreSQL.
	// - Must be 1 to 63 letters or numbers.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	MasterUsername *string `field:"optional" json:"masterUsername" yaml:"masterUsername"`
	// The password for the master user. The password can include any printable ASCII character except "/", """, or "@".
	//
	// *Amazon Aurora*
	//
	// Not applicable. The password for the master user is managed by the DB cluster.
	//
	// *MariaDB*
	//
	// Constraints: Must contain from 8 to 41 characters.
	//
	// *Microsoft SQL Server*
	//
	// Constraints: Must contain from 8 to 128 characters.
	//
	// *MySQL*
	//
	// Constraints: Must contain from 8 to 41 characters.
	//
	// *Oracle*
	//
	// Constraints: Must contain from 8 to 30 characters.
	//
	// *PostgreSQL*
	//
	// Constraints: Must contain from 8 to 128 characters.
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
	// The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.
	//
	// For more information about this setting, including limitations that apply to it, see [Managing capacity automatically with Amazon RDS storage autoscaling](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling) in the *Amazon RDS User Guide* .
	//
	// This setting doesn't apply to RDS Custom.
	MaxAllocatedStorage *float64 `field:"optional" json:"maxAllocatedStorage" yaml:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.
	//
	// To disable collection of Enhanced Monitoring metrics, specify 0. The default is 0.
	//
	// If `MonitoringRoleArn` is specified, then you must set `MonitoringInterval` to a value other than 0.
	//
	// This setting doesn't apply to RDS Custom.
	//
	// Valid Values: `0, 1, 5, 10, 15, 30, 60`.
	MonitoringInterval *float64 `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs.
	//
	// For example, `arn:aws:iam:123456789012:role/emaccess` . For information on creating a monitoring role, see [Setting Up and Enabling Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the *Amazon RDS User Guide* .
	//
	// If `MonitoringInterval` is set to a value other than 0, then you must supply a `MonitoringRoleArn` value.
	//
	// This setting doesn't apply to RDS Custom.
	MonitoringRoleArn *string `field:"optional" json:"monitoringRoleArn" yaml:"monitoringRoleArn"`
	// Specifies whether the database instance is a Multi-AZ DB instance deployment.
	//
	// You can't set the `AvailabilityZone` parameter if the `MultiAZ` parameter is set to true.
	//
	// Currently, you can't use AWS CloudFormation to create a Multi-AZ DB cluster deployment.
	//
	// For more information, see [Multi-AZ deployments for high availability](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.MultiAZ.html) in the *Amazon RDS User Guide* .
	//
	// *Amazon Aurora*
	//
	// Not applicable. Amazon Aurora storage is replicated across all of the Availability Zones and doesn't require the `MultiAZ` option to be set.
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// Indicates that the DB instance should be associated with the specified option group.
	//
	// Permanent options, such as the TDE option for Oracle Advanced Security TDE, can't be removed from an option group. Also, that option group can't be removed from a DB instance once it is associated with a DB instance.
	OptionGroupName *string `field:"optional" json:"optionGroupName" yaml:"optionGroupName"`
	// The AWS KMS key identifier for encryption of Performance Insights data.
	//
	// The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
	//
	// If you do not specify a value for `PerformanceInsightsKMSKeyId` , then Amazon RDS uses your default KMS key. There is a default KMS key for your AWS account. Your AWS account has a different default KMS key for each AWS Region.
	//
	// For information about enabling Performance Insights, see [EnablePerformanceInsights](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-enableperformanceinsights) .
	PerformanceInsightsKmsKeyId *string `field:"optional" json:"performanceInsightsKmsKeyId" yaml:"performanceInsightsKmsKeyId"`
	// The amount of time, in days, to retain Performance Insights data. Valid values are 7 or 731 (2 years).
	//
	// For information about enabling Performance Insights, see [EnablePerformanceInsights](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-enableperformanceinsights) .
	PerformanceInsightsRetentionPeriod *float64 `field:"optional" json:"performanceInsightsRetentionPeriod" yaml:"performanceInsightsRetentionPeriod"`
	// The port number on which the database accepts connections.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The port number is managed by the DB cluster.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled, using the `BackupRetentionPeriod` parameter.
	//
	// For more information, see [Backup Window](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow) in the *Amazon RDS User Guide.*
	//
	// Constraints:
	//
	// - Must be in the format `hh24:mi-hh24:mi` .
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The daily time range for creating automated backups is managed by the DB cluster.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week. To see the time blocks available, see [Adjusting the Preferred DB Instance Maintenance Window](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow) in the *Amazon RDS User Guide.*
	//
	// > This property applies when AWS CloudFormation initially creates the DB instance. If you use AWS CloudFormation to update the DB instance, those updates are applied immediately.
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.
	//
	// This setting doesn't apply to RDS Custom.
	ProcessorFeatures interface{} `field:"optional" json:"processorFeatures" yaml:"processorFeatures"`
	// A value that specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance.
	//
	// For more information, see [Fault Tolerance for an Aurora DB Cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.FaultTolerance) in the *Amazon Aurora User Guide* .
	//
	// This setting doesn't apply to RDS Custom.
	//
	// Default: 1
	//
	// Valid Values: 0 - 15.
	PromotionTier *float64 `field:"optional" json:"promotionTier" yaml:"promotionTier"`
	// Indicates whether the DB instance is an internet-facing instance.
	//
	// If you specify `true` , AWS CloudFormation creates an instance with a publicly resolvable DNS name, which resolves to a public IP address. If you specify false, AWS CloudFormation creates an internal instance with a DNS name that resolves to a private IP address.
	//
	// The default behavior value depends on your VPC setup and the database subnet group. For more information, see the `PubliclyAccessible` parameter in [`CreateDBInstance`](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) in the *Amazon RDS API Reference* .
	//
	// If this resource has a public IP address and is also in a VPC that is defined in the same template, you must use the *DependsOn* attribute to declare a dependency on the VPC-gateway attachment. For more information, see [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
	//
	// > If you specify DBSecurityGroups, AWS CloudFormation ignores this property. To specify a security group and this property, you must use a VPC security group. For more information about Amazon RDS and VPC, see [Using Amazon RDS with Amazon VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.html) in the *Amazon RDS User Guide* .
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// If you want to create a read replica DB instance, specify the ID of the source DB instance.
	//
	// Each DB instance can have a limited number of read replicas. For more information, see [Working with Read Replicas](https://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/USER_ReadRepl.html) in the *Amazon RDS User Guide* .
	//
	// For information about constraints that apply to DB instance identifiers, see [Naming constraints in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon RDS User Guide* .
	//
	// The `SourceDBInstanceIdentifier` property determines whether a DB instance is a read replica. If you remove the `SourceDBInstanceIdentifier` property from your template and then update your stack, AWS CloudFormation deletes the Read Replica and creates a new DB instance (not a read replica).
	//
	// > - If you specify a source DB instance that uses VPC security groups, we recommend that you specify the `VPCSecurityGroups` property. If you don't specify the property, the read replica inherits the value of the `VPCSecurityGroups` property from the source DB when you create the replica. However, if you update the stack, AWS CloudFormation reverts the replica's `VPCSecurityGroups` property to the default value because it's not defined in the stack's template. This change might cause unexpected issues.
	// > - Read replicas don't support deletion policies. AWS CloudFormation ignores any deletion policy that's associated with a read replica.
	// > - If you specify `SourceDBInstanceIdentifier` , don't specify the `DBSnapshotIdentifier` property. You can't create a read replica from a snapshot.
	// > - Don't set the `BackupRetentionPeriod` , `DBName` , `MasterUsername` , `MasterUserPassword` , and `PreferredBackupWindow` properties. The database attributes are inherited from the source DB instance, and backups are disabled for read replicas.
	// > - If the source DB instance is in a different region than the read replica, specify the source region in `SourceRegion` , and specify an ARN for a valid DB instance in `SourceDBInstanceIdentifier` . For more information, see [Constructing a Amazon RDS Amazon Resource Name (ARN)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html#USER_Tagging.ARN) in the *Amazon RDS User Guide* .
	// > - For DB instances in Amazon Aurora clusters, don't specify this property. Amazon RDS automatically assigns writer and reader DB instances.
	SourceDbInstanceIdentifier *string `field:"optional" json:"sourceDbInstanceIdentifier" yaml:"sourceDbInstanceIdentifier"`
	// The ID of the region that contains the source DB instance for the read replica.
	SourceRegion *string `field:"optional" json:"sourceRegion" yaml:"sourceRegion"`
	// A value that indicates whether the DB instance is encrypted. By default, it isn't encrypted.
	//
	// If you specify the `KmsKeyId` property, then you must enable encryption.
	//
	// If you specify the `SourceDBInstanceIdentifier` property, don't specify this property. The value is inherited from the source DB instance, and if the DB instance is encrypted, the specified `KmsKeyId` property is used.
	//
	// If you specify the `SnapshotIdentifier` and the specified snapshot is encrypted, don't specify this property. The value is inherited from the snapshot, and the specified `KmsKeyId` property is used.
	//
	// If you specify the `SnapshotIdentifier` and the specified snapshot isn't encrypted, you can use this property to specify that the restored DB instance is encrypted. Specify the `KmsKeyId` property for the KMS key to use for encryption. If you don't want the restored DB instance to be encrypted, then don't set this property or set it to `false` .
	//
	// *Amazon Aurora*
	//
	// Not applicable. The encryption for DB instances is managed by the DB cluster.
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// Specifies the storage type to be associated with the DB instance.
	//
	// Valid values: `standard | gp2 | io1`
	//
	// The `standard` value is also known as magnetic.
	//
	// If you specify `io1` , you must also include a value for the `Iops` parameter.
	//
	// Default: `io1` if the `Iops` parameter is specified, otherwise `standard`
	//
	// For more information, see [Amazon RDS DB Instance Storage](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html) in the *Amazon RDS User Guide* .
	//
	// *Amazon Aurora*
	//
	// Not applicable. Aurora data is stored in the cluster volume, which is a single, virtual volume that uses solid state drives (SSDs).
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Tags to assign to the DB instance.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The time zone of the DB instance.
	//
	// The time zone parameter is currently supported only by [Microsoft SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone) .
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// A value that indicates whether the DB instance class of the DB instance uses its default processor features.
	//
	// This setting doesn't apply to RDS Custom.
	UseDefaultProcessorFeatures interface{} `field:"optional" json:"useDefaultProcessorFeatures" yaml:"useDefaultProcessorFeatures"`
	// A list of the VPC security group IDs to assign to the DB instance.
	//
	// The list can include both the physical IDs of existing VPC security groups and references to [AWS::EC2::SecurityGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html) resources created in the template.
	//
	// If you plan to update the resource, don't specify VPC security groups in a shared VPC.
	//
	// If you set `VPCSecurityGroups` , you must not set [`DBSecurityGroups`](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsecuritygroups) , and vice versa.
	//
	// > You can migrate a DB instance in your stack from an RDS DB security group to a VPC security group, but keep the following in mind:
	// >
	// > - You can't revert to using an RDS security group after you establish a VPC security group membership.
	// > - When you migrate your DB instance to VPC security groups, if your stack update rolls back because the DB instance update fails or because an update fails in another AWS CloudFormation resource, the rollback fails because it can't revert to an RDS security group.
	// > - To use the properties that are available when you use a VPC security group, you must recreate the DB instance. If you don't, AWS CloudFormation submits only the property values that are listed in the [`DBSecurityGroups`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsecuritygroups) property.
	//
	// To avoid this situation, migrate your DB instance to using VPC security groups only when that is the only change in your stack template.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The associated list of EC2 VPC security groups is managed by the DB cluster. If specified, the setting must match the DB cluster setting.
	VpcSecurityGroups *[]*string `field:"optional" json:"vpcSecurityGroups" yaml:"vpcSecurityGroups"`
}

