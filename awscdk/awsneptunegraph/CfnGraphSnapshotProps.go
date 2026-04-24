package awsneptunegraph

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGraphSnapshot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGraphSnapshotProps := &CfnGraphSnapshotProps{
//   	SnapshotName: jsii.String("snapshotName"),
//
//   	// the properties below are optional
//   	GraphIdentifier: jsii.String("graphIdentifier"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graphsnapshot.html
//
type CfnGraphSnapshotProps struct {
	// The snapshot name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graphsnapshot.html#cfn-neptunegraph-graphsnapshot-snapshotname
	//
	SnapshotName *string `field:"required" json:"snapshotName" yaml:"snapshotName"`
	// The unique identifier of the Neptune Analytics graph to create the snapshot from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graphsnapshot.html#cfn-neptunegraph-graphsnapshot-graphidentifier
	//
	GraphIdentifier *string `field:"optional" json:"graphIdentifier" yaml:"graphIdentifier"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graphsnapshot.html#cfn-neptunegraph-graphsnapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

