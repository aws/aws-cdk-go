package awsamplifyuibuilder


// Stores the configuration information for a visual helper element for a form.
//
// A sectional element can be a header, a text block, or a divider. These elements are static and not associated with any data.
//
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
	// The type of sectional element.
	//
	// Valid values are `Heading` , `Text` , and `Divider` .
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies the size of the font for a `Heading` sectional element.
	//
	// Valid values are `1 | 2 | 3 | 4 | 5 | 6` .
	Level *float64 `field:"optional" json:"level" yaml:"level"`
	// Specifies the orientation for a `Divider` sectional element.
	//
	// Valid values are `horizontal` or `vertical` .
	Orientation *string `field:"optional" json:"orientation" yaml:"orientation"`
	// Specifies the position of the text in a field for a `Text` sectional element.
	Position interface{} `field:"optional" json:"position" yaml:"position"`
	// The text for a `Text` sectional element.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

