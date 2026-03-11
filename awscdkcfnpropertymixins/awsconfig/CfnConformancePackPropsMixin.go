package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A conformance pack is a collection of AWS Config rules and remediation actions that can be easily deployed in an account and a region.
//
// ConformancePack creates a service linked role in your account. The service linked role is created only when the role does not exist in your account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var templateSsmDocumentDetails interface{}
//
//   cfnConformancePackPropsMixin := awscdkcfnpropertymixins.Aws_config.NewCfnConformancePackPropsMixin(&CfnConformancePackMixinProps{
//   	ConformancePackInputParameters: []interface{}{
//   		&ConformancePackInputParameterProperty{
//   			ParameterName: jsii.String("parameterName"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	ConformancePackName: jsii.String("conformancePackName"),
//   	DeliveryS3Bucket: jsii.String("deliveryS3Bucket"),
//   	DeliveryS3KeyPrefix: jsii.String("deliveryS3KeyPrefix"),
//   	TemplateBody: jsii.String("templateBody"),
//   	TemplateS3Uri: jsii.String("templateS3Uri"),
//   	TemplateSsmDocumentDetails: templateSsmDocumentDetails,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html
//
type CfnConformancePackPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnConformancePackMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConformancePackPropsMixin
type jsiiProxy_CfnConformancePackPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnConformancePackPropsMixin) Props() *CfnConformancePackMixinProps {
	var returns *CfnConformancePackMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePackPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Config::ConformancePack`.
func NewCfnConformancePackPropsMixin(props *CfnConformancePackMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnConformancePackPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConformancePackPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConformancePackPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConformancePackPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Config::ConformancePack`.
func NewCfnConformancePackPropsMixin_Override(c CfnConformancePackPropsMixin, props *CfnConformancePackMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConformancePackPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnConformancePackPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConformancePackPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConformancePackPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConformancePackPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnConformancePackPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConformancePackPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnConformancePackPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

