package awsec2


// Describes the port range for a Verified Access endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portRangeProperty := &PortRangeProperty{
//   	FromPort: jsii.Number(123),
//   	ToPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-portrange.html
//
type CfnVerifiedAccessEndpoint_PortRangeProperty struct {
	// The start of the port range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-portrange.html#cfn-ec2-verifiedaccessendpoint-portrange-fromport
	//
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// The end of the port range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-portrange.html#cfn-ec2-verifiedaccessendpoint-portrange-toport
	//
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

