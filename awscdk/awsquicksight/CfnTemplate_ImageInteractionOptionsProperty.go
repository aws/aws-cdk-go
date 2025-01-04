package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var availabilityStatus interface{}
//
//   imageInteractionOptionsProperty := &ImageInteractionOptionsProperty{
//   	ImageMenuOption: &ImageMenuOptionProperty{
//   		AvailabilityStatus: availabilityStatus,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imageinteractionoptions.html
//
type CfnTemplate_ImageInteractionOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imageinteractionoptions.html#cfn-quicksight-template-imageinteractionoptions-imagemenuoption
	//
	ImageMenuOption interface{} `field:"optional" json:"imageMenuOption" yaml:"imageMenuOption"`
}

