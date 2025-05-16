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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html
//
type CfnNetworkInsightsPathProps struct {
	// The protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The ID or ARN of the source.
	//
	// If the resource is in another account, you must specify an ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-source
	//
	Source *string `field:"required" json:"source" yaml:"source"`
	// The ID or ARN of the destination.
	//
	// If the resource is in another account, you must specify an ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-destination
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The IP address of the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-destinationip
	//
	DestinationIp *string `field:"optional" json:"destinationIp" yaml:"destinationIp"`
	// The destination port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-destinationport
	//
	DestinationPort *float64 `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// Scopes the analysis to network paths that match specific filters at the destination.
	//
	// If you specify this parameter, you can't specify the parameter for the destination IP address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-filteratdestination
	//
	FilterAtDestination interface{} `field:"optional" json:"filterAtDestination" yaml:"filterAtDestination"`
	// Scopes the analysis to network paths that match specific filters at the source.
	//
	// If you specify this parameter, you can't specify the parameters for the source IP address or the destination port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-filteratsource
	//
	FilterAtSource interface{} `field:"optional" json:"filterAtSource" yaml:"filterAtSource"`
	// The IP address of the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-sourceip
	//
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// The tags to add to the path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html#cfn-ec2-networkinsightspath-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

