package awsmemorydb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSnapshotPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSnapshotMixinProps := &CfnSnapshotMixinProps{
//   	ClusterName: jsii.String("clusterName"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	SnapshotName: jsii.String("snapshotName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-snapshot.html
//
type CfnSnapshotMixinProps struct {
	// The name of the cluster from which the snapshot was taken.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-snapshot.html#cfn-memorydb-snapshot-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The ID of the KMS key used to encrypt the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-snapshot.html#cfn-memorydb-snapshot-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-snapshot.html#cfn-memorydb-snapshot-snapshotname
	//
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// A list of tags to be added to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-snapshot.html#cfn-memorydb-snapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

