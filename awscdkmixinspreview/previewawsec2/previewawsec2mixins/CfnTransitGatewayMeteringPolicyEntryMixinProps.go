package previewawsec2mixins


// Properties for CfnTransitGatewayMeteringPolicyEntryPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransitGatewayMeteringPolicyEntryMixinProps := &CfnTransitGatewayMeteringPolicyEntryMixinProps{
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	DestinationPortRange: jsii.String("destinationPortRange"),
//   	DestinationTransitGatewayAttachmentId: jsii.String("destinationTransitGatewayAttachmentId"),
//   	DestinationTransitGatewayAttachmentType: jsii.String("destinationTransitGatewayAttachmentType"),
//   	MeteredAccount: jsii.String("meteredAccount"),
//   	PolicyRuleNumber: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	SourceCidrBlock: jsii.String("sourceCidrBlock"),
//   	SourcePortRange: jsii.String("sourcePortRange"),
//   	SourceTransitGatewayAttachmentId: jsii.String("sourceTransitGatewayAttachmentId"),
//   	SourceTransitGatewayAttachmentType: jsii.String("sourceTransitGatewayAttachmentType"),
//   	TransitGatewayMeteringPolicyId: jsii.String("transitGatewayMeteringPolicyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html
//
type CfnTransitGatewayMeteringPolicyEntryMixinProps struct {
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
	// The AWS account ID to which the metered traffic is attributed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-meteredaccount
	//
	MeteredAccount *string `field:"optional" json:"meteredAccount" yaml:"meteredAccount"`
	// The rule number of the metering policy entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-policyrulenumber
	//
	PolicyRuleNumber *float64 `field:"optional" json:"policyRuleNumber" yaml:"policyRuleNumber"`
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
	// The ID of the transit gateway metering policy for which the entry is being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html#cfn-ec2-transitgatewaymeteringpolicyentry-transitgatewaymeteringpolicyid
	//
	TransitGatewayMeteringPolicyId *string `field:"optional" json:"transitGatewayMeteringPolicyId" yaml:"transitGatewayMeteringPolicyId"`
}

