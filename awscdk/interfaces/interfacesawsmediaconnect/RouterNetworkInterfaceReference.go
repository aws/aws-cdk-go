package interfacesawsmediaconnect


// A reference to a RouterNetworkInterface resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routerNetworkInterfaceReference := &RouterNetworkInterfaceReference{
//   	RouterNetworkInterfaceArn: jsii.String("routerNetworkInterfaceArn"),
//   }
//
type RouterNetworkInterfaceReference struct {
	// The Arn of the RouterNetworkInterface resource.
	RouterNetworkInterfaceArn *string `field:"required" json:"routerNetworkInterfaceArn" yaml:"routerNetworkInterfaceArn"`
}

