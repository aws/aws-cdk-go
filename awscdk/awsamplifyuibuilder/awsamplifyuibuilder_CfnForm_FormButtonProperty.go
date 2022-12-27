package awsamplifyuibuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formButtonProperty := &formButtonProperty{
//   	children: jsii.String("children"),
//   	excluded: jsii.Boolean(false),
//   	position: &fieldPositionProperty{
//   		below: jsii.String("below"),
//   		fixed: jsii.String("fixed"),
//   		rightOf: jsii.String("rightOf"),
//   	},
//   }
//
type CfnForm_FormButtonProperty struct {
	// `CfnForm.FormButtonProperty.Children`.
	Children *string `field:"optional" json:"children" yaml:"children"`
	// `CfnForm.FormButtonProperty.Excluded`.
	Excluded interface{} `field:"optional" json:"excluded" yaml:"excluded"`
	// `CfnForm.FormButtonProperty.Position`.
	Position interface{} `field:"optional" json:"position" yaml:"position"`
}

