package awsamplifyuibuilder


// Describes the configuration for the default input values to display for a field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldInputConfigProperty := &fieldInputConfigProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	defaultChecked: jsii.Boolean(false),
//   	defaultCountryCode: jsii.String("defaultCountryCode"),
//   	defaultValue: jsii.String("defaultValue"),
//   	descriptiveText: jsii.String("descriptiveText"),
//   	maxValue: jsii.Number(123),
//   	minValue: jsii.Number(123),
//   	name: jsii.String("name"),
//   	placeholder: jsii.String("placeholder"),
//   	readOnly: jsii.Boolean(false),
//   	required: jsii.Boolean(false),
//   	step: jsii.Number(123),
//   	value: jsii.String("value"),
//   	valueMappings: &valueMappingsProperty{
//   		values: []interface{}{
//   			&valueMappingProperty{
//   				value: &formInputValuePropertyProperty{
//   					value: jsii.String("value"),
//   				},
//
//   				// the properties below are optional
//   				displayValue: &formInputValuePropertyProperty{
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnForm_FieldInputConfigProperty struct {
	// The input type for the field.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies whether a field has a default value.
	DefaultChecked interface{} `field:"optional" json:"defaultChecked" yaml:"defaultChecked"`
	// The default country code for a phone number.
	DefaultCountryCode *string `field:"optional" json:"defaultCountryCode" yaml:"defaultCountryCode"`
	// The default value for the field.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The text to display to describe the field.
	DescriptiveText *string `field:"optional" json:"descriptiveText" yaml:"descriptiveText"`
	// The maximum value to display for the field.
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// The minimum value to display for the field.
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
	// The name of the field.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The text to display as a placeholder for the field.
	Placeholder *string `field:"optional" json:"placeholder" yaml:"placeholder"`
	// Specifies a read only field.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Specifies a field that requires input.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// The stepping increment for a numeric value in a field.
	Step *float64 `field:"optional" json:"step" yaml:"step"`
	// The value for the field.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// The information to use to customize the input fields with data at runtime.
	ValueMappings interface{} `field:"optional" json:"valueMappings" yaml:"valueMappings"`
}

