package awsquicksight


// The data point menu options of a dashboard.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPointMenuLabelOptionProperty := &DataPointMenuLabelOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datapointmenulabeloption.html
//
type CfnDashboard_DataPointMenuLabelOptionProperty struct {
	// The status of the data point menu options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datapointmenulabeloption.html#cfn-quicksight-dashboard-datapointmenulabeloption-availabilitystatus
	//
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

