package awsec2


// Suppresses the specified device included in the block device mapping of the AMI.
//
// To suppress a device, specify an empty string.
//
// `NoDevice` is a property of the [Amazon EC2 BlockDeviceMapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html) property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   noDeviceProperty := &noDeviceProperty{
//   }
//
type CfnInstance_NoDeviceProperty struct {
}

