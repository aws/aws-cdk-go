package previewawsamplifyuibuildermixins


// The `ValueMappings` property specifies the data binding configuration for a value map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var formInputValuePropertyProperty_ FormInputValuePropertyProperty
//
//   valueMappingsProperty := &ValueMappingsProperty{
//   	BindingProperties: map[string]interface{}{
//   		"bindingPropertiesKey": &FormInputBindingPropertiesValueProperty{
//   			"bindingProperties": &FormInputBindingPropertiesValuePropertiesProperty{
//   				"model": jsii.String("model"),
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	Values: []interface{}{
//   		&ValueMappingProperty{
//   			DisplayValue: &FormInputValuePropertyProperty{
//   				BindingProperties: &FormInputValuePropertyBindingPropertiesProperty{
//   					Field: jsii.String("field"),
//   					Property: jsii.String("property"),
//   				},
//   				Concat: []interface{}{
//   					formInputValuePropertyProperty_,
//   				},
//   				Value: jsii.String("value"),
//   			},
//   			Value: &FormInputValuePropertyProperty{
//   				BindingProperties: &FormInputValuePropertyBindingPropertiesProperty{
//   					Field: jsii.String("field"),
//   					Property: jsii.String("property"),
//   				},
//   				Concat: []interface{}{
//   					formInputValuePropertyProperty_,
//   				},
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-valuemappings.html
//
type CfnFormPropsMixin_ValueMappingsProperty struct {
	// The information to bind fields to data at runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-valuemappings.html#cfn-amplifyuibuilder-form-valuemappings-bindingproperties
	//
	BindingProperties interface{} `field:"optional" json:"bindingProperties" yaml:"bindingProperties"`
	// The value and display value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-valuemappings.html#cfn-amplifyuibuilder-form-valuemappings-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

