package awswellarchitected

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awswellarchitected/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::WellArchitected::AgentOrganizationalProfile.
//
// An Agent Organizational Profile is an Agent Profile scoped to an AWS Organization; the organization is inferred from the caller's identity, and the profile does not carry a single-account execution role or aggregation configuration. Only usable in the AWS Organizations management or delegated-admin account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAgentOrganizationalProfilePropsMixin := awscdkcfnpropertymixins.Aws_wellarchitected.NewCfnAgentOrganizationalProfilePropsMixin(&CfnAgentOrganizationalProfileMixinProps{
//   	BusinessOverview: jsii.String("businessOverview"),
//   	DeletionProtection: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	Name: jsii.String("name"),
//   	Pillars: []*string{
//   		jsii.String("pillars"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentorganizationalprofile.html
//
type CfnAgentOrganizationalProfilePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAgentOrganizationalProfileMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAgentOrganizationalProfilePropsMixin
type jsiiProxy_CfnAgentOrganizationalProfilePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAgentOrganizationalProfilePropsMixin) Props() *CfnAgentOrganizationalProfileMixinProps {
	var returns *CfnAgentOrganizationalProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgentOrganizationalProfilePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WellArchitected::AgentOrganizationalProfile`.
func NewCfnAgentOrganizationalProfilePropsMixin(props *CfnAgentOrganizationalProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAgentOrganizationalProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAgentOrganizationalProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAgentOrganizationalProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentOrganizationalProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WellArchitected::AgentOrganizationalProfile`.
func NewCfnAgentOrganizationalProfilePropsMixin_Override(c CfnAgentOrganizationalProfilePropsMixin, props *CfnAgentOrganizationalProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentOrganizationalProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAgentOrganizationalProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAgentOrganizationalProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentOrganizationalProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAgentOrganizationalProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentOrganizationalProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAgentOrganizationalProfilePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAgentOrganizationalProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

