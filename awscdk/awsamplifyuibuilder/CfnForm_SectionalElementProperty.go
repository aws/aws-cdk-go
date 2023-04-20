package awsamplifyuibuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sectionalElementProperty := &SectionalElementProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Excluded: jsii.Boolean(false),
//   	Level: jsii.Number(123),
//   	Orientation: jsii.String("orientation"),
//   	Position: &FieldPositionProperty{
//   		Below: jsii.String("below"),
//   		Fixed: jsii.String("fixed"),
//   		RightOf: jsii.String("rightOf"),
//   	},
//   	Text: jsii.String("text"),
//   }
//
type CfnForm_SectionalElementProperty struct {
	// `CfnForm.SectionalElementProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnForm.SectionalElementProperty.Excluded`.
	Excluded interface{} `field:"optional" json:"excluded" yaml:"excluded"`
	// `CfnForm.SectionalElementProperty.Level`.
	Level *float64 `field:"optional" json:"level" yaml:"level"`
	// `CfnForm.SectionalElementProperty.Orientation`.
	Orientation *string `field:"optional" json:"orientation" yaml:"orientation"`
	// `CfnForm.SectionalElementProperty.Position`.
	Position interface{} `field:"optional" json:"position" yaml:"position"`
	// `CfnForm.SectionalElementProperty.Text`.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

