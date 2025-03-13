package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs"
)

// A Volume that uses an AWS Elastic File System (EFS);
//
// this volume can grow and shrink as needed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fileSystem fileSystem
//
//   efsVolume := awscdk.Aws_batch.NewEfsVolume(&EfsVolumeOptions{
//   	ContainerPath: jsii.String("containerPath"),
//   	FileSystem: fileSystem,
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AccessPointId: jsii.String("accessPointId"),
//   	EnableTransitEncryption: jsii.Boolean(false),
//   	Readonly: jsii.Boolean(false),
//   	RootDirectory: jsii.String("rootDirectory"),
//   	TransitEncryptionPort: jsii.Number(123),
//   	UseJobRole: jsii.Boolean(false),
//   })
//
type EfsVolume interface {
	EcsVolume
	// The Amazon EFS access point ID to use.
	//
	// If an access point is specified, `rootDirectory` must either be omitted or set to `/`
	// which enforces the path set on the EFS access point.
	// If an access point is used, `enableTransitEncryption` must be `true`.
	// See: https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html
	//
	// Default: - no accessPointId.
	//
	AccessPointId() *string
	// The path on the container that this volume will be mounted to.
	ContainerPath() *string
	// Enables encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server.
	// See: https://docs.aws.amazon.com/efs/latest/ug/encryption-in-transit.html
	//
	// Default: false.
	//
	EnableTransitEncryption() *bool
	// The EFS File System that supports this volume.
	FileSystem() awsefs.IFileSystem
	// The name of this volume.
	Name() *string
	// Whether or not the container has readonly access to this volume.
	// Default: false.
	//
	Readonly() *bool
	// The directory within the Amazon EFS file system to mount as the root directory inside the host.
	//
	// If this parameter is omitted, the root of the Amazon EFS volume is used instead.
	// Specifying `/` has the same effect as omitting this parameter.
	// The maximum length is 4,096 characters.
	// Default: - root of the EFS File System.
	//
	RootDirectory() *string
	// The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server.
	//
	// The value must be between 0 and 65,535.
	// See: https://docs.aws.amazon.com/efs/latest/ug/efs-mount-helper.html
	//
	// Default: - chosen by the EFS Mount Helper.
	//
	TransitEncryptionPort() *float64
	// Whether or not to use the AWS Batch job IAM role defined in a job definition when mounting the Amazon EFS file system.
	//
	// If specified, `enableTransitEncryption` must be `true`.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/efs-volumes.html#efs-volume-accesspoints
	//
	// Default: false.
	//
	UseJobRole() *bool
}

// The jsii proxy struct for EfsVolume
type jsiiProxy_EfsVolume struct {
	jsiiProxy_EcsVolume
}

func (j *jsiiProxy_EfsVolume) AccessPointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsVolume) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsVolume) EnableTransitEncryption() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableTransitEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsVolume) FileSystem() awsefs.IFileSystem {
	var returns awsefs.IFileSystem
	_jsii_.Get(
		j,
		"fileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsVolume) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsVolume) RootDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsVolume) TransitEncryptionPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transitEncryptionPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsVolume) UseJobRole() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"useJobRole",
		&returns,
	)
	return returns
}


func NewEfsVolume(options *EfsVolumeOptions) EfsVolume {
	_init_.Initialize()

	if err := validateNewEfsVolumeParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_EfsVolume{}

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.EfsVolume",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewEfsVolume_Override(e EfsVolume, options *EfsVolumeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.EfsVolume",
		[]interface{}{options},
		e,
	)
}

// Creates a Volume that uses an AWS Elastic File System (EFS);
//
// this volume can grow and shrink as needed.
// See: https://docs.aws.amazon.com/batch/latest/userguide/efs-volumes.html
//
func EfsVolume_Efs(options *EfsVolumeOptions) EfsVolume {
	_init_.Initialize()

	if err := validateEfsVolume_EfsParameters(options); err != nil {
		panic(err)
	}
	var returns EfsVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.EfsVolume",
		"efs",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a Host volume.
//
// This volume will persist on the host at the specified `hostPath`.
// If the `hostPath` is not specified, Docker will choose the host path. In this case,
// the data may not persist after the containers that use it stop running.
func EfsVolume_Host(options *HostVolumeOptions) HostVolume {
	_init_.Initialize()

	if err := validateEfsVolume_HostParameters(options); err != nil {
		panic(err)
	}
	var returns HostVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.EfsVolume",
		"host",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Returns true if x is an EfsVolume, false otherwise.
func EfsVolume_IsEfsVolume(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEfsVolume_IsEfsVolumeParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.EfsVolume",
		"isEfsVolume",
		[]interface{}{x},
		&returns,
	)

	return returns
}

