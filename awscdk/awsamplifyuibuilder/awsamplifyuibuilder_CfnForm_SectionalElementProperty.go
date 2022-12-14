package awsamplifyuibuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sectionalElementProperty := &sectionalElementProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	level: jsii.Number(123),
//   	orientation: jsii.String("orientation"),
//   	position: &fieldPositionProperty{
//   		below: jsii.String("below"),
//   		fixed: jsii.String("fixed"),
//   		rightOf: jsii.String("rightOf"),
//   	},
//   	text: jsii.String("text"),
//   }
//
type CfnForm_SectionalElementProperty struct {
	// `CfnForm.SectionalElementProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnForm.SectionalElementProperty.Level`.
	Level *float64 `field:"optional" json:"level" yaml:"level"`
	// `CfnForm.SectionalElementProperty.Orientation`.
	Orientation *string `field:"optional" json:"orientation" yaml:"orientation"`
	// `CfnForm.SectionalElementProperty.Position`.
	Position interface{} `field:"optional" json:"position" yaml:"position"`
	// `CfnForm.SectionalElementProperty.Text`.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

