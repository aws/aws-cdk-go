package awssagemaker


// The specification of a Jupyter kernel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kernelSpecProperty := &KernelSpecProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DisplayName: jsii.String("displayName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-kernelspec.html
//
type CfnAppImageConfig_KernelSpecProperty struct {
	// The name of the Jupyter kernel in the image.
	//
	// This value is case sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-kernelspec.html#cfn-sagemaker-appimageconfig-kernelspec-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The display name of the kernel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-kernelspec.html#cfn-sagemaker-appimageconfig-kernelspec-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

