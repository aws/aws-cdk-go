package awsgreengrassv2


// Contains information about a device that Linux processes in a container can access.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaDeviceMountProperty := &lambdaDeviceMountProperty{
//   	addGroupOwner: jsii.Boolean(false),
//   	path: jsii.String("path"),
//   	permission: jsii.String("permission"),
//   }
//
type CfnComponentVersion_LambdaDeviceMountProperty struct {
	// Whether or not to add the component's system user as an owner of the device.
	//
	// Default: `false`.
	AddGroupOwner interface{} `field:"optional" json:"addGroupOwner" yaml:"addGroupOwner"`
	// The mount path for the device in the file system.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The permission to access the device: read/only ( `ro` ) or read/write ( `rw` ).
	//
	// Default: `ro`.
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

