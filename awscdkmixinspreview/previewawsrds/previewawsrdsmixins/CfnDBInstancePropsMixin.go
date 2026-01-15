package previewawsrdsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrds/previewawsrdsmixins/internal"
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
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDBInstancePropsMixin := awscdkmixinspreview.Mixins.NewCfnDBInstancePropsMixin(&CfnDBInstanceMixinProps{
//   	AdditionalStorageVolumes: []interface{}{
//   		&AdditionalStorageVolumeProperty{
//   			AllocatedStorage: jsii.String("allocatedStorage"),
//   			Iops: jsii.Number(123),
//   			MaxAllocatedStorage: jsii.Number(123),
//   			StorageThroughput: jsii.Number(123),
//   			StorageType: jsii.String("storageType"),
//   			VolumeName: jsii.String("volumeName"),
//   		},
//   	},
//   	AllocatedStorage: jsii.String("allocatedStorage"),
//   	AllowMajorVersionUpgrade: jsii.Boolean(false),
//   	ApplyImmediately: jsii.Boolean(false),
//   	AssociatedRoles: []interface{}{
//   		&DBInstanceRoleProperty{
//   			FeatureName: jsii.String("featureName"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	AutomaticBackupReplicationKmsKeyId: jsii.String("automaticBackupReplicationKmsKeyId"),
//   	AutomaticBackupReplicationRegion: jsii.String("automaticBackupReplicationRegion"),
//   	AutomaticBackupReplicationRetentionPeriod: jsii.Number(123),
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	BackupRetentionPeriod: jsii.Number(123),
//   	BackupTarget: jsii.String("backupTarget"),
//   	CaCertificateIdentifier: jsii.String("caCertificateIdentifier"),
//   	CertificateRotationRestart: jsii.Boolean(false),
//   	CharacterSetName: jsii.String("characterSetName"),
//   	CopyTagsToSnapshot: jsii.Boolean(false),
//   	CustomIamInstanceProfile: jsii.String("customIamInstanceProfile"),
//   	DatabaseInsightsMode: jsii.String("databaseInsightsMode"),
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
//   	DbSystemId: jsii.String("dbSystemId"),
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
//   	Engine: jsii.String("engine"),
//   	EngineLifecycleSupport: jsii.String("engineLifecycleSupport"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	Iops: jsii.Number(123),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LicenseModel: jsii.String("licenseModel"),
//   	ManageMasterUserPassword: jsii.Boolean(false),
//   	MasterUserAuthenticationType: jsii.String("masterUserAuthenticationType"),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html
//
type CfnDBInstancePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDBInstanceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDBInstancePropsMixin
type jsiiProxy_CfnDBInstancePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDBInstancePropsMixin) Props() *CfnDBInstanceMixinProps {
	var returns *CfnDBInstanceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstancePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RDS::DBInstance`.
func NewCfnDBInstancePropsMixin(props *CfnDBInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDBInstancePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDBInstancePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDBInstancePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBInstancePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RDS::DBInstance`.
func NewCfnDBInstancePropsMixin_Override(c CfnDBInstancePropsMixin, props *CfnDBInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBInstancePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDBInstancePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBInstancePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBInstancePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBInstancePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBInstancePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBInstancePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstancePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

