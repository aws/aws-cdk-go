package awseks


// Network configuration of pods run on-premises with EKS Hybrid Nodes.
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
	// Specifies the list of remote pod CIDRs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-remotepodnetwork.html#cfn-eks-cluster-remotepodnetwork-cidrs
	//
	Cidrs *[]*string `field:"required" json:"cidrs" yaml:"cidrs"`
}

