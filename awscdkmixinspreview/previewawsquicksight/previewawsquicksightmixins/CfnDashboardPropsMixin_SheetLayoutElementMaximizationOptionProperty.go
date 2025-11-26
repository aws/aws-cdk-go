package previewawsquicksightmixins


// The sheet layout maximization options of a dashbaord.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sheetLayoutElementMaximizationOptionProperty := &SheetLayoutElementMaximizationOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetlayoutelementmaximizationoption.html
//
type CfnDashboardPropsMixin_SheetLayoutElementMaximizationOptionProperty struct {
	// The status of the sheet layout maximization options of a dashbaord.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetlayoutelementmaximizationoption.html#cfn-quicksight-dashboard-sheetlayoutelementmaximizationoption-availabilitystatus
	//
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

