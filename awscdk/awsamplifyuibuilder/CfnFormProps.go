package awsamplifyuibuilder


// Properties for defining a `CfnForm`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFormProps := &CfnFormProps{
//   	DataType: &FormDataTypeConfigProperty{
//   		DataSourceType: jsii.String("dataSourceType"),
//   		DataTypeName: jsii.String("dataTypeName"),
//   	},
//   	Fields: map[string]interface{}{
//   		"fieldsKey": &FieldConfigProperty{
//   			"excluded": jsii.Boolean(false),
//   			"inputType": &FieldInputConfigProperty{
//   				"type": jsii.String("type"),
//
//   				// the properties below are optional
//   				"defaultChecked": jsii.Boolean(false),
//   				"defaultCountryCode": jsii.String("defaultCountryCode"),
//   				"defaultValue": jsii.String("defaultValue"),
//   				"descriptiveText": jsii.String("descriptiveText"),
//   				"fileUploaderConfig": &FileUploaderFieldConfigProperty{
//   					"acceptedFileTypes": []*string{
//   						jsii.String("acceptedFileTypes"),
//   					},
//   					"accessLevel": jsii.String("accessLevel"),
//
//   					// the properties below are optional
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
//   				"value": jsii.String("value"),
//   				"valueMappings": &ValueMappingsProperty{
//   					"values": []interface{}{
//   						&ValueMappingProperty{
//   							"value": &FormInputValuePropertyProperty{
//   								"value": jsii.String("value"),
//   							},
//
//   							// the properties below are optional
//   							"displayValue": &FormInputValuePropertyProperty{
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
//   					"type": jsii.String("type"),
//
//   					// the properties below are optional
//   					"numValues": []interface{}{
//   						jsii.Number(123),
//   					},
//   					"strValues": []*string{
//   						jsii.String("strValues"),
//   					},
//   					"validationMessage": jsii.String("validationMessage"),
//   				},
//   			},
//   		},
//   	},
//   	FormActionType: jsii.String("formActionType"),
//   	Name: jsii.String("name"),
//   	SchemaVersion: jsii.String("schemaVersion"),
//   	SectionalElements: map[string]interface{}{
//   		"sectionalElementsKey": &SectionalElementProperty{
//   			"type": jsii.String("type"),
//
//   			// the properties below are optional
//   			"excluded": jsii.Boolean(false),
//   			"level": jsii.Number(123),
//   			"orientation": jsii.String("orientation"),
//   			"position": &FieldPositionProperty{
//   				"below": jsii.String("below"),
//   				"fixed": jsii.String("fixed"),
//   				"rightOf": jsii.String("rightOf"),
//   			},
//   			"text": jsii.String("text"),
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
//
//   	// the properties below are optional
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
//   	EnvironmentName: jsii.String("environmentName"),
//   	LabelDecorator: jsii.String("labelDecorator"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnFormProps struct {
	// The type of data source to use to create the form.
	DataType interface{} `field:"required" json:"dataType" yaml:"dataType"`
	// The configuration information for the form's fields.
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// Specifies whether to perform a create or update action on the form.
	FormActionType *string `field:"required" json:"formActionType" yaml:"formActionType"`
	// The name of the form.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema version of the form.
	SchemaVersion *string `field:"required" json:"schemaVersion" yaml:"schemaVersion"`
	// The configuration information for the visual helper elements for the form.
	//
	// These elements are not associated with any data.
	SectionalElements interface{} `field:"required" json:"sectionalElements" yaml:"sectionalElements"`
	// The configuration for the form's style.
	Style interface{} `field:"required" json:"style" yaml:"style"`
	// The unique ID of the Amplify app associated with the form.
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// The `FormCTA` object that stores the call to action configuration for the form.
	Cta interface{} `field:"optional" json:"cta" yaml:"cta"`
	// The name of the backend environment that is a part of the Amplify app.
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// Specifies an icon or decoration to display on the form.
	LabelDecorator *string `field:"optional" json:"labelDecorator" yaml:"labelDecorator"`
	// One or more key-value pairs to use when tagging the form data.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

