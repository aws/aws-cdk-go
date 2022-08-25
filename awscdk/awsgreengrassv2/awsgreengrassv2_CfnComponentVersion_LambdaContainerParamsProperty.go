package awsgreengrassv2


// Contains information about a container in which AWS Lambda functions run on Greengrass core devices.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaContainerParamsProperty := &lambdaContainerParamsProperty{
//   	devices: []interface{}{
//   		&lambdaDeviceMountProperty{
//   			addGroupOwner: jsii.Boolean(false),
//   			path: jsii.String("path"),
//   			permission: jsii.String("permission"),
//   		},
//   	},
//   	memorySizeInKb: jsii.Number(123),
//   	mountRoSysfs: jsii.Boolean(false),
//   	volumes: []interface{}{
//   		&lambdaVolumeMountProperty{
//   			addGroupOwner: jsii.Boolean(false),
//   			destinationPath: jsii.String("destinationPath"),
//   			permission: jsii.String("permission"),
//   			sourcePath: jsii.String("sourcePath"),
//   		},
//   	},
//   }
//
type CfnComponentVersion_LambdaContainerParamsProperty struct {
	// The list of system devices that the container can access.
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// The memory size of the container, expressed in kilobytes.
	//
	// Default: `16384` (16 MB).
	MemorySizeInKb *float64 `field:"optional" json:"memorySizeInKb" yaml:"memorySizeInKb"`
	// Whether or not the container can read information from the device's `/sys` folder.
	//
	// Default: `false`.
	MountRoSysfs interface{} `field:"optional" json:"mountRoSysfs" yaml:"mountRoSysfs"`
	// The list of volumes that the container can access.
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

