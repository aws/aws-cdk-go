package mixinsawsamplifyuibuilder


// The `FormInputValueProperty` property specifies the configuration for an input field on a form.
//
// Use `FormInputValueProperty` to specify the values to render or bind by default.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var formInputValuePropertyProperty_ FormInputValuePropertyProperty
//
//   formInputValuePropertyProperty := &FormInputValuePropertyProperty{
//   	BindingProperties: &FormInputValuePropertyBindingPropertiesProperty{
//   		Field: jsii.String("field"),
//   		Property: jsii.String("property"),
//   	},
//   	Concat: []interface{}{
//   		formInputValuePropertyProperty_,
//   	},
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-forminputvalueproperty.html
//
type CfnFormPropsMixin_FormInputValuePropertyProperty struct {
	// The information to bind fields to data at runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-forminputvalueproperty.html#cfn-amplifyuibuilder-form-forminputvalueproperty-bindingproperties
	//
	BindingProperties interface{} `field:"optional" json:"bindingProperties" yaml:"bindingProperties"`
	// A list of form properties to concatenate to create the value to assign to this field property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-forminputvalueproperty.html#cfn-amplifyuibuilder-form-forminputvalueproperty-concat
	//
	Concat interface{} `field:"optional" json:"concat" yaml:"concat"`
	// The value to assign to the input field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-forminputvalueproperty.html#cfn-amplifyuibuilder-form-forminputvalueproperty-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

