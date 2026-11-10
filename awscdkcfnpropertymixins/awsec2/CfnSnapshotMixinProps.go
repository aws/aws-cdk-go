package awsec2


// Properties for CfnSnapshotPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSnapshotMixinProps := &CfnSnapshotMixinProps{
//   	Description: jsii.String("description"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VolumeId: jsii.String("volumeId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-snapshot.html
//
type CfnSnapshotMixinProps struct {
	// A description for the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-snapshot.html#cfn-ec2-snapshot-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Any tags assigned to the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-snapshot.html#cfn-ec2-snapshot-tags
	//
	Tags *[]*CfnSnapshotPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the Amazon EBS volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-snapshot.html#cfn-ec2-snapshot-volumeid
	//
	VolumeId *string `field:"optional" json:"volumeId" yaml:"volumeId"`
}

