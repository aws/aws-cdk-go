package awsamplifyuibuilder


// The `ValueMappings` property specifies the data binding configuration for a value map.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-valuemappings.html
//
type CfnForm_ValueMappingsProperty struct {
	// The value and display value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-valuemappings.html#cfn-amplifyuibuilder-form-valuemappings-values
	//
	Values interface{} `field:"required" json:"values" yaml:"values"`
}

