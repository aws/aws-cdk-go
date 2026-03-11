package awsdsql

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdsql/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DSQL::Cluster` resource specifies an cluster. You can use this resource to create, modify, and manage clusters.
//
// This resource supports both single-Region clusters and multi-Region clusters through the `MultiRegionProperties` parameter.
//
// > Creating multi-Region clusters requires additional IAM permissions beyond those needed for single-Region clusters. > - The witness Region specified in `multiRegionProperties.witnessRegion` cannot be the same as the cluster's Region.
//
// *Required permissions*
//
// - **dsql:CreateCluster** - Required to create a cluster.
//
// Resources: `arn:aws:dsql:region:account-id:cluster/*`
// - **dsql:TagResource** - Permission to add tags to a resource.
//
// Resources: `arn:aws:dsql:region:account-id:cluster/*`
// - **dsql:PutMultiRegionProperties** - Permission to configure multi-Region properties for a cluster.
//
// Resources: `arn:aws:dsql:region:account-id:cluster/*`
// - **dsql:AddPeerCluster** - When specifying `multiRegionProperties.clusters` , permission to add peer clusters.
//
// Resources:
//
// - Local cluster: `arn:aws:dsql:region:account-id:cluster/*`
// - Each peer cluster: exact ARN of each specified peer cluster
// - **dsql:PutWitnessRegion** - When specifying `multiRegionProperties.witnessRegion` , permission to set a witness Region. This permission is checked both in the cluster Region and in the witness Region.
//
// Resources: `arn:aws:dsql:region:account-id:cluster/*`
//
// Condition Keys: `dsql:WitnessRegion` (matching the specified witness region).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnClusterPropsMixin := awscdkcfnpropertymixins.Aws_dsql.NewCfnClusterPropsMixin(&CfnClusterMixinProps{
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	KmsEncryptionKey: jsii.String("kmsEncryptionKey"),
//   	MultiRegionProperties: &MultiRegionPropertiesProperty{
//   		Clusters: []*string{
//   			jsii.String("clusters"),
//   		},
//   		WitnessRegion: jsii.String("witnessRegion"),
//   	},
//   	PolicyDocument: jsii.String("policyDocument"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dsql-cluster.html
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


// Create a mixin to apply properties to `AWS::DSQL::Cluster`.
func NewCfnClusterPropsMixin(props *CfnClusterMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dsql.CfnClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DSQL::Cluster`.
func NewCfnClusterPropsMixin_Override(c CfnClusterPropsMixin, props *CfnClusterMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dsql.CfnClusterPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_dsql.CfnClusterPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_dsql.CfnClusterPropsMixin",
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

