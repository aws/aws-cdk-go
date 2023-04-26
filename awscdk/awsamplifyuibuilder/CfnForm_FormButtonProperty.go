package awsamplifyuibuilder


// The `FormButton` property specifies the configuration for a button UI element that is a part of a form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formButtonProperty := &FormButtonProperty{
//   	Children: jsii.String("children"),
//   	Excluded: jsii.Boolean(false),
//   	Position: &FieldPositionProperty{
//   		Below: jsii.String("below"),
//   		Fixed: jsii.String("fixed"),
//   		RightOf: jsii.String("rightOf"),
//   	},
//   }
//
type CfnForm_FormButtonProperty struct {
	// Describes the button's properties.
	Children *string `field:"optional" json:"children" yaml:"children"`
	// Specifies whether the button is visible on the form.
	Excluded interface{} `field:"optional" json:"excluded" yaml:"excluded"`
	// The position of the button.
	Position interface{} `field:"optional" json:"position" yaml:"position"`
}

