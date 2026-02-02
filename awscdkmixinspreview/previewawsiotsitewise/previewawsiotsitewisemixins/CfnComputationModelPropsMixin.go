package previewawsiotsitewisemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiotsitewise/previewawsiotsitewisemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a computation model with a configuration and data binding.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var computationModelDataBindingValueProperty_ ComputationModelDataBindingValueProperty
//
//   cfnComputationModelPropsMixin := awscdkmixinspreview.Mixins.NewCfnComputationModelPropsMixin(&CfnComputationModelMixinProps{
//   	ComputationModelConfiguration: &ComputationModelConfigurationProperty{
//   		AnomalyDetection: &AnomalyDetectionComputationModelConfigurationProperty{
//   			InputProperties: jsii.String("inputProperties"),
//   			ResultProperty: jsii.String("resultProperty"),
//   		},
//   	},
//   	ComputationModelDataBinding: map[string]interface{}{
//   		"computationModelDataBindingKey": &ComputationModelDataBindingValueProperty{
//   			"assetModelProperty": &AssetModelPropertyBindingValueProperty{
//   				"assetModelId": jsii.String("assetModelId"),
//   				"propertyId": jsii.String("propertyId"),
//   			},
//   			"assetProperty": &AssetPropertyBindingValueProperty{
//   				"assetId": jsii.String("assetId"),
//   				"propertyId": jsii.String("propertyId"),
//   			},
//   			"list": []interface{}{
//   				computationModelDataBindingValueProperty_,
//   			},
//   		},
//   	},
//   	ComputationModelDescription: jsii.String("computationModelDescription"),
//   	ComputationModelName: jsii.String("computationModelName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-computationmodel.html
//
type CfnComputationModelPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnComputationModelMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnComputationModelPropsMixin
type jsiiProxy_CfnComputationModelPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnComputationModelPropsMixin) Props() *CfnComputationModelMixinProps {
	var returns *CfnComputationModelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComputationModelPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTSiteWise::ComputationModel`.
func NewCfnComputationModelPropsMixin(props *CfnComputationModelMixinProps, options *mixins.CfnPropertyMixinOptions) CfnComputationModelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnComputationModelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnComputationModelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnComputationModelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTSiteWise::ComputationModel`.
func NewCfnComputationModelPropsMixin_Override(c CfnComputationModelPropsMixin, props *CfnComputationModelMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnComputationModelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnComputationModelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnComputationModelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnComputationModelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComputationModelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnComputationModelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnComputationModelPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnComputationModelPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

