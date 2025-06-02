package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInstanceSnapshot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceSnapshotProps := &CfnInstanceSnapshotProps{
//   	InstanceName: jsii.String("instanceName"),
//   	InstanceSnapshotName: jsii.String("instanceSnapshotName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instancesnapshot.html
//
type CfnInstanceSnapshotProps struct {
	// The instance from which the snapshot was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instancesnapshot.html#cfn-lightsail-instancesnapshot-instancename
	//
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// The name of the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instancesnapshot.html#cfn-lightsail-instancesnapshot-instancesnapshotname
	//
	InstanceSnapshotName *string `field:"required" json:"instanceSnapshotName" yaml:"instanceSnapshotName"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instancesnapshot.html#cfn-lightsail-instancesnapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

