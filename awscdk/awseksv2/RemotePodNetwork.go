package awseksv2


// Remote network configuration for pods on hybrid nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   remotePodNetwork := &RemotePodNetwork{
//   	Cidrs: []*string{
//   		jsii.String("cidrs"),
//   	},
//   }
//
type RemotePodNetwork struct {
	// IPv4 CIDR blocks for the remote pod network.
	Cidrs *[]*string `field:"required" json:"cidrs" yaml:"cidrs"`
}

