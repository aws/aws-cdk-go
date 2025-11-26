package previewawssagemakermixins


// The specification of a Jupyter kernel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kernelSpecProperty := &KernelSpecProperty{
//   	DisplayName: jsii.String("displayName"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-kernelspec.html
//
type CfnAppImageConfigPropsMixin_KernelSpecProperty struct {
	// The display name of the kernel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-kernelspec.html#cfn-sagemaker-appimageconfig-kernelspec-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the Jupyter kernel in the image.
	//
	// This value is case sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-kernelspec.html#cfn-sagemaker-appimageconfig-kernelspec-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

