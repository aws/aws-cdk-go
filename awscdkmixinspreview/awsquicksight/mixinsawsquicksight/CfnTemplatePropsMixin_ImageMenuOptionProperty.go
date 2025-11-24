package mixinsawsquicksight


// The menu options for the interactions of an image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageMenuOptionProperty := &ImageMenuOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagemenuoption.html
//
type CfnTemplatePropsMixin_ImageMenuOptionProperty struct {
	// The availability status of the image menu.
	//
	// If the value of this property is set to `ENABLED` , dashboard readers can interact with the image menu.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagemenuoption.html#cfn-quicksight-template-imagemenuoption-availabilitystatus
	//
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

