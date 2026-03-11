package awsdax

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdax/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a DAX cluster.
//
// All nodes in the cluster run the same DAX caching software.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var tags interface{}
//
//   cfnClusterPropsMixin := awscdkcfnpropertymixins.Aws_dax.NewCfnClusterPropsMixin(&CfnClusterMixinProps{
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	ClusterEndpointEncryptionType: jsii.String("clusterEndpointEncryptionType"),
//   	ClusterName: jsii.String("clusterName"),
//   	Description: jsii.String("description"),
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	NetworkType: jsii.String("networkType"),
//   	NodeType: jsii.String("nodeType"),
//   	NotificationTopicArn: jsii.String("notificationTopicArn"),
//   	ParameterGroupName: jsii.String("parameterGroupName"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	ReplicationFactor: jsii.Number(123),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SseSpecification: &SSESpecificationProperty{
//   		SseEnabled: jsii.Boolean(false),
//   	},
//   	SubnetGroupName: jsii.String("subnetGroupName"),
//   	Tags: tags,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html
//
type CfnClusterPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnClusterMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnClusterPropsMixin
type jsiiProxy_CfnClusterPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnClusterPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DAX::Cluster`.
func NewCfnClusterPropsMixin(props *CfnClusterMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dax.CfnClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DAX::Cluster`.
func NewCfnClusterPropsMixin_Override(c CfnClusterPropsMixin, props *CfnClusterMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dax.CfnClusterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnClusterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClusterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_dax.CfnClusterPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_dax.CfnClusterPropsMixin",
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

