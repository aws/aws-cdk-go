package awsamplifyuibuilder


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
	// `CfnForm.FormButtonProperty.Children`.
	Children *string `field:"optional" json:"children" yaml:"children"`
	// `CfnForm.FormButtonProperty.Excluded`.
	Excluded interface{} `field:"optional" json:"excluded" yaml:"excluded"`
	// `CfnForm.FormButtonProperty.Position`.
	Position interface{} `field:"optional" json:"position" yaml:"position"`
}

