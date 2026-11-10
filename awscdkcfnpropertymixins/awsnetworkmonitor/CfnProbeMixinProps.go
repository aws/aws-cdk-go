package awsnetworkmonitor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnProbePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnProbeMixinProps := &CfnProbeMixinProps{
//   	Destination: jsii.String("destination"),
//   	DestinationPort: jsii.Number(123),
//   	MonitorName: jsii.String("monitorName"),
//   	PacketSize: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	SourceArn: jsii.String("sourceArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmonitor-probe.html
//
type CfnProbeMixinProps struct {
	// The literal IPv4 or IPv6 address the probe sends traffic to.
	//
	// The service rejects a domain name and a CIDR block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmonitor-probe.html#cfn-networkmonitor-probe-destination
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The port the probe sends traffic to.
	//
	// Applies to the TCP protocol only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmonitor-probe.html#cfn-networkmonitor-probe-destinationport
	//
	DestinationPort *float64 `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// The name of the monitor the probe belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmonitor-probe.html#cfn-networkmonitor-probe-monitorname
	//
	MonitorName *string `field:"optional" json:"monitorName" yaml:"monitorName"`
	// The size, in bytes, of the packets the probe sends.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmonitor-probe.html#cfn-networkmonitor-probe-packetsize
	//
	PacketSize *float64 `field:"optional" json:"packetSize" yaml:"packetSize"`
	// The protocol the probe uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmonitor-probe.html#cfn-networkmonitor-probe-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The ARN of the subnet the probe sends traffic from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmonitor-probe.html#cfn-networkmonitor-probe-sourcearn
	//
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
	// The tags for the probe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmonitor-probe.html#cfn-networkmonitor-probe-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

