package mixinsawsquicksight


// Determines whether or not hidden fields are visible on exported dashbaords.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   exportWithHiddenFieldsOptionProperty := &ExportWithHiddenFieldsOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-exportwithhiddenfieldsoption.html
//
type CfnDashboardPropsMixin_ExportWithHiddenFieldsOptionProperty struct {
	// The status of the export with hidden fields options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-exportwithhiddenfieldsoption.html#cfn-quicksight-dashboard-exportwithhiddenfieldsoption-availabilitystatus
	//
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

