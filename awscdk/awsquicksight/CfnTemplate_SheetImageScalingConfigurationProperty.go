package awsquicksight


// Determines how the image is scaled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetImageScalingConfigurationProperty := &SheetImageScalingConfigurationProperty{
//   	ScalingType: jsii.String("scalingType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimagescalingconfiguration.html
//
type CfnTemplate_SheetImageScalingConfigurationProperty struct {
	// The scaling option to use when fitting the image inside the container.
	//
	// Valid values are defined as follows:
	//
	// - `SCALE_TO_WIDTH` : The image takes up the entire width of the container. The image aspect ratio is preserved.
	// - `SCALE_TO_HEIGHT` : The image takes up the entire height of the container. The image aspect ratio is preserved.
	// - `SCALE_TO_CONTAINER` : The image takes up the entire width and height of the container. The image aspect ratio is not preserved.
	// - `SCALE_NONE` : The image is displayed in its original size and is not scaled to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimagescalingconfiguration.html#cfn-quicksight-template-sheetimagescalingconfiguration-scalingtype
	//
	ScalingType *string `field:"optional" json:"scalingType" yaml:"scalingType"`
}

