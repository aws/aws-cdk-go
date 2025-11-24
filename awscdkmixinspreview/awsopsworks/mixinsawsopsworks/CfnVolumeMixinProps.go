package mixinsawsopsworks


// Properties for CfnVolumePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVolumeMixinProps := &CfnVolumeMixinProps{
//   	Ec2VolumeId: jsii.String("ec2VolumeId"),
//   	MountPoint: jsii.String("mountPoint"),
//   	Name: jsii.String("name"),
//   	StackId: jsii.String("stackId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html
//
type CfnVolumeMixinProps struct {
	// The Amazon EC2 volume ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-ec2volumeid
	//
	Ec2VolumeId *string `field:"optional" json:"ec2VolumeId" yaml:"ec2VolumeId"`
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
	// The stack ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-stackid
	//
	StackId *string `field:"optional" json:"stackId" yaml:"stackId"`
}

