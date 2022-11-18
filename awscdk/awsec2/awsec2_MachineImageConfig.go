package awsec2


// Configuration for a machine image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userData userData
//
//   machineImageConfig := &machineImageConfig{
//   	imageId: jsii.String("imageId"),
//   	osType: awscdk.Aws_ec2.operatingSystemType_LINUX,
//   	userData: userData,
//   }
//
type MachineImageConfig struct {
	// The AMI ID of the image to use.
	ImageId *string `field:"required" json:"imageId" yaml:"imageId"`
	// Operating system type for this image.
	OsType OperatingSystemType `field:"required" json:"osType" yaml:"osType"`
	// Initial UserData for this image.
	UserData UserData `field:"required" json:"userData" yaml:"userData"`
}

