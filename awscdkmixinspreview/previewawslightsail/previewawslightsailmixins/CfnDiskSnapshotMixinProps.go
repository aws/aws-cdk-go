package previewawslightsailmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDiskSnapshotPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDiskSnapshotMixinProps := &CfnDiskSnapshotMixinProps{
//   	DiskName: jsii.String("diskName"),
//   	DiskSnapshotName: jsii.String("diskSnapshotName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disksnapshot.html
//
type CfnDiskSnapshotMixinProps struct {
	// The unique name of the disk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disksnapshot.html#cfn-lightsail-disksnapshot-diskname
	//
	DiskName *string `field:"optional" json:"diskName" yaml:"diskName"`
	// The name of the disk snapshot ( `my-disk-snapshot` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disksnapshot.html#cfn-lightsail-disksnapshot-disksnapshotname
	//
	DiskSnapshotName *string `field:"optional" json:"diskSnapshotName" yaml:"diskSnapshotName"`
	// The tag keys and optional values for the resource.
	//
	// For more information about tags in Lightsail, see the [Amazon Lightsail Developer Guide](https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-tags) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disksnapshot.html#cfn-lightsail-disksnapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

