package previewawsquicksightmixins


// The data point tooltip options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataPointTooltipOptionProperty := &DataPointTooltipOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datapointtooltipoption.html
//
type CfnDashboardPropsMixin_DataPointTooltipOptionProperty struct {
	// The status of the data point tool tip options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datapointtooltipoption.html#cfn-quicksight-dashboard-datapointtooltipoption-availabilitystatus
	//
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

