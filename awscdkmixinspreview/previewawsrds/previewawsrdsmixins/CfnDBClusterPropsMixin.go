package previewawsrdsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrds/previewawsrdsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::RDS::DBCluster` resource creates an Amazon Aurora DB cluster or Multi-AZ DB cluster.
//
// For more information about creating an Aurora DB cluster, see [Creating an Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.CreateInstance.html) in the *Amazon Aurora User Guide* .
//
// For more information about creating a Multi-AZ DB cluster, see [Creating a Multi-AZ DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/create-multi-az-db-cluster.html) in the *Amazon RDS User Guide* .
//
// > You can only create this resource in AWS Regions where Amazon Aurora or Multi-AZ DB clusters are supported.
//
// *Updating DB clusters*
//
// When properties labeled " *Update requires:* [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) " are updated, AWS CloudFormation first creates a replacement DB cluster, then changes references from other dependent resources to point to the replacement DB cluster, and finally deletes the old DB cluster.
//
// > We highly recommend that you take a snapshot of the database before updating the stack. If you don't, you lose the data when AWS CloudFormation replaces your DB cluster. To preserve your data, perform the following procedure:
// >
// > - Deactivate any applications that are using the DB cluster so that there's no activity on the DB instance.
// > - Create a snapshot of the DB cluster. For more information, see [Creating a DB cluster snapshot](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_CreateSnapshotCluster.html) .
// > - If you want to restore your DB cluster using a DB cluster snapshot, modify the updated template with your DB cluster changes and add the `SnapshotIdentifier` property with the ID of the DB cluster snapshot that you want to use.
// >
// > After you restore a DB cluster with a `SnapshotIdentifier` property, you must specify the same `SnapshotIdentifier` property for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the DB cluster snapshot again, and the data in the database is not changed. However, if you don't specify the `SnapshotIdentifier` property, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB cluster is restored from the specified `SnapshotIdentifier` property, and the original DB cluster is deleted.
// > - Update the stack.
//
// Currently, when you are updating the stack for an Aurora Serverless DB cluster, you can't include changes to any other properties when you specify one of the following properties: `PreferredBackupWindow` , `PreferredMaintenanceWindow` , and `Port` . This limitation doesn't apply to provisioned DB clusters.
//
// For more information about updating other properties of this resource, see `[ModifyDBCluster](https://docs.aws.amazon.com//AmazonRDS/latest/APIReference/API_ModifyDBCluster.html)` . For more information about updating stacks, see [AWS CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html) .
//
// *Deleting DB clusters*
//
// The default `DeletionPolicy` for `AWS::RDS::DBCluster` resources is `Snapshot` . For more information about how AWS CloudFormation deletes resources, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDBClusterPropsMixin := awscdkmixinspreview.Mixins.NewCfnDBClusterPropsMixin(&CfnDBClusterMixinProps{
//   	AllocatedStorage: jsii.Number(123),
//   	AssociatedRoles: []interface{}{
//   		&DBClusterRoleProperty{
//   			FeatureName: jsii.String("featureName"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	BacktrackWindow: jsii.Number(123),
//   	BackupRetentionPeriod: jsii.Number(123),
//   	ClusterScalabilityType: jsii.String("clusterScalabilityType"),
//   	CopyTagsToSnapshot: jsii.Boolean(false),
//   	DatabaseInsightsMode: jsii.String("databaseInsightsMode"),
//   	DatabaseName: jsii.String("databaseName"),
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	DbClusterInstanceClass: jsii.String("dbClusterInstanceClass"),
//   	DbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	DbInstanceParameterGroupName: jsii.String("dbInstanceParameterGroupName"),
//   	DbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	DbSystemId: jsii.String("dbSystemId"),
//   	DeleteAutomatedBackups: jsii.Boolean(false),
//   	DeletionProtection: jsii.Boolean(false),
//   	Domain: jsii.String("domain"),
//   	DomainIamRoleName: jsii.String("domainIamRoleName"),
//   	EnableCloudwatchLogsExports: []*string{
//   		jsii.String("enableCloudwatchLogsExports"),
//   	},
//   	EnableGlobalWriteForwarding: jsii.Boolean(false),
//   	EnableHttpEndpoint: jsii.Boolean(false),
//   	EnableIamDatabaseAuthentication: jsii.Boolean(false),
//   	EnableLocalWriteForwarding: jsii.Boolean(false),
//   	Engine: jsii.String("engine"),
//   	EngineLifecycleSupport: jsii.String("engineLifecycleSupport"),
//   	EngineMode: jsii.String("engineMode"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   	Iops: jsii.Number(123),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ManageMasterUserPassword: jsii.Boolean(false),
//   	MasterUserAuthenticationType: jsii.String("masterUserAuthenticationType"),
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
//   		SecondsUntilAutoPause: jsii.Number(123),
//   	},
//   	SnapshotIdentifier: jsii.String("snapshotIdentifier"),
//   	SourceDbClusterIdentifier: jsii.String("sourceDbClusterIdentifier"),
//   	SourceDbClusterResourceId: jsii.String("sourceDbClusterResourceId"),
//   	SourceRegion: jsii.String("sourceRegion"),
//   	StorageEncrypted: jsii.Boolean(false),
//   	StorageType: jsii.String("storageType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseLatestRestorableTime: jsii.Boolean(false),
//   	VpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html
//
type CfnDBClusterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDBClusterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDBClusterPropsMixin
type jsiiProxy_CfnDBClusterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDBClusterPropsMixin) Props() *CfnDBClusterMixinProps {
	var returns *CfnDBClusterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RDS::DBCluster`.
func NewCfnDBClusterPropsMixin(props *CfnDBClusterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDBClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDBClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDBClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RDS::DBCluster`.
func NewCfnDBClusterPropsMixin_Override(c CfnDBClusterPropsMixin, props *CfnDBClusterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBClusterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDBClusterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBClusterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBClusterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBClusterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBClusterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBClusterPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDBClusterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

