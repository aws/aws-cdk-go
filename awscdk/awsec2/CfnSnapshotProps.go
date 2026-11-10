package awsec2


// Properties for defining a `CfnSnapshot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSnapshotProps := &CfnSnapshotProps{
//   	VolumeId: jsii.String("volumeId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-snapshot.html
//
type CfnSnapshotProps struct {
	// The ID of the Amazon EBS volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-snapshot.html#cfn-ec2-snapshot-volumeid
	//
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// A description for the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-snapshot.html#cfn-ec2-snapshot-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Any tags assigned to the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-snapshot.html#cfn-ec2-snapshot-tags
	//
	Tags *[]*CfnSnapshot_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

