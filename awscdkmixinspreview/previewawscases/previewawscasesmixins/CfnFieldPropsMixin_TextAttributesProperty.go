package previewawscasesmixins


// Field attributes for Text field type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   textAttributesProperty := &TextAttributesProperty{
//   	IsMultiline: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-field-textattributes.html
//
type CfnFieldPropsMixin_TextAttributesProperty struct {
	// Attribute that defines rendering component and validation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-field-textattributes.html#cfn-cases-field-textattributes-ismultiline
	//
	IsMultiline interface{} `field:"optional" json:"isMultiline" yaml:"isMultiline"`
}

