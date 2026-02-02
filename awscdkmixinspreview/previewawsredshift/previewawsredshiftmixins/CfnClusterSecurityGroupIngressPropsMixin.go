package previewawsredshiftmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsredshift/previewawsredshiftmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds an inbound (ingress) rule to an Amazon Redshift security group.
//
// Depending on whether the application accessing your cluster is running on the Internet or an Amazon EC2 instance, you can authorize inbound access to either a Classless Interdomain Routing (CIDR)/Internet Protocol (IP) range or to an Amazon EC2 security group. You can add as many as 20 ingress rules to an Amazon Redshift security group.
//
// If you authorize access to an Amazon EC2 security group, specify *EC2SecurityGroupName* and *EC2SecurityGroupOwnerId* . The Amazon EC2 security group and Amazon Redshift cluster must be in the same AWS Region .
//
// If you authorize access to a CIDR/IP address range, specify *CIDRIP* . For an overview of CIDR blocks, see the Wikipedia article on [Classless Inter-Domain Routing](https://docs.aws.amazon.com/http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) .
//
// You must also associate the security group with a cluster so that clients running on these IP addresses or the EC2 instance are authorized to connect to the cluster. For information about managing security groups, go to [Working with Security Groups](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html) in the *Amazon Redshift Cluster Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterSecurityGroupIngressPropsMixin := awscdkmixinspreview.Mixins.NewCfnClusterSecurityGroupIngressPropsMixin(&CfnClusterSecurityGroupIngressMixinProps{
//   	Cidrip: jsii.String("cidrip"),
//   	ClusterSecurityGroupName: jsii.String("clusterSecurityGroupName"),
//   	Ec2SecurityGroupName: jsii.String("ec2SecurityGroupName"),
//   	Ec2SecurityGroupOwnerId: jsii.String("ec2SecurityGroupOwnerId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html
//
type CfnClusterSecurityGroupIngressPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnClusterSecurityGroupIngressMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnClusterSecurityGroupIngressPropsMixin
type jsiiProxy_CfnClusterSecurityGroupIngressPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngressPropsMixin) Props() *CfnClusterSecurityGroupIngressMixinProps {
	var returns *CfnClusterSecurityGroupIngressMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngressPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Redshift::ClusterSecurityGroupIngress`.
func NewCfnClusterSecurityGroupIngressPropsMixin(props *CfnClusterSecurityGroupIngressMixinProps, options *mixins.CfnPropertyMixinOptions) CfnClusterSecurityGroupIngressPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClusterSecurityGroupIngressPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClusterSecurityGroupIngressPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_redshift.mixins.CfnClusterSecurityGroupIngressPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Redshift::ClusterSecurityGroupIngress`.
func NewCfnClusterSecurityGroupIngressPropsMixin_Override(c CfnClusterSecurityGroupIngressPropsMixin, props *CfnClusterSecurityGroupIngressMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_redshift.mixins.CfnClusterSecurityGroupIngressPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnClusterSecurityGroupIngressPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClusterSecurityGroupIngressPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_redshift.mixins.CfnClusterSecurityGroupIngressPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClusterSecurityGroupIngressPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_redshift.mixins.CfnClusterSecurityGroupIngressPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngressPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngressPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

