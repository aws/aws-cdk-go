package awseksv2


// Remote network configuration for hybrid nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   remoteNodeNetwork := &RemoteNodeNetwork{
//   	Cidrs: []*string{
//   		jsii.String("cidrs"),
//   	},
//   }
//
type RemoteNodeNetwork struct {
	// IPv4 CIDR blocks for the remote node network.
	Cidrs *[]*string `field:"required" json:"cidrs" yaml:"cidrs"`
}

