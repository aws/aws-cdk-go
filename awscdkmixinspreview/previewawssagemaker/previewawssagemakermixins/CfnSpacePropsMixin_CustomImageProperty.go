package previewawssagemakermixins


// A custom SageMaker AI image.
//
// For more information, see [Bring your own SageMaker AI image](https://docs.aws.amazon.com/sagemaker/latest/dg/studio-byoi.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customImageProperty := &CustomImageProperty{
//   	AppImageConfigName: jsii.String("appImageConfigName"),
//   	ImageName: jsii.String("imageName"),
//   	ImageVersionNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-customimage.html
//
type CfnSpacePropsMixin_CustomImageProperty struct {
	// The name of the AppImageConfig.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-customimage.html#cfn-sagemaker-space-customimage-appimageconfigname
	//
	AppImageConfigName *string `field:"optional" json:"appImageConfigName" yaml:"appImageConfigName"`
	// The name of the CustomImage.
	//
	// Must be unique to your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-customimage.html#cfn-sagemaker-space-customimage-imagename
	//
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// The version number of the CustomImage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-customimage.html#cfn-sagemaker-space-customimage-imageversionnumber
	//
	ImageVersionNumber *float64 `field:"optional" json:"imageVersionNumber" yaml:"imageVersionNumber"`
}

