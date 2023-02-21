package awssynthetics


// If this canary is to test an endpoint in a VPC, this structure contains information about the subnet and security groups of the VPC endpoint.
//
// For more information, see [Running a Canary in a VPC](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCConfigProperty := &vPCConfigProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnCanary_VPCConfigProperty struct {
	// The IDs of the security groups for this canary.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The IDs of the subnets where this canary is to run.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the VPC where this canary is to run.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

