package awspersonalize

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awspersonalize/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A deployment of a solution version that provides real-time recommendations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnCampaignPropsMixin := awscdkcfnpropertymixins.Aws_personalize.NewCfnCampaignPropsMixin(&CfnCampaignMixinProps{
//   	CampaignConfig: &CampaignConfigProperty{
//   		EnableMetadataWithRecommendations: jsii.Boolean(false),
//   		ItemExplorationConfig: map[string]*string{
//   			"itemExplorationConfigKey": jsii.String("itemExplorationConfig"),
//   		},
//   		RankingInfluence: map[string]*f64{
//   			"rankingInfluenceKey": jsii.Number(123),
//   		},
//   		SyncWithLatestSolutionVersion: jsii.Boolean(false),
//   	},
//   	MinProvisionedTps: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	SolutionVersionArn: jsii.String("solutionVersionArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-campaign.html
//
type CfnCampaignPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnCampaignMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCampaignPropsMixin
type jsiiProxy_CfnCampaignPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnCampaignPropsMixin) Props() *CfnCampaignMixinProps {
	var returns *CfnCampaignMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaignPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Personalize::Campaign`.
func NewCfnCampaignPropsMixin(props *CfnCampaignMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnCampaignPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCampaignPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCampaignPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnCampaignPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Personalize::Campaign`.
func NewCfnCampaignPropsMixin_Override(c CfnCampaignPropsMixin, props *CfnCampaignMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnCampaignPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnCampaignPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCampaignPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnCampaignPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCampaignPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnCampaignPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCampaignPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCampaignPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

