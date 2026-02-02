package previewawsrdsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrds/previewawsrdsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::RDS::DBSecurityGroupIngress` resource enables ingress to a DB security group using one of two forms of authorization.
//
// First, you can add EC2 or VPC security groups to the DB security group if the application using the database is running on EC2 or VPC instances. Second, IP ranges are available if the application accessing your database is running on the Internet.
//
// This type supports updates. For more information about updating stacks, see [AWS CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html) .
//
// For details about the settings for DB security group ingress, see [AuthorizeDBSecurityGroupIngress](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_AuthorizeDBSecurityGroupIngress.html) .
//
// > EC2-Classic was retired on August 15, 2022. If you haven't migrated from EC2-Classic to a VPC, we recommend that you migrate as soon as possible. For more information, see [Migrate from EC2-Classic to a VPC](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html) in the *Amazon EC2 User Guide* , the blog [EC2-Classic Networking is Retiring – Here’s How to Prepare](https://docs.aws.amazon.com/aws/ec2-classic-is-retiring-heres-how-to-prepare/) , and [Moving a DB instance not in a VPC into a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html) in the *Amazon RDS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDBSecurityGroupIngressPropsMixin := awscdkmixinspreview.Mixins.NewCfnDBSecurityGroupIngressPropsMixin(&CfnDBSecurityGroupIngressMixinProps{
//   	Cidrip: jsii.String("cidrip"),
//   	DbSecurityGroupName: jsii.String("dbSecurityGroupName"),
//   	Ec2SecurityGroupId: jsii.String("ec2SecurityGroupId"),
//   	Ec2SecurityGroupName: jsii.String("ec2SecurityGroupName"),
//   	Ec2SecurityGroupOwnerId: jsii.String("ec2SecurityGroupOwnerId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsecuritygroupingress.html
//
type CfnDBSecurityGroupIngressPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDBSecurityGroupIngressMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDBSecurityGroupIngressPropsMixin
type jsiiProxy_CfnDBSecurityGroupIngressPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDBSecurityGroupIngressPropsMixin) Props() *CfnDBSecurityGroupIngressMixinProps {
	var returns *CfnDBSecurityGroupIngressMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSecurityGroupIngressPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RDS::DBSecurityGroupIngress`.
func NewCfnDBSecurityGroupIngressPropsMixin(props *CfnDBSecurityGroupIngressMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDBSecurityGroupIngressPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDBSecurityGroupIngressPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDBSecurityGroupIngressPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBSecurityGroupIngressPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RDS::DBSecurityGroupIngress`.
func NewCfnDBSecurityGroupIngressPropsMixin_Override(c CfnDBSecurityGroupIngressPropsMixin, props *CfnDBSecurityGroupIngressMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBSecurityGroupIngressPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDBSecurityGroupIngressPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBSecurityGroupIngressPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBSecurityGroupIngressPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBSecurityGroupIngressPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBSecurityGroupIngressPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBSecurityGroupIngressPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDBSecurityGroupIngressPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

