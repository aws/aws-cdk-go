package previewawsquicksightmixins


// The drill down options for data points in a dashbaord.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataPointDrillUpDownOptionProperty := &DataPointDrillUpDownOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datapointdrillupdownoption.html
//
type CfnDashboardPropsMixin_DataPointDrillUpDownOptionProperty struct {
	// The status of the drill down options of data points.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datapointdrillupdownoption.html#cfn-quicksight-dashboard-datapointdrillupdownoption-availabilitystatus
	//
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

