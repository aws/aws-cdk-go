package awsredshift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a new Amazon Redshift security group. You use security groups to control access to non-VPC clusters.
//
// For information about managing security groups, go to [Amazon Redshift Cluster Security Groups](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html) in the *Amazon Redshift Cluster Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnClusterSecurityGroupPropsMixin := awscdkcfnpropertymixins.Aws_redshift.NewCfnClusterSecurityGroupPropsMixin(&CfnClusterSecurityGroupMixinProps{
//   	Description: jsii.String("description"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html
//
type CfnClusterSecurityGroupPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnClusterSecurityGroupMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnClusterSecurityGroupPropsMixin
type jsiiProxy_CfnClusterSecurityGroupPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnClusterSecurityGroupPropsMixin) Props() *CfnClusterSecurityGroupMixinProps {
	var returns *CfnClusterSecurityGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Redshift::ClusterSecurityGroup`.
func NewCfnClusterSecurityGroupPropsMixin(props *CfnClusterSecurityGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnClusterSecurityGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClusterSecurityGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClusterSecurityGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnClusterSecurityGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Redshift::ClusterSecurityGroup`.
func NewCfnClusterSecurityGroupPropsMixin_Override(c CfnClusterSecurityGroupPropsMixin, props *CfnClusterSecurityGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnClusterSecurityGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnClusterSecurityGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClusterSecurityGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnClusterSecurityGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClusterSecurityGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_redshift.CfnClusterSecurityGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

