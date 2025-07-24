package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSnapshot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSnapshotProps := &CfnSnapshotProps{
//   	Name: jsii.String("name"),
//   	VolumeId: jsii.String("volumeId"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-snapshot.html
//
type CfnSnapshotProps struct {
	// The name of the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-snapshot.html#cfn-fsx-snapshot-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the volume that the snapshot is of.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-snapshot.html#cfn-fsx-snapshot-volumeid
	//
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// A list of `Tag` values, with a maximum of 50 elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-snapshot.html#cfn-fsx-snapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

