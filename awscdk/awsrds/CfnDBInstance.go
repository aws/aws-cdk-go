package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::RDS::DBInstance` resource creates an Amazon DB instance.
//
// The new DB instance can be an RDS DB instance, or it can be a DB instance in an Aurora DB cluster.
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
// > After you restore a DB instance with a `DBSnapshotIdentifier` property, you can delete the `DBSnapshotIdentifier` property. When you specify this property for an update, the DB instance is not restored from the DB snapshot again, and the data in the database is not changed. However, if you don't specify the `DBSnapshotIdentifier` property, an empty DB instance is created, and the original DB instance is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB instance is restored from the specified `DBSnapshotIdentifier` property, and the original DB instance is deleted.
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
//   cfnDBInstance := awscdk.Aws_rds.NewCfnDBInstance(this, jsii.String("MyCfnDBInstance"), &CfnDBInstanceProps{
//   	AllocatedStorage: jsii.String("allocatedStorage"),
//   	AllowMajorVersionUpgrade: jsii.Boolean(false),
//   	AssociatedRoles: []interface{}{
//   		&DBInstanceRoleProperty{
//   			FeatureName: jsii.String("featureName"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	AutomaticBackupReplicationRegion: jsii.String("automaticBackupReplicationRegion"),
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	BackupRetentionPeriod: jsii.Number(123),
//   	CaCertificateIdentifier: jsii.String("caCertificateIdentifier"),
//   	CertificateDetails: &CertificateDetailsProperty{
//   		CaIdentifier: jsii.String("caIdentifier"),
//   		ValidTill: jsii.String("validTill"),
//   	},
//   	CertificateRotationRestart: jsii.Boolean(false),
//   	CharacterSetName: jsii.String("characterSetName"),
//   	CopyTagsToSnapshot: jsii.Boolean(false),
//   	CustomIamInstanceProfile: jsii.String("customIamInstanceProfile"),
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	DbClusterSnapshotIdentifier: jsii.String("dbClusterSnapshotIdentifier"),
//   	DbInstanceClass: jsii.String("dbInstanceClass"),
//   	DbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   	DbName: jsii.String("dbName"),
//   	DbParameterGroupName: jsii.String("dbParameterGroupName"),
//   	DbSecurityGroups: []*string{
//   		jsii.String("dbSecurityGroups"),
//   	},
//   	DbSnapshotIdentifier: jsii.String("dbSnapshotIdentifier"),
//   	DbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	DedicatedLogVolume: jsii.Boolean(false),
//   	DeleteAutomatedBackups: jsii.Boolean(false),
//   	DeletionProtection: jsii.Boolean(false),
//   	Domain: jsii.String("domain"),
//   	DomainAuthSecretArn: jsii.String("domainAuthSecretArn"),
//   	DomainDnsIps: []*string{
//   		jsii.String("domainDnsIps"),
//   	},
//   	DomainFqdn: jsii.String("domainFqdn"),
//   	DomainIamRoleName: jsii.String("domainIamRoleName"),
//   	DomainOu: jsii.String("domainOu"),
//   	EnableCloudwatchLogsExports: []*string{
//   		jsii.String("enableCloudwatchLogsExports"),
//   	},
//   	EnableIamDatabaseAuthentication: jsii.Boolean(false),
//   	EnablePerformanceInsights: jsii.Boolean(false),
//   	Endpoint: &EndpointProperty{
//   		Address: jsii.String("address"),
//   		HostedZoneId: jsii.String("hostedZoneId"),
//   		Port: jsii.String("port"),
//   	},
//   	Engine: jsii.String("engine"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	Iops: jsii.Number(123),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LicenseModel: jsii.String("licenseModel"),
//   	ManageMasterUserPassword: jsii.Boolean(false),
//   	MasterUsername: jsii.String("masterUsername"),
//   	MasterUserPassword: jsii.String("masterUserPassword"),
//   	MasterUserSecret: &MasterUserSecretProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	MaxAllocatedStorage: jsii.Number(123),
//   	MonitoringInterval: jsii.Number(123),
//   	MonitoringRoleArn: jsii.String("monitoringRoleArn"),
//   	MultiAz: jsii.Boolean(false),
//   	NcharCharacterSetName: jsii.String("ncharCharacterSetName"),
//   	NetworkType: jsii.String("networkType"),
//   	OptionGroupName: jsii.String("optionGroupName"),
//   	PerformanceInsightsKmsKeyId: jsii.String("performanceInsightsKmsKeyId"),
//   	PerformanceInsightsRetentionPeriod: jsii.Number(123),
//   	Port: jsii.String("port"),
//   	PreferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	ProcessorFeatures: []interface{}{
//   		&ProcessorFeatureProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	PromotionTier: jsii.Number(123),
//   	PubliclyAccessible: jsii.Boolean(false),
//   	ReplicaMode: jsii.String("replicaMode"),
//   	RestoreTime: jsii.String("restoreTime"),
//   	SourceDbClusterIdentifier: jsii.String("sourceDbClusterIdentifier"),
//   	SourceDbInstanceAutomatedBackupsArn: jsii.String("sourceDbInstanceAutomatedBackupsArn"),
//   	SourceDbInstanceIdentifier: jsii.String("sourceDbInstanceIdentifier"),
//   	SourceDbiResourceId: jsii.String("sourceDbiResourceId"),
//   	SourceRegion: jsii.String("sourceRegion"),
//   	StorageEncrypted: jsii.Boolean(false),
//   	StorageThroughput: jsii.Number(123),
//   	StorageType: jsii.String("storageType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TdeCredentialArn: jsii.String("tdeCredentialArn"),
//   	TdeCredentialPassword: jsii.String("tdeCredentialPassword"),
//   	Timezone: jsii.String("timezone"),
//   	UseDefaultProcessorFeatures: jsii.Boolean(false),
//   	UseLatestRestorableTime: jsii.Boolean(false),
//   	VpcSecurityGroups: []*string{
//   		jsii.String("vpcSecurityGroups"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html
//
type CfnDBInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The amount of storage in gibibytes (GiB) to be initially allocated for the database instance.
	AllocatedStorage() *string
	SetAllocatedStorage(val *string)
	// A value that indicates whether major version upgrades are allowed.
	AllowMajorVersionUpgrade() interface{}
	SetAllowMajorVersionUpgrade(val interface{})
	// The AWS Identity and Access Management (IAM) roles associated with the DB instance.
	AssociatedRoles() interface{}
	SetAssociatedRoles(val interface{})
	// The CA identifier of the CA certificate used for the DB instance's server certificate.
	AttrCertificateDetailsCaIdentifier() *string
	// The expiration date of the DB instance’s server certificate.
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
	// The destination region for the backup replication of the DB instance.
	AutomaticBackupReplicationRegion() *string
	SetAutomaticBackupReplicationRegion(val *string)
	// A value that indicates whether minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	// The Availability Zone (AZ) where the database will be created.
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	// The number of days for which automated backups are retained.
	BackupRetentionPeriod() *float64
	SetBackupRetentionPeriod(val *float64)
	// The identifier of the CA certificate for this DB instance.
	CaCertificateIdentifier() *string
	SetCaCertificateIdentifier(val *string)
	// The details of the DB instance's server certificate.
	CertificateDetails() interface{}
	SetCertificateDetails(val interface{})
	// Specifies whether the DB instance is restarted when you rotate your SSL/TLS certificate.
	CertificateRotationRestart() interface{}
	SetCertificateRotationRestart(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// For supported engines, indicates that the DB instance should be associated with the specified character set.
	CharacterSetName() *string
	SetCharacterSetName(val *string)
	// Specifies whether to copy tags from the DB instance to snapshots of the DB instance.
	CopyTagsToSnapshot() interface{}
	SetCopyTagsToSnapshot(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance.
	CustomIamInstanceProfile() *string
	SetCustomIamInstanceProfile(val *string)
	// The identifier of the DB cluster that the instance will belong to.
	DbClusterIdentifier() *string
	SetDbClusterIdentifier(val *string)
	// The identifier for the Multi-AZ DB cluster snapshot to restore from.
	DbClusterSnapshotIdentifier() *string
	SetDbClusterSnapshotIdentifier(val *string)
	// The compute and memory capacity of the DB instance, for example `db.m5.large` . Not all DB instance classes are available in all AWS Regions , or for all database engines. For the full list of DB instance classes, and availability for your engine, see [DB instance classes](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the *Amazon RDS User Guide* or [Aurora DB instance classes](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.DBInstanceClass.html) in the *Amazon Aurora User Guide* .
	DbInstanceClass() *string
	SetDbInstanceClass(val *string)
	// A name for the DB instance.
	DbInstanceIdentifier() *string
	SetDbInstanceIdentifier(val *string)
	// The meaning of this parameter differs according to the database engine you use.
	DbName() *string
	SetDbName(val *string)
	// The name of an existing DB parameter group or a reference to an [AWS::RDS::DBParameterGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html) resource created in the template.
	DbParameterGroupName() *string
	SetDbParameterGroupName(val *string)
	// A list of the DB security groups to assign to the DB instance.
	DbSecurityGroups() *[]*string
	SetDbSecurityGroups(val *[]*string)
	// The name or Amazon Resource Name (ARN) of the DB snapshot that's used to restore the DB instance.
	DbSnapshotIdentifier() *string
	SetDbSnapshotIdentifier(val *string)
	// A DB subnet group to associate with the DB instance.
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	// Indicates whether the DB instance has a dedicated log volume (DLV) enabled.
	DedicatedLogVolume() interface{}
	SetDedicatedLogVolume(val interface{})
	// A value that indicates whether to remove automated backups immediately after the DB instance is deleted.
	DeleteAutomatedBackups() interface{}
	SetDeleteAutomatedBackups(val interface{})
	// A value that indicates whether the DB instance has deletion protection enabled.
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	// The Active Directory directory ID to create the DB instance in.
	Domain() *string
	SetDomain(val *string)
	// The ARN for the Secrets Manager secret with the credentials for the user joining the domain.
	DomainAuthSecretArn() *string
	SetDomainAuthSecretArn(val *string)
	// The IPv4 DNS IP addresses of your primary and secondary Active Directory domain controllers.
	DomainDnsIps() *[]*string
	SetDomainDnsIps(val *[]*string)
	// The fully qualified domain name (FQDN) of an Active Directory domain.
	DomainFqdn() *string
	SetDomainFqdn(val *string)
	// The name of the IAM role to use when making API calls to the Directory Service.
	DomainIamRoleName() *string
	SetDomainIamRoleName(val *string)
	// The Active Directory organizational unit for your DB instance to join.
	DomainOu() *string
	SetDomainOu(val *string)
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	EnableCloudwatchLogsExports() *[]*string
	SetEnableCloudwatchLogsExports(val *[]*string)
	// A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	EnableIamDatabaseAuthentication() interface{}
	SetEnableIamDatabaseAuthentication(val interface{})
	// Specifies whether to enable Performance Insights for the DB instance.
	EnablePerformanceInsights() interface{}
	SetEnablePerformanceInsights(val interface{})
	// The connection endpoint for the DB instance.
	Endpoint() interface{}
	SetEndpoint(val interface{})
	// The name of the database engine to use for this DB instance.
	Engine() *string
	SetEngine(val *string)
	// The version number of the database engine to use.
	EngineVersion() *string
	SetEngineVersion(val *string)
	// The number of I/O operations per second (IOPS) that the database provisions.
	Iops() *float64
	SetIops(val *float64)
	// The ARN of the AWS KMS key that's used to encrypt the DB instance, such as `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef` .
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	// License model information for this DB instance.
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
	LogicalId() *string
	// Specifies whether to manage the master user password with AWS Secrets Manager.
	ManageMasterUserPassword() interface{}
	SetManageMasterUserPassword(val interface{})
	// The master user name for the DB instance.
	MasterUsername() *string
	SetMasterUsername(val *string)
	// The password for the master user.
	//
	// The password can include any printable ASCII character except "/", """, or "@".
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	// The secret managed by RDS in AWS Secrets Manager for the master user password.
	MasterUserSecret() interface{}
	SetMasterUserSecret(val interface{})
	// The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.
	MaxAllocatedStorage() *float64
	SetMaxAllocatedStorage(val *float64)
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.
	MonitoringInterval() *float64
	SetMonitoringInterval(val *float64)
	// The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs.
	MonitoringRoleArn() *string
	SetMonitoringRoleArn(val *string)
	// Specifies whether the database instance is a Multi-AZ DB instance deployment.
	MultiAz() interface{}
	SetMultiAz(val interface{})
	// The name of the NCHAR character set for the Oracle DB instance.
	NcharCharacterSetName() *string
	SetNcharCharacterSetName(val *string)
	// The network type of the DB instance.
	NetworkType() *string
	SetNetworkType(val *string)
	// The tree node.
	Node() constructs.Node
	// Indicates that the DB instance should be associated with the specified option group.
	OptionGroupName() *string
	SetOptionGroupName(val *string)
	// The AWS KMS key identifier for encryption of Performance Insights data.
	PerformanceInsightsKmsKeyId() *string
	SetPerformanceInsightsKmsKeyId(val *string)
	// The number of days to retain Performance Insights data.
	PerformanceInsightsRetentionPeriod() *float64
	SetPerformanceInsightsRetentionPeriod(val *float64)
	// The port number on which the database accepts connections.
	Port() *string
	SetPort(val *string)
	// The daily time range during which automated backups are created if automated backups are enabled, using the `BackupRetentionPeriod` parameter.
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	// The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.
	ProcessorFeatures() interface{}
	SetProcessorFeatures(val interface{})
	// The order of priority in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance.
	PromotionTier() *float64
	SetPromotionTier(val *float64)
	// Indicates whether the DB instance is an internet-facing instance.
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The open mode of an Oracle read replica.
	ReplicaMode() *string
	SetReplicaMode(val *string)
	// The date and time to restore from.
	RestoreTime() *string
	SetRestoreTime(val *string)
	// The identifier of the Multi-AZ DB cluster that will act as the source for the read replica.
	SourceDbClusterIdentifier() *string
	SetSourceDbClusterIdentifier(val *string)
	// The Amazon Resource Name (ARN) of the replicated automated backups from which to restore, for example, `arn:aws:rds:us-east-1:123456789012:auto-backup:ab-L2IJCEXJP7XQ7HOJ4SIEXAMPLE` .
	SourceDbInstanceAutomatedBackupsArn() *string
	SetSourceDbInstanceAutomatedBackupsArn(val *string)
	// If you want to create a read replica DB instance, specify the ID of the source DB instance.
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
	Stack() awscdk.Stack
	// A value that indicates whether the DB instance is encrypted.
	//
	// By default, it isn't encrypted.
	StorageEncrypted() interface{}
	SetStorageEncrypted(val interface{})
	// Specifies the storage throughput value for the DB instance.
	//
	// This setting applies only to the `gp3` storage type.
	StorageThroughput() *float64
	SetStorageThroughput(val *float64)
	// The storage type to associate with the DB instance.
	StorageType() *string
	SetStorageType(val *string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// An optional array of key-value pairs to apply to this DB instance.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// Deprecated: this property has been deprecated.
	TdeCredentialArn() *string
	// Deprecated: this property has been deprecated.
	SetTdeCredentialArn(val *string)
	// Deprecated: this property has been deprecated.
	TdeCredentialPassword() *string
	// Deprecated: this property has been deprecated.
	SetTdeCredentialPassword(val *string)
	// The time zone of the DB instance.
	Timezone() *string
	SetTimezone(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Specifies whether the DB instance class of the DB instance uses its default processor features.
	UseDefaultProcessorFeatures() interface{}
	SetUseDefaultProcessorFeatures(val interface{})
	// Specifies whether the DB instance is restored from the latest backup time.
	UseLatestRestorableTime() interface{}
	SetUseLatestRestorableTime(val interface{})
	// A list of the VPC security group IDs to assign to the DB instance.
	VpcSecurityGroups() *[]*string
	SetVpcSecurityGroups(val *[]*string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBInstance
type jsiiProxy_CfnDBInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
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

func (j *jsiiProxy_CfnDBInstance) AutomaticBackupReplicationRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticBackupReplicationRegion",
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

func (j *jsiiProxy_CfnDBInstance) DedicatedLogVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedLogVolume",
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

func (j *jsiiProxy_CfnDBInstance) DomainAuthSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAuthSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DomainDnsIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDnsIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DomainFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdn",
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

func (j *jsiiProxy_CfnDBInstance) DomainOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainOu",
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

func (j *jsiiProxy_CfnDBInstance) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_CfnDBInstance) SourceDbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifier",
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

func (j *jsiiProxy_CfnDBInstance) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) TdeCredentialArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeCredentialArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) TdeCredentialPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeCredentialPassword",
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

func (j *jsiiProxy_CfnDBInstance) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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


func NewCfnDBInstance(scope constructs.Construct, id *string, props *CfnDBInstanceProps) CfnDBInstance {
	_init_.Initialize()

	if err := validateNewCfnDBInstanceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDBInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnDBInstance_Override(c CfnDBInstance, scope constructs.Construct, id *string, props *CfnDBInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.CfnDBInstance",
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

func (j *jsiiProxy_CfnDBInstance)SetAutomaticBackupReplicationRegion(val *string) {
	_jsii_.Set(
		j,
		"automaticBackupReplicationRegion",
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

func (j *jsiiProxy_CfnDBInstance)SetDedicatedLogVolume(val interface{}) {
	if err := j.validateSetDedicatedLogVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedLogVolume",
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

func (j *jsiiProxy_CfnDBInstance)SetDomainAuthSecretArn(val *string) {
	_jsii_.Set(
		j,
		"domainAuthSecretArn",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDomainDnsIps(val *[]*string) {
	_jsii_.Set(
		j,
		"domainDnsIps",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetDomainFqdn(val *string) {
	_jsii_.Set(
		j,
		"domainFqdn",
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

func (j *jsiiProxy_CfnDBInstance)SetDomainOu(val *string) {
	_jsii_.Set(
		j,
		"domainOu",
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

func (j *jsiiProxy_CfnDBInstance)SetSourceDbClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"sourceDbClusterIdentifier",
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

func (j *jsiiProxy_CfnDBInstance)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetTdeCredentialArn(val *string) {
	_jsii_.Set(
		j,
		"tdeCredentialArn",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance)SetTdeCredentialPassword(val *string) {
	_jsii_.Set(
		j,
		"tdeCredentialPassword",
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
func CfnDBInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBInstance_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnDBInstance_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBInstance_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBInstance",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnDBInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CfnDBInstance",
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
		"aws-cdk-lib.aws_rds.CfnDBInstance",
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

func (c *jsiiProxy_CfnDBInstance) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnDBInstance) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
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

func (c *jsiiProxy_CfnDBInstance) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
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

func (c *jsiiProxy_CfnDBInstance) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnDBInstance) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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

