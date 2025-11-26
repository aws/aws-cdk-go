package previewawslightsailmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnInstanceSnapshotPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceSnapshotMixinProps := &CfnInstanceSnapshotMixinProps{
//   	InstanceName: jsii.String("instanceName"),
//   	InstanceSnapshotName: jsii.String("instanceSnapshotName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instancesnapshot.html
//
type CfnInstanceSnapshotMixinProps struct {
	// The name the user gave the instance ( `Amazon_Linux_2023-1` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instancesnapshot.html#cfn-lightsail-instancesnapshot-instancename
	//
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// The name of the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instancesnapshot.html#cfn-lightsail-instancesnapshot-instancesnapshotname
	//
	InstanceSnapshotName *string `field:"optional" json:"instanceSnapshotName" yaml:"instanceSnapshotName"`
	// The tag keys and optional values for the resource.
	//
	// For more information about tags in Lightsail, see the [Amazon Lightsail Developer Guide](https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-tags) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instancesnapshot.html#cfn-lightsail-instancesnapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

