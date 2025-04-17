package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Volume that can be mounted to a container that uses ECS.
//
// Example:
//   var myFileSystem iFileSystem
//   var myJobRole role
//
//   myFileSystem.GrantRead(myJobRole)
//
//   jobDefn := batch.NewEcsJobDefinition(this, jsii.String("JobDefn"), &EcsJobDefinitionProps{
//   	Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("containerDefn"), &EcsEc2ContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("public.ecr.aws/amazonlinux/amazonlinux:latest")),
//   		Memory: cdk.Size_Mebibytes(jsii.Number(2048)),
//   		Cpu: jsii.Number(256),
//   		Volumes: []ecsVolume{
//   			batch.*ecsVolume_Efs(&EfsVolumeOptions{
//   				Name: jsii.String("myVolume"),
//   				FileSystem: myFileSystem,
//   				ContainerPath: jsii.String("/Volumes/myVolume"),
//   				UseJobRole: jsii.Boolean(true),
//   			}),
//   		},
//   		JobRole: myJobRole,
//   	}),
//   })
//
type EcsVolume interface {
	// The path on the container that this volume will be mounted to.
	ContainerPath() *string
	// The name of this volume.
	Name() *string
	// Whether or not the container has readonly access to this volume.
	// Default: false.
	//
	Readonly() *bool
}

// The jsii proxy struct for EcsVolume
type jsiiProxy_EcsVolume struct {
	_ byte // padding
}

func (j *jsiiProxy_EcsVolume) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsVolume) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


func NewEcsVolume_Override(e EcsVolume, options *EcsVolumeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.EcsVolume",
		[]interface{}{options},
		e,
	)
}

// Creates a Volume that uses an AWS Elastic File System (EFS);
//
// this volume can grow and shrink as needed.
// See: https://docs.aws.amazon.com/batch/latest/userguide/efs-volumes.html
//
func EcsVolume_Efs(options *EfsVolumeOptions) EfsVolume {
	_init_.Initialize()

	if err := validateEcsVolume_EfsParameters(options); err != nil {
		panic(err)
	}
	var returns EfsVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.EcsVolume",
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
func EcsVolume_Host(options *HostVolumeOptions) HostVolume {
	_init_.Initialize()

	if err := validateEcsVolume_HostParameters(options); err != nil {
		panic(err)
	}
	var returns HostVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.EcsVolume",
		"host",
		[]interface{}{options},
		&returns,
	)

	return returns
}

