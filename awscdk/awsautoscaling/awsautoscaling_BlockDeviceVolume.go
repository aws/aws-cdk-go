package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Describes a block device mapping for an EC2 instance or Auto Scaling group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockDeviceVolume := awscdk.Aws_autoscaling.BlockDeviceVolume_Ebs(jsii.Number(123), &EbsDeviceOptions{
//   	DeleteOnTermination: jsii.Boolean(false),
//   	Encrypted: jsii.Boolean(false),
//   	Iops: jsii.Number(123),
//   	VolumeType: awscdk.*Aws_autoscaling.EbsDeviceVolumeType_STANDARD,
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
		"monocdk.aws_autoscaling.BlockDeviceVolume",
		[]interface{}{ebsDevice, virtualName},
		&j,
	)

	return &j
}

// Experimental.
func NewBlockDeviceVolume_Override(b BlockDeviceVolume, ebsDevice *EbsDeviceProps, virtualName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_autoscaling.BlockDeviceVolume",
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
		"monocdk.aws_autoscaling.BlockDeviceVolume",
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
		"monocdk.aws_autoscaling.BlockDeviceVolume",
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
		"monocdk.aws_autoscaling.BlockDeviceVolume",
		"ephemeral",
		[]interface{}{volumeIndex},
		&returns,
	)

	return returns
}

// Supresses a volume mapping.
// Experimental.
func BlockDeviceVolume_NoDevice() BlockDeviceVolume {
	_init_.Initialize()

	var returns BlockDeviceVolume

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.BlockDeviceVolume",
		"noDevice",
		nil, // no parameters
		&returns,
	)

	return returns
}

