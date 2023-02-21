package awsimagebuilder


// Defines block device mappings for the instance used to configure your image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceBlockDeviceMappingProperty := &instanceBlockDeviceMappingProperty{
//   	deviceName: jsii.String("deviceName"),
//   	ebs: &ebsInstanceBlockDeviceSpecificationProperty{
//   		deleteOnTermination: jsii.Boolean(false),
//   		encrypted: jsii.Boolean(false),
//   		iops: jsii.Number(123),
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   		snapshotId: jsii.String("snapshotId"),
//   		throughput: jsii.Number(123),
//   		volumeSize: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//   	},
//   	noDevice: jsii.String("noDevice"),
//   	virtualName: jsii.String("virtualName"),
//   }
//
type CfnImageRecipe_InstanceBlockDeviceMappingProperty struct {
	// The device to which these mappings apply.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Use to manage Amazon EBS-specific configuration for this mapping.
	Ebs interface{} `field:"optional" json:"ebs" yaml:"ebs"`
	// Enter an empty string to remove a mapping from the parent image.
	//
	// The following is an example of an empty string value in the `NoDevice` field.
	//
	// `NoDevice:""`.
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Manages the instance ephemeral devices.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

