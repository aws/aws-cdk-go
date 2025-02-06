package awsec2


// The options for cidr type endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cidrOptionsProperty := &CidrOptionsProperty{
//   	Cidr: jsii.String("cidr"),
//   	PortRanges: []interface{}{
//   		&PortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	Protocol: jsii.String("protocol"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-cidroptions.html
//
type CfnVerifiedAccessEndpoint_CidrOptionsProperty struct {
	// The IP address range, in CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-cidroptions.html#cfn-ec2-verifiedaccessendpoint-cidroptions-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The list of port range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-cidroptions.html#cfn-ec2-verifiedaccessendpoint-cidroptions-portranges
	//
	PortRanges interface{} `field:"optional" json:"portRanges" yaml:"portRanges"`
	// The IP protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-cidroptions.html#cfn-ec2-verifiedaccessendpoint-cidroptions-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The IDs of the subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-cidroptions.html#cfn-ec2-verifiedaccessendpoint-cidroptions-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

