package awsquicksight


// Determines the color that is applied to a particular data value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataColorProperty := &DataColorProperty{
//   	Color: jsii.String("color"),
//   	DataValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datacolor.html
//
type CfnDashboard_DataColorProperty struct {
	// The color that is applied to the data value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datacolor.html#cfn-quicksight-dashboard-datacolor-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The data value that the color is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datacolor.html#cfn-quicksight-dashboard-datacolor-datavalue
	//
	DataValue *float64 `field:"optional" json:"dataValue" yaml:"dataValue"`
}

