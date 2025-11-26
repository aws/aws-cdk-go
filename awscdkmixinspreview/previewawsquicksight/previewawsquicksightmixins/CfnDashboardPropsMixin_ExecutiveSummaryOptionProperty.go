package previewawsquicksightmixins


// Data stories sharing option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   executiveSummaryOptionProperty := &ExecutiveSummaryOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-executivesummaryoption.html
//
type CfnDashboardPropsMixin_ExecutiveSummaryOptionProperty struct {
	// Availability status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-executivesummaryoption.html#cfn-quicksight-dashboard-executivesummaryoption-availabilitystatus
	//
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

