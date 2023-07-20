package awsec2


// Properties for defining a `CfnTrafficMirrorFilterRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficMirrorFilterRuleProps := &CfnTrafficMirrorFilterRuleProps{
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	RuleAction: jsii.String("ruleAction"),
//   	RuleNumber: jsii.Number(123),
//   	SourceCidrBlock: jsii.String("sourceCidrBlock"),
//   	TrafficDirection: jsii.String("trafficDirection"),
//   	TrafficMirrorFilterId: jsii.String("trafficMirrorFilterId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DestinationPortRange: &TrafficMirrorPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   	Protocol: jsii.Number(123),
//   	SourcePortRange: &TrafficMirrorPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html
//
type CfnTrafficMirrorFilterRuleProps struct {
	// The destination CIDR block to assign to the Traffic Mirror rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-destinationcidrblock
	//
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The action to take on the filtered traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-ruleaction
	//
	RuleAction *string `field:"required" json:"ruleAction" yaml:"ruleAction"`
	// The number of the Traffic Mirror rule.
	//
	// This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-rulenumber
	//
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// The source CIDR block to assign to the Traffic Mirror rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-sourcecidrblock
	//
	SourceCidrBlock *string `field:"required" json:"sourceCidrBlock" yaml:"sourceCidrBlock"`
	// The type of traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-trafficdirection
	//
	TrafficDirection *string `field:"required" json:"trafficDirection" yaml:"trafficDirection"`
	// The ID of the filter that this rule is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-trafficmirrorfilterid
	//
	TrafficMirrorFilterId *string `field:"required" json:"trafficMirrorFilterId" yaml:"trafficMirrorFilterId"`
	// The description of the Traffic Mirror rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The destination port range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-destinationportrange
	//
	DestinationPortRange interface{} `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// The protocol, for example UDP, to assign to the Traffic Mirror rule.
	//
	// For information about the protocol value, see [Protocol Numbers](https://docs.aws.amazon.com/https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-protocol
	//
	Protocol *float64 `field:"optional" json:"protocol" yaml:"protocol"`
	// The source port range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-sourceportrange
	//
	SourcePortRange interface{} `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
}

