package awsamplifyuibuilder


// The `FieldConfig` property specifies the configuration information for a field in a table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var formInputValuePropertyProperty_ formInputValuePropertyProperty
//
//   fieldConfigProperty := &FieldConfigProperty{
//   	Excluded: jsii.Boolean(false),
//   	InputType: &FieldInputConfigProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		DefaultChecked: jsii.Boolean(false),
//   		DefaultCountryCode: jsii.String("defaultCountryCode"),
//   		DefaultValue: jsii.String("defaultValue"),
//   		DescriptiveText: jsii.String("descriptiveText"),
//   		FileUploaderConfig: &FileUploaderFieldConfigProperty{
//   			AcceptedFileTypes: []*string{
//   				jsii.String("acceptedFileTypes"),
//   			},
//   			AccessLevel: jsii.String("accessLevel"),
//
//   			// the properties below are optional
//   			IsResumable: jsii.Boolean(false),
//   			MaxFileCount: jsii.Number(123),
//   			MaxSize: jsii.Number(123),
//   			ShowThumbnails: jsii.Boolean(false),
//   		},
//   		IsArray: jsii.Boolean(false),
//   		MaxValue: jsii.Number(123),
//   		MinValue: jsii.Number(123),
//   		Name: jsii.String("name"),
//   		Placeholder: jsii.String("placeholder"),
//   		ReadOnly: jsii.Boolean(false),
//   		Required: jsii.Boolean(false),
//   		Step: jsii.Number(123),
//   		Value: jsii.String("value"),
//   		ValueMappings: &ValueMappingsProperty{
//   			Values: []interface{}{
//   				&ValueMappingProperty{
//   					Value: &formInputValuePropertyProperty{
//   						BindingProperties: &FormInputValuePropertyBindingPropertiesProperty{
//   							Property: jsii.String("property"),
//
//   							// the properties below are optional
//   							Field: jsii.String("field"),
//   						},
//   						Concat: []interface{}{
//   							formInputValuePropertyProperty_,
//   						},
//   						Value: jsii.String("value"),
//   					},
//
//   					// the properties below are optional
//   					DisplayValue: &formInputValuePropertyProperty{
//   						BindingProperties: &FormInputValuePropertyBindingPropertiesProperty{
//   							Property: jsii.String("property"),
//
//   							// the properties below are optional
//   							Field: jsii.String("field"),
//   						},
//   						Concat: []interface{}{
//   							formInputValuePropertyProperty_,
//   						},
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			BindingProperties: map[string]interface{}{
//   				"bindingPropertiesKey": &FormInputBindingPropertiesValueProperty{
//   					"bindingProperties": &FormInputBindingPropertiesValuePropertiesProperty{
//   						"model": jsii.String("model"),
//   					},
//   					"type": jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   	Label: jsii.String("label"),
//   	Position: &FieldPositionProperty{
//   		Below: jsii.String("below"),
//   		Fixed: jsii.String("fixed"),
//   		RightOf: jsii.String("rightOf"),
//   	},
//   	Validations: []interface{}{
//   		&FieldValidationConfigurationProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			NumValues: []interface{}{
//   				jsii.Number(123),
//   			},
//   			StrValues: []*string{
//   				jsii.String("strValues"),
//   			},
//   			ValidationMessage: jsii.String("validationMessage"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldconfig.html
//
type CfnForm_FieldConfigProperty struct {
	// Specifies whether to hide a field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldconfig.html#cfn-amplifyuibuilder-form-fieldconfig-excluded
	//
	Excluded interface{} `field:"optional" json:"excluded" yaml:"excluded"`
	// Describes the configuration for the default input value to display for a field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldconfig.html#cfn-amplifyuibuilder-form-fieldconfig-inputtype
	//
	InputType interface{} `field:"optional" json:"inputType" yaml:"inputType"`
	// The label for the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldconfig.html#cfn-amplifyuibuilder-form-fieldconfig-label
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Specifies the field position.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldconfig.html#cfn-amplifyuibuilder-form-fieldconfig-position
	//
	Position interface{} `field:"optional" json:"position" yaml:"position"`
	// The validations to perform on the value in the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldconfig.html#cfn-amplifyuibuilder-form-fieldconfig-validations
	//
	Validations interface{} `field:"optional" json:"validations" yaml:"validations"`
}

