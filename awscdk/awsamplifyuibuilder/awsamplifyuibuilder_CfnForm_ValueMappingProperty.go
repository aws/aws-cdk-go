package awsamplifyuibuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   valueMappingProperty := &valueMappingProperty{
//   	value: &formInputValuePropertyProperty{
//   		value: jsii.String("value"),
//   	},
//
//   	// the properties below are optional
//   	displayValue: &formInputValuePropertyProperty{
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnForm_ValueMappingProperty struct {
	// `CfnForm.ValueMappingProperty.Value`.
	Value interface{} `field:"required" json:"value" yaml:"value"`
	// `CfnForm.ValueMappingProperty.DisplayValue`.
	DisplayValue interface{} `field:"optional" json:"displayValue" yaml:"displayValue"`
}

