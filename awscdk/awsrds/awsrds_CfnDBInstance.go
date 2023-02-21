package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::RDS::DBInstance`.
//
// The `AWS::RDS::DBInstance` resource creates an Amazon DB instance. The new DB instance can be an RDS DB instance, or it can be a DB instance in an Aurora DB cluster.
//
// For more information about creating an RDS DB instance, see [Creating an Amazon RDS DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateDBInstance.html) in the *Amazon RDS User Guide* .
//
// For more information about creating a DB instance in an Aurora DB cluster, see [Creating an Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.CreateInstance.html) in the *Amazon Aurora User Guide* .
//
// If you import an existing DB instance, and the template configuration doesn't match the actual configuration of the DB instance, AWS CloudFormation applies the changes in the template during the import operation.
//
// > If a DB instance is deleted or replaced during an update, AWS CloudFormation deletes all automated snapshots. However, it retains manual DB snapshots. During an update that requires replacement, you can apply a stack policy to prevent DB instances from being replaced. For more information, see [Prevent Updates to Stack Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html) .
//
// *Updating DB instances*
//
// When properties labeled " *Update requires:* [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) " are updated, AWS CloudFormation first creates a replacement DB instance, then changes references from other dependent resources to point to the replacement DB instance, and finally deletes the old DB instance.
//
// > We highly recommend that you take a snapshot of the database before updating the stack. If you don't, you lose the data when AWS CloudFormation replaces your DB instance. To preserve your data, perform the following procedure:
// >
// > - Deactivate any applications that are using the DB instance so that there's no activity on the DB instance.
// > - Create a snapshot of the DB instance. For more information, see [Creating a DB Snapshot](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateSnapshot.html) .
// > - If you want to restore your instance using a DB snapshot, modify the updated template with your DB instance changes and add the `DBSnapshotIdentifier` property with the ID of the DB snapshot that you want to use.
// >
// > After you restore a DB instance with a `DBSnapshotIdentifier` property, you must specify the same `DBSnapshotIdentifier` property for any future updates to the DB instance. When you specify this property for an update, the DB instance is not restored from the DB snapshot again, and the data in the database is not changed. However, if you don't specify the `DBSnapshotIdentifier` property, an empty DB instance is created, and the original DB instance is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB instance is restored from the specified `DBSnapshotIdentifier` property, and the original DB instance is deleted.
// > - Update the stack.
//
// For more information about updating other properties of this resource, see `[ModifyDBInstance](https://docs.aws.amazon.com//AmazonRDS/latest/APIReference/API_ModifyDBInstance.html)` . For more information about updating stacks, see [AWS CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html) .
//
// *Deleting DB instances*
//
// For DB instances that are part of an Aurora DB cluster, you can set a deletion policy for your DB instance to control how AWS CloudFormation handles the DB instance when the stack is deleted. For Amazon RDS DB instances, you can choose to *retain* the DB instance, to *delete* the DB instance, or to *create a snapshot* of the DB instance. The default AWS CloudFormation behavior depends on the `DBClusterIdentifier` property:
//
// - For `AWS::RDS::DBInstance` resources that don't specify the `DBClusterIdentifier` property, AWS CloudFormation saves a snapshot of the DB instance.
// - For `AWS::RDS::DBInstance` resources that do specify the `DBClusterIdentifier` property, AWS CloudFormation deletes the DB instance.
//
// For more information, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBInstance := awscdk.Aws_rds.NewCfnDBInstance(this, jsii.String("MyCfnDBInstance"), &cfnDBInstanceProps{
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
//   	certificateDetails: &certificateDetailsProperty{
//   		caIdentifier: jsii.String("caIdentifier"),
//   		validTill: jsii.String("validTill"),
//   	},
//   	certificateRotationRestart: jsii.Boolean(false),
//   	characterSetName: jsii.String("characterSetName"),
//   	copyTagsToSnapshot: jsii.Boolean(false),
//   	customIamInstanceProfile: jsii.String("customIamInstanceProfile"),
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	dbClusterSnapshotIdentifier: jsii.String("dbClusterSnapshotIdentifier"),
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
//   	endpoint: &endpointProperty{
//   		address: jsii.String("address"),
//   		hostedZoneId: jsii.String("hostedZoneId"),
//   		port: jsii.String("port"),
//   	},
//   	engine: jsii.String("engine"),
//   	engineVersion: jsii.String("engineVersion"),
//   	iops: jsii.Number(123),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	licenseModel: jsii.String("licenseModel"),
//   	manageMasterUserPassword: jsii.Boolean(false),
//   	masterUsername: jsii.String("masterUsername"),
//   	masterUserPassword: jsii.String("masterUserPassword"),
//   	masterUserSecret: &masterUserSecretProperty{
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   		secretArn: jsii.String("secretArn"),
//   	},
//   	maxAllocatedStorage: jsii.Number(123),
//   	monitoringInterval: jsii.Number(123),
//   	monitoringRoleArn: jsii.String("monitoringRoleArn"),
//   	multiAz: jsii.Boolean(false),
//   	ncharCharacterSetName: jsii.String("ncharCharacterSetName"),
//   	networkType: jsii.String("networkType"),
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
//   	replicaMode: jsii.String("replicaMode"),
//   	restoreTime: jsii.String("restoreTime"),
//   	sourceDbInstanceAutomatedBackupsArn: jsii.String("sourceDbInstanceAutomatedBackupsArn"),
//   	sourceDbInstanceIdentifier: jsii.String("sourceDbInstanceIdentifier"),
//   	sourceDbiResourceId: jsii.String("sourceDbiResourceId"),
//   	sourceRegion: jsii.String("sourceRegion"),
//   	storageEncrypted: jsii.Boolean(false),
//   	storageThroughput: jsii.Number(123),
//   	storageType: jsii.String("storageType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timezone: jsii.String("timezone"),
//   	useDefaultProcessorFeatures: jsii.Boolean(false),
//   	useLatestRestorableTime: jsii.Boolean(false),
//   	vpcSecurityGroups: []*string{
//   		jsii.String("vpcSecurityGroups"),
//   	},
//   })
//
type CfnDBInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The amount of storage in gibibytes (GiB) to be initially allocated for the database instance.
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
	AllocatedStorage() *string
	SetAllocatedStorage(val *string)
	// A value that indicates whether major version upgrades are allowed.
	//
	// Changing this parameter doesn't result in an outage and the change is asynchronously applied as soon as possible.
	//
	// Constraints: Major version upgrades must be allowed when specifying a value for the `EngineVersion` parameter that is a different major version than the DB instance's current version.
	AllowMajorVersionUpgrade() interface{}
	SetAllowMajorVersionUpgrade(val interface{})
	// The AWS Identity and Access Management (IAM) roles associated with the DB instance.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The associated roles are managed by the DB cluster.
	AssociatedRoles() interface{}
	SetAssociatedRoles(val interface{})
	AttrCertificateDetailsCaIdentifier() *string
	AttrCertificateDetailsValidTill() *string
	// The Amazon Resource Name (ARN) for the DB instance.
	AttrDbInstanceArn() *string
	// The AWS Region-unique, immutable identifier for the DB instance.
	//
	// This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB instance is accessed.
	AttrDbiResourceId() *string
	// The Oracle system ID (Oracle SID) for a container database (CDB).
	//
	// The Oracle SID is also the name of the CDB.
	//
	// This setting is valid for RDS Custom only.
	AttrDbSystemId() *string
	// The connection endpoint for the database. For example: `mystack-mydb-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`.
	//
	// For Aurora Serverless DB clusters, the connection endpoint only applies to the DB cluster.
	AttrEndpointAddress() *string
	// The ID that Amazon Route 53 assigns when you create a hosted zone.
	AttrEndpointHostedZoneId() *string
	// The port number on which the database accepts connections.
	//
	// For example: `3306`.
	AttrEndpointPort() *string
	// The Amazon Resource Name (ARN) of the secret.
	AttrMasterUserSecretSecretArn() *string
	// A value that indicates whether minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	//
	// By default, minor engine upgrades are applied automatically.
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
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
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
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
	BackupRetentionPeriod() *float64
	SetBackupRetentionPeriod(val *float64)
	// The identifier of the CA certificate for this DB instance.
	//
	// > Specifying or updating this property triggers a reboot.
	//
	// For more information about CA certificate identifiers for RDS DB engines, see [Rotating Your SSL/TLS Certificate](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL-certificate-rotation.html) in the *Amazon RDS User Guide* .
	//
	// For more information about CA certificate identifiers for Aurora DB engines, see [Rotating Your SSL/TLS Certificate](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL-certificate-rotation.html) in the *Amazon Aurora User Guide* .
	CaCertificateIdentifier() *string
	SetCaCertificateIdentifier(val *string)
	// The details of the DB instance's server certificate.
	CertificateDetails() interface{}
	SetCertificateDetails(val interface{})
	// A value that indicates whether the DB instance is restarted when you rotate your SSL/TLS certificate.
	//
	// By default, the DB instance is restarted when you rotate your SSL/TLS certificate. The certificate is not updated until the DB instance is restarted.
	//
	// > Set this parameter only if you are *not* using SSL/TLS to connect to the DB instance.
	//
	// If you are using SSL/TLS to connect to the DB instance, follow the appropriate instructions for your DB engine to rotate your SSL/TLS certificate:
	//
	// - For more information about rotating your SSL/TLS certificate for RDS DB engines, see [Rotating Your SSL/TLS Certificate.](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL-certificate-rotation.html) in the *Amazon RDS User Guide.*
	// - For more information about rotating your SSL/TLS certificate for Aurora DB engines, see [Rotating Your SSL/TLS Certificate](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL-certificate-rotation.html) in the *Amazon Aurora User Guide* .
	//
	// This setting doesn't apply to RDS Custom.
	CertificateRotationRestart() interface{}
	SetCertificateRotationRestart(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// For supported engines, indicates that the DB instance should be associated with the specified character set.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The character set is managed by the DB cluster. For more information, see [AWS::RDS::DBCluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html) .
	CharacterSetName() *string
	SetCharacterSetName(val *string)
	// A value that indicates whether to copy tags from the DB instance to snapshots of the DB instance.
	//
	// By default, tags are not copied.
	//
	// *Amazon Aurora*
	//
	// Not applicable. Copying tags to snapshots is managed by the DB cluster. Setting this value for an Aurora DB instance has no effect on the DB cluster setting.
	CopyTagsToSnapshot() interface{}
	SetCopyTagsToSnapshot(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance.
	//
	// The instance profile must meet the following requirements:
	//
	// - The profile must exist in your account.
	// - The profile must have an IAM role that Amazon EC2 has permissions to assume.
	// - The instance profile name and the associated IAM role name must start with the prefix `AWSRDSCustom` .
	//
	// For the list of permissions required for the IAM role, see [Configure IAM and your VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc) in the *Amazon RDS User Guide* .
	//
	// This setting is required for RDS Custom.
	CustomIamInstanceProfile() *string
	SetCustomIamInstanceProfile(val *string)
	// The identifier of the DB cluster that the instance will belong to.
	DbClusterIdentifier() *string
	SetDbClusterIdentifier(val *string)
	// The identifier for the RDS for MySQL Multi-AZ DB cluster snapshot to restore from.
	//
	// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html) in the *Amazon RDS User Guide* .
	//
	// Constraints:
	//
	// - Must match the identifier of an existing Multi-AZ DB cluster snapshot.
	// - Can't be specified when `DBSnapshotIdentifier` is specified.
	// - Must be specified when `DBSnapshotIdentifier` isn't specified.
	// - If you are restoring from a shared manual Multi-AZ DB cluster snapshot, the `DBClusterSnapshotIdentifier` must be the ARN of the shared snapshot.
	// - Can't be the identifier of an Aurora DB cluster snapshot.
	// - Can't be the identifier of an RDS for PostgreSQL Multi-AZ DB cluster snapshot.
	DbClusterSnapshotIdentifier() *string
	SetDbClusterSnapshotIdentifier(val *string)
	// The compute and memory capacity of the DB instance, for example, `db.m4.large` . Not all DB instance classes are available in all AWS Regions, or for all database engines.
	//
	// For the full list of DB instance classes, and availability for your engine, see [DB Instance Class](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the *Amazon RDS User Guide.* For more information about DB instance class pricing and AWS Region support for DB instance classes, see [Amazon RDS Pricing](https://docs.aws.amazon.com/rds/pricing/) .
	DbInstanceClass() *string
	SetDbInstanceClass(val *string)
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation converts it to lowercase. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the DB instance. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// For information about constraints that apply to DB instance identifiers, see [Naming constraints in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon RDS User Guide* .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DbInstanceIdentifier() *string
	SetDbInstanceIdentifier(val *string)
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
	// - Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).
	// - Must contain 1 to 63 characters.
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
	DbName() *string
	SetDbName(val *string)
	// The name of an existing DB parameter group or a reference to an [AWS::RDS::DBParameterGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html) resource created in the template.
	//
	// To list all of the available DB parameter group names, use the following command:
	//
	// `aws rds describe-db-parameter-groups --query "DBParameterGroups[].DBParameterGroupName" --output text`
	//
	// > If any of the data members of the referenced parameter group are changed during an update, the DB instance might need to be restarted, which causes some interruption. If the parameter group contains static parameters, whether they were changed or not, an update triggers a reboot.
	//
	// If you don't specify a value for `DBParameterGroupName` property, the default DB parameter group for the specified engine and engine version is used.
	DbParameterGroupName() *string
	SetDbParameterGroupName(val *string)
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
	DbSecurityGroups() *[]*string
	SetDbSecurityGroups(val *[]*string)
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
	DbSnapshotIdentifier() *string
	SetDbSnapshotIdentifier(val *string)
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
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	// A value that indicates whether to remove automated backups immediately after the DB instance is deleted.
	//
	// This parameter isn't case-sensitive. The default is to remove automated backups immediately after the DB instance is deleted.
	//
	// *Amazon Aurora*
	//
	// Not applicable. When you delete a DB cluster, all automated backups for that DB cluster are deleted and can't be recovered. Manual DB cluster snapshots of the DB cluster are not deleted.
	DeleteAutomatedBackups() interface{}
	SetDeleteAutomatedBackups(val interface{})
	// A value that indicates whether the DB instance has deletion protection enabled.
	//
	// The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled. For more information, see [Deleting a DB Instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html) .
	//
	// *Amazon Aurora*
	//
	// Not applicable. You can enable or disable deletion protection for the DB cluster. For more information, see `CreateDBCluster` . DB instances in a DB cluster can be deleted even when deletion protection is enabled for the DB cluster.
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	// The Active Directory directory ID to create the DB instance in.
	//
	// Currently, only Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active Directory Domain.
	//
	// For more information, see [Kerberos Authentication](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html) in the *Amazon RDS User Guide* .
	Domain() *string
	SetDomain(val *string)
	// Specify the name of the IAM role to be used when making API calls to the Directory Service.
	//
	// This setting doesn't apply to RDS Custom.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The domain is managed by the DB cluster.
	DomainIamRoleName() *string
	SetDomainIamRoleName(val *string)
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
	EnableCloudwatchLogsExports() *[]*string
	SetEnableCloudwatchLogsExports(val *[]*string)
	// A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	//
	// By default, mapping is disabled.
	//
	// This property is supported for RDS for MariaDB, RDS for MySQL, and RDS for PostgreSQL. For more information, see [IAM Database Authentication for MariaDB, MySQL, and PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html) in the *Amazon RDS User Guide.*
	//
	// *Amazon Aurora*
	//
	// Not applicable. Mapping AWS IAM accounts to database accounts is managed by the DB cluster.
	EnableIamDatabaseAuthentication() interface{}
	SetEnableIamDatabaseAuthentication(val interface{})
	// A value that indicates whether to enable Performance Insights for the DB instance.
	//
	// For more information, see [Using Amazon Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the *Amazon RDS User Guide* .
	//
	// This setting doesn't apply to RDS Custom.
	EnablePerformanceInsights() interface{}
	SetEnablePerformanceInsights(val interface{})
	// Specifies the connection endpoint.
	//
	// > The endpoint might not be shown for instances whose status is `creating` .
	Endpoint() interface{}
	SetEndpoint(val interface{})
	// The name of the database engine that you want to use for this DB instance.
	//
	// > When you are creating a DB instance, the `Engine` property is required.
	//
	// Valid Values:
	//
	// - `aurora` (for MySQL 5.6-compatible Aurora)
	// - `aurora-mysql` (for MySQL 5.7-compatible Aurora)
	// - `aurora-postgresql`
	// - `custom-oracle-ee`
	// - `custom-oracle-ee-cdb`
	// - `custom-sqlserver-ee`
	// - `custom-sqlserver-se`
	// - `custom-sqlserver-web`
	// - `mariadb`
	// - `mysql`
	// - `oracle-ee`
	// - `oracle-ee-cdb`
	// - `oracle-se2`
	// - `oracle-se2-cdb`
	// - `postgres`
	// - `sqlserver-ee`
	// - `sqlserver-se`
	// - `sqlserver-ex`
	// - `sqlserver-web`.
	Engine() *string
	SetEngine(val *string)
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
	EngineVersion() *string
	SetEngineVersion(val *string)
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	//
	// If you specify this property, you must follow the range of allowed ratios of your requested IOPS rate to the amount of storage that you allocate (IOPS to allocated storage). For example, you can provision an Oracle database instance with 1000 IOPS and 200 GiB of storage (a ratio of 5:1), or specify 2000 IOPS with 200 GiB of storage (a ratio of 10:1). For more information, see [Amazon RDS Provisioned IOPS Storage to Improve Performance](https://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/CHAP_Storage.html#USER_PIOPS) in the *Amazon RDS User Guide* .
	//
	// > If you specify `io1` for the `StorageType` property, then you must also specify the `Iops` property.
	Iops() *float64
	SetIops(val *float64)
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
	KmsKeyId() *string
	SetKmsKeyId(val *string)
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
	LicenseModel() *string
	SetLicenseModel(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// A value that indicates whether to manage the master user password with AWS Secrets Manager.
	//
	// For more information, see [Password management with AWS Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide.*
	//
	// Constraints:
	//
	// - Can't manage the master user password with AWS Secrets Manager if `MasterUserPassword` is specified.
	ManageMasterUserPassword() interface{}
	SetManageMasterUserPassword(val interface{})
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
	MasterUsername() *string
	SetMasterUsername(val *string)
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
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	// Contains the secret managed by RDS in AWS Secrets Manager for the master user password.
	//
	// For more information, see [Password management with AWS Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide.*
	MasterUserSecret() interface{}
	SetMasterUserSecret(val interface{})
	// The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.
	//
	// For more information about this setting, including limitations that apply to it, see [Managing capacity automatically with Amazon RDS storage autoscaling](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling) in the *Amazon RDS User Guide* .
	//
	// This setting doesn't apply to RDS Custom.
	//
	// *Amazon Aurora*
	//
	// Not applicable. Storage is managed by the DB cluster.
	MaxAllocatedStorage() *float64
	SetMaxAllocatedStorage(val *float64)
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.
	//
	// To disable collection of Enhanced Monitoring metrics, specify 0. The default is 0.
	//
	// If `MonitoringRoleArn` is specified, then you must set `MonitoringInterval` to a value other than 0.
	//
	// This setting doesn't apply to RDS Custom.
	//
	// Valid Values: `0, 1, 5, 10, 15, 30, 60`.
	MonitoringInterval() *float64
	SetMonitoringInterval(val *float64)
	// The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs.
	//
	// For example, `arn:aws:iam:123456789012:role/emaccess` . For information on creating a monitoring role, see [Setting Up and Enabling Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the *Amazon RDS User Guide* .
	//
	// If `MonitoringInterval` is set to a value other than 0, then you must supply a `MonitoringRoleArn` value.
	//
	// This setting doesn't apply to RDS Custom.
	MonitoringRoleArn() *string
	SetMonitoringRoleArn(val *string)
	// Specifies whether the database instance is a Multi-AZ DB instance deployment.
	//
	// You can't set the `AvailabilityZone` parameter if the `MultiAZ` parameter is set to true.
	//
	// For more information, see [Multi-AZ deployments for high availability](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.MultiAZ.html) in the *Amazon RDS User Guide* .
	//
	// *Amazon Aurora*
	//
	// Not applicable. Amazon Aurora storage is replicated across all of the Availability Zones and doesn't require the `MultiAZ` option to be set.
	MultiAz() interface{}
	SetMultiAz(val interface{})
	// The name of the NCHAR character set for the Oracle DB instance.
	//
	// This parameter doesn't apply to RDS Custom.
	NcharCharacterSetName() *string
	SetNcharCharacterSetName(val *string)
	// The network type of the DB instance.
	//
	// Valid values:
	//
	// - `IPV4`
	// - `DUAL`
	//
	// The network type is determined by the `DBSubnetGroup` specified for the DB instance. A `DBSubnetGroup` can support only the IPv4 protocol or the IPv4 and IPv6 protocols ( `DUAL` ).
	//
	// For more information, see [Working with a DB instance in a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the *Amazon RDS User Guide.*
	NetworkType() *string
	SetNetworkType(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Indicates that the DB instance should be associated with the specified option group.
	//
	// Permanent options, such as the TDE option for Oracle Advanced Security TDE, can't be removed from an option group. Also, that option group can't be removed from a DB instance once it is associated with a DB instance.
	OptionGroupName() *string
	SetOptionGroupName(val *string)
	// The AWS KMS key identifier for encryption of Performance Insights data.
	//
	// The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
	//
	// If you do not specify a value for `PerformanceInsightsKMSKeyId` , then Amazon RDS uses your default KMS key. There is a default KMS key for your AWS account. Your AWS account has a different default KMS key for each AWS Region.
	//
	// For information about enabling Performance Insights, see [EnablePerformanceInsights](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-enableperformanceinsights) .
	PerformanceInsightsKmsKeyId() *string
	SetPerformanceInsightsKmsKeyId(val *string)
	// The amount of time, in days, to retain Performance Insights data. Valid values are 7 or 731 (2 years).
	//
	// For information about enabling Performance Insights, see [EnablePerformanceInsights](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-enableperformanceinsights) .
	PerformanceInsightsRetentionPeriod() *float64
	SetPerformanceInsightsRetentionPeriod(val *float64)
	// The port number on which the database accepts connections.
	//
	// *Amazon Aurora*
	//
	// Not applicable. The port number is managed by the DB cluster.
	Port() *string
	SetPort(val *string)
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
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week. To see the time blocks available, see [Adjusting the Preferred DB Instance Maintenance Window](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow) in the *Amazon RDS User Guide.*
	//
	// > This property applies when AWS CloudFormation initially creates the DB instance. If you use AWS CloudFormation to update the DB instance, those updates are applied immediately.
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	// The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.
	//
	// This setting doesn't apply to RDS Custom.
	//
	// *Amazon Aurora*
	//
	// Not applicable.
	ProcessorFeatures() interface{}
	SetProcessorFeatures(val interface{})
	// A value that specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance.
	//
	// For more information, see [Fault Tolerance for an Aurora DB Cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.FaultTolerance) in the *Amazon Aurora User Guide* .
	//
	// This setting doesn't apply to RDS Custom.
	//
	// Default: 1
	//
	// Valid Values: 0 - 15.
	PromotionTier() *float64
	SetPromotionTier(val *float64)
	// Indicates whether the DB instance is an internet-facing instance.
	//
	// If you specify true, AWS CloudFormation creates an instance with a publicly resolvable DNS name, which resolves to a public IP address. If you specify false, AWS CloudFormation creates an internal instance with a DNS name that resolves to a private IP address.
	//
	// The default behavior value depends on your VPC setup and the database subnet group. For more information, see the `PubliclyAccessible` parameter in the [CreateDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) in the *Amazon RDS API Reference* .
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The open mode of an Oracle read replica.
	//
	// For more information, see [Working with Oracle Read Replicas for Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html) in the *Amazon RDS User Guide* .
	//
	// This setting is only supported in RDS for Oracle.
	//
	// Default: `open-read-only`
	//
	// Valid Values: `open-read-only` or `mounted`.
	ReplicaMode() *string
	SetReplicaMode(val *string)
	// The date and time to restore from.
	//
	// Valid Values: Value must be a time in Universal Coordinated Time (UTC) format
	//
	// Constraints:
	//
	// - Must be before the latest restorable time for the DB instance
	// - Can't be specified if the `UseLatestRestorableTime` parameter is enabled
	//
	// Example: `2009-09-07T23:45:00Z`.
	RestoreTime() *string
	SetRestoreTime(val *string)
	// The Amazon Resource Name (ARN) of the replicated automated backups from which to restore, for example, `arn:aws:rds:useast-1:123456789012:auto-backup:ab-L2IJCEXJP7XQ7HOJ4SIEXAMPLE` .
	//
	// This setting doesn't apply to RDS Custom.
	SourceDbInstanceAutomatedBackupsArn() *string
	SetSourceDbInstanceAutomatedBackupsArn(val *string)
	// If you want to create a read replica DB instance, specify the ID of the source DB instance.
	//
	// Each DB instance can have a limited number of read replicas. For more information, see [Working with Read Replicas](https://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/USER_ReadRepl.html) in the *Amazon RDS User Guide* .
	//
	// For information about constraints that apply to DB instance identifiers, see [Naming constraints in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon RDS User Guide* .
	//
	// The `SourceDBInstanceIdentifier` property determines whether a DB instance is a read replica. If you remove the `SourceDBInstanceIdentifier` property from your template and then update your stack, AWS CloudFormation promotes the Read Replica to a standalone DB instance.
	//
	// > - If you specify a source DB instance that uses VPC security groups, we recommend that you specify the `VPCSecurityGroups` property. If you don't specify the property, the read replica inherits the value of the `VPCSecurityGroups` property from the source DB when you create the replica. However, if you update the stack, AWS CloudFormation reverts the replica's `VPCSecurityGroups` property to the default value because it's not defined in the stack's template. This change might cause unexpected issues.
	// > - Read replicas don't support deletion policies. AWS CloudFormation ignores any deletion policy that's associated with a read replica.
	// > - If you specify `SourceDBInstanceIdentifier` , don't specify the `DBSnapshotIdentifier` property. You can't create a read replica from a snapshot.
	// > - Don't set the `BackupRetentionPeriod` , `DBName` , `MasterUsername` , `MasterUserPassword` , and `PreferredBackupWindow` properties. The database attributes are inherited from the source DB instance, and backups are disabled for read replicas.
	// > - If the source DB instance is in a different region than the read replica, specify the source region in `SourceRegion` , and specify an ARN for a valid DB instance in `SourceDBInstanceIdentifier` . For more information, see [Constructing a Amazon RDS Amazon Resource Name (ARN)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html#USER_Tagging.ARN) in the *Amazon RDS User Guide* .
	// > - For DB instances in Amazon Aurora clusters, don't specify this property. Amazon RDS automatically assigns writer and reader DB instances.
	SourceDbInstanceIdentifier() *string
	SetSourceDbInstanceIdentifier(val *string)
	// The resource ID of the source DB instance from which to restore.
	SourceDbiResourceId() *string
	SetSourceDbiResourceId(val *string)
	// The ID of the region that contains the source DB instance for the read replica.
	SourceRegion() *string
	SetSourceRegion(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
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
	StorageEncrypted() interface{}
	SetStorageEncrypted(val interface{})
	// Specifies the storage throughput value for the DB instance. This setting applies only to the `gp3` storage type.
	//
	// This setting doesn't apply to RDS Custom or Amazon Aurora.
	StorageThroughput() *float64
	SetStorageThroughput(val *float64)
	// Specifies the storage type to be associated with the DB instance.
	//
	// Valid values: `gp2 | gp3 | io1 | standard`
	//
	// The `standard` value is also known as magnetic.
	//
	// If you specify `io1` or `gp3` , you must also include a value for the `Iops` parameter.
	//
	// Default: `io1` if the `Iops` parameter is specified, otherwise `gp2`
	//
	// For more information, see [Amazon RDS DB Instance Storage](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html) in the *Amazon RDS User Guide* .
	//
	// *Amazon Aurora*
	//
	// Not applicable. Aurora data is stored in the cluster volume, which is a single, virtual volume that uses solid state drives (SSDs).
	StorageType() *string
	SetStorageType(val *string)
	// An optional array of key-value pairs to apply to this DB instance.
	Tags() awscdk.TagManager
	// The time zone of the DB instance.
	//
	// The time zone parameter is currently supported only by [Microsoft SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone) .
	Timezone() *string
	SetTimezone(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// A value that indicates whether the DB instance class of the DB instance uses its default processor features.
	//
	// This setting doesn't apply to RDS Custom.
	UseDefaultProcessorFeatures() interface{}
	SetUseDefaultProcessorFeatures(val interface{})
	// A value that indicates whether the DB instance is restored from the latest backup time.
	//
	// By default, the DB instance isn't restored from the latest backup time.
	//
	// Constraints: Can't be specified if the `RestoreTime` parameter is provided.
	UseLatestRestorableTime() interface{}
	SetUseLatestRestorableTime(val interface{})
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
	VpcSecurityGroups() *[]*string
	SetVpcSecurityGroups(val *[]*string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBInstance
type jsiiProxy_CfnDBInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBInstance) AllocatedStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AllowMajorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AssociatedRoles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AttrCertificateDetailsCaIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCertificateDetailsCaIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AttrCertificateDetailsValidTill() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCertificateDetailsValidTill",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AttrDbInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDbInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AttrDbiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDbiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AttrDbSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDbSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AttrEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AttrEndpointHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AttrEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AttrMasterUserSecretSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMasterUserSecretSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CaCertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CertificateDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CertificateRotationRestart() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateRotationRestart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CustomIamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customIamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbClusterSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DeleteAutomatedBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DomainIamRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) EnableCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) EnableIamDatabaseAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIamDatabaseAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) EnablePerformanceInsights() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceInsights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Endpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) ManageMasterUserPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) MasterUserSecret() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"masterUserSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) MaxAllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) MonitoringInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) MonitoringRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) NcharCharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharCharacterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) PerformanceInsightsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) PerformanceInsightsRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) ProcessorFeatures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processorFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) PromotionTier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"promotionTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) ReplicaMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) RestoreTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) SourceDbInstanceAutomatedBackupsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceAutomatedBackupsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) SourceDbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) SourceDbiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) SourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) StorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) UseDefaultProcessorFeatures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultProcessorFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) UseLatestRestorableTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) VpcSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroups",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::DBInstance`.
func NewCfnDBInstance(scope awscdk.Construct, id *string, props *CfnDBInstanceProps) CfnDBInstance {
	_init_.Initialize()

	if err := validateNewCfnDBInstanceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDBInstance{}

	_jsii_.Create(
		"monocdk.aws_rds.CfnDBInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::DBInstance`.
