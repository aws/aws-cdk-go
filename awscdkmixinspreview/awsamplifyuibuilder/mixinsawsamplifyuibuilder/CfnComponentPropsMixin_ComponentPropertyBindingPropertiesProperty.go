package mixinsawsamplifyuibuilder


// The `ComponentPropertyBindingProperties` property specifies a component property to associate with a binding property.
//
// This enables exposed properties on the top level component to propagate data to the component's property values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   componentPropertyBindingPropertiesProperty := &ComponentPropertyBindingPropertiesProperty{
//   	Field: jsii.String("field"),
//   	Property: jsii.String("property"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentpropertybindingproperties.html
//
type CfnComponentPropsMixin_ComponentPropertyBindingPropertiesProperty struct {
	// The data field to bind the property to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentpropertybindingproperties.html#cfn-amplifyuibuilder-component-componentpropertybindingproperties-field
	//
	Field *string `field:"optional" json:"field" yaml:"field"`
	// The component property to bind to the data field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentpropertybindingproperties.html#cfn-amplifyuibuilder-component-componentpropertybindingproperties-property
	//
	Property *string `field:"optional" json:"property" yaml:"property"`
}

