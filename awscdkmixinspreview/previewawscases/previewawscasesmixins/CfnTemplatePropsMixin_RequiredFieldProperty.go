package previewawscasesmixins


// List of fields that must have a value provided to create a case.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requiredFieldProperty := &RequiredFieldProperty{
//   	FieldId: jsii.String("fieldId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-template-requiredfield.html
//
type CfnTemplatePropsMixin_RequiredFieldProperty struct {
	// Unique identifier of a field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-template-requiredfield.html#cfn-cases-template-requiredfield-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
}

