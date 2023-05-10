package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnNetworkInsightsPath`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkInsightsPathProps := &CfnNetworkInsightsPathProps{
//   	Protocol: jsii.String("protocol"),
//   	Source: jsii.String("source"),
//
//   	// the properties below are optional
//   	Destination: jsii.String("destination"),
//   	DestinationIp: jsii.String("destinationIp"),
//   	DestinationPort: jsii.Number(123),
//   	FilterAtDestination: &PathFilterProperty{
//   		DestinationAddress: jsii.String("destinationAddress"),
//   		DestinationPortRange: &FilterPortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			ToPort: jsii.Number(123),
//   		},
//   		SourceAddress: jsii.String("sourceAddress"),
//   		SourcePortRange: &FilterPortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	FilterAtSource: &PathFilterProperty{
//   		DestinationAddress: jsii.String("destinationAddress"),
//   		DestinationPortRange: &FilterPortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			ToPort: jsii.Number(123),
//   		},
//   		SourceAddress: jsii.String("sourceAddress"),
//   		SourcePortRange: &FilterPortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	SourceIp: jsii.String("sourceIp"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnNetworkInsightsPathProps struct {
	// The protocol.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The ID or ARN of the source.
	//
	// If the resource is in another account, you must specify an ARN.
	Source *string `field:"required" json:"source" yaml:"source"`
	// The ID or ARN of the destination.
	//
	// If the resource is in another account, you must specify an ARN.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The IP address of the destination.
	DestinationIp *string `field:"optional" json:"destinationIp" yaml:"destinationIp"`
	// The destination port.
	DestinationPort *float64 `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// `AWS::EC2::NetworkInsightsPath.FilterAtDestination`.
	FilterAtDestination interface{} `field:"optional" json:"filterAtDestination" yaml:"filterAtDestination"`
	// `AWS::EC2::NetworkInsightsPath.FilterAtSource`.
	FilterAtSource interface{} `field:"optional" json:"filterAtSource" yaml:"filterAtSource"`
	// The IP address of the source.
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// The tags to add to the path.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

