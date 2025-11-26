package previewawsgrafanamixins


// The configuration settings for in-bound network access to your workspace.
//
// When this is configured, only listed IP addresses and VPC endpoints will be able to access your workspace. Standard Grafana authentication and authorization are still required.
//
// Access is granted to a caller that is in either the IP address list or the VPC endpoint list - they do not need to be in both.
//
// If this is not configured, or is removed, then all IP addresses and VPC endpoints are allowed. Standard Grafana authentication and authorization are still required.
//
// > While both `prefixListIds` and `vpceIds` are required, you can pass in an empty array of strings for either parameter if you do not want to allow any of that type.
// >
// > If both are passed as empty arrays, no traffic is allowed to the workspace, because only *explicitly* allowed connections are accepted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkAccessControlProperty := &NetworkAccessControlProperty{
//   	PrefixListIds: []*string{
//   		jsii.String("prefixListIds"),
//   	},
//   	VpceIds: []*string{
//   		jsii.String("vpceIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-networkaccesscontrol.html
//
type CfnWorkspacePropsMixin_NetworkAccessControlProperty struct {
	// An array of prefix list IDs.
	//
	// A prefix list is a list of CIDR ranges of IP addresses. The IP addresses specified are allowed to access your workspace. If the list is not included in the configuration (passed an empty array) then no IP addresses are allowed to access the workspace. You create a prefix list using the Amazon VPC console.
	//
	// Prefix list IDs have the format `pl- *1a2b3c4d*` .
	//
	// For more information about prefix lists, see [Group CIDR blocks using managed prefix lists](https://docs.aws.amazon.com/vpc/latest/userguide/managed-prefix-lists.html) in the *Amazon Virtual Private Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-networkaccesscontrol.html#cfn-grafana-workspace-networkaccesscontrol-prefixlistids
	//
	PrefixListIds *[]*string `field:"optional" json:"prefixListIds" yaml:"prefixListIds"`
	// An array of Amazon VPC endpoint IDs for the workspace.
	//
	// You can create VPC endpoints to your Amazon Managed Grafana workspace for access from within a VPC. If a `NetworkAccessConfiguration` is specified then only VPC endpoints specified here are allowed to access the workspace. If you pass in an empty array of strings, then no VPCs are allowed to access the workspace.
	//
	// VPC endpoint IDs have the format `vpce- *1a2b3c4d*` .
	//
	// For more information about creating an interface VPC endpoint, see [Interface VPC endpoints](https://docs.aws.amazon.com/grafana/latest/userguide/VPC-endpoints) in the *Amazon Managed Grafana User Guide* .
	//
	// > The only VPC endpoints that can be specified here are interface VPC endpoints for Grafana workspaces (using the `com.amazonaws.[region].grafana-workspace` service endpoint). Other VPC endpoints are ignored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-networkaccesscontrol.html#cfn-grafana-workspace-networkaccesscontrol-vpceids
	//
	VpceIds *[]*string `field:"optional" json:"vpceIds" yaml:"vpceIds"`
}

