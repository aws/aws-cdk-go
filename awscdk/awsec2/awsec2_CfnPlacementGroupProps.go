package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPlacementGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPlacementGroupProps := &cfnPlacementGroupProps{
//   	partitionCount: jsii.Number(123),
//   	spreadLevel: jsii.String("spreadLevel"),
//   	strategy: jsii.String("strategy"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPlacementGroupProps struct {
	// `AWS::EC2::PlacementGroup.PartitionCount`.
	PartitionCount *float64 `field:"optional" json:"partitionCount" yaml:"partitionCount"`
	// `AWS::EC2::PlacementGroup.SpreadLevel`.
	SpreadLevel *string `field:"optional" json:"spreadLevel" yaml:"spreadLevel"`
	// The placement strategy.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
	// `AWS::EC2::PlacementGroup.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

