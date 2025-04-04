package awsquicksight


// The context menu options for a visual's interactions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contextMenuOptionProperty := &ContextMenuOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-contextmenuoption.html
//
type CfnAnalysis_ContextMenuOptionProperty struct {
	// The availability status of the context menu options.
	//
	// If the value of this property is set to `ENABLED` , dashboard readers can interact with the context menu.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-contextmenuoption.html#cfn-quicksight-analysis-contextmenuoption-availabilitystatus
	//
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

