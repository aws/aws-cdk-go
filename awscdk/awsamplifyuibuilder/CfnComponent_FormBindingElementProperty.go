package awsamplifyuibuilder


// Describes how to bind a component property to form data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formBindingElementProperty := &FormBindingElementProperty{
//   	Element: jsii.String("element"),
//   	Property: jsii.String("property"),
//   }
//
type CfnComponent_FormBindingElementProperty struct {
	// The name of the component to retrieve a value from.
	Element *string `field:"required" json:"element" yaml:"element"`
	// The property to retrieve a value from.
	Property *string `field:"required" json:"property" yaml:"property"`
}

