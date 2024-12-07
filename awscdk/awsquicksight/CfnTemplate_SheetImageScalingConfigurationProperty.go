package awsquicksight


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimagescalingconfiguration.html#cfn-quicksight-template-sheetimagescalingconfiguration-scalingtype
	//
	ScalingType *string `field:"optional" json:"scalingType" yaml:"scalingType"`
}

