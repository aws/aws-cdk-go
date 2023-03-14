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
//   cfnNetworkInsightsPathProps := &cfnNetworkInsightsPathProps{
//   	destination: jsii.String("destination"),
//   	protocol: jsii.String("protocol"),
//   	source: jsii.String("source"),
//
//   	// the properties below are optional
//   	destinationIp: jsii.String("destinationIp"),
//   	destinationPort: jsii.Number(123),
//   	sourceIp: jsii.String("sourceIp"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

