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
//   cfnPlacementGroupProps := &CfnPlacementGroupProps{
//   	PartitionCount: jsii.Number(123),
//   	SpreadLevel: jsii.String("spreadLevel"),
//   	Strategy: jsii.String("strategy"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html
//
type CfnPlacementGroupProps struct {
	// The number of partitions.
	//
	// Valid only when *Strategy* is set to `partition` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html#cfn-ec2-placementgroup-partitioncount
	//
	PartitionCount *float64 `field:"optional" json:"partitionCount" yaml:"partitionCount"`
	// Determines how placement groups spread instances.
	//
	// - Host – You can use `host` only with Outpost placement groups.
	// - Rack – No usage restrictions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html#cfn-ec2-placementgroup-spreadlevel
	//
	SpreadLevel *string `field:"optional" json:"spreadLevel" yaml:"spreadLevel"`
	// The placement strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html#cfn-ec2-placementgroup-strategy
	//
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
	// The tags to apply to the new placement group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html#cfn-ec2-placementgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

