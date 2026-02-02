package previewawsmemorydbmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmemorydb/previewawsmemorydbmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a cluster .
//
// All nodes in the cluster run the same protocol-compliant engine software.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPropsMixin := awscdkmixinspreview.Mixins.NewCfnClusterPropsMixin(&CfnClusterMixinProps{
//   	AclName: jsii.String("aclName"),
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	ClusterEndpoint: &EndpointProperty{
//   		Address: jsii.String("address"),
//   		Port: jsii.Number(123),
//   	},
//   	ClusterName: jsii.String("clusterName"),
//   	DataTiering: jsii.String("dataTiering"),
//   	Description: jsii.String("description"),
//   	Engine: jsii.String("engine"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	FinalSnapshotName: jsii.String("finalSnapshotName"),
//   	IpDiscovery: jsii.String("ipDiscovery"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MaintenanceWindow: jsii.String("maintenanceWindow"),
//   	MultiRegionClusterName: jsii.String("multiRegionClusterName"),
//   	NetworkType: jsii.String("networkType"),
//   	NodeType: jsii.String("nodeType"),
//   	NumReplicasPerShard: jsii.Number(123),
//   	NumShards: jsii.Number(123),
//   	ParameterGroupName: jsii.String("parameterGroupName"),
//   	Port: jsii.Number(123),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SnapshotArns: []*string{
//   		jsii.String("snapshotArns"),
//   	},
//   	SnapshotName: jsii.String("snapshotName"),
//   	SnapshotRetentionLimit: jsii.Number(123),
//   	SnapshotWindow: jsii.String("snapshotWindow"),
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   	SnsTopicStatus: jsii.String("snsTopicStatus"),
//   	SubnetGroupName: jsii.String("subnetGroupName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TlsEnabled: jsii.Boolean(false),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-cluster.html
//
type CfnClusterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnClusterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
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


// Create a mixin to apply properties to `AWS::MemoryDB::Cluster`.
func NewCfnClusterPropsMixin(props *CfnClusterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_memorydb.mixins.CfnClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MemoryDB::Cluster`.
func NewCfnClusterPropsMixin_Override(c CfnClusterPropsMixin, props *CfnClusterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_memorydb.mixins.CfnClusterPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_memorydb.mixins.CfnClusterPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_memorydb.mixins.CfnClusterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClusterPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

