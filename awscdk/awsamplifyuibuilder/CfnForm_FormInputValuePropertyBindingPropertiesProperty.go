package awsamplifyuibuilder


// Associates a form property to a binding property.
//
// This enables exposed properties on the top level form to propagate data to the form's property values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formInputValuePropertyBindingPropertiesProperty := &FormInputValuePropertyBindingPropertiesProperty{
//   	Property: jsii.String("property"),
//
//   	// the properties below are optional
//   	Field: jsii.String("field"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-forminputvaluepropertybindingproperties.html
//
type CfnForm_FormInputValuePropertyBindingPropertiesProperty struct {
	// The form property to bind to the data field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-forminputvaluepropertybindingproperties.html#cfn-amplifyuibuilder-form-forminputvaluepropertybindingproperties-property
	//
	Property *string `field:"required" json:"property" yaml:"property"`
	// The data field to bind the property to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-forminputvaluepropertybindingproperties.html#cfn-amplifyuibuilder-form-forminputvaluepropertybindingproperties-field
	//
	Field *string `field:"optional" json:"field" yaml:"field"`
}

