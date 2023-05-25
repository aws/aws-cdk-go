package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
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
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//   	BlockDevices: []blockDevice{
//   		&blockDevice{
//   			DeviceName: jsii.String("gp3-volume"),
//   			Volume: autoscaling.BlockDeviceVolume_Ebs(jsii.Number(15), &EbsDeviceOptions{
//   				VolumeType: autoscaling.EbsDeviceVolumeType_GP3,
//   				Throughput: jsii.Number(125),
//   			}),
//   		},
//   	},
//   })
//
type BlockDeviceVolume interface {
	// EBS device info.
	EbsDevice() *EbsDeviceProps
	// Virtual device name.
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


func NewBlockDeviceVolume(ebsDevice *EbsDeviceProps, virtualName *string) BlockDeviceVolume {
	_init_.Initialize()

	if err := validateNewBlockDeviceVolumeParameters(ebsDevice); err != nil {
		panic(err)
	}
	j := jsiiProxy_BlockDeviceVolume{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.BlockDeviceVolume",
		[]interface{}{ebsDevice, virtualName},
		&j,
	)

	return &j
}

func NewBlockDeviceVolume_Override(b BlockDeviceVolume, ebsDevice *EbsDeviceProps, virtualName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.BlockDeviceVolume",
		[]interface{}{ebsDevice, virtualName},
		b,
	)
}

// Creates a new Elastic Block Storage device.
func BlockDeviceVolume_Ebs(volumeSize *float64, options *EbsDeviceOptions) BlockDeviceVolume {
	_init_.Initialize()

	if err := validateBlockDeviceVolume_EbsParameters(volumeSize, options); err != nil {
		panic(err)
	}
	var returns BlockDeviceVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.BlockDeviceVolume",
		"ebs",
		[]interface{}{volumeSize, options},
		&returns,
	)

	return returns
}

// Creates a new Elastic Block Storage device from an existing snapshot.
func BlockDeviceVolume_EbsFromSnapshot(snapshotId *string, options *EbsDeviceSnapshotOptions) BlockDeviceVolume {
	_init_.Initialize()

	if err := validateBlockDeviceVolume_EbsFromSnapshotParameters(snapshotId, options); err != nil {
		panic(err)
	}
	var returns BlockDeviceVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.BlockDeviceVolume",
		"ebsFromSnapshot",
		[]interface{}{snapshotId, options},
		&returns,
	)

	return returns
}

// Creates a virtual, ephemeral device.
//
// The name will be in the form ephemeral{volumeIndex}.
func BlockDeviceVolume_Ephemeral(volumeIndex *float64) BlockDeviceVolume {
	_init_.Initialize()

	if err := validateBlockDeviceVolume_EphemeralParameters(volumeIndex); err != nil {
		panic(err)
	}
	var returns BlockDeviceVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.BlockDeviceVolume",
		"ephemeral",
		[]interface{}{volumeIndex},
		&returns,
	)

	return returns
}

// Supresses a volume mapping.
func BlockDeviceVolume_NoDevice() BlockDeviceVolume {
	_init_.Initialize()

	var returns BlockDeviceVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.BlockDeviceVolume",
		"noDevice",
		nil, // no parameters
		&returns,
	)

	return returns
}

