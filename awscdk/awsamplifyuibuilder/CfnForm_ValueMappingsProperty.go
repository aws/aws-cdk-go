package awsamplifyuibuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   valueMappingsProperty := &ValueMappingsProperty{
//   	Values: []interface{}{
//   		&ValueMappingProperty{
//   			Value: &FormInputValuePropertyProperty{
//   				Value: jsii.String("value"),
//   			},
//
//   			// the properties below are optional
//   			DisplayValue: &FormInputValuePropertyProperty{
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnForm_ValueMappingsProperty struct {
	// `CfnForm.ValueMappingsProperty.Values`.
	Values interface{} `field:"required" json:"values" yaml:"values"`
}

