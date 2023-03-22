package awsappmesh


// An object that represents the specification of a virtual service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualServiceSpecProperty := &virtualServiceSpecProperty{
//   	provider: &virtualServiceProviderProperty{
//   		virtualNode: &virtualNodeServiceProviderProperty{
//   			virtualNodeName: jsii.String("virtualNodeName"),
//   		},
//   		virtualRouter: &virtualRouterServiceProviderProperty{
//   			virtualRouterName: jsii.String("virtualRouterName"),
//   		},
//   	},
//   }
//
type CfnVirtualService_VirtualServiceSpecProperty struct {
	// The App Mesh object that is acting as the provider for a virtual service.
	//
	// You can specify a single virtual node or virtual router.
	Provider interface{} `field:"optional" json:"provider" yaml:"provider"`
}

