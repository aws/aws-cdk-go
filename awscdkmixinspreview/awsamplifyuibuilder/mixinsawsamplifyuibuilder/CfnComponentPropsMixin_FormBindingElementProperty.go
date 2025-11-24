package mixinsawsamplifyuibuilder


// Describes how to bind a component property to form data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   formBindingElementProperty := &FormBindingElementProperty{
//   	Element: jsii.String("element"),
//   	Property: jsii.String("property"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-formbindingelement.html
//
type CfnComponentPropsMixin_FormBindingElementProperty struct {
	// The name of the component to retrieve a value from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-formbindingelement.html#cfn-amplifyuibuilder-component-formbindingelement-element
	//
	Element *string `field:"optional" json:"element" yaml:"element"`
	// The property to retrieve a value from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-formbindingelement.html#cfn-amplifyuibuilder-component-formbindingelement-property
	//
	Property *string `field:"optional" json:"property" yaml:"property"`
}

