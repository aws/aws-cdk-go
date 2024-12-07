package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageInteractionOptionsProperty := &ImageInteractionOptionsProperty{
//   	ImageMenuOption: &ImageMenuOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-imageinteractionoptions.html
//
type CfnDashboard_ImageInteractionOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-imageinteractionoptions.html#cfn-quicksight-dashboard-imageinteractionoptions-imagemenuoption
	//
	ImageMenuOption interface{} `field:"optional" json:"imageMenuOption" yaml:"imageMenuOption"`
}

