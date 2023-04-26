package awsamplifyuibuilder


// The `SectionalElement` property specifies the configuration information for a visual helper element for a form.
//
// A sectional element can be a header, a text block, or a divider. These elements are static and not associated with any data.
//
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
	// The type of sectional element.
	//
	// Valid values are `Heading` , `Text` , and `Divider` .
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnForm.SectionalElementProperty.Excluded`.
	Excluded interface{} `field:"optional" json:"excluded" yaml:"excluded"`
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