func NewCfnDBInstance_Override(c CfnDBInstance, scope awscdk.Construct, id *string, props *CfnDBInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_rds.CfnDBInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetAllocatedStorage(val *string) {
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetAllowMajorVersionUpgrade(val interface{}) {
	if err := j.validateSetAllowMajorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMajorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetAssociatedRoles(val interface{}) {
	if err := j.validateSetAssociatedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatedRoles",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetBackupRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetCaCertificateIdentifier(val *string) {
	_jsii_.Set(
		j,
		"caCertificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetCertificateDetails(val interface{}) {
	if err := j.validateSetCertificateDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateDetails",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetCertificateRotationRestart(val interface{}) {
	if err := j.validateSetCertificateRotationRestartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateRotationRestart",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetCharacterSetName(val *string) {
	_jsii_.Set(
		j,
		"characterSetName",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetCopyTagsToSnapshot(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetCustomIamInstanceProfile(val *string) {
	_jsii_.Set(
		j,
		"customIamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDbClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDbClusterSnapshotIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbClusterSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDbInstanceClass(val *string) {
	_jsii_.Set(
		j,
		"dbInstanceClass",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDbInstanceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbInstanceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDbName(val *string) {
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDbParameterGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDbSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"dbSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDbSnapshotIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDbSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDeleteAutomatedBackups(val interface{}) {
	if err := j.validateSetDeleteAutomatedBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAutomatedBackups",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDomainIamRoleName(val *string) {
	_jsii_.Set(
		j,
		"domainIamRoleName",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetEnableCloudwatchLogsExports(val *[]*string) {
	_jsii_.Set(
		j,
		"enableCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetEnableIamDatabaseAuthentication(val interface{}) {
	if err := j.validateSetEnableIamDatabaseAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIamDatabaseAuthentication",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetEnablePerformanceInsights(val interface{}) {
	if err := j.validateSetEnablePerformanceInsightsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePerformanceInsights",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetEndpoint(val interface{}) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetEngine(val *string) {
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetIops(val *float64) {
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetLicenseModel(val *string) {
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetManageMasterUserPassword(val interface{}) {
	if err := j.validateSetManageMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageMasterUserPassword",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetMasterUsername(val *string) {
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetMasterUserPassword(val *string) {
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetMasterUserSecret(val interface{}) {
	if err := j.validateSetMasterUserSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserSecret",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetMaxAllocatedStorage(val *float64) {
	_jsii_.Set(
		j,
		"maxAllocatedStorage",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetMonitoringInterval(val *float64) {
	_jsii_.Set(
		j,
		"monitoringInterval",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetMonitoringRoleArn(val *string) {
	_jsii_.Set(
		j,
		"monitoringRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetMultiAz(val interface{}) {
	if err := j.validateSetMultiAzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetNcharCharacterSetName(val *string) {
	_jsii_.Set(
		j,
		"ncharCharacterSetName",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetNetworkType(val *string) {
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetOptionGroupName(val *string) {
	_jsii_.Set(
		j,
		"optionGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetPerformanceInsightsKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"performanceInsightsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetPerformanceInsightsRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"performanceInsightsRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetPort(val *string) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetPreferredBackupWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetProcessorFeatures(val interface{}) {
	if err := j.validateSetProcessorFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorFeatures",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetPromotionTier(val *float64) {
	_jsii_.Set(
		j,
		"promotionTier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetReplicaMode(val *string) {
	_jsii_.Set(
		j,
		"replicaMode",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetRestoreTime(val *string) {
	_jsii_.Set(
		j,
		"restoreTime",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetSourceDbInstanceAutomatedBackupsArn(val *string) {
	_jsii_.Set(
		j,
		"sourceDbInstanceAutomatedBackupsArn",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetSourceDbInstanceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"sourceDbInstanceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetSourceDbiResourceId(val *string) {
	_jsii_.Set(
		j,
		"sourceDbiResourceId",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetSourceRegion(val *string) {
	_jsii_.Set(
		j,
		"sourceRegion",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetStorageEncrypted(val interface{}) {
	if err := j.validateSetStorageEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetStorageThroughput(val *float64) {
	_jsii_.Set(
		j,
		"storageThroughput",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetStorageType(val *string) {
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetTimezone(val *string) {
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetUseDefaultProcessorFeatures(val interface{}) {
	if err := j.validateSetUseDefaultProcessorFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDefaultProcessorFeatures",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetUseLatestRestorableTime(val interface{}) {
	if err := j.validateSetUseLatestRestorableTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLatestRestorableTime",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetVpcSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcSecurityGroups",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnDBInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBInstance_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.CfnDBInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDBInstance_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnDBInstance_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.CfnDBInstance",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDBInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.CfnDBInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBInstance_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_rds.CfnDBInstance",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBInstance) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDBInstance) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDBInstance) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDBInstance) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDBInstance) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDBInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDBInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDBInstance) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDBInstance) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDBInstance) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDBInstance) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBInstance) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDBInstance) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDBInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

