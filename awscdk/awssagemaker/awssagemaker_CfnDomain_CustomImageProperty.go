package awssagemaker


// A custom SageMaker image.
//
// For more information, see [Bring your own SageMaker image](https://docs.aws.amazon.com/sagemaker/latest/dg/studio-byoi.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customImageProperty := &customImageProperty{
//   	appImageConfigName: jsii.String("appImageConfigName"),
//   	imageName: jsii.String("imageName"),
//
//   	// the properties below are optional
//   	imageVersionNumber: jsii.Number(123),
//   }
//
type CfnDomain_CustomImageProperty struct {
	// The name of the AppImageConfig.
	AppImageConfigName *string `field:"required" json:"appImageConfigName" yaml:"appImageConfigName"`
	// The name of the CustomImage.
	//
	// Must be unique to your account.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The version number of the CustomImage.
	ImageVersionNumber *float64 `field:"optional" json:"imageVersionNumber" yaml:"imageVersionNumber"`
}

