package awsamplifyuibuilder


// Describes the configuration information for a field in a table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldConfigProperty := &fieldConfigProperty{
//   	excluded: jsii.Boolean(false),
//   	inputType: &fieldInputConfigProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		defaultChecked: jsii.Boolean(false),
//   		defaultCountryCode: jsii.String("defaultCountryCode"),
//   		defaultValue: jsii.String("defaultValue"),
//   		descriptiveText: jsii.String("descriptiveText"),
//   		isArray: jsii.Boolean(false),
//   		maxValue: jsii.Number(123),
//   		minValue: jsii.Number(123),
//   		name: jsii.String("name"),
//   		placeholder: jsii.String("placeholder"),
//   		readOnly: jsii.Boolean(false),
//   		required: jsii.Boolean(false),
//   		step: jsii.Number(123),
//   		value: jsii.String("value"),
//   		valueMappings: &valueMappingsProperty{
//   			values: []interface{}{
//   				&valueMappingProperty{
//   					value: &formInputValuePropertyProperty{
//   						value: jsii.String("value"),
//   					},
//
//   					// the properties below are optional
//   					displayValue: &formInputValuePropertyProperty{
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	label: jsii.String("label"),
//   	position: &fieldPositionProperty{
//   		below: jsii.String("below"),
//   		fixed: jsii.String("fixed"),
//   		rightOf: jsii.String("rightOf"),
//   	},
//   	validations: []interface{}{
//   		&fieldValidationConfigurationProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			numValues: []interface{}{
//   				jsii.Number(123),
//   			},
//   			strValues: []*string{
//   				jsii.String("strValues"),
//   			},
//   			validationMessage: jsii.String("validationMessage"),
//   		},
//   	},
//   }
//
type CfnForm_FieldConfigProperty struct {
	// Specifies whether to hide a field.
	Excluded interface{} `field:"optional" json:"excluded" yaml:"excluded"`
	// Describes the configuration for the default input value to display for a field.
	InputType interface{} `field:"optional" json:"inputType" yaml:"inputType"`
	// The label for the field.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Specifies the field position.
	Position interface{} `field:"optional" json:"position" yaml:"position"`
	// The validations to perform on the value in the field.
	Validations interface{} `field:"optional" json:"validations" yaml:"validations"`
}

