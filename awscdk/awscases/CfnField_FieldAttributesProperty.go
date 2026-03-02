package awscases


// Union of field attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldAttributesProperty := &FieldAttributesProperty{
//   	Text: &TextAttributesProperty{
//   		IsMultiline: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-field-fieldattributes.html
//
type CfnField_FieldAttributesProperty struct {
	// Field attributes for Text field type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-field-fieldattributes.html#cfn-cases-field-fieldattributes-text
	//
	Text interface{} `field:"optional" json:"text" yaml:"text"`
}

