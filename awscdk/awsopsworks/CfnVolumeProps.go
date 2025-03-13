package awsopsworks


// Properties for defining a `CfnVolume`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVolumeProps := &CfnVolumeProps{
//   	Ec2VolumeId: jsii.String("ec2VolumeId"),
//   	StackId: jsii.String("stackId"),
//
//   	// the properties below are optional
//   	MountPoint: jsii.String("mountPoint"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html
//
type CfnVolumeProps struct {
	// The Amazon EC2 volume ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-ec2volumeid
	//
	Ec2VolumeId *string `field:"required" json:"ec2VolumeId" yaml:"ec2VolumeId"`
	// The stack ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-stackid
	//
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
	// The volume mount point.
	//
	// For example, "/mnt/disk1".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-mountpoint
	//
	MountPoint *string `field:"optional" json:"mountPoint" yaml:"mountPoint"`
	// The volume name.
	//
	// Volume names are a maximum of 128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

