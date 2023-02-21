package awsamplifyuibuilder


// Represents the data binding configuration for a value map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   valueMappingsProperty := &valueMappingsProperty{
//   	values: []interface{}{
//   		&valueMappingProperty{
//   			value: &formInputValuePropertyProperty{
//   				value: jsii.String("value"),
//   			},
//
//   			// the properties below are optional
//   			displayValue: &formInputValuePropertyProperty{
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnForm_ValueMappingsProperty struct {
	// The value and display value pairs.
	Values interface{} `field:"required" json:"values" yaml:"values"`
}

