package awsamplifyuibuilder


// Associates a complex object with a display value.
//
// Use `ValueMapping` to store how to represent complex objects when they are displayed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   valueMappingProperty := &ValueMappingProperty{
//   	Value: &FormInputValuePropertyProperty{
//   		Value: jsii.String("value"),
//   	},
//
//   	// the properties below are optional
//   	DisplayValue: &FormInputValuePropertyProperty{
//   		Value: jsii.String("value"),
//   	},
//   }
//
type CfnForm_ValueMappingProperty struct {
	// The complex object.
	Value interface{} `field:"required" json:"value" yaml:"value"`
	// The value to display for the complex object.
	DisplayValue interface{} `field:"optional" json:"displayValue" yaml:"displayValue"`
}

