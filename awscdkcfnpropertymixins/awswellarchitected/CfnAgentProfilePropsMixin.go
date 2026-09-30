package awswellarchitected

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awswellarchitected/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::WellArchitected::AgentProfile.
//
// An Agent Profile defines the execution role, account/region scope (via aggregation configuration access roles), and Well-Architected pillar focus used to generate Well-Architected Agent recommendations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAgentProfilePropsMixin := awscdkcfnpropertymixins.Aws_wellarchitected.NewCfnAgentProfilePropsMixin(&CfnAgentProfileMixinProps{
//   	AggregationConfiguration: []interface{}{
//   		&AggregationConfigurationProperty{
//   			AccessRoleArn: jsii.String("accessRoleArn"),
//   			AccountId: jsii.String("accountId"),
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   	},
//   	BusinessOverview: jsii.String("businessOverview"),
//   	DeletionProtection: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofile.html
//
type CfnAgentProfilePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAgentProfileMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAgentProfilePropsMixin
type jsiiProxy_CfnAgentProfilePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAgentProfilePropsMixin) Props() *CfnAgentProfileMixinProps {
	var returns *CfnAgentProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgentProfilePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WellArchitected::AgentProfile`.
func NewCfnAgentProfilePropsMixin(props *CfnAgentProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAgentProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAgentProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAgentProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WellArchitected::AgentProfile`.
func NewCfnAgentProfilePropsMixin_Override(c CfnAgentProfilePropsMixin, props *CfnAgentProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAgentProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAgentProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAgentProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAgentProfilePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAgentProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

