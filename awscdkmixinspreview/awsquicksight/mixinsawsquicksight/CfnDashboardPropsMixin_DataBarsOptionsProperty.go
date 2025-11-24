package mixinsawsquicksight


// The options for data bars.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataBarsOptionsProperty := &DataBarsOptionsProperty{
//   	FieldId: jsii.String("fieldId"),
//   	NegativeColor: jsii.String("negativeColor"),
//   	PositiveColor: jsii.String("positiveColor"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-databarsoptions.html
//
type CfnDashboardPropsMixin_DataBarsOptionsProperty struct {
	// The field ID for the data bars options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-databarsoptions.html#cfn-quicksight-dashboard-databarsoptions-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
	// The color of the negative data bar.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-databarsoptions.html#cfn-quicksight-dashboard-databarsoptions-negativecolor
	//
	NegativeColor *string `field:"optional" json:"negativeColor" yaml:"negativeColor"`
	// The color of the positive data bar.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-databarsoptions.html#cfn-quicksight-dashboard-databarsoptions-positivecolor
	//
	PositiveColor *string `field:"optional" json:"positiveColor" yaml:"positiveColor"`
}

