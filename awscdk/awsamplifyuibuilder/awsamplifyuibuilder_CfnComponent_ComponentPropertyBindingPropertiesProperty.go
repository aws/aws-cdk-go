package awsamplifyuibuilder


// The `ComponentPropertyBindingProperties` property specifies a component property to associate with a binding property.
//
// This enables exposed properties on the top level component to propagate data to the component's property values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentPropertyBindingPropertiesProperty := &componentPropertyBindingPropertiesProperty{
//   	property: jsii.String("property"),
//
//   	// the properties below are optional
//   	field: jsii.String("field"),
//   }
//
type CfnComponent_ComponentPropertyBindingPropertiesProperty struct {
	// The component property to bind to the data field.
	Property *string `field:"required" json:"property" yaml:"property"`
	// The data field to bind the property to.
	Field *string `field:"optional" json:"field" yaml:"field"`
}

