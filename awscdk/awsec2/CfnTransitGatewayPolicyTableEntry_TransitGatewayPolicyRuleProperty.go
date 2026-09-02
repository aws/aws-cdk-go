package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayPolicyRuleProperty := &TransitGatewayPolicyRuleProperty{
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	DestinationPortRange: jsii.String("destinationPortRange"),
//   	Protocol: jsii.String("protocol"),
//   	SourceCidrBlock: jsii.String("sourceCidrBlock"),
//   	SourcePortRange: jsii.String("sourcePortRange"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaypolicytableentry-transitgatewaypolicyrule.html
//
type CfnTransitGatewayPolicyTableEntry_TransitGatewayPolicyRuleProperty struct {
	// The destination CIDR block for the transit gateway policy rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaypolicytableentry-transitgatewaypolicyrule.html#cfn-ec2-transitgatewaypolicytableentry-transitgatewaypolicyrule-destinationcidrblock
	//
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The destination port range for the transit gateway policy rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaypolicytableentry-transitgatewaypolicyrule.html#cfn-ec2-transitgatewaypolicytableentry-transitgatewaypolicyrule-destinationportrange
	//
	DestinationPortRange *string `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// The protocol for the transit gateway policy rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaypolicytableentry-transitgatewaypolicyrule.html#cfn-ec2-transitgatewaypolicytableentry-transitgatewaypolicyrule-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The source CIDR block for the transit gateway policy rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaypolicytableentry-transitgatewaypolicyrule.html#cfn-ec2-transitgatewaypolicytableentry-transitgatewaypolicyrule-sourcecidrblock
	//
	SourceCidrBlock *string `field:"optional" json:"sourceCidrBlock" yaml:"sourceCidrBlock"`
	// The source port range for the transit gateway policy rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaypolicytableentry-transitgatewaypolicyrule.html#cfn-ec2-transitgatewaypolicytableentry-transitgatewaypolicyrule-sourceportrange
	//
	SourcePortRange *string `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
}

