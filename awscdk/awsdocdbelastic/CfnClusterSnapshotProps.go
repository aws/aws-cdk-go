package awsdocdbelastic

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnClusterSnapshot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterSnapshotProps := &CfnClusterSnapshotProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	SnapshotName: jsii.String("snapshotName"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-clustersnapshot.html
//
type CfnClusterSnapshotProps struct {
	// The ARN of the elastic cluster of which to create a snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-clustersnapshot.html#cfn-docdbelastic-clustersnapshot-clusterarn
	//
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// The name of the elastic cluster snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-clustersnapshot.html#cfn-docdbelastic-clustersnapshot-snapshotname
	//
	SnapshotName *string `field:"required" json:"snapshotName" yaml:"snapshotName"`
	// An array of key-value pairs to apply to the elastic cluster snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-clustersnapshot.html#cfn-docdbelastic-clustersnapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

