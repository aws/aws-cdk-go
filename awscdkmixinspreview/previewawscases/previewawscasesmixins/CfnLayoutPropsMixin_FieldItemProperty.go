package previewawscasesmixins


// Object for field related information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fieldItemProperty := &FieldItemProperty{
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-fielditem.html
//
type CfnLayoutPropsMixin_FieldItemProperty struct {
	// Unique identifier of a field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-fielditem.html#cfn-cases-layout-fielditem-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

