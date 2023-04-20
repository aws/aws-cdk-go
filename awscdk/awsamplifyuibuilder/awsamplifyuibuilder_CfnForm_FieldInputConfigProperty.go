package awsamplifyuibuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldInputConfigProperty := &FieldInputConfigProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	DefaultChecked: jsii.Boolean(false),
//   	DefaultCountryCode: jsii.String("defaultCountryCode"),
//   	DefaultValue: jsii.String("defaultValue"),
//   	DescriptiveText: jsii.String("descriptiveText"),
//   	IsArray: jsii.Boolean(false),
//   	MaxValue: jsii.Number(123),
//   	MinValue: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Placeholder: jsii.String("placeholder"),
//   	ReadOnly: jsii.Boolean(false),
//   	Required: jsii.Boolean(false),
//   	Step: jsii.Number(123),
//   	Value: jsii.String("value"),
//   	ValueMappings: &ValueMappingsProperty{
//   		Values: []interface{}{
//   			&ValueMappingProperty{
//   				Value: &FormInputValuePropertyProperty{
//   					Value: jsii.String("value"),
//   				},
//
//   				// the properties below are optional
//   				DisplayValue: &FormInputValuePropertyProperty{
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnForm_FieldInputConfigProperty struct {
	// `CfnForm.FieldInputConfigProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnForm.FieldInputConfigProperty.DefaultChecked`.
	DefaultChecked interface{} `field:"optional" json:"defaultChecked" yaml:"defaultChecked"`
	// `CfnForm.FieldInputConfigProperty.DefaultCountryCode`.
	DefaultCountryCode *string `field:"optional" json:"defaultCountryCode" yaml:"defaultCountryCode"`
	// `CfnForm.FieldInputConfigProperty.DefaultValue`.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// `CfnForm.FieldInputConfigProperty.DescriptiveText`.
	DescriptiveText *string `field:"optional" json:"descriptiveText" yaml:"descriptiveText"`
	// `CfnForm.FieldInputConfigProperty.IsArray`.
	IsArray interface{} `field:"optional" json:"isArray" yaml:"isArray"`
	// `CfnForm.FieldInputConfigProperty.MaxValue`.
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// `CfnForm.FieldInputConfigProperty.MinValue`.
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
	// `CfnForm.FieldInputConfigProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnForm.FieldInputConfigProperty.Placeholder`.
	Placeholder *string `field:"optional" json:"placeholder" yaml:"placeholder"`
	// `CfnForm.FieldInputConfigProperty.ReadOnly`.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// `CfnForm.FieldInputConfigProperty.Required`.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// `CfnForm.FieldInputConfigProperty.Step`.
	Step *float64 `field:"optional" json:"step" yaml:"step"`
	// `CfnForm.FieldInputConfigProperty.Value`.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// `CfnForm.FieldInputConfigProperty.ValueMappings`.
	ValueMappings interface{} `field:"optional" json:"valueMappings" yaml:"valueMappings"`
}

