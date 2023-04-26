package awsamplifyuibuilder


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
	// `CfnComponent.FormBindingElementProperty.Element`.
	Element *string `field:"required" json:"element" yaml:"element"`
	// `CfnComponent.FormBindingElementProperty.Property`.
	Property *string `field:"required" json:"property" yaml:"property"`
}

