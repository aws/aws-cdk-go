package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnNetworkInsightsPath`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkInsightsPathProps := &CfnNetworkInsightsPathProps{
//   	Destination: jsii.String("destination"),
//   	Protocol: jsii.String("protocol"),
//   	Source: jsii.String("source"),
//
//   	// the properties below are optional
//   	DestinationIp: jsii.String("destinationIp"),
//   	DestinationPort: jsii.Number(123),
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
	// The AWS resource that is the destination of the path.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The protocol.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The AWS resource that is the source of the path.
	Source *string `field:"required" json:"source" yaml:"source"`
	// The IP address of the AWS resource that is the destination of the path.
	DestinationIp *string `field:"optional" json:"destinationIp" yaml:"destinationIp"`
	// The destination port.
	DestinationPort *float64 `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// The IP address of the AWS resource that is the source of the path.
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// The tags to add to the path.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

