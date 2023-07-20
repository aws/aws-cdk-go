package awsquicksight


// The field label type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldLabelTypeProperty := &FieldLabelTypeProperty{
//   	FieldId: jsii.String("fieldId"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fieldlabeltype.html
//
type CfnTemplate_FieldLabelTypeProperty struct {
	// Indicates the field that is targeted by the field label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fieldlabeltype.html#cfn-quicksight-template-fieldlabeltype-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
	// The visibility of the field label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fieldlabeltype.html#cfn-quicksight-template-fieldlabeltype-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

