package awseks


// Configuration for network access to the Argo CD capability's managed API server endpoint.
//
// By default, the Argo CD server is accessible via a public endpoint. You can optionally specify one or more VPC endpoint IDs to enable private connectivity from your VPCs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkAccessProperty := &NetworkAccessProperty{
//   	VpceIds: []*string{
//   		jsii.String("vpceIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-networkaccess.html
//
type CfnCapability_NetworkAccessProperty struct {
	// A list of VPC endpoint IDs to associate with the managed Argo CD API server endpoint.
	//
	// Each VPC endpoint provides private connectivity from a specific VPC to the Argo CD server. You can specify multiple VPC endpoint IDs to enable access from multiple VPCs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-networkaccess.html#cfn-eks-capability-networkaccess-vpceids
	//
	VpceIds *[]*string `field:"optional" json:"vpceIds" yaml:"vpceIds"`
}

