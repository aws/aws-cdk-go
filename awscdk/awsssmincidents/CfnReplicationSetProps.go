package awsssmincidents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnReplicationSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationSetProps := &CfnReplicationSetProps{
//   	Regions: []interface{}{
//   		&ReplicationRegionProperty{
//   			RegionConfiguration: &RegionConfigurationProperty{
//   				SseKmsKeyId: jsii.String("sseKmsKeyId"),
//   			},
//   			RegionName: jsii.String("regionName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	DeletionProtected: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-replicationset.html
//
type CfnReplicationSetProps struct {
	// Specifies the Regions of the replication set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-replicationset.html#cfn-ssmincidents-replicationset-regions
	//
	Regions interface{} `field:"required" json:"regions" yaml:"regions"`
	// Determines if the replication set deletion protection is enabled or not.
	//
	// If deletion protection is enabled, you can't delete the last Region in the replication set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-replicationset.html#cfn-ssmincidents-replicationset-deletionprotected
	//
	DeletionProtected interface{} `field:"optional" json:"deletionProtected" yaml:"deletionProtected"`
	// A list of tags to add to the replication set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-replicationset.html#cfn-ssmincidents-replicationset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

