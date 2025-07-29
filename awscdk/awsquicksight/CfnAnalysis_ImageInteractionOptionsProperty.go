package awsquicksight


// The general image interactions setup for image publish options.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imageinteractionoptions.html
//
type CfnAnalysis_ImageInteractionOptionsProperty struct {
	// The menu options for the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imageinteractionoptions.html#cfn-quicksight-analysis-imageinteractionoptions-imagemenuoption
	//
	ImageMenuOption interface{} `field:"optional" json:"imageMenuOption" yaml:"imageMenuOption"`
}

