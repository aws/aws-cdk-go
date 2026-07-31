package awsmediaconnectalpha


// Attributes for importing an existing Router Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   routerInputAttributes := &RouterInputAttributes{
//   	RouterInputArn: jsii.String("routerInputArn"),
//
//   	// the properties below are optional
//   	IpAddress: jsii.String("ipAddress"),
//   	RouterInputId: jsii.String("routerInputId"),
//   }
//
// Experimental.
type RouterInputAttributes struct {
	// The Amazon Resource Name (ARN) of the router input.
	// Experimental.
	RouterInputArn *string `field:"required" json:"routerInputArn" yaml:"routerInputArn"`
	// The IP address that the router input uses to ingest content.
	// Default: - accessing `ipAddress` on the imported input throws; only provide when available.
	//
	// Experimental.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// The unique identifier of the router input.
	// Default: - accessing `routerInputId` on the imported input throws; only provide when available.
	//
	// Experimental.
	RouterInputId *string `field:"optional" json:"routerInputId" yaml:"routerInputId"`
}

