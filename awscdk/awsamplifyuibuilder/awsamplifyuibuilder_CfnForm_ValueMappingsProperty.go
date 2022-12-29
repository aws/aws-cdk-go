package awsamplifyuibuilder


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
	// `CfnForm.ValueMappingsProperty.Values`.
	Values interface{} `field:"required" json:"values" yaml:"values"`
}

