package awsec2


// Properties for defining a `CfnTransitGatewayMeteringPolicyEntry`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayMeteringPolicyEntryProps := &CfnTransitGatewayMeteringPolicyEntryProps{
//   	MeteredAccount: jsii.String("meteredAccount"),
//   	PolicyRuleNumber: jsii.Number(123),
//   	TransitGatewayMeteringPolicyId: jsii.String("transitGatewayMeteringPolicyId"),
//
//   	// the properties below are optional
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	DestinationPortRange: jsii.String("destinationPortRange"),
//   	DestinationTransitGatewayAttachmentId: jsii.String("destinationTransitGatewayAttachmentId"),
//   	DestinationTransitGatewayAttachmentType: jsii.String("destinationTransitGatewayAttachmentType"),
//   	Protocol: jsii.String("protocol"),
//   	SourceCidrBlock: jsii.String("sourceCidrBlock"),
//   	SourcePortRange: jsii.String("sourcePortRange"),
//   	SourceTransitGatewayAttachmentId: jsii.String("sourceTransitGatewayAttachmentId"),
//   	SourceTransitGatewayAttachmentType: jsii.String("sourceTransitGatewayAttachmentType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html
//
type CfnTransitGatewayMeteringPolicyEntryProps struct {
	// The AWS account ID to which the metered traffic is attributed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-meteredaccount
	//
	MeteredAccount *string `field:"required" json:"meteredAccount" yaml:"meteredAccount"`
	// The rule number of the metering policy entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-policyrulenumber
	//
	PolicyRuleNumber *float64 `field:"required" json:"policyRuleNumber" yaml:"policyRuleNumber"`
	// The ID of the transit gateway metering policy for which the entry is being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-transitgatewaymeteringpolicyid
	//
	TransitGatewayMeteringPolicyId *string `field:"required" json:"transitGatewayMeteringPolicyId" yaml:"transitGatewayMeteringPolicyId"`
	// Describes an IPv4 CIDR block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-destinationcidrblock
	//
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// Describes a range of ports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-destinationportrange
	//
	DestinationPortRange *string `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// The ID of the source attachment through which traffic leaves a transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-destinationtransitgatewayattachmentid
	//
	DestinationTransitGatewayAttachmentId *string `field:"optional" json:"destinationTransitGatewayAttachmentId" yaml:"destinationTransitGatewayAttachmentId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-destinationtransitgatewayattachmenttype
	//
	DestinationTransitGatewayAttachmentType *string `field:"optional" json:"destinationTransitGatewayAttachmentType" yaml:"destinationTransitGatewayAttachmentType"`
	// The protocol of the traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Describes an IPv4 CIDR block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-sourcecidrblock
	//
	SourceCidrBlock *string `field:"optional" json:"sourceCidrBlock" yaml:"sourceCidrBlock"`
	// Describes a range of ports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-sourceportrange
	//
	SourcePortRange *string `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
	// The ID of the source attachment through which traffic enters a transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-sourcetransitgatewayattachmentid
	//
	SourceTransitGatewayAttachmentId *string `field:"optional" json:"sourceTransitGatewayAttachmentId" yaml:"sourceTransitGatewayAttachmentId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-sourcetransitgatewayattachmenttype
	//
	SourceTransitGatewayAttachmentType *string `field:"optional" json:"sourceTransitGatewayAttachmentType" yaml:"sourceTransitGatewayAttachmentType"`
}

