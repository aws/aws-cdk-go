package awsdatabrew

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdatabrew/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a new AWS Glue DataBrew transformation recipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnRecipePropsMixin := awscdkcfnpropertymixins.Aws_databrew.NewCfnRecipePropsMixin(&CfnRecipeMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Steps: []interface{}{
//   		&RecipeStepProperty{
//   			Action: &ActionProperty{
//   				Operation: jsii.String("operation"),
//   				Parameters: map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   			},
//   			ConditionExpressions: []interface{}{
//   				&ConditionExpressionProperty{
//   					Condition: jsii.String("condition"),
//   					TargetColumn: jsii.String("targetColumn"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-recipe.html
//
type CfnRecipePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRecipeMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRecipePropsMixin
type jsiiProxy_CfnRecipePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRecipePropsMixin) Props() *CfnRecipeMixinProps {
	var returns *CfnRecipeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataBrew::Recipe`.
func NewCfnRecipePropsMixin(props *CfnRecipeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRecipePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRecipePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRecipePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRecipePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataBrew::Recipe`.
func NewCfnRecipePropsMixin_Override(c CfnRecipePropsMixin, props *CfnRecipeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRecipePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRecipePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRecipePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRecipePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRecipePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRecipePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRecipePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRecipePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

