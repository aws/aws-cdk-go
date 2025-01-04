package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var availabilityStatus interface{}
//
//   imageMenuOptionProperty := &ImageMenuOptionProperty{
//   	AvailabilityStatus: availabilityStatus,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagemenuoption.html
//
type CfnAnalysis_ImageMenuOptionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagemenuoption.html#cfn-quicksight-analysis-imagemenuoption-availabilitystatus
	//
	AvailabilityStatus interface{} `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

