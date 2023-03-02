package awssagemaker


// The specification of a Jupyter kernel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kernelSpecProperty := &kernelSpecProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	displayName: jsii.String("displayName"),
//   }
//
type CfnAppImageConfig_KernelSpecProperty struct {
	// The name of the Jupyter kernel in the image.
	//
	// This value is case sensitive.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The display name of the kernel.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

