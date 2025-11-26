package previewawssyntheticsmixins


// If this canary is to test an endpoint in a VPC, this structure contains information about the subnet and security groups of the VPC endpoint.
//
// For more information, see [Running a Canary in a VPC](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vPCConfigProperty := &VPCConfigProperty{
//   	Ipv6AllowedForDualStack: jsii.Boolean(false),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-vpcconfig.html
//
type CfnCanaryPropsMixin_VPCConfigProperty struct {
	// Set this to `true` to allow outbound IPv6 traffic on VPC canaries that are connected to dual-stack subnets.
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-vpcconfig.html#cfn-synthetics-canary-vpcconfig-ipv6allowedfordualstack
	//
	Ipv6AllowedForDualStack interface{} `field:"optional" json:"ipv6AllowedForDualStack" yaml:"ipv6AllowedForDualStack"`
	// The IDs of the security groups for this canary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-vpcconfig.html#cfn-synthetics-canary-vpcconfig-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The IDs of the subnets where this canary is to run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-vpcconfig.html#cfn-synthetics-canary-vpcconfig-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the VPC where this canary is to run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-vpcconfig.html#cfn-synthetics-canary-vpcconfig-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

