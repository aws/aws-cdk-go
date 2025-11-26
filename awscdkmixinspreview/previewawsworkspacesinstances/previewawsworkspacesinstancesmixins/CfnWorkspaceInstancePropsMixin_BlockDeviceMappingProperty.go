package previewawsworkspacesinstancesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   blockDeviceMappingProperty := &BlockDeviceMappingProperty{
//   	DeviceName: jsii.String("deviceName"),
//   	Ebs: &EbsBlockDeviceProperty{
//   		Encrypted: jsii.Boolean(false),
//   		Iops: jsii.Number(123),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		Throughput: jsii.Number(123),
//   		VolumeSize: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	NoDevice: jsii.String("noDevice"),
//   	VirtualName: jsii.String("virtualName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-blockdevicemapping.html
//
type CfnWorkspaceInstancePropsMixin_BlockDeviceMappingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-blockdevicemapping.html#cfn-workspacesinstances-workspaceinstance-blockdevicemapping-devicename
	//
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-blockdevicemapping.html#cfn-workspacesinstances-workspaceinstance-blockdevicemapping-ebs
	//
	Ebs interface{} `field:"optional" json:"ebs" yaml:"ebs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-blockdevicemapping.html#cfn-workspacesinstances-workspaceinstance-blockdevicemapping-nodevice
	//
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-blockdevicemapping.html#cfn-workspacesinstances-workspaceinstance-blockdevicemapping-virtualname
	//
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

