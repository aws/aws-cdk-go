package awsmediaconnectalpha


// Attributes for importing an existing Router Network Interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   routerNetworkInterfaceAttributes := &RouterNetworkInterfaceAttributes{
//   	RouterNetworkInterfaceArn: jsii.String("routerNetworkInterfaceArn"),
//
//   	// the properties below are optional
//   	RouterNetworkInterfaceId: jsii.String("routerNetworkInterfaceId"),
//   }
//
// Experimental.
type RouterNetworkInterfaceAttributes struct {
	// The Amazon Resource Name (ARN) of the router network interface.
	// Experimental.
	RouterNetworkInterfaceArn *string `field:"required" json:"routerNetworkInterfaceArn" yaml:"routerNetworkInterfaceArn"`
	// The unique identifier of the router network interface.
	// Default: - accessing `routerNetworkInterfaceId` on the imported interface throws; only provide when available.
	//
	// Experimental.
	RouterNetworkInterfaceId *string `field:"optional" json:"routerNetworkInterfaceId" yaml:"routerNetworkInterfaceId"`
}

