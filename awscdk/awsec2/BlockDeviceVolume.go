package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Describes a block device mapping for an EC2 instance or Auto Scaling group.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//   var machineImage iMachineImage
//
//
//   ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//
//   	// ...
//
//   	BlockDevices: []blockDevice{
//   		&blockDevice{
//   			DeviceName: jsii.String("/dev/sda1"),
//   			Volume: ec2.BlockDeviceVolume_Ebs(jsii.Number(50)),
//   		},
//   		&blockDevice{
//   			DeviceName: jsii.String("/dev/sdm"),
//   			Volume: ec2.BlockDeviceVolume_*Ebs(jsii.Number(100)),
//   		},
//   	},
//   })
//
// Experimental.
type BlockDeviceVolume interface {
	// EBS device info.
	// Experimental.
	EbsDevice() *EbsDeviceProps
	// Virtual device name.
	// Experimental.
	VirtualName() *string
}

// The jsii proxy struct for BlockDeviceVolume
type jsiiProxy_BlockDeviceVolume struct {
	_ byte // padding
}

func (j *jsiiProxy_BlockDeviceVolume) EbsDevice() *EbsDeviceProps {
	var returns *EbsDeviceProps
	_jsii_.Get(
		j,
		"ebsDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockDeviceVolume) VirtualName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualName",
		&returns,
	)
	return returns
}


// Experimental.
func NewBlockDeviceVolume(ebsDevice *EbsDeviceProps, virtualName *string) BlockDeviceVolume {
	_init_.Initialize()

	if err := validateNewBlockDeviceVolumeParameters(ebsDevice); err != nil {
		panic(err)
	}
	j := jsiiProxy_BlockDeviceVolume{}

	_jsii_.Create(
		"monocdk.aws_ec2.BlockDeviceVolume",
		[]interface{}{ebsDevice, virtualName},
		&j,
	)

	return &j
}

// Experimental.
func NewBlockDeviceVolume_Override(b BlockDeviceVolume, ebsDevice *EbsDeviceProps, virtualName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.BlockDeviceVolume",
		[]interface{}{ebsDevice, virtualName},
		b,
	)
}

// Creates a new Elastic Block Storage device.
// Experimental.
func BlockDeviceVolume_Ebs(volumeSize *float64, options *EbsDeviceOptions) BlockDeviceVolume {
	_init_.Initialize()

	if err := validateBlockDeviceVolume_EbsParameters(volumeSize, options); err != nil {
		panic(err)
	}
	var returns BlockDeviceVolume

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.BlockDeviceVolume",
		"ebs",
		[]interface{}{volumeSize, options},
		&returns,
	)

	return returns
}

// Creates a new Elastic Block Storage device from an existing snapshot.
// Experimental.
func BlockDeviceVolume_EbsFromSnapshot(snapshotId *string, options *EbsDeviceSnapshotOptions) BlockDeviceVolume {
	_init_.Initialize()

	if err := validateBlockDeviceVolume_EbsFromSnapshotParameters(snapshotId, options); err != nil {
		panic(err)
	}
	var returns BlockDeviceVolume

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.BlockDeviceVolume",
		"ebsFromSnapshot",
		[]interface{}{snapshotId, options},
		&returns,
	)

	return returns
}

// Creates a virtual, ephemeral device.
//
// The name will be in the form ephemeral{volumeIndex}.
// Experimental.
func BlockDeviceVolume_Ephemeral(volumeIndex *float64) BlockDeviceVolume {
	_init_.Initialize()

	if err := validateBlockDeviceVolume_EphemeralParameters(volumeIndex); err != nil {
		panic(err)
	}
	var returns BlockDeviceVolume

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.BlockDeviceVolume",
		"ephemeral",
		[]interface{}{volumeIndex},
		&returns,
	)

	return returns
}

