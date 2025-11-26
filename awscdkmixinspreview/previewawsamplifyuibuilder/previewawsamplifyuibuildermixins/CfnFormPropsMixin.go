package previewawsamplifyuibuildermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsamplifyuibuilder/previewawsamplifyuibuildermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::AmplifyUIBuilder::Form resource specifies all of the information that is required to create a form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var formInputValuePropertyProperty_ FormInputValuePropertyProperty
//
//   cfnFormPropsMixin := awscdkmixinspreview.Mixins.NewCfnFormPropsMixin(&CfnFormMixinProps{
//   	AppId: jsii.String("appId"),
//   	Cta: &FormCTAProperty{
//   		Cancel: &FormButtonProperty{
//   			Children: jsii.String("children"),
//   			Excluded: jsii.Boolean(false),
//   			Position: &FieldPositionProperty{
//   				Below: jsii.String("below"),
//   				Fixed: jsii.String("fixed"),
//   				RightOf: jsii.String("rightOf"),
//   			},
//   		},
//   		Clear: &FormButtonProperty{
//   			Children: jsii.String("children"),
//   			Excluded: jsii.Boolean(false),
//   			Position: &FieldPositionProperty{
//   				Below: jsii.String("below"),
//   				Fixed: jsii.String("fixed"),
//   				RightOf: jsii.String("rightOf"),
//   			},
//   		},
//   		Position: jsii.String("position"),
//   		Submit: &FormButtonProperty{
//   			Children: jsii.String("children"),
//   			Excluded: jsii.Boolean(false),
//   			Position: &FieldPositionProperty{
//   				Below: jsii.String("below"),
//   				Fixed: jsii.String("fixed"),
//   				RightOf: jsii.String("rightOf"),
//   			},
//   		},
//   	},
//   	DataType: &FormDataTypeConfigProperty{
//   		DataSourceType: jsii.String("dataSourceType"),
//   		DataTypeName: jsii.String("dataTypeName"),
//   	},
//   	EnvironmentName: jsii.String("environmentName"),
//   	Fields: map[string]interface{}{
//   		"fieldsKey": &FieldConfigProperty{
//   			"excluded": jsii.Boolean(false),
//   			"inputType": &FieldInputConfigProperty{
//   				"defaultChecked": jsii.Boolean(false),
//   				"defaultCountryCode": jsii.String("defaultCountryCode"),
//   				"defaultValue": jsii.String("defaultValue"),
//   				"descriptiveText": jsii.String("descriptiveText"),
//   				"fileUploaderConfig": &FileUploaderFieldConfigProperty{
//   					"acceptedFileTypes": []*string{
//   						jsii.String("acceptedFileTypes"),
//   					},
//   					"accessLevel": jsii.String("accessLevel"),
//   					"isResumable": jsii.Boolean(false),
//   					"maxFileCount": jsii.Number(123),
//   					"maxSize": jsii.Number(123),
//   					"showThumbnails": jsii.Boolean(false),
//   				},
//   				"isArray": jsii.Boolean(false),
//   				"maxValue": jsii.Number(123),
//   				"minValue": jsii.Number(123),
//   				"name": jsii.String("name"),
//   				"placeholder": jsii.String("placeholder"),
//   				"readOnly": jsii.Boolean(false),
//   				"required": jsii.Boolean(false),
//   				"step": jsii.Number(123),
//   				"type": jsii.String("type"),
//   				"value": jsii.String("value"),
//   				"valueMappings": &ValueMappingsProperty{
//   					"bindingProperties": map[string]interface{}{
//   						"bindingPropertiesKey": &FormInputBindingPropertiesValueProperty{
//   							"bindingProperties": &FormInputBindingPropertiesValuePropertiesProperty{
//   								"model": jsii.String("model"),
//   							},
//   							"type": jsii.String("type"),
//   						},
//   					},
//   					"values": []interface{}{
//   						&ValueMappingProperty{
//   							"displayValue": &FormInputValuePropertyProperty{
//   								"bindingProperties": &FormInputValuePropertyBindingPropertiesProperty{
//   									"field": jsii.String("field"),
//   									"property": jsii.String("property"),
//   								},
//   								"concat": []interface{}{
//   									formInputValuePropertyProperty_,
//   								},
//   								"value": jsii.String("value"),
//   							},
//   							"value": &FormInputValuePropertyProperty{
//   								"bindingProperties": &FormInputValuePropertyBindingPropertiesProperty{
//   									"field": jsii.String("field"),
//   									"property": jsii.String("property"),
//   								},
//   								"concat": []interface{}{
//   									formInputValuePropertyProperty_,
//   								},
//   								"value": jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			"label": jsii.String("label"),
//   			"position": &FieldPositionProperty{
//   				"below": jsii.String("below"),
//   				"fixed": jsii.String("fixed"),
//   				"rightOf": jsii.String("rightOf"),
//   			},
//   			"validations": []interface{}{
//   				&FieldValidationConfigurationProperty{
//   					"numValues": []interface{}{
//   						jsii.Number(123),
//   					},
//   					"strValues": []*string{
//   						jsii.String("strValues"),
//   					},
//   					"type": jsii.String("type"),
//   					"validationMessage": jsii.String("validationMessage"),
//   				},
//   			},
//   		},
//   	},
//   	FormActionType: jsii.String("formActionType"),
//   	LabelDecorator: jsii.String("labelDecorator"),
//   	Name: jsii.String("name"),
//   	SchemaVersion: jsii.String("schemaVersion"),
//   	SectionalElements: map[string]interface{}{
//   		"sectionalElementsKey": &SectionalElementProperty{
//   			"excluded": jsii.Boolean(false),
//   			"level": jsii.Number(123),
//   			"orientation": jsii.String("orientation"),
//   			"position": &FieldPositionProperty{
//   				"below": jsii.String("below"),
//   				"fixed": jsii.String("fixed"),
//   				"rightOf": jsii.String("rightOf"),
//   			},
//   			"text": jsii.String("text"),
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	Style: &FormStyleProperty{
//   		HorizontalGap: &FormStyleConfigProperty{
//   			TokenReference: jsii.String("tokenReference"),
//   			Value: jsii.String("value"),
//   		},
//   		OuterPadding: &FormStyleConfigProperty{
//   			TokenReference: jsii.String("tokenReference"),
//   			Value: jsii.String("value"),
//   		},
//   		VerticalGap: &FormStyleConfigProperty{
//   			TokenReference: jsii.String("tokenReference"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html
//
type CfnFormPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFormMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFormPropsMixin
type jsiiProxy_CfnFormPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFormPropsMixin) Props() *CfnFormMixinProps {
	var returns *CfnFormMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFormPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AmplifyUIBuilder::Form`.
func NewCfnFormPropsMixin(props *CfnFormMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFormPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFormPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFormPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_amplifyuibuilder.mixins.CfnFormPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AmplifyUIBuilder::Form`.
func NewCfnFormPropsMixin_Override(c CfnFormPropsMixin, props *CfnFormMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_amplifyuibuilder.mixins.CfnFormPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFormPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFormPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_amplifyuibuilder.mixins.CfnFormPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFormPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_amplifyuibuilder.mixins.CfnFormPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFormPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnFormPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

