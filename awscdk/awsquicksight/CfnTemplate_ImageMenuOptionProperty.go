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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagemenuoption.html
//
type CfnTemplate_ImageMenuOptionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagemenuoption.html#cfn-quicksight-template-imagemenuoption-availabilitystatus
	//
	AvailabilityStatus interface{} `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

