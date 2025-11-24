package mixinsawsredshift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsredshift/mixinsawsredshift/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a cluster. A *cluster* is a fully managed data warehouse that consists of a set of compute nodes.
//
// To create a cluster in Virtual Private Cloud (VPC), you must provide a cluster subnet group name. The cluster subnet group identifies the subnets of your VPC that Amazon Redshift uses when creating the cluster. For more information about managing clusters, go to [Amazon Redshift Clusters](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html) in the *Amazon Redshift Cluster Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var namespaceResourcePolicy interface{}
//
//   cfnClusterPropsMixin := awscdkmixinspreview.Mixins.NewCfnClusterPropsMixin(&CfnClusterMixinProps{
//   	AllowVersionUpgrade: jsii.Boolean(false),
//   	AquaConfigurationStatus: jsii.String("aquaConfigurationStatus"),
//   	AutomatedSnapshotRetentionPeriod: jsii.Number(123),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneRelocation: jsii.Boolean(false),
//   	AvailabilityZoneRelocationStatus: jsii.String("availabilityZoneRelocationStatus"),
//   	Classic: jsii.Boolean(false),
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	ClusterParameterGroupName: jsii.String("clusterParameterGroupName"),
//   	ClusterSecurityGroups: []*string{
//   		jsii.String("clusterSecurityGroups"),
//   	},
//   	ClusterSubnetGroupName: jsii.String("clusterSubnetGroupName"),
//   	ClusterType: jsii.String("clusterType"),
//   	ClusterVersion: jsii.String("clusterVersion"),
//   	DbName: jsii.String("dbName"),
//   	DeferMaintenance: jsii.Boolean(false),
//   	DeferMaintenanceDuration: jsii.Number(123),
//   	DeferMaintenanceEndTime: jsii.String("deferMaintenanceEndTime"),
//   	DeferMaintenanceStartTime: jsii.String("deferMaintenanceStartTime"),
//   	DestinationRegion: jsii.String("destinationRegion"),
//   	ElasticIp: jsii.String("elasticIp"),
//   	Encrypted: jsii.Boolean(false),
//   	Endpoint: &EndpointProperty{
//   		Address: jsii.String("address"),
//   		Port: jsii.String("port"),
//   	},
//   	EnhancedVpcRouting: jsii.Boolean(false),
//   	HsmClientCertificateIdentifier: jsii.String("hsmClientCertificateIdentifier"),
//   	HsmConfigurationIdentifier: jsii.String("hsmConfigurationIdentifier"),
//   	IamRoles: []*string{
//   		jsii.String("iamRoles"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LoggingProperties: &LoggingPropertiesProperty{
//   		BucketName: jsii.String("bucketName"),
//   		LogDestinationType: jsii.String("logDestinationType"),
//   		LogExports: []*string{
//   			jsii.String("logExports"),
//   		},
//   		S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	},
//   	MaintenanceTrackName: jsii.String("maintenanceTrackName"),
//   	ManageMasterPassword: jsii.Boolean(false),
//   	ManualSnapshotRetentionPeriod: jsii.Number(123),
//   	MasterPasswordSecretKmsKeyId: jsii.String("masterPasswordSecretKmsKeyId"),
//   	MasterUsername: jsii.String("masterUsername"),
//   	MasterUserPassword: jsii.String("masterUserPassword"),
//   	MultiAz: jsii.Boolean(false),
//   	NamespaceResourcePolicy: namespaceResourcePolicy,
//   	NodeType: jsii.String("nodeType"),
//   	NumberOfNodes: jsii.Number(123),
//   	OwnerAccount: jsii.String("ownerAccount"),
//   	Port: jsii.Number(123),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	PubliclyAccessible: jsii.Boolean(false),
//   	ResourceAction: jsii.String("resourceAction"),
//   	RevisionTarget: jsii.String("revisionTarget"),
//   	RotateEncryptionKey: jsii.Boolean(false),
//   	SnapshotClusterIdentifier: jsii.String("snapshotClusterIdentifier"),
//   	SnapshotCopyGrantName: jsii.String("snapshotCopyGrantName"),
//   	SnapshotCopyManual: jsii.Boolean(false),
//   	SnapshotCopyRetentionPeriod: jsii.Number(123),
//   	SnapshotIdentifier: jsii.String("snapshotIdentifier"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html
//
type CfnClusterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnClusterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnClusterPropsMixin
type jsiiProxy_CfnClusterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnClusterPropsMixin) Props() *CfnClusterMixinProps {
	var returns *CfnClusterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Redshift::Cluster`.
func NewCfnClusterPropsMixin(props *CfnClusterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_redshift.mixins.CfnClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Redshift::Cluster`.
func NewCfnClusterPropsMixin_Override(c CfnClusterPropsMixin, props *CfnClusterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_redshift.mixins.CfnClusterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnClusterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClusterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_redshift.mixins.CfnClusterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClusterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_redshift.mixins.CfnClusterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClusterPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnClusterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

