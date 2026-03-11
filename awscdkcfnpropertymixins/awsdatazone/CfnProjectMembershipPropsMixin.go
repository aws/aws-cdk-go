package awsdatazone

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DataZone::ProjectMembership` resource adds a member to an Amazon DataZone project.
//
// Project members consume assets from the Amazon DataZone catalog and produce new assets using one or more analytical workflows.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnProjectMembershipPropsMixin := awscdkcfnpropertymixins.Aws_datazone.NewCfnProjectMembershipPropsMixin(&CfnProjectMembershipMixinProps{
//   	Designation: jsii.String("designation"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Member: &MemberProperty{
//   		GroupIdentifier: jsii.String("groupIdentifier"),
//   		UserIdentifier: jsii.String("userIdentifier"),
//   	},
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectmembership.html
//
type CfnProjectMembershipPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnProjectMembershipMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnProjectMembershipPropsMixin
type jsiiProxy_CfnProjectMembershipPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnProjectMembershipPropsMixin) Props() *CfnProjectMembershipMixinProps {
	var returns *CfnProjectMembershipMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProjectMembershipPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataZone::ProjectMembership`.
func NewCfnProjectMembershipPropsMixin(props *CfnProjectMembershipMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnProjectMembershipPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnProjectMembershipPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProjectMembershipPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectMembershipPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataZone::ProjectMembership`.
func NewCfnProjectMembershipPropsMixin_Override(c CfnProjectMembershipPropsMixin, props *CfnProjectMembershipMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectMembershipPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnProjectMembershipPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProjectMembershipPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectMembershipPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProjectMembershipPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnProjectMembershipPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProjectMembershipPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnProjectMembershipPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

