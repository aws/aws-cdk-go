package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Associates a subnet with a network ACL. For more information, see [ReplaceNetworkAclAssociation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceNetworkAclAssociation.html) in the *Amazon EC2 API Reference* .
//
// When `AWS::EC2::SubnetNetworkAclAssociation` resources are created during create or update operations, AWS CloudFormation adopts existing resources that share the same key properties (the properties that contribute to uniquely identify the resource). However, if the operation fails and rolls back, AWS CloudFormation deletes the previously out-of-band resources. You can protect against this behavior by using `Retain` deletion policies. For more information, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSubnetNetworkAclAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnSubnetNetworkAclAssociationPropsMixin(&CfnSubnetNetworkAclAssociationMixinProps{
//   	NetworkAclId: jsii.String("networkAclId"),
//   	SubnetId: jsii.String("subnetId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetnetworkaclassociation.html
//
type CfnSubnetNetworkAclAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSubnetNetworkAclAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSubnetNetworkAclAssociationPropsMixin
type jsiiProxy_CfnSubnetNetworkAclAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSubnetNetworkAclAssociationPropsMixin) Props() *CfnSubnetNetworkAclAssociationMixinProps {
	var returns *CfnSubnetNetworkAclAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubnetNetworkAclAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::SubnetNetworkAclAssociation`.
func NewCfnSubnetNetworkAclAssociationPropsMixin(props *CfnSubnetNetworkAclAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSubnetNetworkAclAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSubnetNetworkAclAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSubnetNetworkAclAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSubnetNetworkAclAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::SubnetNetworkAclAssociation`.
func NewCfnSubnetNetworkAclAssociationPropsMixin_Override(c CfnSubnetNetworkAclAssociationPropsMixin, props *CfnSubnetNetworkAclAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSubnetNetworkAclAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSubnetNetworkAclAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSubnetNetworkAclAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSubnetNetworkAclAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSubnetNetworkAclAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSubnetNetworkAclAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSubnetNetworkAclAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSubnetNetworkAclAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

