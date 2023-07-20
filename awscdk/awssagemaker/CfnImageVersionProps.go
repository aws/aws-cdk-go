package awssagemaker


// Properties for defining a `CfnImageVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageVersionProps := &CfnImageVersionProps{
//   	BaseImage: jsii.String("baseImage"),
//   	ImageName: jsii.String("imageName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html
//
type CfnImageVersionProps struct {
	// The container image that the SageMaker image version is based on.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 255.
	//
	// *Pattern* : `.*`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-baseimage
	//
	BaseImage *string `field:"required" json:"baseImage" yaml:"baseImage"`
	// The name of the parent image.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 63.
	//
	// *Pattern* : `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-imagename
	//
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
}

