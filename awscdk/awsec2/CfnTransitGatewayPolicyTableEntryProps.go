package awsec2


// Properties for defining a `CfnTransitGatewayPolicyTableEntry`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayPolicyTableEntryProps := &CfnTransitGatewayPolicyTableEntryProps{
//   	PolicyRule: &TransitGatewayPolicyRuleProperty{
//   		DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   		DestinationPortRange: jsii.String("destinationPortRange"),
//   		Protocol: jsii.String("protocol"),
//   		SourceCidrBlock: jsii.String("sourceCidrBlock"),
//   		SourcePortRange: jsii.String("sourcePortRange"),
//   	},
//   	PolicyRuleNumber: jsii.String("policyRuleNumber"),
//   	TargetRouteTableId: jsii.String("targetRouteTableId"),
//   	TransitGatewayPolicyTableId: jsii.String("transitGatewayPolicyTableId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypolicytableentry.html
//
type CfnTransitGatewayPolicyTableEntryProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypolicytableentry.html#cfn-ec2-transitgatewaypolicytableentry-policyrule
	//
	PolicyRule interface{} `field:"required" json:"policyRule" yaml:"policyRule"`
	// The rule number for the policy table entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypolicytableentry.html#cfn-ec2-transitgatewaypolicytableentry-policyrulenumber
	//
	PolicyRuleNumber *string `field:"required" json:"policyRuleNumber" yaml:"policyRuleNumber"`
	// The ID of the target route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypolicytableentry.html#cfn-ec2-transitgatewaypolicytableentry-targetroutetableid
	//
	TargetRouteTableId *string `field:"required" json:"targetRouteTableId" yaml:"targetRouteTableId"`
	// The ID of the transit gateway policy table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypolicytableentry.html#cfn-ec2-transitgatewaypolicytableentry-transitgatewaypolicytableid
	//
	TransitGatewayPolicyTableId *string `field:"required" json:"transitGatewayPolicyTableId" yaml:"transitGatewayPolicyTableId"`
}

