package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quickSuiteActionsOptionProperty := &QuickSuiteActionsOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-quicksuiteactionsoption.html
//
type CfnDashboard_QuickSuiteActionsOptionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-quicksuiteactionsoption.html#cfn-quicksight-dashboard-quicksuiteactionsoption-availabilitystatus
	//
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

