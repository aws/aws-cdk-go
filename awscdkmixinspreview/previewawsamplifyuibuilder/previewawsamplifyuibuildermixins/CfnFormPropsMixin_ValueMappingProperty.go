package previewawsamplifyuibuildermixins


// The `ValueMapping` property specifies the association between a complex object and a display value.
//
// Use `ValueMapping` to store how to represent complex objects when they are displayed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var formInputValuePropertyProperty_ FormInputValuePropertyProperty
//
//   valueMappingProperty := &ValueMappingProperty{
//   	DisplayValue: &FormInputValuePropertyProperty{
//   		BindingProperties: &FormInputValuePropertyBindingPropertiesProperty{
//   			Field: jsii.String("field"),
//   			Property: jsii.String("property"),
//   		},
//   		Concat: []interface{}{
//   			formInputValuePropertyProperty_,
//   		},
//   		Value: jsii.String("value"),
//   	},
//   	Value: &FormInputValuePropertyProperty{
//   		BindingProperties: &FormInputValuePropertyBindingPropertiesProperty{
//   			Field: jsii.String("field"),
//   			Property: jsii.String("property"),
//   		},
//   		Concat: []interface{}{
//   			formInputValuePropertyProperty_,
//   		},
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-valuemapping.html
//
type CfnFormPropsMixin_ValueMappingProperty struct {
	// The value to display for the complex object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-valuemapping.html#cfn-amplifyuibuilder-form-valuemapping-displayvalue
	//
	DisplayValue interface{} `field:"optional" json:"displayValue" yaml:"displayValue"`
	// The complex object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-valuemapping.html#cfn-amplifyuibuilder-form-valuemapping-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

