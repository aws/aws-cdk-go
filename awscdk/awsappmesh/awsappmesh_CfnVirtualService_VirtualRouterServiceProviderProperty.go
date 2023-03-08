package awsappmesh


// An object that represents a virtual node service provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualRouterServiceProviderProperty := &virtualRouterServiceProviderProperty{
//   	virtualRouterName: jsii.String("virtualRouterName"),
//   }
//
type CfnVirtualService_VirtualRouterServiceProviderProperty struct {
	// The name of the virtual router that is acting as a service provider.
	VirtualRouterName *string `field:"required" json:"virtualRouterName" yaml:"virtualRouterName"`
}

