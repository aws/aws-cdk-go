package awsamplifyuibuilder


// The `ValueMappings` property specifies the data binding configuration for a value map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var formInputValuePropertyProperty_ FormInputValuePropertyProperty
//
//   valueMappingsProperty := &ValueMappingsProperty{
//   	Values: []interface{}{
//   		&ValueMappingProperty{
//   			Value: &FormInputValuePropertyProperty{
//   				BindingProperties: &FormInputValuePropertyBindingPropertiesProperty{
//   					Property: jsii.String("property"),
//
//   					// the properties below are optional
//   					Field: jsii.String("field"),
//   				},
//   				Concat: []interface{}{
//   					formInputValuePropertyProperty_,
//   				},
//   				Value: jsii.String("value"),
//   			},
//
//   			// the properties below are optional
//   			DisplayValue: &FormInputValuePropertyProperty{
//   				BindingProperties: &FormInputValuePropertyBindingPropertiesProperty{
//   					Property: jsii.String("property"),
//
//   					// the properties below are optional
//   					Field: jsii.String("field"),
//   				},
//   				Concat: []interface{}{
//   					formInputValuePropertyProperty_,
//   				},
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	BindingProperties: map[string]interface{}{
//   		"bindingPropertiesKey": &FormInputBindingPropertiesValueProperty{
//   			"bindingProperties": &FormInputBindingPropertiesValuePropertiesProperty{
//   				"model": jsii.String("model"),
//   			},
//   			"type": jsii.String("type"),
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
	// The information to bind fields to data at runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-valuemappings.html#cfn-amplifyuibuilder-form-valuemappings-bindingproperties
	//
	BindingProperties interface{} `field:"optional" json:"bindingProperties" yaml:"bindingProperties"`
}

