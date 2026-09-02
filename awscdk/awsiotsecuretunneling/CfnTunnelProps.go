package awsiotsecuretunneling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTunnel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTunnelProps := &CfnTunnelProps{
//   	Description: jsii.String("description"),
//   	DestinationConfig: &DestinationConfigProperty{
//   		Services: []*string{
//   			jsii.String("services"),
//   		},
//
//   		// the properties below are optional
//   		ThingName: jsii.String("thingName"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeoutConfig: &TimeoutConfigProperty{
//   		MaxLifetimeTimeoutMinutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsecuretunneling-tunnel.html
//
type CfnTunnelProps struct {
	// A short text description of the tunnel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsecuretunneling-tunnel.html#cfn-iotsecuretunneling-tunnel-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The destination configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsecuretunneling-tunnel.html#cfn-iotsecuretunneling-tunnel-destinationconfig
	//
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// A collection of tag metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsecuretunneling-tunnel.html#cfn-iotsecuretunneling-tunnel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Tunnel timeout configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsecuretunneling-tunnel.html#cfn-iotsecuretunneling-tunnel-timeoutconfig
	//
	TimeoutConfig interface{} `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
}

