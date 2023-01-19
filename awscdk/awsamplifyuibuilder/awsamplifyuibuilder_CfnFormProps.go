package awsamplifyuibuilder


// Properties for defining a `CfnForm`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFormProps := &cfnFormProps{
//   	dataType: &formDataTypeConfigProperty{
//   		dataSourceType: jsii.String("dataSourceType"),
//   		dataTypeName: jsii.String("dataTypeName"),
//   	},
//   	fields: map[string]interface{}{
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
//   	formActionType: jsii.String("formActionType"),
//   	name: jsii.String("name"),
//   	schemaVersion: jsii.String("schemaVersion"),
//   	sectionalElements: map[string]interface{}{
//   		"sectionalElementsKey": &SectionalElementProperty{
//   			"type": jsii.String("type"),
//
//   			// the properties below are optional
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
//   	style: &formStyleProperty{
//   		horizontalGap: &formStyleConfigProperty{
//   			tokenReference: jsii.String("tokenReference"),
//   			value: jsii.String("value"),
//   		},
//   		outerPadding: &formStyleConfigProperty{
//   			tokenReference: jsii.String("tokenReference"),
//   			value: jsii.String("value"),
//   		},
//   		verticalGap: &formStyleConfigProperty{
//   			tokenReference: jsii.String("tokenReference"),
//   			value: jsii.String("value"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	appId: jsii.String("appId"),
//   	cta: &formCTAProperty{
//   		cancel: &formButtonProperty{
//   			children: jsii.String("children"),
//   			excluded: jsii.Boolean(false),
//   			position: &fieldPositionProperty{
//   				below: jsii.String("below"),
//   				fixed: jsii.String("fixed"),
//   				rightOf: jsii.String("rightOf"),
//   			},
//   		},
//   		clear: &formButtonProperty{
//   			children: jsii.String("children"),
//   			excluded: jsii.Boolean(false),
//   			position: &fieldPositionProperty{
//   				below: jsii.String("below"),
//   				fixed: jsii.String("fixed"),
//   				rightOf: jsii.String("rightOf"),
//   			},
//   		},
//   		position: jsii.String("position"),
//   		submit: &formButtonProperty{
//   			children: jsii.String("children"),
//   			excluded: jsii.Boolean(false),
//   			position: &fieldPositionProperty{
//   				below: jsii.String("below"),
//   				fixed: jsii.String("fixed"),
//   				rightOf: jsii.String("rightOf"),
//   			},
//   		},
//   	},
//   	environmentName: jsii.String("environmentName"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnFormProps struct {
	// The type of data source to use to create the form.
	DataType interface{} `field:"required" json:"dataType" yaml:"dataType"`
	// Stores the information about the form's fields.
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// The operation to perform on the specified form.
	FormActionType *string `field:"required" json:"formActionType" yaml:"formActionType"`
	// The name of the form.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema version of the form when it was imported.
	SchemaVersion *string `field:"required" json:"schemaVersion" yaml:"schemaVersion"`
	// Stores the visual helper elements for the form that are not associated with any data.
	SectionalElements interface{} `field:"required" json:"sectionalElements" yaml:"sectionalElements"`
	// Stores the configuration for the form's style.
	Style interface{} `field:"required" json:"style" yaml:"style"`
	// The unique ID of the Amplify app associated with the form.
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// Stores the call to action configuration for the form.
	Cta interface{} `field:"optional" json:"cta" yaml:"cta"`
	// The name of the backend environment that is a part of the Amplify app.
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// One or more key-value pairs to use when tagging the form.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

