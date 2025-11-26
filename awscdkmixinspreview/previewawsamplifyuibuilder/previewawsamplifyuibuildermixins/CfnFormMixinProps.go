package previewawsamplifyuibuildermixins


// Properties for CfnFormPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var formInputValuePropertyProperty_ FormInputValuePropertyProperty
//
//   cfnFormMixinProps := &CfnFormMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html
//
type CfnFormMixinProps struct {
	// The unique ID of the Amplify app associated with the form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html#cfn-amplifyuibuilder-form-appid
	//
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// The `FormCTA` object that stores the call to action configuration for the form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html#cfn-amplifyuibuilder-form-cta
	//
	Cta interface{} `field:"optional" json:"cta" yaml:"cta"`
	// The type of data source to use to create the form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html#cfn-amplifyuibuilder-form-datatype
	//
	DataType interface{} `field:"optional" json:"dataType" yaml:"dataType"`
	// The name of the backend environment that is a part of the Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html#cfn-amplifyuibuilder-form-environmentname
	//
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// The configuration information for the form's fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html#cfn-amplifyuibuilder-form-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// Specifies whether to perform a create or update action on the form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html#cfn-amplifyuibuilder-form-formactiontype
	//
	FormActionType *string `field:"optional" json:"formActionType" yaml:"formActionType"`
	// Specifies an icon or decoration to display on the form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html#cfn-amplifyuibuilder-form-labeldecorator
	//
	LabelDecorator *string `field:"optional" json:"labelDecorator" yaml:"labelDecorator"`
	// The name of the form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html#cfn-amplifyuibuilder-form-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The schema version of the form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html#cfn-amplifyuibuilder-form-schemaversion
	//
	SchemaVersion *string `field:"optional" json:"schemaVersion" yaml:"schemaVersion"`
	// The configuration information for the visual helper elements for the form.
	//
	// These elements are not associated with any data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html#cfn-amplifyuibuilder-form-sectionalelements
	//
	SectionalElements interface{} `field:"optional" json:"sectionalElements" yaml:"sectionalElements"`
	// The configuration for the form's style.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html#cfn-amplifyuibuilder-form-style
	//
	Style interface{} `field:"optional" json:"style" yaml:"style"`
	// One or more key-value pairs to use when tagging the form data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-form.html#cfn-amplifyuibuilder-form-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

