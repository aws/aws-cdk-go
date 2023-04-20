package awsamplifyuibuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   					Value: &FormInputValuePropertyProperty{
//   						Value: jsii.String("value"),
//   					},
//
//   					// the properties below are optional
//   					DisplayValue: &FormInputValuePropertyProperty{
//   						Value: jsii.String("value"),
//   					},
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
type CfnForm_FieldConfigProperty struct {
	// `CfnForm.FieldConfigProperty.Excluded`.
	Excluded interface{} `field:"optional" json:"excluded" yaml:"excluded"`
	// `CfnForm.FieldConfigProperty.InputType`.
	InputType interface{} `field:"optional" json:"inputType" yaml:"inputType"`
	// `CfnForm.FieldConfigProperty.Label`.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// `CfnForm.FieldConfigProperty.Position`.
	Position interface{} `field:"optional" json:"position" yaml:"position"`
	// `CfnForm.FieldConfigProperty.Validations`.
	Validations interface{} `field:"optional" json:"validations" yaml:"validations"`
}

