package awscdkeksv2alpha


// Remote network configuration for pods on hybrid nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeksv2alpha"
//
//   remotePodNetwork := &RemotePodNetwork{
//   	Cidrs: []*string{
//   		jsii.String("cidrs"),
//   	},
//   }
//
// Experimental.
type RemotePodNetwork struct {
	// IPv4 CIDR blocks for the remote pod network.
	// Experimental.
	Cidrs *[]*string `field:"required" json:"cidrs" yaml:"cidrs"`
}

