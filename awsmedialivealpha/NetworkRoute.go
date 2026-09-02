package awsmedialivealpha


// A route for a MediaLive Network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   networkRoute := &NetworkRoute{
//   	Cidr: jsii.String("cidr"),
//   	Gateway: jsii.String("gateway"),
//   }
//
// Experimental.
type NetworkRoute struct {
	// The CIDR block for the route.
	// Experimental.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// The gateway for the route.
	// Experimental.
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
}

