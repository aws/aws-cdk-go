package previewawsec2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTrafficMirrorFilterRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTrafficMirrorFilterRuleMixinProps := &CfnTrafficMirrorFilterRuleMixinProps{
//   	Description: jsii.String("description"),
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	DestinationPortRange: &TrafficMirrorPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   	Protocol: jsii.Number(123),
//   	RuleAction: jsii.String("ruleAction"),
//   	RuleNumber: jsii.Number(123),
//   	SourceCidrBlock: jsii.String("sourceCidrBlock"),
//   	SourcePortRange: &TrafficMirrorPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrafficDirection: jsii.String("trafficDirection"),
//   	TrafficMirrorFilterId: jsii.String("trafficMirrorFilterId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html
//
type CfnTrafficMirrorFilterRuleMixinProps struct {
	// The description of the Traffic Mirror rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The destination CIDR block to assign to the Traffic Mirror rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-destinationcidrblock
	//
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
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
	// The action to take on the filtered traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-ruleaction
	//
	RuleAction *string `field:"optional" json:"ruleAction" yaml:"ruleAction"`
	// The number of the Traffic Mirror rule.
	//
	// This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-rulenumber
	//
	RuleNumber *float64 `field:"optional" json:"ruleNumber" yaml:"ruleNumber"`
	// The source CIDR block to assign to the Traffic Mirror rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-sourcecidrblock
	//
	SourceCidrBlock *string `field:"optional" json:"sourceCidrBlock" yaml:"sourceCidrBlock"`
	// The source port range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-sourceportrange
	//
	SourcePortRange interface{} `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
	// Tags on Traffic Mirroring filter rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-trafficdirection
	//
	TrafficDirection *string `field:"optional" json:"trafficDirection" yaml:"trafficDirection"`
	// The ID of the filter that this rule is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html#cfn-ec2-trafficmirrorfilterrule-trafficmirrorfilterid
	//
	TrafficMirrorFilterId *string `field:"optional" json:"trafficMirrorFilterId" yaml:"trafficMirrorFilterId"`
}

