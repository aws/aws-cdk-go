package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an AWS Glue usage profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnUsageProfilePropsMixin := awscdkcfnpropertymixins.Aws_glue.NewCfnUsageProfilePropsMixin(&CfnUsageProfileMixinProps{
//   	Configuration: &ProfileConfigurationProperty{
//   		JobConfiguration: map[string]interface{}{
//   			"jobConfigurationKey": &ConfigurationObjectProperty{
//   				"allowedValues": []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				"defaultValue": jsii.String("defaultValue"),
//   				"maxValue": jsii.String("maxValue"),
//   				"minValue": jsii.String("minValue"),
//   			},
//   		},
//   		SessionConfiguration: map[string]interface{}{
//   			"sessionConfigurationKey": &ConfigurationObjectProperty{
//   				"allowedValues": []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				"defaultValue": jsii.String("defaultValue"),
//   				"maxValue": jsii.String("maxValue"),
//   				"minValue": jsii.String("minValue"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-usageprofile.html
//
type CfnUsageProfilePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnUsageProfileMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUsageProfilePropsMixin
type jsiiProxy_CfnUsageProfilePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnUsageProfilePropsMixin) Props() *CfnUsageProfileMixinProps {
	var returns *CfnUsageProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUsageProfilePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::UsageProfile`.
func NewCfnUsageProfilePropsMixin(props *CfnUsageProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnUsageProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUsageProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUsageProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnUsageProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::UsageProfile`.
func NewCfnUsageProfilePropsMixin_Override(c CfnUsageProfilePropsMixin, props *CfnUsageProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnUsageProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnUsageProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUsageProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnUsageProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUsageProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnUsageProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUsageProfilePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnUsageProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

