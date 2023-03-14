package awsimagebuilder


// Defines a custom base AMI and block device mapping configurations of an instance used for building and testing container images.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceConfigurationProperty := &instanceConfigurationProperty{
//   	blockDeviceMappings: []interface{}{
//   		&instanceBlockDeviceMappingProperty{
//   			deviceName: jsii.String("deviceName"),
//   			ebs: &ebsInstanceBlockDeviceSpecificationProperty{
//   				deleteOnTermination: jsii.Boolean(false),
//   				encrypted: jsii.Boolean(false),
//   				iops: jsii.Number(123),
//   				kmsKeyId: jsii.String("kmsKeyId"),
//   				snapshotId: jsii.String("snapshotId"),
//   				throughput: jsii.Number(123),
//   				volumeSize: jsii.Number(123),
//   				volumeType: jsii.String("volumeType"),
//   			},
//   			noDevice: jsii.String("noDevice"),
//   			virtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	image: jsii.String("image"),
//   }
//
type CfnContainerRecipe_InstanceConfigurationProperty struct {
	// Defines the block devices to attach for building an instance from this Image Builder AMI.
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// The AMI ID to use as the base image for a container build and test instance.
	//
	// If not specified, Image Builder will use the appropriate ECS-optimized AMI as a base image.
	Image *string `field:"optional" json:"image" yaml:"image"`
}

