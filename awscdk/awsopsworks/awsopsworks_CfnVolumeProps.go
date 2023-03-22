package awsopsworks


// Properties for defining a `CfnVolume`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVolumeProps := &cfnVolumeProps{
//   	ec2VolumeId: jsii.String("ec2VolumeId"),
//   	stackId: jsii.String("stackId"),
//
//   	// the properties below are optional
//   	mountPoint: jsii.String("mountPoint"),
//   	name: jsii.String("name"),
//   }
//
type CfnVolumeProps struct {
	// The Amazon EC2 volume ID.
	Ec2VolumeId *string `field:"required" json:"ec2VolumeId" yaml:"ec2VolumeId"`
	// The stack ID.
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
	// The volume mount point.
	//
	// For example, "/mnt/disk1".
	MountPoint *string `field:"optional" json:"mountPoint" yaml:"mountPoint"`
	// The volume name.
	//
	// Volume names are a maximum of 128 characters.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

