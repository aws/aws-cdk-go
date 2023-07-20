package awsquicksight


// The options that determine how a date axis is displayed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateAxisOptionsProperty := &DateAxisOptionsProperty{
//   	MissingDateVisibility: jsii.String("missingDateVisibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dateaxisoptions.html
//
type CfnDashboard_DateAxisOptionsProperty struct {
	// Determines whether or not missing dates are displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dateaxisoptions.html#cfn-quicksight-dashboard-dateaxisoptions-missingdatevisibility
	//
	MissingDateVisibility *string `field:"optional" json:"missingDateVisibility" yaml:"missingDateVisibility"`
}

