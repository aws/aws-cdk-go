package previewawsquicksightmixins


// The option that specifies individual data values for labels.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataPathLabelTypeProperty := &DataPathLabelTypeProperty{
//   	FieldId: jsii.String("fieldId"),
//   	FieldValue: jsii.String("fieldValue"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datapathlabeltype.html
//
type CfnTemplatePropsMixin_DataPathLabelTypeProperty struct {
	// The field ID of the field that the data label needs to be applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datapathlabeltype.html#cfn-quicksight-template-datapathlabeltype-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
	// The actual value of the field that is labeled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datapathlabeltype.html#cfn-quicksight-template-datapathlabeltype-fieldvalue
	//
	FieldValue *string `field:"optional" json:"fieldValue" yaml:"fieldValue"`
	// The visibility of the data label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datapathlabeltype.html#cfn-quicksight-template-datapathlabeltype-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

