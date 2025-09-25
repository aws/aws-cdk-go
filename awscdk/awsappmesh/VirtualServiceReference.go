package awsappmesh


// A reference to a VirtualService resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualServiceReference := &VirtualServiceReference{
//   	VirtualServiceArn: jsii.String("virtualServiceArn"),
//   	VirtualServiceId: jsii.String("virtualServiceId"),
//   }
//
type VirtualServiceReference struct {
	// The ARN of the VirtualService resource.
	VirtualServiceArn *string `field:"required" json:"virtualServiceArn" yaml:"virtualServiceArn"`
	// The Id of the VirtualService resource.
	VirtualServiceId *string `field:"required" json:"virtualServiceId" yaml:"virtualServiceId"`
}

