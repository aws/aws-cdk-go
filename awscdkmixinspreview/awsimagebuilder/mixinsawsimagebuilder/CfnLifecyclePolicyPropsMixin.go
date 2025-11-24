package mixinsawsimagebuilder

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsimagebuilder/mixinsawsimagebuilder/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a lifecycle policy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLifecyclePolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnLifecyclePolicyPropsMixin(&CfnLifecyclePolicyMixinProps{
//   	Description: jsii.String("description"),
//   	ExecutionRole: jsii.String("executionRole"),
//   	Name: jsii.String("name"),
//   	PolicyDetails: []interface{}{
//   		&PolicyDetailProperty{
//   			Action: &ActionProperty{
//   				IncludeResources: &IncludeResourcesProperty{
//   					Amis: jsii.Boolean(false),
//   					Containers: jsii.Boolean(false),
//   					Snapshots: jsii.Boolean(false),
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			ExclusionRules: &ExclusionRulesProperty{
//   				Amis: &AmiExclusionRulesProperty{
//   					IsPublic: jsii.Boolean(false),
//   					LastLaunched: &LastLaunchedProperty{
//   						Unit: jsii.String("unit"),
//   						Value: jsii.Number(123),
//   					},
//   					Regions: []*string{
//   						jsii.String("regions"),
//   					},
//   					SharedAccounts: []*string{
//   						jsii.String("sharedAccounts"),
//   					},
//   					TagMap: map[string]*string{
//   						"tagMapKey": jsii.String("tagMap"),
//   					},
//   				},
//   				TagMap: map[string]*string{
//   					"tagMapKey": jsii.String("tagMap"),
//   				},
//   			},
//   			Filter: &FilterProperty{
//   				RetainAtLeast: jsii.Number(123),
//   				Type: jsii.String("type"),
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ResourceSelection: &ResourceSelectionProperty{
//   		Recipes: []interface{}{
//   			&RecipeSelectionProperty{
//   				Name: jsii.String("name"),
//   				SemanticVersion: jsii.String("semanticVersion"),
//   			},
//   		},
//   		TagMap: map[string]*string{
//   			"tagMapKey": jsii.String("tagMap"),
//   		},
//   	},
//   	ResourceType: jsii.String("resourceType"),
//   	Status: jsii.String("status"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-lifecyclepolicy.html
//
type CfnLifecyclePolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLifecyclePolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLifecyclePolicyPropsMixin
type jsiiProxy_CfnLifecyclePolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLifecyclePolicyPropsMixin) Props() *CfnLifecyclePolicyMixinProps {
	var returns *CfnLifecyclePolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ImageBuilder::LifecyclePolicy`.
func NewCfnLifecyclePolicyPropsMixin(props *CfnLifecyclePolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLifecyclePolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLifecyclePolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLifecyclePolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ImageBuilder::LifecyclePolicy`.
func NewCfnLifecyclePolicyPropsMixin_Override(c CfnLifecyclePolicyPropsMixin, props *CfnLifecyclePolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLifecyclePolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLifecyclePolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLifecyclePolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLifecyclePolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLifecyclePolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

