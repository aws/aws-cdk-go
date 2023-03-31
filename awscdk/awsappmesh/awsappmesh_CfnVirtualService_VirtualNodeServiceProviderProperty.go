package awsappmesh


// An object that represents a virtual node service provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualNodeServiceProviderProperty := &virtualNodeServiceProviderProperty{
//   	virtualNodeName: jsii.String("virtualNodeName"),
//   }
//
type CfnVirtualService_VirtualNodeServiceProviderProperty struct {
	// The name of the virtual node that is acting as a service provider.
	VirtualNodeName *string `field:"required" json:"virtualNodeName" yaml:"virtualNodeName"`
}

